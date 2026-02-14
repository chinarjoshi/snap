.PHONY: antlr test build clean

antlr:
	antlr -Dlanguage=Go -visitor -package workout -o libs/workout Workoutliner.g4
	antlr -Dlanguage=Swift -visitor -o swift/Sources/WorkoutParser/Parser Workoutliner.g4

test:
	go test ./...
	cd swift && swift test

build:
	go build -o snap ./cmd/snap

clean:
	rm -f snap
	rm -f libs/workout/*.interp libs/workout/*.tokens
	rm -f swift/Sources/WorkoutParser/Parser/*.interp swift/Sources/WorkoutParser/Parser/*.tokens
