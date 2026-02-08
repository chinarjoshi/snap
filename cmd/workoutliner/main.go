package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chijoshi/workoutliner/libs/workoutliner"
)

func main() {
	flag.Parse()

	input, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(transform(string(input)))
}

func readInput() ([]byte, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.ReadFile(args[0])
	}
	return io.ReadAll(os.Stdin)
}

func transform(input string) string {
	paragraphs := strings.Split(input, "\n\n")
	result := make([]string, 0, len(paragraphs))

	for _, para := range paragraphs {
		transformed := transformParagraph(para)
		result = append(result, transformed)
	}

	return strings.Join(result, "\n\n")
}

func containsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func transformParagraph(para string) string {
	if !containsDigit(para) {
		return para
	}

	parseResult, err := workoutliner.Parse(para)
	if err != nil {
		return para
	}

	if !parseResult.HasExercises() {
		return para
	}

	return formatResult(parseResult)
}

func formatResult(result *workoutliner.ParseResult) string {
	var parts []string

	if len(result.ProseLines) > 0 {
		parts = append(parts, strings.Join(result.ProseLines, "\n"))
	}

	tableStr := formatExercises(result.Exercises)
	parts = append(parts, tableStr)

	return strings.Join(parts, "\n\n")
}

func formatExercises(exercises []workoutliner.Exercise) string {
	if len(exercises) == 0 {
		return ""
	}

	maxSets := maxSetsInExercises(exercises)
	if maxSets == 0 {
		maxSets = 1
	}

	maxNameLen := maxExerciseNameLen(exercises)

	var sb strings.Builder

	// Header
	sb.WriteString("| ")
	sb.WriteString(padRight("", maxNameLen))
	sb.WriteString("|")
	for i := 0; i < maxSets; i++ {
		sb.WriteString(fmt.Sprintf(" %-5d |", i))
	}
	sb.WriteString("\n")

	// Separator
	sb.WriteString("|")
	sb.WriteString(strings.Repeat("-", maxNameLen+2))
	sb.WriteString("|")
	for i := 0; i < maxSets; i++ {
		sb.WriteString("-------|")
	}
	sb.WriteString("\n")

	// Data rows
	for _, ex := range exercises {
		sb.WriteString("| ")
		sb.WriteString(padRight(ex.Name, maxNameLen))
		sb.WriteString("|")
		for i := 0; i < maxSets; i++ {
			if i < len(ex.Sets) {
				sb.WriteString(" ")
				sb.WriteString(padRight(formatSet(ex.Sets[i]), 5))
				sb.WriteString(" |")
			} else {
				sb.WriteString("       |")
			}
		}
		sb.WriteString("\n")
	}

	// Notes as description list
	var notes []string
	for _, ex := range exercises {
		note := collectNotes(ex.Sets)
		if note != "" {
			notes = append(notes, fmt.Sprintf("- %s :: %s", ex.Name, note))
		}
	}
	if len(notes) > 0 {
		sb.WriteString("\n")
		sb.WriteString(strings.Join(notes, "\n"))
	}

	return strings.TrimSuffix(sb.String(), "\n")
}

func formatSet(s workoutliner.Set) string {
	return fmt.Sprintf("%d@%d", s.Reps, s.Weight)
}

func maxSetsInExercises(exercises []workoutliner.Exercise) int {
	max := 0
	for _, ex := range exercises {
		if len(ex.Sets) > max {
			max = len(ex.Sets)
		}
	}
	return max
}

func maxExerciseNameLen(exercises []workoutliner.Exercise) int {
	max := 0
	for _, ex := range exercises {
		if len(ex.Name) > max {
			max = len(ex.Name)
		}
	}
	return max
}


func collectNotes(sets []workoutliner.Set) string {
	var notes []string
	for _, s := range sets {
		if s.Note != "" {
			notes = append(notes, s.Note)
		}
	}
	return strings.Join(notes, ", ")
}

func padRight(s string, length int) string {
	return fmt.Sprintf("%-*s", length, s)
}
