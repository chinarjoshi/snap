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

func TestTransformTableMerge(t *testing.T) {
	input := "|      | 0     | 1     | 2     |\n|-------|-------|-------|-------|\n| Squat| 8@225 | 8@225 | 8@225 |\nbench 3x8 155"
	result := transform(input)

	if !strings.Contains(result, "| Squat") {
		t.Error("Expected Squat in merged table")
	}
	if !strings.Contains(result, "| Bench") {
		t.Error("Expected Bench in merged table")
	}
	if !strings.Contains(result, "8@225") {
		t.Error("Expected 8@225 in merged table")
	}
	if !strings.Contains(result, "8@155") {
		t.Error("Expected 8@155 in merged table")
	}
}

func TestTransformBodyweight(t *testing.T) {
	input := "dips 8 8 8"
	result := transform(input)

	if !strings.Contains(result, "| Dips") {
		t.Error("Expected Dips in table")
	}
	// Should show just "8" not "8@0"
	if strings.Contains(result, "@") {
		t.Errorf("Bodyweight should not have @ symbol, got: %s", result)
	}
	count := strings.Count(result, "| 8")
	if count < 3 {
		t.Errorf("Expected at least 3 sets of 8, got: %s", result)
	}
}

func TestTransformSupersetBasic(t *testing.T) {
	input := "dips* press* 8 135 8 135 8 135"
	result := transform(input)

	if !strings.Contains(result, "| Dips*") {
		t.Errorf("Expected Dips* in table, got: %s", result)
	}
	if !strings.Contains(result, "| Press*") {
		t.Errorf("Expected Press* in table, got: %s", result)
	}
	// Dips gets odd-indexed values (8, 8, 8 - bodyweight)
	// Press gets even-indexed values (135, 135, 135 - weights)
}

func TestTransformSupersetMixedBodyweightWeighted(t *testing.T) {
	input := "dips* shoulder press* 8 8x135 8 8x125 8 8x100"
	result := transform(input)

	if !strings.Contains(result, "| Dips*") {
		t.Errorf("Expected Dips* in table, got: %s", result)
	}
	if !strings.Contains(result, "| Shoulder Press*") {
		t.Errorf("Expected Shoulder Press* in table, got: %s", result)
	}
	// Dips should have bodyweight (no @)
	// Shoulder Press should have weights
	if !strings.Contains(result, "8@135") {
		t.Errorf("Expected 8@135 for press, got: %s", result)
	}
	if !strings.Contains(result, "8@125") {
		t.Errorf("Expected 8@125 for press, got: %s", result)
	}
	if !strings.Contains(result, "8@100") {
		t.Errorf("Expected 8@100 for press, got: %s", result)
	}
}

func TestTransformSupersetOddSets(t *testing.T) {
	input := "dips* press* 8 135 8 135 8"
	result := transform(input)

	// With 5 sets: indices 0,2,4 go to dips (3 sets), indices 1,3 go to press (2 sets)
	if !strings.Contains(result, "| Dips*") {
		t.Errorf("Expected Dips* in table, got: %s", result)
	}
	if !strings.Contains(result, "| Press*") {
		t.Errorf("Expected Press* in table, got: %s", result)
	}
}

func TestTransformMultiLineBasic(t *testing.T) {
	input := "Bench\n8x135\n8\n8\n8"
	result := transform(input)

	if !strings.Contains(result, "| Bench") {
		t.Errorf("Expected Bench in table, got: %s", result)
	}
	// Should have 4 sets of 8@135 (first set explicit, next 3 inherit weight)
	count := strings.Count(result, "8@135")
	if count != 4 {
		t.Errorf("Expected 4 sets of 8@135, got %d in: %s", count, result)
	}
}

func TestTransformMultiLineInlinePlusContinuation(t *testing.T) {
	input := "shoulder press 7x100\n6\n5"
	result := transform(input)

	if !strings.Contains(result, "| Shoulder Press") {
		t.Errorf("Expected Shoulder Press in table, got: %s", result)
	}
	if !strings.Contains(result, "7@100") {
		t.Errorf("Expected 7@100, got: %s", result)
	}
	if !strings.Contains(result, "6@100") {
		t.Errorf("Expected 6@100, got: %s", result)
	}
	if !strings.Contains(result, "5@100") {
		t.Errorf("Expected 5@100, got: %s", result)
	}
}

func TestTransformMultiLineBodyweight(t *testing.T) {
	input := "Dips\n8\n8\n8"
	result := transform(input)

	if !strings.Contains(result, "| Dips") {
		t.Errorf("Expected Dips in table, got: %s", result)
	}
	// Bodyweight - no @ symbol
	if strings.Contains(result, "@") {
		t.Errorf("Expected bodyweight (no @), got: %s", result)
	}
}

func TestTransformMultiLineMixedExercises(t *testing.T) {
	input := "Bench\n8x135\nshoulder press 7x100\n6"
	result := transform(input)

	if !strings.Contains(result, "| Bench") {
		t.Errorf("Expected Bench in table, got: %s", result)
	}
	if !strings.Contains(result, "| Shoulder Press") {
		t.Errorf("Expected Shoulder Press in table, got: %s", result)
	}
	if !strings.Contains(result, "8@135") {
		t.Errorf("Expected 8@135, got: %s", result)
	}
	if !strings.Contains(result, "7@100") {
		t.Errorf("Expected 7@100, got: %s", result)
	}
	if !strings.Contains(result, "6@100") {
		t.Errorf("Expected 6@100, got: %s", result)
	}
}
