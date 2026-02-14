#include <string.h>
#include <stdlib.h>
#include <ctype.h>
#include <stdio.h>

#include <tree_sitter/api.h>
#include "workout.h"

/* ---- tree-sitter language extern ---- */
extern const TSLanguage *tree_sitter_workout(void);

/* ---- Constants ---- */
#define DEFAULT_REPS 8
#define DEFAULT_SETS 3
#define REP_THRESHOLD 20

/* ---- Token Result types ---- */
typedef enum {
    TR_SETS,          /* fully resolved sets */
    TR_PENDING_REPS,  /* reps waiting for a weight */
    TR_WEIGHT,        /* a single weight value */
    TR_REPS,          /* a single rep value */
    TR_NOTE           /* a note string */
} TokenResultKind;

typedef struct {
    TokenResultKind kind;
    union {
        struct {
            WorkoutSet items[WORKOUT_MAX_SETS];
            int count;
        } sets;
        struct {
            int items[WORKOUT_MAX_SETS];
            int count;
        } pending_reps;
        int weight;
        int reps;
        char note[WORKOUT_NOTE_MAX];
    } data;
} TokenResult;

#define MAX_TOKEN_RESULTS 128

/* ---- Helper: extract text from a node ---- */
static void node_text(TSNode node, const char *src, char *buf, int bufsize) {
    uint32_t start = ts_node_start_byte(node);
    uint32_t end = ts_node_end_byte(node);
    int len = (int)(end - start);
    if (len >= bufsize) len = bufsize - 1;
    memcpy(buf, src + start, len);
    buf[len] = '\0';
}

/* ---- Helper: parse an integer from a node ---- */
static int node_int(TSNode node, const char *src) {
    char buf[64];
    node_text(node, src, buf, sizeof(buf));
    return atoi(buf);
}

/* ---- Helper: title case ---- */
static void title_case(const char *input, char *out, int outsize) {
    int i = 0, o = 0;
    int word_start = 1;
    while (input[i] && o < outsize - 1) {
        if (input[i] == ' ' || input[i] == '\t') {
            out[o++] = ' ';
            word_start = 1;
            i++;
            /* skip extra spaces */
            while (input[i] == ' ' || input[i] == '\t') i++;
        } else {
            if (word_start) {
                out[o++] = (char)toupper((unsigned char)input[i]);
                word_start = 0;
            } else {
                out[o++] = (char)tolower((unsigned char)input[i]);
            }
            i++;
        }
    }
    out[o] = '\0';
}

/* ---- Helper: join notes ---- */
static void join_notes(char *existing, const char *new_note, int bufsize) {
    if (existing[0] == '\0') {
        strncpy(existing, new_note, bufsize - 1);
        existing[bufsize - 1] = '\0';
    } else if (new_note[0] != '\0') {
        int len = (int)strlen(existing);
        if (len < bufsize - 2) {
            existing[len] = ' ';
            strncpy(existing + len + 1, new_note, bufsize - len - 2);
            existing[bufsize - 1] = '\0';
        }
    }
}

/* ---- Helper: add a set to an exercise ---- */
static void exercise_add_set(Exercise *ex, int reps, int weight, const char *note) {
    if (ex->set_count >= WORKOUT_MAX_SETS) return;
    WorkoutSet *s = &ex->sets[ex->set_count++];
    s->reps = reps;
    s->weight = weight;
    s->note[0] = '\0';
    if (note && note[0]) {
        strncpy(s->note, note, WORKOUT_NOTE_MAX - 1);
        s->note[WORKOUT_NOTE_MAX - 1] = '\0';
    }
}

/* ---- Helper: check if a node type matches ---- */
static int is_type(TSNode node, const char *type) {
    return strcmp(ts_node_type(node), type) == 0;
}

