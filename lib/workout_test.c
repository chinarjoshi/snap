#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "workout.h"

static int tests_run = 0;

#define RUN_TEST(name) do { \
    printf("  %-55s", #name); \
    name(); \
    tests_run++; \
    printf("PASS\n"); \
} while(0)

/* ---- Helper: find exercise by name ---- */
static Exercise *find_exercise(WorkoutResult *r, const char *name) {
    for (int i = 0; i < r->exercise_count; i++) {
        if (strcmp(r->exercises[i].name, name) == 0) return &r->exercises[i];
    }
    return NULL;
}

/* ==================== Basic Parsing ==================== */

/* 1. testParsePureWorkout */
static void test_parse_pure_workout(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 8 8 135\nbench 3x8 95", &r) == 1);
    assert(r.exercise_count == 2);

    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(squat->set_count == 3);
    assert(squat->sets[0].reps == 8);
    assert(squat->sets[0].weight == 135);
    assert(squat->sets[1].reps == 8);
    assert(squat->sets[1].weight == 135);
    assert(squat->sets[2].reps == 8);
    assert(squat->sets[2].weight == 135);

    Exercise *bench = find_exercise(&r, "Bench");
    assert(bench != NULL);
    assert(bench->set_count == 3);
    assert(bench->sets[0].reps == 8);
    assert(bench->sets[0].weight == 95);
    assert(bench->sets[1].reps == 8);
    assert(bench->sets[1].weight == 95);
    assert(bench->sets[2].reps == 8);
    assert(bench->sets[2].weight == 95);
}

/* 2. testParseProseBeforeWorkout */
static void test_parse_prose_before_workout(void) {
    WorkoutResult r;
    assert(workout_parse("ate a banana beforehand\nsquat 8 8 8 135", &r) == 1);

    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(r.prose_count >= 1);
    int found = 0;
    for (int i = 0; i < r.prose_count; i++) {
        if (strcmp(r.prose_lines[i], "ate a banana beforehand") == 0) found = 1;
    }
    assert(found);
}

/* 3. testParseMultipleProseLines */
static void test_parse_multiple_prose_lines(void) {
    WorkoutResult r;
    assert(workout_parse("morning workout\nfelt strong\nsquat 8 8 8 135", &r) == 1);

    int found_morning = 0, found_strong = 0;
    for (int i = 0; i < r.prose_count; i++) {
        if (strcmp(r.prose_lines[i], "morning workout") == 0) found_morning = 1;
        if (strcmp(r.prose_lines[i], "felt strong") == 0) found_strong = 1;
    }
    assert(found_morning);
    assert(found_strong);

    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
}

/* 4. testParseInlineNotes */
static void test_parse_inline_notes(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 8 8 135 light headed\nbench 3x8 95", &r) == 1);

    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);

    int found_note = 0;
    for (int i = 0; i < squat->set_count; i++) {
        if (strstr(squat->sets[i].note, "light headed") != NULL) {
            found_note = 1;
            break;
        }
    }
    assert(found_note);
}

/* 5. testParseMultiplierSyntax (4 sub-cases) */
static void test_parse_multiplier_syntax(void) {
    /* 5a: "squat 3x8 135" */
    {
        WorkoutResult r;
        assert(workout_parse("squat 3x8 135", &r) == 1);
        Exercise *squat = find_exercise(&r, "Squat");
        assert(squat != NULL);
        assert(squat->sets[0].reps == 8);
        assert(squat->sets[0].weight == 135);
    }

    /* 5b: "squat 3x8x135" */
    {
        WorkoutResult r;
        assert(workout_parse("squat 3x8x135", &r) == 1);
        Exercise *squat = find_exercise(&r, "Squat");
        assert(squat != NULL);
        assert(squat->sets[0].reps == 8);
        assert(squat->sets[0].weight == 135);
    }

    /* 5c: "squat 8 by 135" */
    {
        WorkoutResult r;
        assert(workout_parse("squat 8 by 135", &r) == 1);
        Exercise *squat = find_exercise(&r, "Squat");
        assert(squat != NULL);
        assert(squat->sets[0].reps == 8);
        assert(squat->sets[0].weight == 135);
    }

    /* 5d: "squat 3 by 8 by 135" */
    {
        WorkoutResult r;
        assert(workout_parse("squat 3 by 8 by 135", &r) == 1);
        Exercise *squat = find_exercise(&r, "Squat");
        assert(squat != NULL);
        assert(squat->sets[0].reps == 8);
        assert(squat->sets[0].weight == 135);
    }
}

