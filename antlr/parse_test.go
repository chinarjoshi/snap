package antlr

import (
	"testing"

	"github.com/chijoshi/workoutliner/parser"
)

func TestParseBasic(t *testing.T) {
	input := "squat 135"
	log, err := Parse(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(log.Exercises) != 1 {
		t.Fatalf("expected 1 exercise, got %d", len(log.Exercises))
	}

	ex := log.Exercises[0]
	if ex.Name != "Squat" {
		t.Errorf("expected name 'Squat', got '%s'", ex.Name)
	}

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
	log, err := Parse(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

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

func TestParseMultiplier(t *testing.T) {
	cases := []struct {
		input    string
		expected []parser.Set
	}{
		{"squat 3x8 135", []parser.Set{{Reps: 8, Weight: 135}, {Reps: 8, Weight: 135}, {Reps: 8, Weight: 135}}},
		{"squat 3x8x135", []parser.Set{{Reps: 8, Weight: 135}, {Reps: 8, Weight: 135}, {Reps: 8, Weight: 135}}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			log, err := Parse(tc.input)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

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
		expected []parser.Set
	}{
		{"press 8 by 35", []parser.Set{{Reps: 8, Weight: 35}, {Reps: 8, Weight: 35}, {Reps: 8, Weight: 35}}},
		{"press 2 by 8 by 35", []parser.Set{{Reps: 8, Weight: 35}, {Reps: 8, Weight: 35}}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			log, err := Parse(tc.input)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

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
	input := "squat 8 135 hard"
	log, err := Parse(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	ex := log.Exercises[0]
	if len(ex.Sets) != 1 {
		t.Fatalf("expected 1 set, got %d", len(ex.Sets))
	}

	if ex.Sets[0].Note != "hard" {
		t.Errorf("expected note 'hard', got '%s'", ex.Sets[0].Note)
	}
}

func TestParseMultipleExercises(t *testing.T) {
	input := "squat 135. bench 185"
	log, err := Parse(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(log.Exercises) != 2 {
		t.Fatalf("expected 2 exercises, got %d", len(log.Exercises))
	}

	if log.Exercises[0].Name != "Squat" {
		t.Errorf("expected 'Squat', got '%s'", log.Exercises[0].Name)
	}
	if log.Exercises[1].Name != "Bench" {
		t.Errorf("expected 'Bench', got '%s'", log.Exercises[1].Name)
	}
}

func TestParseTrailingReps(t *testing.T) {
	input := "leg press 8 315 7 250 7"
	log, err := Parse(input)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	ex := log.Exercises[0]
	expected := []parser.Set{
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