/* ---- Walk a numeric_token node and produce a TokenResult ---- */
static TokenResult walk_numeric_token(TSNode node, const char *src) {
    TokenResult tr;
    memset(&tr, 0, sizeof(tr));

    /* numeric_token wraps: by_expr | multiplier | number */
    /* by_expr wraps: three_part_by | two_part_by */
    /* multiplier wraps: full_multiplier | partial_multiplier */

    /* Drill into the child to find the actual node */
    uint32_t ncount = ts_node_named_child_count(node);
    TSNode child = (ncount > 0) ? ts_node_named_child(node, 0) : node;

    /* by_expr wraps another level */
    if (is_type(child, "by_expr")) {
        uint32_t cn = ts_node_named_child_count(child);
        if (cn > 0) child = ts_node_named_child(child, 0);
    }
    /* multiplier wraps another level */
    if (is_type(child, "multiplier")) {
        uint32_t cn = ts_node_named_child_count(child);
        if (cn > 0) child = ts_node_named_child(child, 0);
    }

    if (is_type(child, "two_part_by")) {
        TSNode reps_node = ts_node_child_by_field_name(child, "reps", 4);
        TSNode weight_node = ts_node_child_by_field_name(child, "weight", 6);
        int reps = ts_node_is_null(reps_node) ? DEFAULT_REPS : node_int(reps_node, src);
        int weight = ts_node_is_null(weight_node) ? 0 : node_int(weight_node, src);
        tr.kind = TR_SETS;
        for (int i = 0; i < DEFAULT_SETS && i < WORKOUT_MAX_SETS; i++) {
            tr.data.sets.items[i].reps = reps;
            tr.data.sets.items[i].weight = weight;
            tr.data.sets.items[i].note[0] = '\0';
        }
        tr.data.sets.count = DEFAULT_SETS;
    } else if (is_type(child, "three_part_by")) {
        TSNode sets_node = ts_node_child_by_field_name(child, "sets", 4);
        TSNode reps_node = ts_node_child_by_field_name(child, "reps", 4);
        TSNode weight_node = ts_node_child_by_field_name(child, "weight", 6);
        int nsets = ts_node_is_null(sets_node) ? DEFAULT_SETS : node_int(sets_node, src);
        int reps = ts_node_is_null(reps_node) ? DEFAULT_REPS : node_int(reps_node, src);
        int weight = ts_node_is_null(weight_node) ? 0 : node_int(weight_node, src);
        tr.kind = TR_SETS;
        if (nsets > WORKOUT_MAX_SETS) nsets = WORKOUT_MAX_SETS;
        for (int i = 0; i < nsets; i++) {
            tr.data.sets.items[i].reps = reps;
            tr.data.sets.items[i].weight = weight;
            tr.data.sets.items[i].note[0] = '\0';
        }
        tr.data.sets.count = nsets;
    } else if (is_type(child, "full_multiplier")) {
        TSNode sets_node = ts_node_child_by_field_name(child, "sets", 4);
        TSNode reps_node = ts_node_child_by_field_name(child, "reps", 4);
        TSNode weight_node = ts_node_child_by_field_name(child, "weight", 6);
        int nsets = ts_node_is_null(sets_node) ? DEFAULT_SETS : node_int(sets_node, src);
        int reps = ts_node_is_null(reps_node) ? DEFAULT_REPS : node_int(reps_node, src);
        int weight = ts_node_is_null(weight_node) ? 0 : node_int(weight_node, src);
        tr.kind = TR_SETS;
        if (nsets > WORKOUT_MAX_SETS) nsets = WORKOUT_MAX_SETS;
        for (int i = 0; i < nsets; i++) {
            tr.data.sets.items[i].reps = reps;
            tr.data.sets.items[i].weight = weight;
            tr.data.sets.items[i].note[0] = '\0';
        }
        tr.data.sets.count = nsets;
    } else if (is_type(child, "partial_multiplier")) {
        TSNode sets_node = ts_node_child_by_field_name(child, "sets", 4);
        TSNode reps_node = ts_node_child_by_field_name(child, "reps", 4);
        int first = ts_node_is_null(sets_node) ? 0 : node_int(sets_node, src);
        int second = ts_node_is_null(reps_node) ? 0 : node_int(reps_node, src);
        if (second > REP_THRESHOLD) {
            /* treat as 1 set of reps@weight */
            tr.kind = TR_SETS;
            tr.data.sets.items[0].reps = first;
            tr.data.sets.items[0].weight = second;
            tr.data.sets.items[0].note[0] = '\0';
            tr.data.sets.count = 1;
        } else {
            /* pending reps waiting for weight */
            tr.kind = TR_PENDING_REPS;
            int count = first;
            if (count > WORKOUT_MAX_SETS) count = WORKOUT_MAX_SETS;
            for (int i = 0; i < count; i++) {
                tr.data.pending_reps.items[i] = second;
            }
            tr.data.pending_reps.count = count;
        }
    } else if (is_type(child, "number")) {
        /* number wraps number_literal */
        TSNode nl = ts_node_named_child(child, 0);
        int n;
        if (!ts_node_is_null(nl)) {
            n = node_int(nl, src);
        } else {
            n = node_int(child, src);
        }
        if (n > REP_THRESHOLD) {
            tr.kind = TR_WEIGHT;
            tr.data.weight = n;
        } else {
            tr.kind = TR_REPS;
            tr.data.reps = n;
        }
    } else {
        /* Fallback: try parsing as a number from the node text */
        char buf[64];
        node_text(child, src, buf, sizeof(buf));
        int n = atoi(buf);
        if (n > REP_THRESHOLD) {
            tr.kind = TR_WEIGHT;
            tr.data.weight = n;
        } else {
            tr.kind = TR_REPS;
            tr.data.reps = n;
        }
    }

    return tr;
}

