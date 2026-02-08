package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chijoshi/workoutliner/parser"
	"github.com/chijoshi/workoutliner/table"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	result, err := parser.Parse("input", input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	log := result.(*parser.WorkoutLog)
	fmt.Print(table.Format(log))
}

func readInput() ([]byte, error) {
	if len(os.Args) > 1 {
		return os.ReadFile(os.Args[1])
	}
	return io.ReadAll(os.Stdin)
}
