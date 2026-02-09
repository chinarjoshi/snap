# Supersets & Bodyweight Support

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Support bodyweight exercises (no weight specified) and supersets (two exercises with alternating sets).

**Architecture:** Extend grammar with `*` token for superset markers, modify visitor to handle bodyweight (weight=0) and superset alternation, update formatters to display correctly.

**Tech Stack:** ANTLR4, Go, Swift

---

## Feature 1: Bodyweight Support

### Input/Output

```
dips 8 8 8
```
→
```
|     | 0 | 1 | 2 |
|-----|---|---|---|
| Dips| 8 | 8 | 8 |
```

### Changes

**Visitor (Go & Swift):** In `buildSetsFromResults`, always add pending reps even if no weight:
```go
// Before
if lastWeight > 0 {
    sets = append(sets, Set{Reps: r, Weight: lastWeight})
}

// After
sets = append(sets, Set{Reps: r, Weight: lastWeight})
```

**Formatter (Go & Swift):** In `formatSet`, check for zero weight:
```go
func formatSet(s Set) string {
    if s.Weight == 0 {
        return fmt.Sprintf("%d", s.Reps)
    }
    return fmt.Sprintf("%d@%d", s.Reps, s.Weight)
}
```

---

## Feature 2: Supersets

### Input/Output

```
dips* shoulder press* 8 8x135 8 8x125 8 8x100
```
→
```
|                 | 0     | 1     | 2     |
|-----------------|-------|-------|-------|
| Dips*           | 8     | 8     | 8     |
| Shoulder Press* | 8@135 | 8@125 | 8@100 |
```

### Grammar Changes (Workoutliner.g4)

```antlr
line
    : exerciseLine NEWLINE?   # ExerciseLineAlt
    | supersetLine NEWLINE?   # SupersetLineAlt
    | proseLine NEWLINE?      # ProseLineAlt
    ;

supersetLine : supersetName supersetName token+ ;

supersetName : WORD* STAR ;

STAR : '*' ;
```

### Visitor Changes

**New `VisitSupersetLineAlt` and `VisitSupersetLine`:**
1. Extract two exercise names from `supersetName` contexts (including `*` suffix)
2. Collect all sets from tokens using existing `buildSetsFromResults`
3. Alternate: even indices (0,2,4...) → first exercise, odd indices (1,3,5...) → second exercise
4. Return two `Exercise` structs

**In `VisitParagraph`:** Handle when visitor returns `[]Exercise` (slice) vs single `Exercise`.

---

## Tasks

### Task 1: Bodyweight Support - Go

**Files:**
- Modify: `libs/workoutliner/visitor_impl.go`
- Modify: `cmd/workoutliner/main.go`
- Modify: `cmd/workoutliner/main_test.go`

**Steps:**
1. In `buildSetsFromResults`, remove `if lastWeight > 0` check
2. In `formatSet`, return just reps if weight is 0
3. Add test: `dips 8 8 8` → displays `8 | 8 | 8` (no @)

### Task 2: Bodyweight Support - Swift

**Files:**
- Modify: `swift/Sources/WorkoutLiner/ParagraphASTVisitor.swift`
- Modify: `swift/Sources/WorkoutLiner/WorkoutLiner.swift`
- Modify: `swift/Tests/WorkoutLinerTests/WorkoutLinerTests.swift`

**Steps:**
1. Same logic changes as Go
2. Add matching test

### Task 3: Superset Grammar

**Files:**
- Modify: `Workoutliner.g4`

**Steps:**
1. Add `STAR : '*' ;` lexer rule
2. Add `supersetName : WORD* STAR ;` parser rule
3. Add `supersetLine : supersetName supersetName token+ ;`
4. Add `supersetLine` alternative to `line` rule
5. Regenerate parsers: `make antlr`

### Task 4: Superset Visitor - Go

**Files:**
- Modify: `libs/workoutliner/visitor_impl.go`

**Steps:**
1. Add `VisitSupersetLineAlt` - delegates to `VisitSupersetLine`
2. Add `VisitSupersetLine`:
   - Extract names from two `supersetName` contexts
   - Build sets from tokens
   - Alternate assignment between two exercises
   - Return `[]Exercise` with both
3. Update `VisitParagraph` to handle `[]Exercise` return type

### Task 5: Superset Visitor - Swift

**Files:**
- Modify: `swift/Sources/WorkoutLiner/ParagraphASTVisitor.swift`

**Steps:**
1. Same logic as Go implementation

### Task 6: Superset Tests

**Files:**
- Modify: `cmd/workoutliner/main_test.go`
- Modify: `swift/Tests/WorkoutLinerTests/WorkoutLinerTests.swift`

**Steps:**
1. Test basic superset: `dips* press* 8 135 8 135` → both exercises in table with `*`
2. Test mixed bodyweight/weighted superset
3. Test odd number of sets (last goes to first exercise)

---

## Edge Cases

- Odd sets in superset → last set goes to first exercise
- All bodyweight superset → both show just reps
- Notes in superset → attach to last set (existing behavior)