/* ---- Walk a token node (numeric_token or note) ---- */
static TokenResult walk_token(TSNode node, const char *src) {
    TokenResult tr;
    memset(&tr, 0, sizeof(tr));

    uint32_t ncount = ts_node_named_child_count(node);
    if (ncount == 0) {
        tr.kind = TR_NOTE;
        node_text(node, src, tr.data.note, WORKOUT_NOTE_MAX);
        return tr;
    }

    TSNode child = ts_node_named_child(node, 0);

    if (is_type(child, "numeric_token") ||
        is_type(child, "by_expr") ||
        is_type(child, "multiplier") ||
        is_type(child, "number") ||
        is_type(child, "two_part_by") ||
        is_type(child, "three_part_by") ||
        is_type(child, "full_multiplier") ||
        is_type(child, "partial_multiplier")) {
        return walk_numeric_token(child, src);
    }

    if (is_type(child, "note")) {
        /* note contains word+ nodes; extract the full text */
        tr.kind = TR_NOTE;
        node_text(child, src, tr.data.note, WORKOUT_NOTE_MAX);
        return tr;
    }

    /* Fallback */
    tr.kind = TR_NOTE;
    node_text(node, src, tr.data.note, WORKOUT_NOTE_MAX);
    return tr;
}

/* ---- Build sets from token results ---- */
static int build_sets(const TokenResult *results, int result_count,
                      int initial_weight,
                      WorkoutSet *out_sets, int max_sets) {
    int set_count = 0;
    int pending_reps[WORKOUT_MAX_SETS];
    int pending_count = 0;
    int last_weight = initial_weight;

    for (int i = 0; i < result_count; i++) {
        const TokenResult *tr = &results[i];

        switch (tr->kind) {
        case TR_SETS:
            for (int j = 0; j < tr->data.sets.count && set_count < max_sets; j++) {
                out_sets[set_count] = tr->data.sets.items[j];
                set_count++;
            }
            if (tr->data.sets.count > 0) {
                last_weight = tr->data.sets.items[tr->data.sets.count - 1].weight;
            }
            break;

        case TR_PENDING_REPS:
            for (int j = 0; j < tr->data.pending_reps.count && pending_count < WORKOUT_MAX_SETS; j++) {
                pending_reps[pending_count++] = tr->data.pending_reps.items[j];
            }
            break;

        case TR_WEIGHT:
            if (pending_count > 0) {
                for (int j = 0; j < pending_count && set_count < max_sets; j++) {
                    out_sets[set_count].reps = pending_reps[j];
                    out_sets[set_count].weight = tr->data.weight;
                    out_sets[set_count].note[0] = '\0';
                    set_count++;
                }
                pending_count = 0;
            } else {
                if (set_count < max_sets) {
                    out_sets[set_count].reps = DEFAULT_REPS;
                    out_sets[set_count].weight = tr->data.weight;
                    out_sets[set_count].note[0] = '\0';
                    set_count++;
                }
            }
            last_weight = tr->data.weight;
            break;

        case TR_REPS:
            if (pending_count < WORKOUT_MAX_SETS) {
                pending_reps[pending_count++] = tr->data.reps;
            }
            break;

        case TR_NOTE:
            /* Resolve pending reps at lastWeight first */
            if (pending_count > 0) {
                for (int j = 0; j < pending_count && set_count < max_sets; j++) {
                    out_sets[set_count].reps = pending_reps[j];
                    out_sets[set_count].weight = last_weight;
                    out_sets[set_count].note[0] = '\0';
                    set_count++;
                }
                pending_count = 0;
            }
            /* Attach note to last set */
            if (set_count > 0) {
                join_notes(out_sets[set_count - 1].note, tr->data.note, WORKOUT_NOTE_MAX);
            }
            break;
        }
    }

    /* Resolve remaining pending reps at last_weight */
    for (int j = 0; j < pending_count && set_count < max_sets; j++) {
        out_sets[set_count].reps = pending_reps[j];
        out_sets[set_count].weight = last_weight;
        out_sets[set_count].note[0] = '\0';
        set_count++;
    }

    return set_count;
}

