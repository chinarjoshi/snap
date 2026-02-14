#include <string.h>
#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#include <tree_sitter/api.h>
#include "recipe.h"

extern const TSLanguage *tree_sitter_recipe(void);

static void title_case(char *dst, const char *src, int max) {
    int i = 0, word_start = 1;
    while (src[i] && i < max - 1) {
        if (src[i] == ' ') {
            dst[i] = ' ';
            word_start = 1;
        } else if (word_start) {
            dst[i] = toupper((unsigned char)src[i]);
            word_start = 0;
        } else {
            dst[i] = tolower((unsigned char)src[i]);
        }
        i++;
    }
    dst[i] = '\0';
}

static void node_text(char *dst, TSNode node, const char *src, int max) {
    uint32_t start = ts_node_start_byte(node);
    uint32_t end = ts_node_end_byte(node);
    int len = end - start;
    if (len >= max) len = max - 1;
    memcpy(dst, src + start, len);
    dst[len] = '\0';
}

static int node_int_val(TSNode node, const char *src) {
    char buf[32];
    node_text(buf, node, src, sizeof(buf));
    return atoi(buf);
}

static void extract_header_name(char *dst, TSNode node, const char *src, int max) {
    uint32_t count = ts_node_named_child_count(node);
    char raw[RECIPE_NAME_MAX] = {0};
    int pos = 0;

    for (uint32_t i = 0; i < count; i++) {
        TSNode child = ts_node_named_child(node, i);
        if (strcmp(ts_node_type(child), "word") == 0) {
            if (pos > 0 && pos < RECIPE_NAME_MAX - 1) raw[pos++] = ' ';
            uint32_t s = ts_node_start_byte(child);
            uint32_t e = ts_node_end_byte(child);
            int len = e - s;
            if (pos + len >= RECIPE_NAME_MAX) len = RECIPE_NAME_MAX - 1 - pos;
            memcpy(raw + pos, src + s, len);
            pos += len;
        }
    }
    raw[pos] = '\0';
    title_case(dst, raw, max);
}

static void extract_quantity(char *dst, TSNode qty_node, const char *src, int max) {
    TSNode min_node = ts_node_child_by_field_name(qty_node, "min", 3);
    TSNode max_node = ts_node_child_by_field_name(qty_node, "max", 3);
    TSNode val_node = ts_node_child_by_field_name(qty_node, "value", 5);

    if (!ts_node_is_null(min_node) && !ts_node_is_null(max_node)) {
        int mn = node_int_val(min_node, src);
        int mx = node_int_val(max_node, src);
        snprintf(dst, max, "%d-%d", mn, mx);
    } else if (!ts_node_is_null(val_node)) {
        node_text(dst, val_node, src, max);
    } else {
        node_text(dst, qty_node, src, max);
    }
}

static void extract_ingredient_text(char *dst, TSNode node, const char *src, int max) {
    uint32_t count = ts_node_named_child_count(node);
    int pos = 0;
    int past_quantity = 0;

    for (uint32_t i = 0; i < count; i++) {
        TSNode child = ts_node_named_child(node, i);
        const char *type = ts_node_type(child);

        if (strcmp(type, "quantity") == 0) {
            past_quantity = 1;
            continue;
        }
        if (!past_quantity) continue;

        if (strcmp(type, "word") == 0 || strcmp(type, "number") == 0) {
            if (pos > 0 && pos < max - 1) dst[pos++] = ' ';
            uint32_t s = ts_node_start_byte(child);
            uint32_t e = ts_node_end_byte(child);
            int len = e - s;
            if (pos + len >= max) len = max - 1 - pos;
            memcpy(dst + pos, src + s, len);
            pos += len;
        }
    }
    dst[pos] = '\0';
}

static int contains_colon_eol(const char *s) {
    const char *p = s;
    while (*p) {
        const char *eol = strchr(p, '\n');
        int len = eol ? (int)(eol - p) : (int)strlen(p);
        if (len > 0) {
            int end = len - 1;
            while (end >= 0 && (p[end] == ' ' || p[end] == '\t')) end--;
            if (end >= 0 && p[end] == ':') return 1;
        }
        if (!eol) break;
        p = eol + 1;
    }
    return 0;
}

static int contains_digit(const char *s) {
    for (int i = 0; s[i]; i++) {
        if (isdigit((unsigned char)s[i])) return 1;
    }
    return 0;
}

static int is_recipe(const RecipeResult *r) {
    if (r->title[0] == '\0') return 0;
    int total = 0;
    for (int i = 0; i < r->section_count; i++) {
        total += r->sections[i].ingredient_count;
    }
    return total >= 2;
}

static RecipeSection *current_or_default(RecipeResult *result) {
    if (result->section_count == 0) {
        result->section_count = 1;
        memset(&result->sections[0], 0, sizeof(RecipeSection));
    }
    return &result->sections[result->section_count - 1];
}

int recipe_parse(const char *input, RecipeResult *result) {
    memset(result, 0, sizeof(RecipeResult));

    if (!input || input[0] == '\0') return 0;
    if (!contains_digit(input) || !contains_colon_eol(input)) return 0;

    TSParser *parser = ts_parser_new();
    ts_parser_set_language(parser, tree_sitter_recipe());

    uint32_t len = (uint32_t)strlen(input);
    TSTree *tree = ts_parser_parse_string(parser, NULL, input, len);
    TSNode root = ts_tree_root_node(tree);

    uint32_t child_count = ts_node_named_child_count(root);
    int title_set = 0;

    for (uint32_t i = 0; i < child_count; i++) {
        TSNode child = ts_node_named_child(root, i);
        const char *type = ts_node_type(child);

        if (strcmp(type, "header_line") == 0) {
            if (!title_set) {
                extract_header_name(result->title, child, input, RECIPE_NAME_MAX);
                title_set = 1;
            } else {
                if (result->section_count < RECIPE_MAX_SECTIONS) {
                    RecipeSection *sec = &result->sections[result->section_count];
                    memset(sec, 0, sizeof(RecipeSection));
                    extract_header_name(sec->name, child, input, RECIPE_NAME_MAX);
                    result->section_count++;
                }
            }
        } else if (strcmp(type, "ingredient_line") == 0) {
            RecipeSection *sec = current_or_default(result);
            if (sec->ingredient_count < RECIPE_MAX_INGREDIENTS) {
                Ingredient *ing = &sec->ingredients[sec->ingredient_count];
                memset(ing, 0, sizeof(Ingredient));

                /* Find quantity child */
                for (uint32_t j = 0; j < ts_node_named_child_count(child); j++) {
                    TSNode nc = ts_node_named_child(child, j);
                    if (strcmp(ts_node_type(nc), "quantity") == 0) {
                        extract_quantity(ing->quantity, nc, input, RECIPE_QTY_MAX);
                        break;
                    }
                }
                extract_ingredient_text(ing->text, child, input, RECIPE_NAME_MAX);
                sec->ingredient_count++;
            }
        } else if (strcmp(type, "text_line") == 0) {
            RecipeSection *sec = current_or_default(result);
            if (sec->instruction_count < RECIPE_MAX_INSTRUCTIONS) {
                node_text(
                    sec->instructions[sec->instruction_count],
                    child, input, RECIPE_TEXT_MAX
                );
                sec->instruction_count++;
            }
        }
    }

    ts_tree_delete(tree);
    ts_parser_delete(parser);

    return is_recipe(result) ? 1 : 0;
}
