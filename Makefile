.PHONY: antlr test build clean

antlr:
	antlr -Dlanguage=Go -visitor -package workoutliner -o libs/workoutliner Workoutliner.g4
	antlr -Dlanguage=Swift -visitor -o swift/Sources/WorkoutLiner/Parser Workoutliner.g4

test:
	go test ./...
	cd swift && swift test

build:
	go build -o workoutliner ./cmd/workoutliner

clean:
	rm -f workoutliner
	rm -f libs/workoutliner/*.interp libs/workoutliner/*.tokens
	rm -f swift/Sources/WorkoutLiner/Parser/*.interp swift/Sources/WorkoutLiner/Parser/*.tokens
