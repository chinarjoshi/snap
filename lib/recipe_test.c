#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "recipe.h"

static RecipeSection *find_section(RecipeResult *r, const char *name) {
    for (int i = 0; i < r->section_count; i++) {
        if (strcmp(r->sections[i].name, name) == 0) return &r->sections[i];
    }
    return NULL;
}

static void test_basic_recipe_with_sections(void) {
    RecipeResult r;
    int ok = recipe_parse(
        "chicken biryani:\n"
        "chicken marinade:\n"
        "1 tbsp jeera\n"
        "1 bay leaf\n"
        "rice water:\n"
        "2 cups rice", &r);
    assert(ok == 1);
    assert(strcmp(r.title, "Chicken Biryani") == 0);
    assert(r.section_count == 2);

    RecipeSection *cm = find_section(&r, "Chicken Marinade");
    assert(cm != NULL);
    assert(cm->ingredient_count == 2);
    assert(strcmp(cm->ingredients[0].quantity, "1") == 0);
    assert(strcmp(cm->ingredients[0].text, "tbsp jeera") == 0);
    assert(strcmp(cm->ingredients[1].quantity, "1") == 0);
    assert(strcmp(cm->ingredients[1].text, "bay leaf") == 0);

    RecipeSection *rw = find_section(&r, "Rice Water");
    assert(rw != NULL);
    assert(rw->ingredient_count == 1);
    assert(strcmp(rw->ingredients[0].quantity, "2") == 0);
    assert(strcmp(rw->ingredients[0].text, "cups rice") == 0);

    printf("  test_basic_recipe_with_sections                       PASS\n");
}

static void test_recipe_with_instructions(void) {
    RecipeResult r;
    int ok = recipe_parse(
        "pasta:\n"
        "2 cups flour\n"
        "1 egg\n"
        "Mix well\n"
        "Knead for 10 min", &r);
    assert(ok == 1);
    assert(strcmp(r.title, "Pasta") == 0);
    assert(r.section_count == 1);
    assert(r.sections[0].ingredient_count == 2);
    assert(r.sections[0].instruction_count == 2);

    printf("  test_recipe_with_instructions                         PASS\n");
}

static void test_range_quantity(void) {
    RecipeResult r;
    int ok = recipe_parse(
        "spice mix:\n"
        "3-5 pieces cardamon\n"
        "8 cloves", &r);
    assert(ok == 1);
    assert(r.section_count == 1);
    assert(r.sections[0].ingredient_count == 2);
    assert(strcmp(r.sections[0].ingredients[0].quantity, "3-5") == 0);
    assert(strcmp(r.sections[0].ingredients[0].text, "pieces cardamon") == 0);
    assert(strcmp(r.sections[0].ingredients[1].quantity, "8") == 0);
    assert(strcmp(r.sections[0].ingredients[1].text, "cloves") == 0);

    printf("  test_range_quantity                                   PASS\n");
}

static void test_not_recipe_prose(void) {
    RecipeResult r;
    assert(recipe_parse("Just some notes today", &r) == 0);
    printf("  test_not_recipe_prose                                 PASS\n");
}

static void test_not_recipe_no_ingredients(void) {
    RecipeResult r;
    assert(recipe_parse("dinner plan:\nGo to the store", &r) == 0);
    printf("  test_not_recipe_no_ingredients                        PASS\n");
}

static void test_not_recipe_empty(void) {
    RecipeResult r;
    assert(recipe_parse("", &r) == 0);
    printf("  test_not_recipe_empty                                 PASS\n");
}

static void test_title_case(void) {
    RecipeResult r;
    int ok = recipe_parse(
        "CHICKEN TIKKA:\n"
        "1 cup yogurt\n"
        "1 tsp turmeric", &r);
    assert(ok == 1);
    assert(strcmp(r.title, "Chicken Tikka") == 0);

    printf("  test_title_case                                       PASS\n");
}

static void test_recipe_no_subsections(void) {
    RecipeResult r;
    int ok = recipe_parse(
        "smoothie:\n"
        "1 banana\n"
        "2 cups milk\n"
        "1 tbsp honey", &r);
    assert(ok == 1);
    assert(strcmp(r.title, "Smoothie") == 0);
    assert(r.section_count == 1);
    assert(r.sections[0].name[0] == '\0');
    assert(r.sections[0].ingredient_count == 3);

    printf("  test_recipe_no_subsections                            PASS\n");
}

static void test_mixed_instructions_and_ingredients(void) {
    RecipeResult r;
    int ok = recipe_parse(
        "soup:\n"
        "1 onion\n"
        "Dice finely\n"
        "2 cups broth\n"
        "Simmer for 20 min", &r);
    assert(ok == 1);
    assert(r.section_count == 1);
    assert(r.sections[0].ingredient_count == 2);
    assert(r.sections[0].instruction_count == 2);

    printf("  test_mixed_instructions_and_ingredients               PASS\n");
}

static void test_workout_not_recipe(void) {
    RecipeResult r;
    assert(recipe_parse("squat 3x8x135", &r) == 0);
    printf("  test_workout_not_recipe                               PASS\n");
}

int main(void) {
    printf("Running recipe parser tests (10 cases)...\n");

    test_basic_recipe_with_sections();
    test_recipe_with_instructions();
    test_range_quantity();
    test_not_recipe_prose();
    test_not_recipe_no_ingredients();
    test_not_recipe_empty();
    test_title_case();
    test_recipe_no_subsections();
    test_mixed_instructions_and_ingredients();
    test_workout_not_recipe();

    printf("\nAll 10 tests passed!\n");
    return 0;
}