/* 6. testParseTitleCase */
static void test_parse_title_case(void) {
    WorkoutResult r;
    assert(workout_parse("BARBELL SQUAT 8 8 8 135", &r) == 1);
    Exercise *ex = find_exercise(&r, "Barbell Squat");
    assert(ex != NULL);
}

/* 7. testParseTrailingReps */
static void test_parse_trailing_reps(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 8 8 135 8 8", &r) == 1);
    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(squat->set_count == 5);
    for (int i = 0; i < 5; i++) {
        assert(squat->sets[i].reps == 8);
        assert(squat->sets[i].weight == 135);
    }
}

/* 8. testParseDefaultSetsForBySyntax */
static void test_parse_default_sets_for_by_syntax(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 by 135", &r) == 1);
    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(squat->set_count == 3);
}

/* 9. testParseDefaultRepsForWeight */
static void test_parse_default_reps_for_weight(void) {
    WorkoutResult r;
    assert(workout_parse("squat 135 135", &r) == 1);
    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(squat->sets[0].reps == 8);
    assert(squat->sets[0].weight == 135);
}

/* 10. testParseMultipleExercises */
static void test_parse_multiple_exercises(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 8 8 135\nbench 3x8 95\npress 8 by 35", &r) == 1);
    assert(r.exercise_count == 3);
    assert(find_exercise(&r, "Squat") != NULL);
    assert(find_exercise(&r, "Bench") != NULL);
    assert(find_exercise(&r, "Press") != NULL);
}

/* ==================== Returns nil ==================== */

/* 11. testParseReturnsNilForEmpty */
static void test_parse_returns_nil_for_empty(void) {
    WorkoutResult r;
    assert(workout_parse("", &r) == 0);
}

/* 12. testParseReturnsNilForPureProse */
static void test_parse_returns_nil_for_pure_prose(void) {
    WorkoutResult r;
    assert(workout_parse("Just some notes today.\nNothing special.", &r) == 0);
}

/* 13. testParseReturnsNilForNoDigits */
static void test_parse_returns_nil_for_no_digits(void) {
    WorkoutResult r;
    assert(workout_parse("Just words here\nno numbers at all", &r) == 0);
}

/* ==================== Bodyweight ==================== */

/* 14. testParseBodyweight */
static void test_parse_bodyweight(void) {
    WorkoutResult r;
    assert(workout_parse("dips 8 8 8\nbench 8x135 8x130", &r) == 1);

    Exercise *dips = find_exercise(&r, "Dips");
    assert(dips != NULL);
    for (int i = 0; i < dips->set_count; i++) {
        assert(dips->sets[i].weight == 0);
    }
}

/* ==================== Supersets ==================== */

/* 15. testParseSupersetBasic */
static void test_parse_superset_basic(void) {
    WorkoutResult r;
    assert(workout_parse("dips* press* 8 135 8 135 8 135", &r) == 1);
    assert(r.exercise_count == 2);
    assert(find_exercise(&r, "Dips*") != NULL);
    assert(find_exercise(&r, "Press*") != NULL);
}

/* 16. testParseSupersetMixedBodyweightWeighted */
static void test_parse_superset_mixed_bw_weighted(void) {
    WorkoutResult r;
    assert(workout_parse("dips* shoulder press* 8 8x135 8 8x125 8 8x100", &r) == 1);

    Exercise *dips = find_exercise(&r, "Dips*");
    assert(dips != NULL);
    assert(dips->set_count == 3);
    for (int i = 0; i < 3; i++) {
        assert(dips->sets[i].weight == 0);
    }

    Exercise *press = find_exercise(&r, "Shoulder Press*");
    assert(press != NULL);
    assert(press->set_count == 3);
    assert(press->sets[0].weight == 135);
    assert(press->sets[1].weight == 125);
    assert(press->sets[2].weight == 100);
}