/* ---- Collect token results from children of a node ---- */
static int collect_token_results(TSNode parent, const char *src,
                                 TokenResult *results, int max_results,
                                 int skip_words) {
    int count = 0;
    uint32_t nchildren = ts_node_named_child_count(parent);

    for (uint32_t i = 0; i < nchildren && count < max_results; i++) {
        TSNode child = ts_node_named_child(parent, i);
        const char *type = ts_node_type(child);

        if (skip_words && strcmp(type, "word") == 0) continue;
        if (strcmp(type, "superset_name") == 0) continue;
        if (strcmp(type, "star") == 0) continue;

        if (strcmp(type, "numeric_token") == 0) {
            results[count++] = walk_numeric_token(child, src);
        } else if (strcmp(type, "token") == 0) {
            results[count++] = walk_token(child, src);
        } else if (strcmp(type, "note") == 0) {
            TokenResult tr;
            memset(&tr, 0, sizeof(tr));
            tr.kind = TR_NOTE;
            node_text(child, src, tr.data.note, WORKOUT_NOTE_MAX);
            results[count++] = tr;
        } else if (strcmp(type, "by_expr") == 0 ||
                   strcmp(type, "multiplier") == 0 ||
                   strcmp(type, "number") == 0 ||
                   strcmp(type, "two_part_by") == 0 ||
                   strcmp(type, "three_part_by") == 0 ||
                   strcmp(type, "full_multiplier") == 0 ||
                   strcmp(type, "partial_multiplier") == 0) {
            results[count++] = walk_numeric_token(child, src);
        }
    }

    return count;
}

/* ---- Check if a TokenResult is numeric (not a note) ---- */
static int is_numeric_result(const TokenResult *tr) {
    return tr->kind != TR_NOTE;
}

/* ---- Extract exercise name from words before the first numeric_token ---- */
static void extract_exercise_name(TSNode exercise_line, const char *src,
                                  char *name, int namesize) {
    /* Find the byte offset of the first numeric_token child */
    uint32_t first_numeric_start = UINT32_MAX;
    uint32_t nchildren = ts_node_named_child_count(exercise_line);
    for (uint32_t i = 0; i < nchildren; i++) {
        TSNode child = ts_node_named_child(exercise_line, i);
        if (is_type(child, "numeric_token")) {
            first_numeric_start = ts_node_start_byte(child);
            break;
        }
    }

    /* Collect word text before the numeric_token */
    char raw[WORKOUT_NAME_MAX];
    raw[0] = '\0';
    int raw_len = 0;

    for (uint32_t i = 0; i < nchildren; i++) {
        TSNode child = ts_node_named_child(exercise_line, i);
        if (!is_type(child, "word")) continue;
        if (ts_node_start_byte(child) >= first_numeric_start) break;

        if (raw_len > 0 && raw_len < WORKOUT_NAME_MAX - 1) {
            raw[raw_len++] = ' ';
        }
        char word_buf[256];
        node_text(child, src, word_buf, sizeof(word_buf));
        int wlen = (int)strlen(word_buf);
        if (raw_len + wlen < WORKOUT_NAME_MAX) {
            memcpy(raw + raw_len, word_buf, wlen);
            raw_len += wlen;
        }
    }
    raw[raw_len] = '\0';

    title_case(raw, name, namesize);
}

