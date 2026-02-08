package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/chijoshi/workoutliner/transform"
)

func main() {
	flag.Parse()

	input, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(transform.Transform(string(input)))
}

func readInput() ([]byte, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.ReadFile(args[0])
	}
	return io.ReadAll(os.Stdin)
}
