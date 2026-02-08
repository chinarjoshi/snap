# Workout Log Transformer Design

## Overview

A Go CLI and Swift library that takes plaintext notes files, detects workout paragraphs, and replaces them in-place with formatted org-mode tables.

**Go:** CLI program for Emacs subprocess integration
**Swift:** Library for iOS app embedding

## Data Flow

1. Split input by `\n\n` into paragraphs
2. For each paragraph:
   - No digits → pass through unchanged
   - Has digits → parse with grammar
3. For parsed paragraphs:
   - Lines with numbers → exercises for table
   - Lines without numbers → prose above table
4. Reassemble paragraphs with `\n\n`
5. Return transformed text

## Grammar

Parses a single paragraph:

```antlr
paragraph    : line+ EOF ;

line
    : exerciseLine    # ExerciseLine
    | proseLine       # ProseLine
    ;

exerciseLine : WORD+ token+ NEWLINE? ;
proseLine    : WORD+ NEWLINE? ;

token
    : byExpr
    | multiplier
    | number
    | note
    ;

byExpr
    : reps=NUMBER BY weight=NUMBER                # TwoPartBy
    | sets=NUMBER BY reps=NUMBER BY weight=NUMBER # ThreePartBy
    ;

multiplier
    : sets=NUMBER X reps=NUMBER                   # PartialMultiplier
    | sets=NUMBER X reps=NUMBER X weight=NUMBER   # FullMultiplier
    ;

number : NUMBER ;
note   : WORD+ ;

NUMBER  : [0-9]+ ;
X       : [xX] ;
BY      : [bB][yY] ;
WORD    : [a-zA-Z]+ ;
NEWLINE : '\n' ;
WS      : [ \t]+ -> skip ;
```

**Grammar handles:** Line classification, token parsing, inline notes

**Code handles:** Paragraph splitting, pre-filter optimization, output formatting

## Table Format

```
| Exercise | Set 1 | Set 2 | Set 3 | Notes |
|----------|-------|-------|-------|-------|
| Squat    | 8@135 | 8@135 | 8@135 | light |
| Bench    | 8@95  | 8@95  | 8@95  |       |
```

- Header: Exercise, Set 1...Set N, Notes (if any exist)
- Set cells: `reps@weight`
- Columns padded for org-mode alignment
- Notes column only if any exercise has notes

## Transformation Example

Input:
```
ate a banana beforehand
squat 8 8 8 135 light headed
bench 3x8 95

Regular notes continue here.
```

Output:
```
ate a banana beforehand

| Exercise | Set 1 | Set 2 | Set 3 | Notes        |
|----------|-------|-------|-------|--------------|
| Squat    | 8@135 | 8@135 | 8@135 | light headed |
| Bench    | 8@95  | 8@95  | 8@95  |              |

Regular notes continue here.
```

**Rules:**
- Prose lines (no numbers) bubble above table
- Inline notes (words after numbers) go in Notes column
- Blank line separates prose from table
- Non-workout paragraphs pass through unchanged

## API & Structure

**Go CLI:**
```
cmd/workoutliner/main.go

Usage:
    workoutliner < notes.txt
    workoutliner notes.txt
```

**Swift Library:**
```swift
import WorkoutLiner

public func transform(_ input: String) -> String
```

**Project structure:**
```
workoutliner/
├── Paragraph.g4
├── go/
│   ├── cmd/workoutliner/
│   │   └── main.go
│   ├── transform.go
│   └── parser/
└── swift/
    ├── Sources/WorkoutLiner/
    │   ├── WorkoutLiner.swift
    │   └── Parser/
    └── Package.swift
```

## Error Handling

- Parse failures → output paragraph unchanged
- Empty input → empty output
- No workout paragraphs → input unchanged
