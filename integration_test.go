package main

import (
	"strings"
	"testing"

	"github.com/chijoshi/workoutliner/parser"
	"github.com/chijoshi/workoutliner/table"
)

func TestFullWorkflow(t *testing.T) {
	input := `squat 8 8 8 135 light headed. leg press 8 315 7 250 7. bench 135 super hard 135 125 115`

	result, err := parser.Parse("test", []byte(input))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	log := result.(*parser.WorkoutLog)

	// Verify exercises parsed
	if len(log.Exercises) != 3 {
		t.Fatalf("expected 3 exercises, got %d", len(log.Exercises))
	}

	// Verify names
	names := []string{"Squat", "Leg Press", "Bench"}
	for i, name := range names {
		if log.Exercises[i].Name != name {
			t.Errorf("exercise %d: expected '%s', got '%s'", i, name, log.Exercises[i].Name)
		}
	}

	// Verify table output
	output := table.Format(log)
	if !strings.Contains(output, "Squat") {
		t.Error("table missing Squat")
	}
	if !strings.Contains(output, "Leg Press") {
		t.Error("table missing Leg Press")
	}
	if !strings.Contains(output, "Bench") {
		t.Error("table missing Bench")
	}
}
