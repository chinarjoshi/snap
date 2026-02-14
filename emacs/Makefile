CFLAGS = -I../vendor/tree-sitter/lib/include -I../vendor/tree-sitter/lib/src -I../grammars/workout/src -I../grammars/recipe/src -I../lib
SRCS = snap.c ../lib/workout.c ../lib/recipe.c ../grammars/workout/src/parser.c ../grammars/recipe/src/parser.c ../vendor/tree-sitter/lib/src/lib.c

# Find emacs-module.h: try Emacs.app, then Homebrew, then system paths
EMACS_INCLUDE := $(shell \
  for d in /Applications/Emacs.app/Contents/Resources/include \
           /opt/homebrew/include \
           /usr/local/include; do \
    [ -f "$$d/emacs-module.h" ] && echo "-I$$d" && break; \
  done)
CFLAGS += $(EMACS_INCLUDE)

ifeq ($(shell uname), Darwin)
  SO = snap-core.dylib
  LDFLAGS = -dynamiclib -undefined dynamic_lookup
else
  SO = snap-core.so
  LDFLAGS = -shared
endif

all: $(SO)

$(SO): $(SRCS)
	$(CC) -fPIC $(CFLAGS) $(LDFLAGS) -o $@ $^

clean:
	rm -f $(SO)
