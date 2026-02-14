.PHONY: grammar test test-c test-workout test-recipe test-swift emacs clean

grammar:
	cd grammars/workout && tree-sitter generate
	cd grammars/recipe && tree-sitter generate

test: test-c test-swift

test-c: test-workout test-recipe

test-workout:
	cc -I vendor/tree-sitter/lib/include -I vendor/tree-sitter/lib/src \
	   -I grammars/workout/src -I lib \
	   grammars/workout/src/parser.c vendor/tree-sitter/lib/src/lib.c lib/workout.c lib/workout_test.c \
	   -o workout_test && ./workout_test && rm workout_test

test-recipe:
	cc -I vendor/tree-sitter/lib/include -I vendor/tree-sitter/lib/src \
	   -I grammars/recipe/src -I lib \
	   grammars/recipe/src/parser.c vendor/tree-sitter/lib/src/lib.c lib/recipe.c lib/recipe_test.c \
	   -o recipe_test && ./recipe_test && rm recipe_test

test-swift:
	swift test

emacs:
	cd emacs && make

clean:
	rm -f workout_test recipe_test
	cd emacs && make clean
	swift package clean
