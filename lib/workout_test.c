#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "workout.h"

static int tests_run = 0;

#define RUN_TEST(name) do { \
    printf("  %-50s", #name); \
    name(); \
    tests_run++; \
    printf("PASS\n"); \
} while(0)

/* ---- Test: Basic exercise ---- */
static void test_basic_exercise(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 8 8 135", &r) == 1);
    assert(r.exercise_count == 1);
    assert(strcmp(r.exercises[0].name, "Squat") == 0);
    assert(r.exercises[0].set_count == 3);
    assert(r.exercises[0].sets[0].reps == 8);
    assert(r.exercises[0].sets[0].weight == 135);
    assert(r.exercises[0].sets[1].reps == 8);
    assert(r.exercises[0].sets[1].weight == 135);
    assert(r.exercises[0].sets[2].reps == 8);
    assert(r.exercises[0].sets[2].weight == 135);
}

/* ---- Test: Multiplier syntax ---- */
static void test_multiplier_syntax(void) {
    WorkoutResult r;
    assert(workout_parse("squat 3x8x135", &r) == 1);
    assert(r.exercises[0].set_count == 3);
    assert(r.exercises[0].sets[0].reps == 8);
    assert(r.exercises[0].sets[0].weight == 135);
}

/* ---- Test: By syntax ---- */
static void test_by_syntax(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 by 135", &r) == 1);
    assert(r.exercises[0].set_count == 3);
    assert(r.exercises[0].sets[0].reps == 8);
    assert(r.exercises[0].sets[0].weight == 135);
}

/* ---- Test: Prose returns 0 ---- */
static void test_prose_returns_zero(void) {
    WorkoutResult r;
    assert(workout_parse("Just some notes today", &r) == 0);
}

/* ---- Test: Empty returns 0 ---- */
static void test_empty_returns_zero(void) {
    WorkoutResult r;
    assert(workout_parse("", &r) == 0);
}

/* ---- Test: Multiple exercises ---- */
static void test_multiple_exercises(void) {
    WorkoutResult r;
    assert(workout_parse("squat 8 8 8 135\nbench 3x8 95", &r) == 1);
    assert(r.exercise_count == 2);
    assert(strcmp(r.exercises[0].name, "Squat") == 0);
    assert(strcmp(r.exercises[1].name, "Bench") == 0);
}

/* ---- Test: Prose within workout ---- */
static void test_prose_within_workout(void) {
    WorkoutResult r;
    assert(workout_parse("ate a banana beforehand\nsquat 8 8 8 135", &r) == 1);
    assert(r.prose_count == 1);
    assert(strcmp(r.prose_lines[0], "ate a banana beforehand") == 0);
}

/* ---- Test: Continuation lines ---- */
static void test_continuation_lines(void) {
    WorkoutResult r;
    assert(workout_parse("Bench\n8x135\n8\n8\n8", &r) == 1);
    assert(r.exercises[0].set_count == 4);
    /* First set: 8 reps at 135 from 8x135 (partial multiplier: 8>REP_THRESHOLD? no; so pending) */
    /* Actually 8x135: 8 is sets field, 135 is reps field. Since 135 > REP_THRESHOLD,
       this is treated as 1 set of reps=8, weight=135 */
    assert(r.exercises[0].sets[0].reps == 8);
    assert(r.exercises[0].sets[0].weight == 135);
    /* Continuation lines carry weight from last set */
    assert(r.exercises[0].sets[1].reps == 8);
    assert(r.exercises[0].sets[1].weight == 135);
    assert(r.exercises[0].sets[2].reps == 8);
    assert(r.exercises[0].sets[2].weight == 135);
    assert(r.exercises[0].sets[3].reps == 8);
    assert(r.exercises[0].sets[3].weight == 135);
}

/* ---- Test: Bodyweight + weighted ---- */
static void test_bodyweight_and_weighted(void) {
    WorkoutResult r;
    assert(workout_parse("dips 8 8 8\nbench 8x135 8x130", &r) == 1);
    assert(r.exercise_count == 2);
    assert(r.exercises[0].sets[0].weight == 0); /* dips bodyweight */
    assert(r.exercises[1].sets[0].weight == 135);
}

/* ---- Test: Superset ---- */
static void test_superset(void) {
    WorkoutResult r;
    assert(workout_parse("dips* press* 8 135 8 135 8 135", &r) == 1);
    assert(r.exercise_count == 2);
    assert(strcmp(r.exercises[0].name, "Dips*") == 0);
    assert(strcmp(r.exercises[1].name, "Press*") == 0);
}

int main(void) {
    printf("Running workout parser tests...\n");

    RUN_TEST(test_basic_exercise);
    RUN_TEST(test_multiplier_syntax);
    RUN_TEST(test_by_syntax);
    RUN_TEST(test_prose_returns_zero);
    RUN_TEST(test_empty_returns_zero);
    RUN_TEST(test_multiple_exercises);
    RUN_TEST(test_prose_within_workout);
    RUN_TEST(test_continuation_lines);
    RUN_TEST(test_bodyweight_and_weighted);
    RUN_TEST(test_superset);

    printf("\nAll %d tests passed!\n", tests_run);
    return 0;
}
