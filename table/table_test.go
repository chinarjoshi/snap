package table

import (
	"strings"
	"testing"

	"github.com/chijoshi/workoutliner/paragraph"
)

func TestTableOutput(t *testing.T) {
	exercises := []paragraph.Exercise{
		{Name: "Squat", Sets: []paragraph.Set{{Reps: 8, Weight: 135}, {Reps: 8, Weight: 135}}},
		{Name: "Bench", Sets: []paragraph.Set{{Reps: 5, Weight: 185}}},
	}

	output := FormatExercises(exercises)

	// Check header row has column numbers
	if !strings.Contains(output, "| 1") {
		t.Error("expected column 1 in header")
	}
	if !strings.Contains(output, "| 2") {
		t.Error("expected column 2 in header")
	}

	// Check exercise rows
	if !strings.Contains(output, "Squat") {
		t.Error("expected 'Squat' in output")
	}
	if !strings.Contains(output, "8@135") {
		t.Error("expected '8@135' in output")
	}
}
