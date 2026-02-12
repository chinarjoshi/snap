package workout

import (
	"testing"
)

func TestParsePureExercise(t *testing.T) {
	result, err := Parse("squat 8 8 8 135")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.ProseLines) != 0 {
		t.Errorf("Expected no prose lines, got %d", len(result.ProseLines))
	}
	if len(result.Exercises) != 1 {
		t.Fatalf("Expected 1 exercise, got %d", len(result.Exercises))
	}
	if result.Exercises[0].Name != "Squat" {
		t.Errorf("Expected 'Squat', got '%s'", result.Exercises[0].Name)
	}
	if len(result.Exercises[0].Sets) != 3 {
		t.Errorf("Expected 3 sets, got %d", len(result.Exercises[0].Sets))
	}
}

func TestParsePureProse(t *testing.T) {
	result, err := Parse("just some words here")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.ProseLines) != 1 {
		t.Errorf("Expected 1 prose line, got %d", len(result.ProseLines))
	}
	if len(result.Exercises) != 0 {
		t.Errorf("Expected no exercises, got %d", len(result.Exercises))
	}
}

func TestParseMixedLines(t *testing.T) {
	result, err := Parse("warmup notes\nsquat 8 8 8 135\nbench 3x8 95")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.ProseLines) != 1 {
		t.Errorf("Expected 1 prose line, got %d", len(result.ProseLines))
	}
	if result.ProseLines[0] != "warmup notes" {
		t.Errorf("Expected 'warmup notes', got '%s'", result.ProseLines[0])
	}
	if len(result.Exercises) != 2 {
		t.Errorf("Expected 2 exercises, got %d", len(result.Exercises))
	}
}

func TestParseInlineNote(t *testing.T) {
	result, err := Parse("squat 8 8 8 135 light headed")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.Exercises) != 1 {
		t.Fatalf("Expected 1 exercise, got %d", len(result.Exercises))
	}

	sets := result.Exercises[0].Sets
	lastSet := sets[len(sets)-1]
	if lastSet.Note != "light headed" {
		t.Errorf("Expected note 'light headed', got '%s'", lastSet.Note)
	}
}

func TestParseMultiplierSyntax(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		numSets  int
		reps     int
		weight   int
	}{
		{"partial 3x8", "squat 3x8 135", 3, 8, 135},
		{"full 3x8x135", "squat 3x8x135", 3, 8, 135},
		{"two part by", "squat 8 by 135", 3, 8, 135},
		{"three part by", "squat 2 by 8 by 135", 2, 8, 135},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}

			if len(result.Exercises) != 1 {
				t.Fatalf("Expected 1 exercise, got %d", len(result.Exercises))
			}

			sets := result.Exercises[0].Sets
			if len(sets) != tt.numSets {
				t.Errorf("Expected %d sets, got %d", tt.numSets, len(sets))
			}
			if sets[0].Reps != tt.reps {
				t.Errorf("Expected %d reps, got %d", tt.reps, sets[0].Reps)
			}
			if sets[0].Weight != tt.weight {
				t.Errorf("Expected %d weight, got %d", tt.weight, sets[0].Weight)
			}
		})
	}
}

func TestParseTrailingReps(t *testing.T) {
	result, err := Parse("squat 135 8 8")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.Exercises) != 1 {
		t.Fatalf("Expected 1 exercise, got %d", len(result.Exercises))
	}

	sets := result.Exercises[0].Sets
	if len(sets) != 3 {
		t.Errorf("Expected 3 sets, got %d", len(sets))
	}

	for i, s := range sets {
		if s.Weight != 135 {
			t.Errorf("Set %d: expected weight 135, got %d", i, s.Weight)
		}
	}
}

func TestParseMultipleExercises(t *testing.T) {
	result, err := Parse("squat 8 8 8 135\nbench 3x8 95\npress 8 by 35")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.Exercises) != 3 {
		t.Fatalf("Expected 3 exercises, got %d", len(result.Exercises))
	}

	names := []string{"Squat", "Bench", "Press"}
	for i, name := range names {
		if result.Exercises[i].Name != name {
			t.Errorf("Exercise %d: expected '%s', got '%s'", i, name, result.Exercises[i].Name)
		}
	}
}

func TestParseMultiWordExerciseName(t *testing.T) {
	result, err := Parse("barbell squat 8 8 8 135")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(result.Exercises) != 1 {
		t.Fatalf("Expected 1 exercise, got %d", len(result.Exercises))
	}

	if result.Exercises[0].Name != "Barbell Squat" {
		t.Errorf("Expected 'Barbell Squat', got '%s'", result.Exercises[0].Name)
	}
}

func TestParseHasExercises(t *testing.T) {
	proseResult, _ := Parse("just words")
	if proseResult.HasExercises() {
		t.Error("Expected HasExercises to be false for prose")
	}

	exerciseResult, _ := Parse("squat 8 8 8 135")
	if !exerciseResult.HasExercises() {
		t.Error("Expected HasExercises to be true for exercise")
	}
}
