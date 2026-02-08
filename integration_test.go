package main

import (
	"strings"
	"testing"

	"github.com/chijoshi/workoutliner/transform"
)

func TestFullWorkflow(t *testing.T) {
	input := "squat 8 8 8 135\nbench 3x8 95\npress 8 by 35"

	output := transform.Transform(input)

	// Verify exercises in table
	if !strings.Contains(output, "Squat") {
		t.Error("table missing Squat")
	}
	if !strings.Contains(output, "Bench") {
		t.Error("table missing Bench")
	}
	if !strings.Contains(output, "Press") {
		t.Error("table missing Press")
	}

	// Verify table format
	if !strings.Contains(output, "|") {
		t.Error("output should be a table")
	}
	if !strings.Contains(output, "8@135") {
		t.Error("table missing set notation")
	}
}
