.PHONY: grammar test test-c test-swift emacs clean

grammar:
	cd grammars/workout && npx tree-sitter generate

test: test-c test-swift

test-c:
	cc -I vendor/tree-sitter/include -I vendor/tree-sitter/src \
	   -I grammars/workout/src -I lib \
	   grammars/workout/src/parser.c vendor/tree-sitter/src/lib.c lib/workout.c lib/workout_test.c \
	   -o workout_test && ./workout_test && rm workout_test

test-swift:
	swift test

emacs:
	cd emacs && make

clean:
	rm -f workout_test
	cd emacs && make clean
	swift package clean
