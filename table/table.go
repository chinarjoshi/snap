package table

import (
	"fmt"
	"strings"

	"github.com/chijoshi/workoutliner/parser"
)

// Format returns a markdown table representation of the workout log
func Format(log *parser.WorkoutLog) string {
	if len(log.Exercises) == 0 {
		return ""
	}

	maxSets := log.MaxSets()
	if maxSets == 0 {
		maxSets = 1
	}

	var sb strings.Builder

	// Header row
	sb.WriteString("|")
	sb.WriteString(pad("", maxExerciseNameLen(log)))
	sb.WriteString(" |")
	for i := 1; i <= maxSets; i++ {
		sb.WriteString(fmt.Sprintf(" Set %d |", i))
	}
	sb.WriteString("\n")

	// Separator row
	sb.WriteString("|")
	sb.WriteString(strings.Repeat("-", maxExerciseNameLen(log)+1))
	sb.WriteString("|")
	for i := 0; i < maxSets; i++ {
		sb.WriteString("---------|")
	}
	sb.WriteString("\n")

	// Data rows
	for _, ex := range log.Exercises {
		sb.WriteString("| ")
		sb.WriteString(padRight(ex.Name, maxExerciseNameLen(log)))
		sb.WriteString("|")
		for i := 0; i < maxSets; i++ {
			if i < len(ex.Sets) {
				sb.WriteString(" ")
				sb.WriteString(formatSet(ex.Sets[i]))
				sb.WriteString(" |")
			} else {
				sb.WriteString("         |")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func formatSet(s parser.Set) string {
	result := fmt.Sprintf("%d @ %d", s.Reps, s.Weight)
	if s.Note != "" {
		result += fmt.Sprintf(" (%s)", s.Note)
	}
	return result
}

func maxExerciseNameLen(log *parser.WorkoutLog) int {
	max := 0
	for _, ex := range log.Exercises {
		if len(ex.Name) > max {
			max = len(ex.Name)
		}
	}
	if max < 8 {
		max = 8
	}
	return max
}

func pad(s string, length int) string {
	return fmt.Sprintf("%*s", length, s)
}

func padRight(s string, length int) string {
	return fmt.Sprintf("%-*s", length, s)
}
