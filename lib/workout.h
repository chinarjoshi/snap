#ifndef WORKOUT_H
#define WORKOUT_H

#define WORKOUT_NOTE_MAX 256
#define WORKOUT_NAME_MAX 256
#define WORKOUT_MAX_SETS 64
#define WORKOUT_MAX_EXERCISES 32
#define WORKOUT_MAX_PROSE 32

typedef struct {
    int reps;
    int weight;
    char note[WORKOUT_NOTE_MAX];
} WorkoutSet;

typedef struct {
    char name[WORKOUT_NAME_MAX];
    WorkoutSet sets[WORKOUT_MAX_SETS];
    int set_count;
} Exercise;

typedef struct {
    Exercise exercises[WORKOUT_MAX_EXERCISES];
    int exercise_count;
    char prose_lines[WORKOUT_MAX_PROSE][WORKOUT_NAME_MAX];
    int prose_count;
} WorkoutResult;

// Returns 1 if input was parsed as a workout, 0 if not.
int workout_parse(const char *input, WorkoutResult *result);

#endif
