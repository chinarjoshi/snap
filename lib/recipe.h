#ifndef RECIPE_H
#define RECIPE_H

#define RECIPE_NAME_MAX 256
#define RECIPE_TEXT_MAX 512
#define RECIPE_QTY_MAX 32
#define RECIPE_MAX_INGREDIENTS 64
#define RECIPE_MAX_SECTIONS 16
#define RECIPE_MAX_INSTRUCTIONS 32

typedef struct {
    char quantity[RECIPE_QTY_MAX];
    char text[RECIPE_NAME_MAX];
} Ingredient;

typedef struct {
    char name[RECIPE_NAME_MAX];
    Ingredient ingredients[RECIPE_MAX_INGREDIENTS];
    int ingredient_count;
    char instructions[RECIPE_MAX_INSTRUCTIONS][RECIPE_TEXT_MAX];
    int instruction_count;
} RecipeSection;

typedef struct {
    char title[RECIPE_NAME_MAX];
    RecipeSection sections[RECIPE_MAX_SECTIONS];
    int section_count;
} RecipeResult;

int recipe_parse(const char *input, RecipeResult *result);

#endif
