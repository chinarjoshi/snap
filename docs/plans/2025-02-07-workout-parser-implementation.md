# Workout Log Parser Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build a PEG parser that converts natural language workout logs into structured tables.

**Architecture:** Use pigeon to generate a Go parser from a PEG grammar. The parser produces an AST (Set, Exercise, WorkoutLog types) which is then formatted as a markdown table for output.

**Tech Stack:** Go, pigeon (PEG parser generator), standard library only for table output.

---

## Task 1: Project Setup

**Files:**
- Create: `go.mod`
- Create: `parser/ast.go`

**Step 1: Initialize Go module**

Run:
```bash
go mod init github.com/chijoshi/workoutliner
```
Expected: Creates `go.mod` file

**Step 2: Install pigeon**

Run:
```bash
go install github.com/mna/pigeon@latest
```
Expected: pigeon binary installed in GOPATH/bin

**Step 3: Create AST types**

Create `parser/ast.go`:

```go
package parser

// Set represents a single set within an exercise
type Set struct {
	Reps   int
	Weight int
	Note   string
}

// Exercise represents a parsed exercise with all its sets
type Exercise struct {
	Name string
	Sets []Set
}

// WorkoutLog represents the complete parsed workout
type WorkoutLog struct {
	Exercises []Exercise
}

// MaxSets returns the maximum number of sets across all exercises
func (w *WorkoutLog) MaxSets() int {
	max := 0
	for _, e := range w.Exercises {
		if len(e.Sets) > max {
			max = len(e.Sets)
		}
	}
	return max
}
```

**Step 4: Verify module compiles**

Run:
```bash
go build ./...
```
Expected: No errors

**Step 5: Commit**

```bash
git add go.mod parser/ast.go
git commit -m "feat: initialize project with AST types"
```

---

## Task 2: Basic PEG Grammar - Exercise Name Only

**Files:**
- Create: `parser/grammar.peg`
- Generate: `parser/grammar.go`
- Create: `parser/parser_test.go`

**Step 1: Write failing test for exercise name parsing**

Create `parser/parser_test.go`:

```go
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
```

**Step 2: Run test to verify it fails**

