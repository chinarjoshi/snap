.PHONY: antlr test build clean

antlr:
	antlr4 -Dlanguage=Go -visitor -package paragraph -o paragraph Paragraph.g4
	antlr4 -Dlanguage=Swift -visitor -o swift/Sources/WorkoutLiner/Parser Paragraph.g4

test:
	go test ./...
	cd swift && swift test

build:
	go build -o workoutliner ./cmd/workoutliner

clean:
	rm -f workoutliner
	rm -f paragraph/*.interp paragraph/*.tokens
	rm -f swift/Sources/WorkoutLiner/Parser/*.interp swift/Sources/WorkoutLiner/Parser/*.tokens