/* 17. testParseSupersetOddSets */
static void test_parse_superset_odd_sets(void) {
    WorkoutResult r;
    assert(workout_parse("dips* press* 8 135 8 135 8", &r) == 1);

    Exercise *dips = find_exercise(&r, "Dips*");
    assert(dips != NULL);
    assert(dips->set_count == 3);

    Exercise *press = find_exercise(&r, "Press*");
    assert(press != NULL);
    assert(press->set_count == 2);
}

/* ==================== Multi-Line ==================== */

/* 18. testParseMultiLineBasic */
static void test_parse_multi_line_basic(void) {
    WorkoutResult r;
    assert(workout_parse("Bench\n8x135\n8\n8\n8", &r) == 1);

    Exercise *bench = find_exercise(&r, "Bench");
    assert(bench != NULL);
    assert(bench->set_count == 4);
    for (int i = 0; i < 4; i++) {
        assert(bench->sets[i].reps == 8);
        assert(bench->sets[i].weight == 135);
    }
}

/* 19. testParseMultiLineInlinePlusContinuation */
static void test_parse_multi_line_inline_plus_continuation(void) {
    WorkoutResult r;
    assert(workout_parse("shoulder press 7x100\n6\n5", &r) == 1);

    Exercise *press = find_exercise(&r, "Shoulder Press");
    assert(press != NULL);
    assert(press->set_count == 3);
    assert(press->sets[0].reps == 7);
    assert(press->sets[0].weight == 100);
    assert(press->sets[1].reps == 6);
    assert(press->sets[1].weight == 100);
    assert(press->sets[2].reps == 5);
    assert(press->sets[2].weight == 100);
}

/* 20. testParseMultiLineBodyweight */
static void test_parse_multi_line_bodyweight(void) {
    WorkoutResult r;
    assert(workout_parse("Dips\n8\n8\n8\nbench 8x135 8x130", &r) == 1);

    Exercise *dips = find_exercise(&r, "Dips");
    assert(dips != NULL);
    for (int i = 0; i < dips->set_count; i++) {
        assert(dips->sets[i].weight == 0);
    }
}

/* 21. testParseMultiLineMixedExercises */
static void test_parse_multi_line_mixed_exercises(void) {
    WorkoutResult r;
    assert(workout_parse("Bench\n8x135\nshoulder press 7x100\n6", &r) == 1);

    Exercise *bench = find_exercise(&r, "Bench");
    assert(bench != NULL);
    assert(bench->set_count == 1);
    assert(bench->sets[0].reps == 8);
    assert(bench->sets[0].weight == 135);

    Exercise *press = find_exercise(&r, "Shoulder Press");
    assert(press != NULL);
    assert(press->set_count == 2);
    assert(press->sets[0].reps == 7);
    assert(press->sets[0].weight == 100);
    assert(press->sets[1].reps == 6);
    assert(press->sets[1].weight == 100);
}

/* ==================== Single-Line Multi-Exercise ==================== */

/* 22. testParseSingleLineMultipleExercises */
static void test_parse_single_line_multiple_exercises(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8x135 bench 8x95 press 8x45", &r) == 1);
    assert(r.exercise_count == 3);

    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(squat->sets[0].weight == 135);

    Exercise *bench = find_exercise(&r, "Bench");
    assert(bench != NULL);
    assert(bench->sets[0].weight == 95);

    Exercise *press = find_exercise(&r, "Press");
    assert(press != NULL);
    assert(press->sets[0].weight == 45);
}

/* 23. testParseSingleLineMultiWordWithModifier */
static void test_parse_single_line_multi_word_with_modifier(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8x135 heavy bench 8x95", &r) == 1);

    assert(find_exercise(&r, "Squat") != NULL);
    assert(find_exercise(&r, "Heavy Bench") != NULL);
}

/* 24. testParseSingleLineMultiWordExercise */
static void test_parse_single_line_multi_word_exercise(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8x135 bench press 8x95", &r) == 1);

    assert(find_exercise(&r, "Squat") != NULL);
    assert(find_exercise(&r, "Bench Press") != NULL);
}

/* ==================== Heuristics ==================== */