Run:
```bash
go test ./parser -v
```
Expected: FAIL (Parse function doesn't exist)

**Step 3: Create minimal PEG grammar**

Create `parser/grammar.peg`:

```peg
{
package parser

import (
	"strings"
)

func toIfaceSlice(v any) []any {
	if v == nil {
		return nil
	}
	return v.([]any)
}

func titleCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}
}

Input <- exercises:ExerciseList EOF {
	return exercises, nil
}

ExerciseList <- first:Exercise rest:(Delimiter Exercise)* {
	exs := []Exercise{first.(Exercise)}
	for _, r := range toIfaceSlice(rest) {
		pair := toIfaceSlice(r)
		exs = append(exs, pair[1].(Exercise))
	}
	return &WorkoutLog{Exercises: exs}, nil
}

Delimiter <- [.\n]+

Exercise <- _ name:ExerciseName data:SetData* _ {
	ex := Exercise{Name: titleCase(name.(string))}
	// For now, ignore set data
	return ex, nil
}

ExerciseName <- chars:(!Digit !Delimiter .)+ {
	var sb strings.Builder
	for _, c := range toIfaceSlice(chars) {
		sb.WriteString(string(toIfaceSlice(c)[2].([]byte)))
	}
	return sb.String(), nil
}

SetData <- _ (Number / NoteText) {
	return nil, nil
}

NoteText <- (!Digit !Delimiter .)+

Number <- Digit+ {
	return string(c.text), nil
}

Digit <- [0-9]

_ <- [ \t]*

EOF <- !.
```

**Step 4: Generate parser from grammar**

Run:
```bash
pigeon -o parser/grammar.go parser/grammar.peg
```
Expected: Creates `parser/grammar.go`

**Step 5: Run test to verify it passes**

Run:
```bash
go test ./parser -v
```
Expected: PASS

**Step 6: Commit**

```bash
git add parser/grammar.peg parser/grammar.go parser/parser_test.go
git commit -m "feat: basic PEG grammar parsing exercise names"
```

---

## Task 3: Parse Single Numbers (Weights with Default Reps)

**Files:**
- Modify: `parser/grammar.peg`
- Modify: `parser/parser_test.go`

**Step 1: Write failing test for weight-only input**

Add to `parser/parser_test.go`:

```go
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
```

**Step 2: Run test to verify it fails**

Run:
```bash
go test ./parser -v -run TestParseWeightOnly
```
Expected: FAIL (Sets slice empty)

**Step 3: Update grammar to parse numbers and build sets**

Update `parser/grammar.peg`:

```peg
{
package parser

import (
	"strconv"
	"strings"
)

const (
	defaultReps = 8
	repThreshold = 20
)

func toIfaceSlice(v any) []any {
	if v == nil {
		return nil
	}
	return v.([]any)
}

func titleCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func isWeight(n int) bool {
	return n > repThreshold
}

type token struct {
	kind  string // "reps", "weight", "note"
	value int
	text  string
}

func buildSets(tokens []token) []Set {
	var sets []Set
	var pendingReps []int
	var lastWeight int

	for _, tok := range tokens {
		switch tok.kind {
		case "weight":
			if len(pendingReps) > 0 {
				// Apply all pending reps to this weight
				for _, r := range pendingReps {
					sets = append(sets, Set{Reps: r, Weight: tok.value})
				}
				pendingReps = nil
			} else {
				// No pending reps, use default
				sets = append(sets, Set{Reps: defaultReps, Weight: tok.value})
			}
			lastWeight = tok.value
		case "reps":
			pendingReps = append(pendingReps, tok.value)
		case "note":
			if len(sets) > 0 {
				sets[len(sets)-1].Note = strings.TrimSpace(tok.text)
			}
		}
	}

	// Handle trailing reps (inherit last weight)
	for _, r := range pendingReps {
		if lastWeight > 0 {
			sets = append(sets, Set{Reps: r, Weight: lastWeight})
		}
	}

	return sets
}
}

Input <- exercises:ExerciseList EOF {
	return exercises, nil
}

ExerciseList <- first:Exercise rest:(Delimiter Exercise)* {
	exs := []Exercise{first.(Exercise)}
	for _, r := range toIfaceSlice(rest) {
		pair := toIfaceSlice(r)
		exs = append(exs, pair[1].(Exercise))
	}
	return &WorkoutLog{Exercises: exs}, nil
}

Delimiter <- [.\n]+

Exercise <- _ name:ExerciseName tokens:Token* _ {
	toks := []token{}
	for _, t := range toIfaceSlice(tokens) {
		if t != nil {
			toks = append(toks, t.(token))
		}
	}
	sets := buildSets(toks)
	return Exercise{Name: titleCase(name.(string)), Sets: sets}, nil
}

ExerciseName <- chars:(!Digit !Delimiter .)+ {
	var sb strings.Builder
	for _, c := range toIfaceSlice(chars) {
		sb.WriteString(string(toIfaceSlice(c)[2].([]byte)))
	}
	return sb.String(), nil
}

Token <- _ t:(Number / NoteText) {
	return t, nil
}

NoteText <- chars:(!Digit !Delimiter .)+ {
	var sb strings.Builder
	for _, c := range toIfaceSlice(chars) {
		sb.WriteString(string(toIfaceSlice(c)[2].([]byte)))
	}
	text := sb.String()
	return token{kind: "note", text: text}, nil
}

Number <- digits:Digit+ {
	s := string(c.text)
	n := toInt(s)
	if isWeight(n) {
		return token{kind: "weight", value: n}, nil
	}
	return token{kind: "reps", value: n}, nil
}

Digit <- [0-9]

_ <- [ \t]*

EOF <- !.
```

**Step 4: Regenerate parser**

Run:
```bash
pigeon -o parser/grammar.go parser/grammar.peg
```

**Step 5: Run test to verify it passes**

Run:
```bash
go test ./parser -v -run TestParseWeightOnly
```
Expected: PASS

**Step 6: Commit**

```bash
git add parser/grammar.peg parser/grammar.go parser/parser_test.go
git commit -m "feat: parse weights with default reps"
```

---

## Task 4: Parse Reps + Weight Pairs

**Files:**
- Modify: `parser/parser_test.go`

**Step 1: Write test for reps + weight**

Add to `parser/parser_test.go`:

```go
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
```

**Step 2: Run tests**

Run:
```bash
go test ./parser -v -run "TestParseRepsAndWeight|TestParseMultipleRepsOneWeight"
```
Expected: PASS (grammar already handles this)

**Step 3: Commit**

```bash
git add parser/parser_test.go
git commit -m "test: add reps + weight parsing tests"
```

---

## Task 5: Parse Trailing Reps (Inherit Weight)

**Files:**
- Modify: `parser/parser_test.go`

**Step 1: Write test for trailing reps**

Add to `parser/parser_test.go`:

```go
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
```

**Step 2: Run test**

Run:
```bash
go test ./parser -v -run TestParseTrailingReps
```
Expected: PASS (already handled by buildSets)

**Step 3: Commit**

```bash
git add parser/parser_test.go
git commit -m "test: add trailing reps inheritance test"
```

---

## Task 6: Parse Multiplier Syntax (3x8, 3x8x135)

**Files:**
- Modify: `parser/grammar.peg`
- Modify: `parser/parser_test.go`

**Step 1: Write failing test for multiplier**

Add to `parser/parser_test.go`:

```go
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
```

**Step 2: Run test to verify it fails**

Run:
```bash
go test ./parser -v -run TestParseMultiplier
```
Expected: FAIL

**Step 3: Add multiplier token type**

Update the token struct and add handling in `parser/grammar.peg`. Add this type inside the code block:

```go
type multiplierToken struct {
	sets   int
	reps   int
	weight int // 0 if not specified
}
```

Update buildSets to handle multipliers:

```go
func buildSets(tokens []any) []Set {
	var sets []Set
	var pendingReps []int
	var lastWeight int

	for _, t := range tokens {
		switch tok := t.(type) {
		case multiplierToken:
			if tok.weight > 0 {
				// Full multiplier: 3x8x135
				for i := 0; i < tok.sets; i++ {
					sets = append(sets, Set{Reps: tok.reps, Weight: tok.weight})
				}
			} else {
				// Partial multiplier: 3x8 (needs weight from next token)
				for i := 0; i < tok.sets; i++ {
					pendingReps = append(pendingReps, tok.reps)
				}
			}
		case token:
			switch tok.kind {
			case "weight":
				if len(pendingReps) > 0 {
					for _, r := range pendingReps {
						sets = append(sets, Set{Reps: r, Weight: tok.value})
					}
					pendingReps = nil
				} else {
					sets = append(sets, Set{Reps: defaultReps, Weight: tok.value})
				}
				lastWeight = tok.value
			case "reps":
				pendingReps = append(pendingReps, tok.value)
			case "note":
				if len(sets) > 0 {
					sets[len(sets)-1].Note = strings.TrimSpace(tok.text)
				}
			}
		}
	}

	for _, r := range pendingReps {
		if lastWeight > 0 {
			sets = append(sets, Set{Reps: r, Weight: lastWeight})
		}
	}

	return sets
}
```

Add multiplier rule before Number in Token:

```peg
Token <- _ t:(Multiplier / Number / NoteText) {
	return t, nil
}

Multiplier <- sets:Digits [xX] reps:Digits weight:([xX] Digits)? {
	s := toInt(sets.(string))
	r := toInt(reps.(string))
	w := 0
	if weight != nil {
		wParts := toIfaceSlice(weight)
		w = toInt(wParts[1].(string))
	}
	return multiplierToken{sets: s, reps: r, weight: w}, nil
}

Digits <- [0-9]+ {
	return string(c.text), nil
}
```

**Step 4: Regenerate and run tests**

Run:
```bash
pigeon -o parser/grammar.go parser/grammar.peg && go test ./parser -v -run TestParseMultiplier
```
Expected: PASS

**Step 5: Commit**

```bash
git add parser/grammar.peg parser/grammar.go parser/parser_test.go
git commit -m "feat: parse multiplier syntax (3x8, 3x8x135)"
```

---

## Task 7: Parse "by" Syntax

**Files:**
- Modify: `parser/grammar.peg`
- Modify: `parser/parser_test.go`

**Step 1: Write failing test for by-syntax**

Add to `parser/parser_test.go`:

```go
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
```

**Step 2: Run test to verify it fails**

Run:
```bash
go test ./parser -v -run TestParseBySyntax
```
Expected: FAIL

**Step 3: Add by-expression rule**

Add byToken type:

```go
type byToken struct {
	sets   int
	reps   int
	weight int
}
```

Add handling in buildSets for byToken:

```go
case byToken:
	for i := 0; i < tok.sets; i++ {
		sets = append(sets, Set{Reps: tok.reps, Weight: tok.weight})
	}
	lastWeight = tok.weight
```

Add ByExpression rule (before Multiplier in Token):

```peg
Token <- _ t:(ByExpression / Multiplier / Number / NoteText) {
	return t, nil
}

ByExpression <- first:Digits _ "by"i _ second:Digits third:(_ "by"i _ Digits)? {
	f := toInt(first.(string))
	s := toInt(second.(string))

	if third != nil {
		// S by R by W format
		parts := toIfaceSlice(third)
		w := toInt(parts[3].(string))
		return byToken{sets: f, reps: s, weight: w}, nil
	}

	// R by W format (default 3 sets)
	return byToken{sets: 3, reps: f, weight: s}, nil
}
```

**Step 4: Regenerate and run tests**

Run:
```bash
pigeon -o parser/grammar.go parser/grammar.peg && go test ./parser -v -run TestParseBySyntax
```
Expected: PASS

**Step 5: Commit**

```bash
git add parser/grammar.peg parser/grammar.go parser/parser_test.go
git commit -m "feat: parse by-syntax (8 by 35, 2 by 8 by 35)"
```

---

## Task 8: Parse Notes

**Files:**
- Modify: `parser/parser_test.go`

**Step 1: Write test for notes**

Add to `parser/parser_test.go`:

```go
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
```

**Step 2: Run test**

Run:
```bash
go test ./parser -v -run TestParseNotes
```
Expected: PASS (already handled)

**Step 3: Commit**

```bash
git add parser/parser_test.go
git commit -m "test: add notes parsing tests"
```

---

## Task 9: Parse Multiple Exercises

**Files:**
- Modify: `parser/parser_test.go`

**Step 1: Write test for multiple exercises**

Add to `parser/parser_test.go`:

```go
func TestParseMultipleExercises(t *testing.T) {
	cases := []struct {
		input         string
		expectedCount int
		expectedNames []string
	}{
		{"squat 135. bench 185", 2, []string{"Squat", "Bench"}},
		{"squat 135\nbench 185", 2, []string{"Squat", "Bench"}},
		{"squat 135. bench 185. deadlift 225", 3, []string{"Squat", "Bench", "Deadlift"}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := Parse("test", []byte(tc.input))
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			log := got.(*WorkoutLog)

			if len(log.Exercises) != tc.expectedCount {
				t.Fatalf("expected %d exercises, got %d", tc.expectedCount, len(log.Exercises))
			}

			for i, name := range tc.expectedNames {
				if log.Exercises[i].Name != name {
					t.Errorf("exercise %d: expected '%s', got '%s'", i, name, log.Exercises[i].Name)
				}
			}
		})
	}
}
```

**Step 2: Run test**

Run:
```bash
go test ./parser -v -run TestParseMultipleExercises
```
Expected: PASS

**Step 3: Commit**

```bash
git add parser/parser_test.go
git commit -m "test: add multiple exercises parsing tests"
```

---

## Task 10: Table Output

**Files:**
- Create: `table/table.go`
- Create: `table/table_test.go`

**Step 1: Write failing test for table output**

Create `table/table_test.go`:

```go
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
	if !strings.Contains(output, "8 @ 135") {
		t.Error("expected '8 @ 135' in output")
	}
}
```

**Step 2: Run test to verify it fails**

Run:
```bash
go test ./table -v
```
Expected: FAIL (table package doesn't exist)

**Step 3: Implement table formatter**

Create `table/table.go`:

```go
package table

import (
	"fmt"
	"strings"

	"github.com/chijoshi/workoutliner/parser"
)

// Format returns a markdown table representation of the workout log
func Format(log *parser.WorkoutLog) string {
	if len(log.Exercises) == 0 {
		return ""
	}

	maxSets := log.MaxSets()
	if maxSets == 0 {
		maxSets = 1
	}

	var sb strings.Builder

	// Header row
	sb.WriteString("|")
	sb.WriteString(pad("", maxExerciseNameLen(log)))
	sb.WriteString(" |")
	for i := 1; i <= maxSets; i++ {
		sb.WriteString(fmt.Sprintf(" Set %d |", i))
	}
	sb.WriteString("\n")

	// Separator row
	sb.WriteString("|")
	sb.WriteString(strings.Repeat("-", maxExerciseNameLen(log)+1))
	sb.WriteString("|")
	for i := 0; i < maxSets; i++ {
		sb.WriteString("---------|")
	}
	sb.WriteString("\n")

	// Data rows
	for _, ex := range log.Exercises {
		sb.WriteString("| ")
		sb.WriteString(padRight(ex.Name, maxExerciseNameLen(log)))
		sb.WriteString("|")
		for i := 0; i < maxSets; i++ {
			if i < len(ex.Sets) {
				sb.WriteString(" ")
				sb.WriteString(formatSet(ex.Sets[i]))
				sb.WriteString(" |")
			} else {
				sb.WriteString("         |")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func formatSet(s parser.Set) string {
	result := fmt.Sprintf("%d @ %d", s.Reps, s.Weight)
	if s.Note != "" {
		result += fmt.Sprintf(" (%s)", s.Note)
	}
	return result
}

func maxExerciseNameLen(log *parser.WorkoutLog) int {
	max := 0
	for _, ex := range log.Exercises {
		if len(ex.Name) > max {
			max = len(ex.Name)
		}
	}
	if max < 8 {
		max = 8
	}
	return max
}

func pad(s string, length int) string {
	return fmt.Sprintf("%*s", length, s)
}

func padRight(s string, length int) string {
	return fmt.Sprintf("%-*s", length, s)
}
```

**Step 4: Run test to verify it passes**

Run:
```bash
go test ./table -v
```
Expected: PASS

**Step 5: Commit**

```bash
git add table/table.go table/table_test.go
git commit -m "feat: add table output formatting"
```

---

## Task 11: CLI Entry Point

**Files:**
- Create: `cmd/workoutliner/main.go`

**Step 1: Create CLI**

Create `cmd/workoutliner/main.go`:

```go
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chijoshi/workoutliner/parser"
	"github.com/chijoshi/workoutliner/table"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	result, err := parser.Parse("input", input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	log := result.(*parser.WorkoutLog)
	fmt.Print(table.Format(log))
}

func readInput() ([]byte, error) {
	if len(os.Args) > 1 {
		return os.ReadFile(os.Args[1])
	}
	return io.ReadAll(os.Stdin)
}
```

**Step 2: Build and test manually**

Run:
```bash
go build -o workoutliner ./cmd/workoutliner
echo "squat 8 8 8 135. bench 5 5 5 185" | ./workoutliner
```
Expected: Table output

**Step 3: Commit**

```bash
git add cmd/workoutliner/main.go
git commit -m "feat: add CLI entry point"
```

---

## Task 12: Integration Test

**Files:**
- Create: `integration_test.go`

**Step 1: Write integration test**

Create `integration_test.go`:

```go
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
```

**Step 2: Run integration test**

Run:
```bash
go test -v -run TestFullWorkflow
```
Expected: PASS

**Step 3: Run all tests**

Run:
```bash
go test ./... -v
```
Expected: All PASS

**Step 4: Commit**

```bash
git add integration_test.go
git commit -m "test: add integration test"
```

---

## Summary

| Task | Description |
|------|-------------|
| 1 | Project setup, AST types |
| 2 | Basic PEG grammar, exercise name parsing |
| 3 | Parse weights with default reps |
| 4 | Parse reps + weight pairs |
| 5 | Parse trailing reps (inherit weight) |
| 6 | Parse multiplier syntax (3x8, 3x8x135) |
| 7 | Parse "by" syntax |
| 8 | Parse notes |
| 9 | Parse multiple exercises |
| 10 | Table output |
| 11 | CLI entry point |
| 12 | Integration test |
