package main

import (
	"strings"
	"testing"
)

func TestTransformPureWorkout(t *testing.T) {
	input := "squat 8 8 8 135\nbench 3x8 95"
	result := transform(input)

	if !strings.Contains(result, "| Squat") {
		t.Error("Expected Squat in table")
	}
	if !strings.Contains(result, "| Bench") {
		t.Error("Expected Bench in table")
	}
	if !strings.Contains(result, "8@135") {
		t.Error("Expected 8@135 in table")
	}
	if !strings.Contains(result, "8@95") {
		t.Error("Expected 8@95 in table")
	}
}

func TestTransformProseBeforeWorkout(t *testing.T) {
	input := "ate a banana beforehand\nsquat 8 8 8 135"
	result := transform(input)

	if !strings.HasPrefix(result, "ate a banana beforehand") {
		t.Errorf("Expected prose at start, got: %s", result)
	}
	if !strings.Contains(result, "| Squat") {
		t.Error("Expected Squat in table")
	}
}

func TestTransformProseAfterWorkout(t *testing.T) {
	input := "squat 8 8 8 135\n\nRegular notes here."
	result := transform(input)

	if !strings.Contains(result, "| Squat") {
		t.Error("Expected Squat in table")
	}
	if !strings.HasSuffix(result, "Regular notes here.") {
		t.Errorf("Expected prose at end, got: %s", result)
	}
}

func TestTransformMultipleProseLines(t *testing.T) {
	input := "morning workout\nfelt strong\nsquat 8 8 8 135"
	result := transform(input)

	if !strings.Contains(result, "morning workout") {
		t.Error("Expected first prose line")
	}
	if !strings.Contains(result, "felt strong") {
		t.Error("Expected second prose line")
	}
	if !strings.Contains(result, "| Squat") {
		t.Error("Expected Squat in table")
	}
}

func TestTransformInlineNotes(t *testing.T) {
	input := "squat 8 8 8 135 light headed\nbench 3x8 95"
	result := transform(input)

	if !strings.Contains(result, "- Squat :: light headed") {
		t.Errorf("Expected note as description list, got: %s", result)
	}
}

func TestTransformPureProse(t *testing.T) {
	input := "Just some notes today.\nNothing special."
	result := transform(input)

	if result != input {
		t.Errorf("Expected pure prose unchanged, got: %s", result)
	}
}

func TestTransformMixedParagraphs(t *testing.T) {
	input := "squat 8 8 8 135\n\nRegular notes here.\n\nbench 3x8 95"
	result := transform(input)

	parts := strings.Split(result, "\n\n")
	if len(parts) != 3 {
		t.Errorf("Expected 3 paragraphs, got %d", len(parts))
	}

	if !strings.Contains(parts[0], "| Squat") {
		t.Error("Expected first paragraph to be squat table")
	}
	if parts[1] != "Regular notes here." {
		t.Errorf("Expected middle paragraph to be prose, got: %s", parts[1])
	}
	if !strings.Contains(parts[2], "| Bench") {
		t.Error("Expected third paragraph to be bench table")
	}
}

func TestTransformMultiplierSyntax(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"squat 3x8 135", "8@135"},
		{"squat 3x8x135", "8@135"},
		{"squat 8 by 135", "8@135"},
		{"squat 3 by 8 by 135", "8@135"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := transform(tt.input)
			if !strings.Contains(result, tt.expected) {
				t.Errorf("Expected %s in result, got: %s", tt.expected, result)
			}
		})
	}
}

func TestTransformEmptyInput(t *testing.T) {
	result := transform("")
	if result != "" {
		t.Errorf("Expected empty result, got: %s", result)
	}
}

func TestTransformNoDigitsParagraphSkipped(t *testing.T) {
	input := "Just words here\nno numbers at all"
	result := transform(input)

	if result != input {
		t.Errorf("Expected unchanged input, got: %s", result)
	}
}

func TestTransformTitleCase(t *testing.T) {
	input := "BARBELL SQUAT 8 8 8 135"
	result := transform(input)

	if !strings.Contains(result, "Barbell Squat") {
		t.Errorf("Expected title-cased exercise name, got: %s", result)
	}
}

func TestTransformTrailingReps(t *testing.T) {
	input := "squat 8 8 8 135 8 8"
	result := transform(input)

	count := strings.Count(result, "8@135")
	if count != 5 {
		t.Errorf("Expected 5 sets of 8@135, got %d", count)
	}
}

func TestTransformDefaultSetsForBySyntax(t *testing.T) {
	input := "squat 8 by 135"
	result := transform(input)

	count := strings.Count(result, "8@135")
	if count != 3 {
		t.Errorf("Expected 3 sets (default) of 8@135, got %d", count)
	}
}

func TestTransformDefaultRepsForWeight(t *testing.T) {
	input := "squat 135"
	result := transform(input)

	if !strings.Contains(result, "8@135") {
		t.Errorf("Expected default 8 reps, got: %s", result)
	}
}

func TestTransformMultipleExercises(t *testing.T) {
	input := "squat 8 8 8 135\nbench 3x8 95\npress 8 by 35"
	result := transform(input)

	if !strings.Contains(result, "| Squat") {
		t.Error("Expected Squat")
	}
	if !strings.Contains(result, "| Bench") {
		t.Error("Expected Bench")
	}
	if !strings.Contains(result, "| Press") {
		t.Error("Expected Press")
	}
}

func TestTransformPreservesBlankLines(t *testing.T) {
	input := "paragraph one\n\nparagraph two"
	result := transform(input)

	if !strings.Contains(result, "\n\n") {
		t.Error("Expected blank line between paragraphs")
	}
}

func TestTransformTableColumnNumbers(t *testing.T) {
	input := "squat 8 8 135"
	result := transform(input)

	if !strings.Contains(result, "| 0") {
		t.Error("Expected column 0 in header")
	}
	if !strings.Contains(result, "| 1") {
		t.Error("Expected column 1 in header")
	}
}