/* ---- Extract superset name (words + "*") ---- */
static void extract_superset_name(TSNode superset_name_node, const char *src,
                                  char *name, int namesize) {
    char raw[WORKOUT_NAME_MAX];
    raw[0] = '\0';
    int raw_len = 0;

    uint32_t nchildren = ts_node_named_child_count(superset_name_node);
    for (uint32_t i = 0; i < nchildren; i++) {
        TSNode child = ts_node_named_child(superset_name_node, i);
        if (is_type(child, "word")) {
            if (raw_len > 0 && raw_len < WORKOUT_NAME_MAX - 1) {
                raw[raw_len++] = ' ';
            }
            char word_buf[256];
            node_text(child, src, word_buf, sizeof(word_buf));
            int wlen = (int)strlen(word_buf);
            if (raw_len + wlen < WORKOUT_NAME_MAX) {
                memcpy(raw + raw_len, word_buf, wlen);
                raw_len += wlen;
            }
        }
    }
    raw[raw_len] = '\0';

    char tc[WORKOUT_NAME_MAX];
    title_case(raw, tc, sizeof(tc));

    /* Append "*" */
    int tclen = (int)strlen(tc);
    if (tclen < namesize - 2) {
        memcpy(name, tc, tclen);
        name[tclen] = '*';
        name[tclen + 1] = '\0';
    } else {
        strncpy(name, tc, namesize - 1);
        name[namesize - 1] = '\0';
    }
}

/* ---- Split exercise line tokens into multiple exercises ---- */
/* Returns number of exercises found */
static int split_into_exercises(const char *first_name,
                                const TokenResult *tokens, int token_count,
                                Exercise *out_exercises, int max_exercises) {
    int ex_count = 0;
    char current_name[WORKOUT_NAME_MAX];
    strncpy(current_name, first_name, WORKOUT_NAME_MAX - 1);
    current_name[WORKOUT_NAME_MAX - 1] = '\0';

    TokenResult *current_tokens = (TokenResult *)calloc(MAX_TOKEN_RESULTS, sizeof(TokenResult));
    if (!current_tokens) return 0;
    int current_token_count = 0;

    for (int i = 0; i < token_count; i++) {
        const TokenResult *tok = &tokens[i];

        /* Check if this note starts a new exercise (note followed by numeric token) */
        if (tok->kind == TR_NOTE &&
            i + 1 < token_count &&
            is_numeric_result(&tokens[i + 1])) {
            /* Flush current exercise */
            if (current_name[0] != '\0' && ex_count < max_exercises) {
                Exercise *ex = &out_exercises[ex_count];
                strncpy(ex->name, current_name, WORKOUT_NAME_MAX - 1);
                ex->name[WORKOUT_NAME_MAX - 1] = '\0';
                ex->set_count = build_sets(current_tokens, current_token_count, 0,
                                           ex->sets, WORKOUT_MAX_SETS);
                ex_count++;
            }
            /* Start new exercise with this note as name */
            title_case(tok->data.note, current_name, WORKOUT_NAME_MAX);
            current_token_count = 0;
        } else {
            if (current_token_count < MAX_TOKEN_RESULTS) {
                current_tokens[current_token_count++] = *tok;
            }
        }
    }

    /* Flush last exercise */
    if (current_name[0] != '\0' && ex_count < max_exercises) {
        Exercise *ex = &out_exercises[ex_count];
        strncpy(ex->name, current_name, WORKOUT_NAME_MAX - 1);
        ex->name[WORKOUT_NAME_MAX - 1] = '\0';
        ex->set_count = build_sets(current_tokens, current_token_count, 0,
                                   ex->sets, WORKOUT_MAX_SETS);
        ex_count++;
    }

    free(current_tokens);
    return ex_count;
}