/* 25. testHeuristicRejectsProseWithNumber (8 sub-cases) */
static void test_heuristic_rejects_prose_with_number(void) {
    WorkoutResult r;
    assert(workout_parse("I ran 5 miles today", &r) == 0);
    assert(workout_parse("ate 3 eggs", &r) == 0);
    assert(workout_parse("rested 10 minutes between sets", &r) == 0);
    assert(workout_parse("dips 8", &r) == 0);
    assert(workout_parse("squat 135", &r) == 0);
    assert(workout_parse("bench 8x135", &r) == 0);
    assert(workout_parse("dips 8 8 8", &r) == 0);
    assert(workout_parse("rested 30 minutes", &r) == 0);
}

/* 26. testHeuristicAcceptsWorkout (4 sub-cases) */
static void test_heuristic_accepts_workout(void) {
    WorkoutResult r;
    assert(workout_parse("squat 135 135 135", &r) == 1);
    assert(workout_parse("bench 8x135 8x130 8x125", &r) == 1);
    assert(workout_parse("squat 135 bench 95", &r) == 1);
    assert(workout_parse("squat 3x8x135", &r) == 1);
}

/* ==================== Notes ==================== */

/* 27. testParseMultiLineNotesPreserved */
static void test_parse_multi_line_notes_preserved(void) {
    WorkoutResult r;
    assert(workout_parse("Squat\n8x225 ouch\n7 dang\n5 fuck", &r) == 1);

    Exercise *squat = find_exercise(&r, "Squat");
    assert(squat != NULL);
    assert(squat->set_count == 3);

    assert(squat->sets[0].reps == 8);
    assert(squat->sets[0].weight == 225);
    assert(strcmp(squat->sets[0].note, "ouch") == 0);

    assert(squat->sets[1].reps == 7);
    assert(squat->sets[1].weight == 225);
    assert(strcmp(squat->sets[1].note, "dang") == 0);

    assert(squat->sets[2].reps == 5);
    assert(squat->sets[2].weight == 225);
    assert(strcmp(squat->sets[2].note, "fuck") == 0);
}

int main(void) {
    printf("Running workout parser tests (27 cases)...\n");

    /* Basic Parsing */
    RUN_TEST(test_parse_pure_workout);
    RUN_TEST(test_parse_prose_before_workout);
    RUN_TEST(test_parse_multiple_prose_lines);
    RUN_TEST(test_parse_inline_notes);
    RUN_TEST(test_parse_multiplier_syntax);
    RUN_TEST(test_parse_title_case);
    RUN_TEST(test_parse_trailing_reps);
    RUN_TEST(test_parse_default_sets_for_by_syntax);
    RUN_TEST(test_parse_default_reps_for_weight);
    RUN_TEST(test_parse_multiple_exercises);

    /* Returns nil */
    RUN_TEST(test_parse_returns_nil_for_empty);
    RUN_TEST(test_parse_returns_nil_for_pure_prose);
    RUN_TEST(test_parse_returns_nil_for_no_digits);

    /* Bodyweight */
    RUN_TEST(test_parse_bodyweight);

    /* Supersets */
    RUN_TEST(test_parse_superset_basic);
    RUN_TEST(test_parse_superset_mixed_bw_weighted);
    RUN_TEST(test_parse_superset_odd_sets);

    /* Multi-Line */
    RUN_TEST(test_parse_multi_line_basic);
    RUN_TEST(test_parse_multi_line_inline_plus_continuation);
    RUN_TEST(test_parse_multi_line_bodyweight);
    RUN_TEST(test_parse_multi_line_mixed_exercises);

    /* Single-Line Multi-Exercise */
    RUN_TEST(test_parse_single_line_multiple_exercises);
    RUN_TEST(test_parse_single_line_multi_word_with_modifier);
    RUN_TEST(test_parse_single_line_multi_word_exercise);

    /* Heuristics */
    RUN_TEST(test_heuristic_rejects_prose_with_number);
    RUN_TEST(test_heuristic_accepts_workout);

    /* Notes */
    RUN_TEST(test_parse_multi_line_notes_preserved);

    printf("\nAll %d tests passed!\n", tests_run);
    return 0;
}
