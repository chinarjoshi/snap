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

func TestParseRepsAndWeight(t *testing.T) {
	input := "squat 8 135"
	got, err := Parse("test", []byte(input))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	log := got.(*WorkoutLog)
	ex := log.Exercises[0]

	if len(ex.Sets) != 1 {
		t.Fatalf("expected 1 set, got %d", len(ex.Sets))
	}

	set := ex.Sets[0]
	if set.Reps != 8 || set.Weight != 135 {
		t.Errorf("expected 8@135, got %d@%d", set.Reps, set.Weight)
	}
}

func TestParseMultipleRepsOneWeight(t *testing.T) {
	input := "squat 8 8 8 135"
	got, err := Parse("test", []byte(input))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	log := got.(*WorkoutLog)
	ex := log.Exercises[0]

	if len(ex.Sets) != 3 {
		t.Fatalf("expected 3 sets, got %d", len(ex.Sets))
	}

	for i, set := range ex.Sets {
		if set.Reps != 8 || set.Weight != 135 {
			t.Errorf("set %d: expected 8@135, got %d@%d", i, set.Reps, set.Weight)
		}
	}
}

func TestParseTrailingReps(t *testing.T) {
	input := "leg press 8 315 7 250 7"
	got, err := Parse("test", []byte(input))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	log := got.(*WorkoutLog)
	ex := log.Exercises[0]

	expected := []Set{
		{Reps: 8, Weight: 315},
		{Reps: 7, Weight: 250},
		{Reps: 7, Weight: 250},
	}

	if len(ex.Sets) != len(expected) {
		t.Fatalf("expected %d sets, got %d", len(expected), len(ex.Sets))
	}

	for i, want := range expected {
		got := ex.Sets[i]
		if got.Reps != want.Reps || got.Weight != want.Weight {
			t.Errorf("set %d: expected %d@%d, got %d@%d", i, want.Reps, want.Weight, got.Reps, got.Weight)
		}
	}
}

func TestParseMultiplier(t *testing.T) {
	cases := []struct {
		input    string
		expected []Set
	}{
		{"squat 3x8 135", []Set{{8, 135, ""}, {8, 135, ""}, {8, 135, ""}}},
		{"squat 3x8x135", []Set{{8, 135, ""}, {8, 135, ""}, {8, 135, ""}}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Parse("test", []byte(tc.input))
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			log := got.(*WorkoutLog)
			ex := log.Exercises[0]

			if len(ex.Sets) != len(tc.expected) {
				t.Fatalf("expected %d sets, got %d", len(tc.expected), len(ex.Sets))
			}

			for i, want := range tc.expected {
				got := ex.Sets[i]
				if got.Reps != want.Reps || got.Weight != want.Weight {
					t.Errorf("set %d: expected %d@%d, got %d@%d", i, want.Reps, want.Weight, got.Reps, got.Weight)
				}
			}
		})
	}
}

func TestParseBySyntax(t *testing.T) {
	cases := []struct {
		input    string
		expected []Set
	}{
		{"press 8 by 35", []Set{{8, 35, ""}, {8, 35, ""}, {8, 35, ""}}},
		{"press 2 by 8 by 35", []Set{{8, 35, ""}, {8, 35, ""}}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Parse("test", []byte(tc.input))
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			log := got.(*WorkoutLog)
			ex := log.Exercises[0]

			if len(ex.Sets) != len(tc.expected) {
				t.Fatalf("expected %d sets, got %d", len(tc.expected), len(ex.Sets))
			}

			for i, want := range tc.expected {
				got := ex.Sets[i]
				if got.Reps != want.Reps || got.Weight != want.Weight {
					t.Errorf("set %d: expected %d@%d, got %d@%d", i, want.Reps, want.Weight, got.Reps, got.Weight)
				}
			}
		})
	}
}

func TestParseNotes(t *testing.T) {
	cases := []struct {
		input    string
		expected []Set
	}{
		{"squat 8 135 hard", []Set{{8, 135, "hard"}}},
		{"bench 135 tough 125", []Set{{8, 135, "tough"}, {8, 125, ""}}},
		{"squat 8 8 8 135 light headed", []Set{{8, 135, ""}, {8, 135, ""}, {8, 135, "light headed"}}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Parse("test", []byte(tc.input))
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			log := got.(*WorkoutLog)
			ex := log.Exercises[0]

			if len(ex.Sets) != len(tc.expected) {
				t.Fatalf("expected %d sets, got %d", len(tc.expected), len(ex.Sets))
			}

			for i, want := range tc.expected {
				got := ex.Sets[i]
				if got.Reps != want.Reps || got.Weight != want.Weight || got.Note != want.Note {
					t.Errorf("set %d: expected %d@%d(%s), got %d@%d(%s)",
						i, want.Reps, want.Weight, want.Note, got.Reps, got.Weight, got.Note)
				}
			}
		})
	}
}