/* ---- Extract prose text from a prose_line node ---- */
static void extract_prose_text(TSNode prose_line, const char *src,
                               char *out, int outsize) {
    out[0] = '\0';
    int len = 0;
    uint32_t nchildren = ts_node_named_child_count(prose_line);
    for (uint32_t i = 0; i < nchildren; i++) {
        TSNode child = ts_node_named_child(prose_line, i);
        if (is_type(child, "word")) {
            if (len > 0 && len < outsize - 1) {
                out[len++] = ' ';
            }
            char word_buf[256];
            node_text(child, src, word_buf, sizeof(word_buf));
            int wlen = (int)strlen(word_buf);
            if (len + wlen < outsize) {
                memcpy(out + len, word_buf, wlen);
                len += wlen;
            }
        }
    }
    out[len] = '\0';
}

/* ---- Build sets from superset tokens (simpler: no pending-reps merging) ---- */
static int build_superset_sets(const TokenResult *results, int result_count,
                               WorkoutSet *out_sets, int max_sets) {
    int set_count = 0;
    for (int i = 0; i < result_count && set_count < max_sets; i++) {
        const TokenResult *tr = &results[i];
        switch (tr->kind) {
        case TR_SETS:
            for (int j = 0; j < tr->data.sets.count && set_count < max_sets; j++) {
                out_sets[set_count++] = tr->data.sets.items[j];
            }
            break;
        case TR_PENDING_REPS:
            for (int j = 0; j < tr->data.pending_reps.count && set_count < max_sets; j++) {
                out_sets[set_count].reps = tr->data.pending_reps.items[j];
                out_sets[set_count].weight = 0;
                out_sets[set_count].note[0] = '\0';
                set_count++;
            }
            break;
        case TR_WEIGHT:
            if (set_count < max_sets) {
                out_sets[set_count].reps = DEFAULT_REPS;
                out_sets[set_count].weight = tr->data.weight;
                out_sets[set_count].note[0] = '\0';
                set_count++;
            }
            break;
        case TR_REPS:
            if (set_count < max_sets) {
                out_sets[set_count].reps = tr->data.reps;
                out_sets[set_count].weight = 0;
                out_sets[set_count].note[0] = '\0';
                set_count++;
            }
            break;
        case TR_NOTE:
            break;
        }
    }
    return set_count;
}

/* ---- Check if input contains a digit (ignoring lines with '|') ---- */
static int contains_digit(const char *input) {
    const char *line_start = input;
    while (*line_start) {
        /* Find end of line */
        const char *line_end = strchr(line_start, '\n');
        if (!line_end) line_end = line_start + strlen(line_start);

        /* Check if line contains '|' */
        int has_pipe = 0;
        for (const char *p = line_start; p < line_end; p++) {
            if (*p == '|') { has_pipe = 1; break; }
        }

        if (!has_pipe) {
            for (const char *p = line_start; p < line_end; p++) {
                if (isdigit((unsigned char)*p)) return 1;
            }
        }

        if (*line_end == '\0') break;
        line_start = line_end + 1;
    }
    return 0;
}

/* ---- isWorkout heuristic ---- */
static int is_workout(const WorkoutResult *result) {
    if (result->exercise_count == 0) return 0;
    int total_sets = 0;
    int has_weight = 0;
    for (int i = 0; i < result->exercise_count; i++) {
        total_sets += result->exercises[i].set_count;
        for (int j = 0; j < result->exercises[i].set_count; j++) {
            if (result->exercises[i].sets[j].weight > 0) {
                has_weight = 1;
            }
        }
    }
    return total_sets >= 2 && has_weight;
}

