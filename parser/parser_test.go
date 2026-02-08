package parser

import (
	"testing"
)

func TestParseExerciseName(t *testing.T) {
	input := "squat 135"
	got, err := Parse("test", []byte(input))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	log, ok := got.(*WorkoutLog)
	if !ok {
		t.Fatalf("expected *WorkoutLog, got %T", got)
	}

	if len(log.Exercises) != 1 {
		t.Fatalf("expected 1 exercise, got %d", len(log.Exercises))
	}

	if log.Exercises[0].Name != "Squat" {
		t.Errorf("expected exercise name 'Squat', got '%s'", log.Exercises[0].Name)
	}
}

func TestParseWeightOnly(t *testing.T) {
	input := "squat 135"
	got, err := Parse("test", []byte(input))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	log := got.(*WorkoutLog)
	if len(log.Exercises) != 1 {
		t.Fatalf("expected 1 exercise, got %d", len(log.Exercises))
	}

	ex := log.Exercises[0]
	if len(ex.Sets) != 1 {
		t.Fatalf("expected 1 set, got %d", len(ex.Sets))
	}

	set := ex.Sets[0]
	if set.Reps != 8 {
		t.Errorf("expected 8 reps (default), got %d", set.Reps)
	}
	if set.Weight != 135 {
		t.Errorf("expected weight 135, got %d", set.Weight)
	}
}
