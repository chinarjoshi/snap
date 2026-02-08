package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	antlrparser "github.com/chijoshi/workoutliner/antlr"
	"github.com/chijoshi/workoutliner/parser"
	"github.com/chijoshi/workoutliner/table"
)

var usePeg = flag.Bool("peg", false, "use PEG parser instead of ANTLR")

func main() {
	flag.Parse()

	input, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	var log *parser.WorkoutLog

	if *usePeg {
		// Use PEG parser (pigeon)
		result, err := parser.Parse("input", input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
			os.Exit(1)
		}
		log = result.(*parser.WorkoutLog)
	} else {
		// Use ANTLR parser (default)
		log, err = antlrparser.Parse(string(input))
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Print(table.Format(log))
}

func readInput() ([]byte, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.ReadFile(args[0])
	}
	return io.ReadAll(os.Stdin)
}
