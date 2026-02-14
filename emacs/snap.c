#include <string.h>
#include <emacs-module.h>
#include "workout.h"
#include "recipe.h"

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

/* Convert a WorkoutSet to alist */
static emacs_value
set_to_elisp(emacs_env *env, const WorkoutSet *s)
{
    emacs_value nil = env->intern(env, "nil");
    emacs_value list = nil;
    list = make_cons(env, make_cons(env, env->intern(env, "note"),
           env->make_string(env, s->note, strlen(s->note))), list);
    list = make_cons(env, make_cons(env, env->intern(env, "weight"),
           env->make_integer(env, s->weight)), list);
    list = make_cons(env, make_cons(env, env->intern(env, "reps"),
           env->make_integer(env, s->reps)), list);
    return list;
}

/* Convert an Exercise to alist */
static emacs_value
exercise_to_elisp(emacs_env *env, const Exercise *ex)
{
    emacs_value nil = env->intern(env, "nil");
    emacs_value sets_vec = make_vector(env, ex->set_count);
    for (int i = 0; i < ex->set_count; i++)
        env->vec_set(env, sets_vec, i, set_to_elisp(env, &ex->sets[i]));

    emacs_value list = nil;
    list = make_cons(env, make_cons(env, env->intern(env, "sets"), sets_vec), list);
    list = make_cons(env, make_cons(env, env->intern(env, "name"),
           env->make_string(env, ex->name, strlen(ex->name))), list);
    return list;
}

/* Convert WorkoutResult to alist */
static emacs_value
workout_to_elisp(emacs_env *env, const WorkoutResult *r)
{
    emacs_value nil = env->intern(env, "nil");

    emacs_value ex_vec = make_vector(env, r->exercise_count);
    for (int i = 0; i < r->exercise_count; i++)
        env->vec_set(env, ex_vec, i, exercise_to_elisp(env, &r->exercises[i]));

    emacs_value prose_vec = make_vector(env, r->prose_count);
    for (int i = 0; i < r->prose_count; i++) {
        const char *line = r->prose_lines[i];
        env->vec_set(env, prose_vec, i, env->make_string(env, line, strlen(line)));
    }

    emacs_value list = nil;
    list = make_cons(env, make_cons(env, env->intern(env, "prose"), prose_vec), list);
    list = make_cons(env, make_cons(env, env->intern(env, "exercises"), ex_vec), list);
    return list;
}

/* Convert Ingredient to alist */
static emacs_value
ingredient_to_elisp(emacs_env *env, const Ingredient *ing)
{
    emacs_value nil = env->intern(env, "nil");
    emacs_value list = nil;
    list = make_cons(env, make_cons(env, env->intern(env, "text"),
           env->make_string(env, ing->text, strlen(ing->text))), list);
    list = make_cons(env, make_cons(env, env->intern(env, "quantity"),
           env->make_string(env, ing->quantity, strlen(ing->quantity))), list);
    return list;
}

/* Convert RecipeSection to alist */
static emacs_value
recipe_section_to_elisp(emacs_env *env, const RecipeSection *sec)
{
    emacs_value nil = env->intern(env, "nil");

    emacs_value ing_vec = make_vector(env, sec->ingredient_count);
    for (int i = 0; i < sec->ingredient_count; i++)
        env->vec_set(env, ing_vec, i, ingredient_to_elisp(env, &sec->ingredients[i]));

    emacs_value instr_vec = make_vector(env, sec->instruction_count);
    for (int i = 0; i < sec->instruction_count; i++) {
        const char *instr = sec->instructions[i];
        env->vec_set(env, instr_vec, i, env->make_string(env, instr, strlen(instr)));
    }

    emacs_value list = nil;
    list = make_cons(env, make_cons(env, env->intern(env, "instructions"), instr_vec), list);
    list = make_cons(env, make_cons(env, env->intern(env, "ingredients"), ing_vec), list);
    list = make_cons(env, make_cons(env, env->intern(env, "name"),
           env->make_string(env, sec->name, strlen(sec->name))), list);
    return list;
}

/* Convert RecipeResult to alist */
static emacs_value
recipe_to_elisp(emacs_env *env, const RecipeResult *r)
{
    emacs_value nil = env->intern(env, "nil");

    emacs_value sec_vec = make_vector(env, r->section_count);
    for (int i = 0; i < r->section_count; i++)
        env->vec_set(env, sec_vec, i, recipe_section_to_elisp(env, &r->sections[i]));

    emacs_value list = nil;
    list = make_cons(env, make_cons(env, env->intern(env, "sections"), sec_vec), list);
    list = make_cons(env, make_cons(env, env->intern(env, "title"),
           env->make_string(env, r->title, strlen(r->title))), list);
    return list;
}

/* (snap-workout STRING) -> alist or nil */
static emacs_value
Fsnap_workout(emacs_env *env, ptrdiff_t nargs,
                    emacs_value *args, void *data)
{
    (void)nargs; (void)data;
    ptrdiff_t len = 0;
    env->copy_string_contents(env, args[0], NULL, &len);
    char buf[len];
    env->copy_string_contents(env, args[0], buf, &len);

    WorkoutResult result;
    memset(&result, 0, sizeof(result));
    if (workout_parse(buf, &result))
        return workout_to_elisp(env, &result);
    return env->intern(env, "nil");
}

/* (snap-recipe STRING) -> alist or nil */
static emacs_value
Fsnap_recipe(emacs_env *env, ptrdiff_t nargs,
                   emacs_value *args, void *data)
{
    (void)nargs; (void)data;
    ptrdiff_t len = 0;
    env->copy_string_contents(env, args[0], NULL, &len);
    char buf[len];
    env->copy_string_contents(env, args[0], buf, &len);

    RecipeResult result;
    memset(&result, 0, sizeof(result));
    if (recipe_parse(buf, &result))
        return recipe_to_elisp(env, &result);
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

    /* Register snap-workout */
    emacs_value wfunc = env->make_function(env, 1, 1, Fsnap_workout,
                                           "Parse workout text.", NULL);
    env->funcall(env, env->intern(env, "defalias"), 2,
                 (emacs_value[]){env->intern(env, "snap-workout"), wfunc});

    /* Register snap-recipe */
    emacs_value rfunc = env->make_function(env, 1, 1, Fsnap_recipe,
                                           "Parse recipe text.", NULL);
    env->funcall(env, env->intern(env, "defalias"), 2,
                 (emacs_value[]){env->intern(env, "snap-recipe"), rfunc});

    /* (provide 'snap-core) */
    env->funcall(env, env->intern(env, "provide"), 1,
                 (emacs_value[]){env->intern(env, "snap-core")});

    return 0;
}