/* ---- Main parse function ---- */
int workout_parse(const char *input, WorkoutResult *result) {
    memset(result, 0, sizeof(*result));

    if (!input || input[0] == '\0') return 0;
    if (!contains_digit(input)) return 0;

    /* Create tree-sitter parser */
    TSParser *parser = ts_parser_new();
    ts_parser_set_language(parser, tree_sitter_workout());

    uint32_t len = (uint32_t)strlen(input);
    TSTree *tree = ts_parser_parse_string(parser, NULL, input, len);
    if (!tree) {
        ts_parser_delete(parser);
        return 0;
    }

    TSNode root = ts_tree_root_node(tree);

    /* The root should be a "paragraph" node */
    if (!is_type(root, "paragraph")) {
        ts_tree_delete(tree);
        ts_parser_delete(parser);
        return 0;
    }

    /* State machine for walking the paragraph */
    Exercise last_exercise;
    int has_last_exercise = 0;
    char pending_prose[WORKOUT_NAME_MAX];
    pending_prose[0] = '\0';

    /* flush helper - implemented inline */
    #define FLUSH() do { \
        if (has_last_exercise && last_exercise.set_count > 0) { \
            if (result->exercise_count < WORKOUT_MAX_EXERCISES) { \
                result->exercises[result->exercise_count++] = last_exercise; \
            } \
        } else if (pending_prose[0] != '\0') { \
            if (result->prose_count < WORKOUT_MAX_PROSE) { \
                strncpy(result->prose_lines[result->prose_count], pending_prose, WORKOUT_NAME_MAX - 1); \
                result->prose_lines[result->prose_count][WORKOUT_NAME_MAX - 1] = '\0'; \
                result->prose_count++; \
            } \
        } \
        has_last_exercise = 0; \
        memset(&last_exercise, 0, sizeof(last_exercise)); \
        pending_prose[0] = '\0'; \
    } while(0)

    /* Iterate over named children of the paragraph (which are line nodes) */
    uint32_t nlines = ts_node_named_child_count(root);
    for (uint32_t li = 0; li < nlines; li++) {
        TSNode line = ts_node_named_child(root, li);
        const char *line_type = ts_node_type(line);

        if (strcmp(line_type, "exercise_line") == 0) {
            FLUSH();

            /* Extract exercise name */
            char name[WORKOUT_NAME_MAX];
            extract_exercise_name(line, input, name, sizeof(name));

            /* Collect token results (heap-allocated to avoid stack overflow) */
            TokenResult *token_results = (TokenResult *)calloc(MAX_TOKEN_RESULTS, sizeof(TokenResult));
            if (!token_results) { ts_tree_delete(tree); ts_parser_delete(parser); return 0; }
            int tr_count = collect_token_results(line, input, token_results, MAX_TOKEN_RESULTS, 1);

            /* Split into multiple exercises if needed */
            Exercise *temp_exercises = (Exercise *)calloc(WORKOUT_MAX_EXERCISES, sizeof(Exercise));
            if (!temp_exercises) { free(token_results); ts_tree_delete(tree); ts_parser_delete(parser); return 0; }
            int ex_count = split_into_exercises(name, token_results, tr_count,
                                                temp_exercises, WORKOUT_MAX_EXERCISES);

            if (ex_count == 1) {
                if (temp_exercises[0].set_count > 0) {
                    last_exercise = temp_exercises[0];
                    has_last_exercise = 1;
                } else {
                    /* Name-only line, treat as pending prose */
                    strncpy(pending_prose, temp_exercises[0].name, WORKOUT_NAME_MAX - 1);
                    pending_prose[WORKOUT_NAME_MAX - 1] = '\0';
                }
            } else if (ex_count > 1) {
                /* Multiple exercises on one line (multi-exercise detection) */
                for (int e = 0; e < ex_count; e++) {
                    if (temp_exercises[e].set_count > 0) {
                        if (e == ex_count - 1) {
                            /* Keep last one as current exercise */
                            last_exercise = temp_exercises[e];
                            has_last_exercise = 1;
                        } else {
                            if (result->exercise_count < WORKOUT_MAX_EXERCISES) {
                                result->exercises[result->exercise_count++] = temp_exercises[e];
                            }
                        }
                    }
                }
            }

            free(token_results);
            free(temp_exercises);

        } else if (strcmp(line_type, "continuation_line") == 0) {
            /* Collect token results from continuation line (heap-allocated) */
            TokenResult *token_results = (TokenResult *)calloc(MAX_TOKEN_RESULTS, sizeof(TokenResult));
            if (!token_results) { ts_tree_delete(tree); ts_parser_delete(parser); return 0; }
            int tr_count = collect_token_results(line, input, token_results, MAX_TOKEN_RESULTS, 0);

            if (has_last_exercise) {
                /* Extend existing exercise with weight carryover */
                int last_weight = 0;
                if (last_exercise.set_count > 0) {
                    last_weight = last_exercise.sets[last_exercise.set_count - 1].weight;
                }
                WorkoutSet new_sets[WORKOUT_MAX_SETS];
                int new_count = build_sets(token_results, tr_count, last_weight,
                                           new_sets, WORKOUT_MAX_SETS);
                for (int j = 0; j < new_count && last_exercise.set_count < WORKOUT_MAX_SETS; j++) {
                    last_exercise.sets[last_exercise.set_count++] = new_sets[j];
                }
            } else if (pending_prose[0] != '\0') {
                /* Start new exercise from pending prose */
                memset(&last_exercise, 0, sizeof(last_exercise));
                title_case(pending_prose, last_exercise.name, WORKOUT_NAME_MAX);
                last_exercise.set_count = build_sets(token_results, tr_count, 0,
                                                     last_exercise.sets, WORKOUT_MAX_SETS);
                has_last_exercise = 1;
                pending_prose[0] = '\0';
            }
            /* else: no context, ignore */

            free(token_results);

        } else if (strcmp(line_type, "superset_line") == 0) {
            FLUSH();

            /* Extract the two superset names */
            char name1[WORKOUT_NAME_MAX], name2[WORKOUT_NAME_MAX];
            name1[0] = name2[0] = '\0';
            int name_idx = 0;
            uint32_t nchildren = ts_node_named_child_count(line);
            for (uint32_t ci = 0; ci < nchildren; ci++) {
                TSNode child = ts_node_named_child(line, ci);
                if (is_type(child, "superset_name")) {
                    if (name_idx == 0) {
                        extract_superset_name(child, input, name1, sizeof(name1));
                    } else {
                        extract_superset_name(child, input, name2, sizeof(name2));
                    }
                    name_idx++;
                }
            }

            /* Collect all token results (heap-allocated) */
            TokenResult *token_results = (TokenResult *)calloc(MAX_TOKEN_RESULTS, sizeof(TokenResult));
            if (!token_results) { ts_tree_delete(tree); ts_parser_delete(parser); return 0; }
            int tr_count = collect_token_results(line, input, token_results, MAX_TOKEN_RESULTS, 1);

            /* Build all sets */
            WorkoutSet all_sets[WORKOUT_MAX_SETS];
            int total = build_superset_sets(token_results, tr_count, all_sets, WORKOUT_MAX_SETS);
            free(token_results);

            /* Alternate between two exercises */
            Exercise ex1, ex2;
            memset(&ex1, 0, sizeof(ex1));
            memset(&ex2, 0, sizeof(ex2));
            strncpy(ex1.name, name1, WORKOUT_NAME_MAX - 1);
            strncpy(ex2.name, name2, WORKOUT_NAME_MAX - 1);

            for (int s = 0; s < total; s++) {
                if (s % 2 == 0) {
                    if (ex1.set_count < WORKOUT_MAX_SETS)
                        ex1.sets[ex1.set_count++] = all_sets[s];
                } else {
                    if (ex2.set_count < WORKOUT_MAX_SETS)
                        ex2.sets[ex2.set_count++] = all_sets[s];
                }
            }

            if (ex1.set_count > 0 && result->exercise_count < WORKOUT_MAX_EXERCISES)
                result->exercises[result->exercise_count++] = ex1;
            if (ex2.set_count > 0 && result->exercise_count < WORKOUT_MAX_EXERCISES)
                result->exercises[result->exercise_count++] = ex2;

        } else if (strcmp(line_type, "prose_line") == 0) {
            FLUSH();
            char prose[WORKOUT_NAME_MAX];
            extract_prose_text(line, input, prose, sizeof(prose));
            if (prose[0] != '\0') {
                strncpy(pending_prose, prose, WORKOUT_NAME_MAX - 1);
                pending_prose[WORKOUT_NAME_MAX - 1] = '\0';
            }
        }
    }

    /* Final flush */
    FLUSH();

    #undef FLUSH

    ts_tree_delete(tree);
    ts_parser_delete(parser);

    return is_workout(result);
}
