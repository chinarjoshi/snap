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
	for _, line := range strings.Split(s, "\n") {
		// Skip lines with table formatting
		if strings.Contains(line, "|") {
			continue
		}
		for _, r := range line {
			if r >= '0' && r <= '9' {
				return true
			}
		}
	}
	return false
}

func transformParagraph(para string) string {
	if !containsDigit(para) {
		return para
	}

	// Check if paragraph starts with an existing table
	lines := strings.Split(para, "\n")
	if len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[0]), "|") {
		return transformTableWithTrailing(lines)
	}

	parseResult, err := workoutliner.Parse(para)
	if err != nil {
		return para
	}

	if !isWorkout(parseResult) {
		return para
	}

	return formatResult(parseResult)
}

func transformTableWithTrailing(lines []string) string {
	var tableLines []string
	var trailingLines []string
	var descListLines []string
	inTable := true

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if inTable {
			if strings.HasPrefix(trimmed, "|") {
				tableLines = append(tableLines, line)
			} else if strings.HasPrefix(trimmed, "- ") && strings.Contains(trimmed, " :: ") {
				// Description list item (notes)
				descListLines = append(descListLines, line)
			} else if trimmed == "" {
				// Empty line between table and desc list, skip
				continue
			} else {
				inTable = false
				trailingLines = append(trailingLines, line)
			}
		} else {
			trailingLines = append(trailingLines, line)
		}
	}

	// If no trailing lines, return original
	if len(trailingLines) == 0 {
		return strings.Join(lines, "\n")
	}

	// Parse existing table into exercises
	existingExercises := parseTableLines(tableLines)

	// Parse existing description list into notes
	existingNotes := parseDescList(descListLines)
	for i := range existingExercises {
		if note, ok := existingNotes[existingExercises[i].Name]; ok {
			for j := range existingExercises[i].Sets {
				if existingExercises[i].Sets[j].Note == "" {
					existingExercises[i].Sets[j].Note = note
					break
				}
			}
		}
	}

	// Parse trailing lines as new exercises
	trailingPara := strings.Join(trailingLines, "\n")
	parseResult, err := workoutliner.Parse(trailingPara)
	if err != nil || !parseResult.HasExercises() {
		return strings.Join(lines, "\n")
	}

	// Merge exercises
	allExercises := append(existingExercises, parseResult.Exercises...)

	// Format combined result
	result := &workoutliner.ParseResult{
		ProseLines: parseResult.ProseLines,
		Exercises:  allExercises,
	}

	return formatResult(result)
}

func parseTableLines(lines []string) []workoutliner.Exercise {
	var exercises []workoutliner.Exercise

	for _, line := range lines {
		// Skip header and separator
		if strings.Contains(line, "---") {
			continue
		}

		cells := strings.Split(line, "|")
		if len(cells) < 3 {
			continue
		}

		// First cell (after empty) is the name
		name := strings.TrimSpace(cells[1])
		if name == "" {
			continue // Skip header row
		}

		var sets []workoutliner.Set
		for _, cell := range cells[2:] {
			cell = strings.TrimSpace(cell)
			if cell == "" {
				continue
			}
			// Parse "8@135" format
			if strings.Contains(cell, "@") {
				parts := strings.Split(cell, "@")
				if len(parts) == 2 {
					reps := parseInt(parts[0])
					weight := parseInt(parts[1])
					if reps > 0 && weight > 0 {
						sets = append(sets, workoutliner.Set{Reps: reps, Weight: weight})
					}
				}
			}
		}

		if len(sets) > 0 {
			exercises = append(exercises, workoutliner.Exercise{Name: name, Sets: sets})
		}
	}

	return exercises
}

func parseDescList(lines []string) map[string]string {
	notes := make(map[string]string)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "- ") {
			line = strings.TrimPrefix(line, "- ")
			if idx := strings.Index(line, " :: "); idx > 0 {
				name := line[:idx]
				note := line[idx+4:]
				notes[name] = note
			}
		}
	}
	return notes
}

func isWorkout(result *workoutliner.ParseResult) bool {
	if !result.HasExercises() {
		return false
	}
	totalSets := 0
	hasWeight := false
	for _, ex := range result.Exercises {
		totalSets += len(ex.Sets)
		for _, s := range ex.Sets {
			if s.Weight > 0 {
				hasWeight = true
			}
		}
	}
	return totalSets >= 2 && hasWeight
}

func parseInt(s string) int {
	var n int
	fmt.Sscanf(strings.TrimSpace(s), "%d", &n)
	return n
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
	if s.Weight == 0 {
		return fmt.Sprintf("%d", s.Reps)
	}
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
	return strings.Join(notes, ". ")
}

func padRight(s string, length int) string {
	return fmt.Sprintf("%-*s", length, s)
}
