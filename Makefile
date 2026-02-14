TS_PREFIX := $(shell brew --prefix tree-sitter 2>/dev/null || echo /opt/homebrew/opt/tree-sitter)

test-c:
	cc -I $(TS_PREFIX)/include -I grammars/workout/src -I lib \
	   grammars/workout/src/parser.c lib/workout.c lib/workout_test.c \
	   -L $(TS_PREFIX)/lib -ltree-sitter \
	   -o workout_test && ./workout_test && rm workout_test

.PHONY: test-c
