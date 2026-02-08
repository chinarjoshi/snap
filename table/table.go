package table

import (
	"fmt"
	"strings"

	"github.com/chijoshi/workoutliner/parser"
)

func Format(log *parser.WorkoutLog) string {
	return FormatExercises(log.Exercises)
}

func FormatExercises(exercises []parser.Exercise) string {
	if len(exercises) == 0 {
		return ""
	}

	maxSets := maxSetsInExercises(exercises)
	if maxSets == 0 {
		maxSets = 1
	}

	hasNotes := exercisesHaveNotes(exercises)
	maxNameLen := maxExerciseNameLenFromSlice(exercises)
	maxNoteLen := maxNoteLenFromSlice(exercises)

	var sb strings.Builder

	sb.WriteString("| ")
	sb.WriteString(padRight("Exercise", maxNameLen))
	sb.WriteString("|")
	for i := 1; i <= maxSets; i++ {
		sb.WriteString(fmt.Sprintf(" Set %d |", i))
	}
	if hasNotes {
		sb.WriteString(" ")
		sb.WriteString(padRight("Notes", maxNoteLen))
		sb.WriteString("|")
	}
	sb.WriteString("\n")

	sb.WriteString("|")
	sb.WriteString(strings.Repeat("-", maxNameLen+2))
	sb.WriteString("|")
	for i := 0; i < maxSets; i++ {
		sb.WriteString("-------|")
	}
	if hasNotes {
		sb.WriteString(strings.Repeat("-", maxNoteLen+2))
		sb.WriteString("|")
	}
	sb.WriteString("\n")

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
		if hasNotes {
			note := collectNotes(ex.Sets)
			sb.WriteString(" ")
			sb.WriteString(padRight(note, maxNoteLen))
			sb.WriteString("|")
		}
		sb.WriteString("\n")
	}

	return strings.TrimSuffix(sb.String(), "\n")
}

func formatSet(s parser.Set) string {
	return fmt.Sprintf("%d@%d", s.Reps, s.Weight)
}

func maxSetsInExercises(exercises []parser.Exercise) int {
	max := 0
	for _, ex := range exercises {
		if len(ex.Sets) > max {
			max = len(ex.Sets)
		}
	}
	return max
}

func maxExerciseNameLenFromSlice(exercises []parser.Exercise) int {
	max := len("Exercise")
	for _, ex := range exercises {
		if len(ex.Name) > max {
			max = len(ex.Name)
		}
	}
	return max
}

func exercisesHaveNotes(exercises []parser.Exercise) bool {
	for _, ex := range exercises {
		for _, s := range ex.Sets {
			if s.Note != "" {
				return true
			}
		}
	}
	return false
}

func maxNoteLenFromSlice(exercises []parser.Exercise) int {
	max := len("Notes")
	for _, ex := range exercises {
		noteLen := len(collectNotes(ex.Sets))
		if noteLen > max {
			max = noteLen
		}
	}
	return max
}

func collectNotes(sets []parser.Set) string {
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
