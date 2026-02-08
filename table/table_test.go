package table

import (
	"strings"
	"testing"

	"github.com/chijoshi/workoutliner/parser"
)

func TestTableOutput(t *testing.T) {
	log := &parser.WorkoutLog{
		Exercises: []parser.Exercise{
			{Name: "Squat", Sets: []parser.Set{{Reps: 8, Weight: 135}, {Reps: 8, Weight: 135}}},
			{Name: "Bench", Sets: []parser.Set{{Reps: 5, Weight: 185}}},
		},
	}

	output := Format(log)

	// Check header row
	if !strings.Contains(output, "Set 1") {
		t.Error("expected 'Set 1' in header")
	}
	if !strings.Contains(output, "Set 2") {
		t.Error("expected 'Set 2' in header")
	}

	// Check exercise rows
	if !strings.Contains(output, "Squat") {
		t.Error("expected 'Squat' in output")
	}
	if !strings.Contains(output, "8@135") {
		t.Error("expected '8@135' in output")
	}
}
