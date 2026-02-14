#include <string.h>
#include <emacs-module.h>
#include "workout.h"

int plugin_is_GPL_compatible;

/* Helper: create a cons cell (key . value) */
static emacs_value
make_cons(emacs_env *env, emacs_value car, emacs_value cdr)
{
    return env->funcall(env, env->intern(env, "cons"), 2,
                        (emacs_value[]){car, cdr});
}

/* Helper: create a vector of given size, initialized to nil */
static emacs_value
make_vector(emacs_env *env, ptrdiff_t size)
{
    return env->funcall(env, env->intern(env, "make-vector"), 2,
                        (emacs_value[]){env->make_integer(env, size),
                                        env->intern(env, "nil")});
}

/* Convert a WorkoutSet to an alist: ((reps . N) (weight . N) (note . "...")) */
static emacs_value
set_to_elisp(emacs_env *env, const WorkoutSet *s)
{
    emacs_value nil = env->intern(env, "nil");

    emacs_value reps_pair  = make_cons(env, env->intern(env, "reps"),
                                       env->make_integer(env, s->reps));
    emacs_value weight_pair = make_cons(env, env->intern(env, "weight"),
                                        env->make_integer(env, s->weight));
    emacs_value note_pair  = make_cons(env, env->intern(env, "note"),
                                       env->make_string(env, s->note,
                                                        strlen(s->note)));

    /* Build list: (note_pair . (weight_pair . (reps_pair . nil))) */
    emacs_value list = nil;
    list = make_cons(env, note_pair, list);
    list = make_cons(env, weight_pair, list);
    list = make_cons(env, reps_pair, list);
    return list;
}

/* Convert an Exercise to an alist: ((name . "...") (sets . [...])) */
static emacs_value
exercise_to_elisp(emacs_env *env, const Exercise *ex)
{
    emacs_value nil = env->intern(env, "nil");

    emacs_value name_pair = make_cons(env, env->intern(env, "name"),
                                      env->make_string(env, ex->name,
                                                       strlen(ex->name)));

    emacs_value sets_vec = make_vector(env, ex->set_count);
    for (int i = 0; i < ex->set_count; i++) {
        env->vec_set(env, sets_vec, i, set_to_elisp(env, &ex->sets[i]));
    }

    emacs_value sets_pair = make_cons(env, env->intern(env, "sets"), sets_vec);

    emacs_value list = nil;
    list = make_cons(env, sets_pair, list);
    list = make_cons(env, name_pair, list);
    return list;
}

/* Convert a WorkoutResult to an alist:
   ((exercises . [...]) (prose . [...])) */
static emacs_value
result_to_elisp(emacs_env *env, const WorkoutResult *r)
{
    emacs_value nil = env->intern(env, "nil");

    /* exercises vector */
    emacs_value ex_vec = make_vector(env, r->exercise_count);
    for (int i = 0; i < r->exercise_count; i++) {
        env->vec_set(env, ex_vec, i,
                     exercise_to_elisp(env, &r->exercises[i]));
    }
    emacs_value ex_pair = make_cons(env, env->intern(env, "exercises"), ex_vec);

    /* prose vector */
    emacs_value prose_vec = make_vector(env, r->prose_count);
    for (int i = 0; i < r->prose_count; i++) {
        const char *line = r->prose_lines[i];
        env->vec_set(env, prose_vec, i,
                     env->make_string(env, line, strlen(line)));
    }
    emacs_value prose_pair = make_cons(env, env->intern(env, "prose"),
                                       prose_vec);

    /* Build the top-level alist */
    emacs_value list = nil;
    list = make_cons(env, prose_pair, list);
    list = make_cons(env, ex_pair, list);
    return list;
}

/* (snap-workout-parse STRING) -> alist or nil */
static emacs_value
Fsnap_workout_parse(emacs_env *env, ptrdiff_t nargs,
                    emacs_value *args, void *data)
{
    (void)nargs;
    (void)data;

    /* Extract the string argument */
    ptrdiff_t len = 0;
    env->copy_string_contents(env, args[0], NULL, &len);

    char buf[len];
    env->copy_string_contents(env, args[0], buf, &len);

    WorkoutResult result;
    memset(&result, 0, sizeof(result));

    if (workout_parse(buf, &result)) {
        return result_to_elisp(env, &result);
    }

    return env->intern(env, "nil");
}

int
emacs_module_init(struct emacs_runtime *runtime)
{
    if (runtime->size < (ptrdiff_t)sizeof(*runtime))
        return 1;

    emacs_env *env = runtime->get_environment(runtime);
    if (env->size < (ptrdiff_t)sizeof(*env))
        return 2;

    emacs_value func = env->make_function(env, 1, 1, Fsnap_workout_parse,
                                          "Parse a workout string into structured data.\n"
                                          "Returns an alist with exercises and prose, or nil.",
                                          NULL);

    emacs_value sym = env->intern(env, "snap-workout-parse");
    env->funcall(env, env->intern(env, "defalias"), 2,
                 (emacs_value[]){sym, func});

    /* (provide 'snap-core) */
    env->funcall(env, env->intern(env, "provide"), 1,
                 (emacs_value[]){env->intern(env, "snap-core")});

    return 0;
}
