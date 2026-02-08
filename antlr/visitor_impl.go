package antlr

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chijoshi/workoutliner/parser"
)

const (
	defaultReps  = 8
	defaultSets  = 3
	repThreshold = 20
)

type WorkoutASTVisitor struct {
	BaseWorkoutVisitor
}

func NewWorkoutASTVisitor() *WorkoutASTVisitor {
	return &WorkoutASTVisitor{}
}

func (v *WorkoutASTVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *WorkoutASTVisitor) VisitWorkoutLog(ctx *WorkoutLogContext) interface{} {
	exercises := []parser.Exercise{}
	for _, exCtx := range ctx.AllExercise() {
		ex := v.Visit(exCtx).(parser.Exercise)
		exercises = append(exercises, ex)
	}
	return &parser.WorkoutLog{Exercises: exercises}
}

func (v *WorkoutASTVisitor) VisitExercise(ctx *ExerciseContext) interface{} {
	name := v.Visit(ctx.ExerciseName()).(string)
	tokens := ctx.AllToken()
	sets := v.buildSets(tokens)
	return parser.Exercise{Name: name, Sets: sets}
}

func (v *WorkoutASTVisitor) VisitExerciseName(ctx *ExerciseNameContext) interface{} {
	words := []string{}
	for _, word := range ctx.AllWORD() {
		words = append(words, word.GetText())
	}
	return titleCase(strings.Join(words, " "))
}

func (v *WorkoutASTVisitor) VisitToken(ctx *TokenContext) interface{} {
	if byExpr := ctx.ByExpr(); byExpr != nil {
		return v.Visit(byExpr)
	}
	if mult := ctx.Multiplier(); mult != nil {
		return v.Visit(mult)
	}
	if num := ctx.Number(); num != nil {
		return v.Visit(num)
	}
	if note := ctx.Note(); note != nil {
		return v.Visit(note)
	}
	return nil
}

type tokenResult struct {
	kind        string
	sets        []parser.Set
	value       int
	text        string
	pendingReps []int
}

func (v *WorkoutASTVisitor) VisitTwoPartBy(ctx *TwoPartByContext) interface{} {
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]parser.Set, defaultSets)
	for i := range sets {
		sets[i] = parser.Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *WorkoutASTVisitor) VisitThreePartBy(ctx *ThreePartByContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]parser.Set, numSets)
	for i := range sets {
		sets[i] = parser.Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *WorkoutASTVisitor) VisitFullMultiplier(ctx *FullMultiplierContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]parser.Set, numSets)
	for i := range sets {
		sets[i] = parser.Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *WorkoutASTVisitor) VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())

	pending := make([]int, numSets)
	for i := range pending {
		pending[i] = reps
	}
	return tokenResult{kind: "pending_reps", pendingReps: pending}
}

func (v *WorkoutASTVisitor) VisitNumber(ctx *NumberContext) interface{} {
	n := toInt(ctx.NUMBER().GetText())
	if n > repThreshold {
		return tokenResult{kind: "weight", value: n}
	}
	return tokenResult{kind: "reps", value: n}
}

func (v *WorkoutASTVisitor) VisitNote(ctx *NoteContext) interface{} {
	words := []string{}
	for _, word := range ctx.AllWORD() {
		words = append(words, word.GetText())
	}
	return tokenResult{kind: "note", text: strings.Join(words, " ")}
}

func (v *WorkoutASTVisitor) buildSets(tokenCtxs []ITokenContext) []parser.Set {
	var sets []parser.Set
	var pendingReps []int
	var lastWeight int

	for _, tokenCtx := range tokenCtxs {
		result := v.Visit(tokenCtx)
		if result == nil {
			continue
		}

		tok := result.(tokenResult)
		switch tok.kind {
		case "sets":
			sets = append(sets, tok.sets...)
			if len(tok.sets) > 0 {
				lastWeight = tok.sets[len(tok.sets)-1].Weight
			}
		case "pending_reps":
			pendingReps = append(pendingReps, tok.pendingReps...)
		case "weight":
			if len(pendingReps) > 0 {
				for _, r := range pendingReps {
					sets = append(sets, parser.Set{Reps: r, Weight: tok.value})
				}
				pendingReps = nil
			} else {
				sets = append(sets, parser.Set{Reps: defaultReps, Weight: tok.value})
			}
			lastWeight = tok.value
		case "reps":
			pendingReps = append(pendingReps, tok.value)
		case "note":
			if len(sets) > 0 {
				sets[len(sets)-1].Note = tok.text
			}
		}
	}

	for _, r := range pendingReps {
		if lastWeight > 0 {
			sets = append(sets, parser.Set{Reps: r, Weight: lastWeight})
		}
	}

	return sets
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
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
