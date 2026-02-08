package workoutliner

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const (
	defaultReps  = 8
	defaultSets  = 3
	repThreshold = 20
)

type ParseResult struct {
	ProseLines []string
	Exercises  []Exercise
}

func (r *ParseResult) HasExercises() bool {
	return len(r.Exercises) > 0
}

type WorkoutlinerASTVisitor struct {
	BaseWorkoutlinerVisitor
}

func NewWorkoutlinerASTVisitor() *WorkoutlinerASTVisitor {
	return &WorkoutlinerASTVisitor{}
}

func (v *WorkoutlinerASTVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *WorkoutlinerASTVisitor) VisitParagraph(ctx *ParagraphContext) interface{} {
	result := &ParseResult{}
	for _, lineCtx := range ctx.AllLine() {
		lineResult := v.Visit(lineCtx)
		if lineResult == nil {
			continue
		}
		switch lr := lineResult.(type) {
		case string:
			result.ProseLines = append(result.ProseLines, lr)
		case Exercise:
			if len(lr.Sets) > 0 {
				result.Exercises = append(result.Exercises, lr)
			} else {
				result.ProseLines = append(result.ProseLines, lr.Name)
			}
		}
	}
	return result
}

func (v *WorkoutlinerASTVisitor) VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{} {
	return v.Visit(ctx.ExerciseLine())
}

func (v *WorkoutlinerASTVisitor) VisitProseLineAlt(ctx *ProseLineAltContext) interface{} {
	return v.Visit(ctx.ProseLine())
}

func (v *WorkoutlinerASTVisitor) VisitExerciseLine(ctx *ExerciseLineContext) interface{} {
	words := ctx.AllWORD()
	numericToken := ctx.NumericToken()
	tokens := ctx.AllToken()

	name := v.extractName(words, numericToken)
	sets := v.buildSetsFromExercise(numericToken, tokens)

	return Exercise{Name: name, Sets: sets}
}

func (v *WorkoutlinerASTVisitor) VisitProseLine(ctx *ProseLineContext) interface{} {
	words := []string{}
	for _, word := range ctx.AllWORD() {
		words = append(words, word.GetText())
	}
	return strings.Join(words, " ")
}

func (v *WorkoutlinerASTVisitor) extractName(words []antlr.TerminalNode, numericToken INumericTokenContext) string {
	firstTokenStart := numericToken.GetStart().GetStart()
	nameWords := []string{}
	for _, w := range words {
		if w.GetSymbol().GetStart() < firstTokenStart {
			nameWords = append(nameWords, w.GetText())
		}
	}
	return titleCase(strings.Join(nameWords, " "))
}

func (v *WorkoutlinerASTVisitor) buildSetsFromExercise(numericToken INumericTokenContext, tokens []ITokenContext) []Set {
	var allTokenResults []tokenResult

	numResult := v.Visit(numericToken)
	if numResult != nil {
		allTokenResults = append(allTokenResults, numResult.(tokenResult))
	}

	for _, tokenCtx := range tokens {
		result := v.Visit(tokenCtx)
		if result != nil {
			allTokenResults = append(allTokenResults, result.(tokenResult))
		}
	}

	return v.buildSetsFromResults(allTokenResults)
}

type tokenResult struct {
	kind        string
	sets        []Set
	value       int
	text        string
	pendingReps []int
}

func (v *WorkoutlinerASTVisitor) VisitNumericToken(ctx *NumericTokenContext) interface{} {
	if byExpr := ctx.ByExpr(); byExpr != nil {
		return v.Visit(byExpr)
	}
	if mult := ctx.Multiplier(); mult != nil {
		return v.Visit(mult)
	}
	if num := ctx.Number(); num != nil {
		return v.Visit(num)
	}
	return nil
}

func (v *WorkoutlinerASTVisitor) VisitToken(ctx *TokenContext) interface{} {
	if numToken := ctx.NumericToken(); numToken != nil {
		return v.Visit(numToken)
	}
	if note := ctx.Note(); note != nil {
		return v.Visit(note)
	}
	return nil
}

func (v *WorkoutlinerASTVisitor) VisitTwoPartBy(ctx *TwoPartByContext) interface{} {
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]Set, defaultSets)
	for i := range sets {
		sets[i] = Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *WorkoutlinerASTVisitor) VisitThreePartBy(ctx *ThreePartByContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]Set, numSets)
	for i := range sets {
		sets[i] = Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *WorkoutlinerASTVisitor) VisitFullMultiplier(ctx *FullMultiplierContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]Set, numSets)
	for i := range sets {
		sets[i] = Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *WorkoutlinerASTVisitor) VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{} {
	first := toInt(ctx.GetSets().GetText())
	second := toInt(ctx.GetReps().GetText())

	// If second number looks like a weight, treat as reps@weight
	if second > repThreshold {
		return tokenResult{kind: "sets", sets: []Set{{Reps: first, Weight: second}}}
	}

	// Otherwise, it's sets x reps waiting for weight
	pending := make([]int, first)
	for i := range pending {
		pending[i] = second
	}
	return tokenResult{kind: "pending_reps", pendingReps: pending}
}

func (v *WorkoutlinerASTVisitor) VisitNumber(ctx *NumberContext) interface{} {
	n := toInt(ctx.NUMBER().GetText())
	if n > repThreshold {
		return tokenResult{kind: "weight", value: n}
	}
	return tokenResult{kind: "reps", value: n}
}

func (v *WorkoutlinerASTVisitor) VisitNote(ctx *NoteContext) interface{} {
	words := []string{}
	for _, word := range ctx.AllWORD() {
		words = append(words, word.GetText())
	}
	return tokenResult{kind: "note", text: strings.Join(words, " ")}
}

func (v *WorkoutlinerASTVisitor) buildSetsFromResults(tokenResults []tokenResult) []Set {
	var sets []Set
	var pendingReps []int
	var lastWeight int

	for _, tok := range tokenResults {
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
				sets[len(sets)-1].Note = joinNotes(sets[len(sets)-1].Note, tok.text)
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

func joinNotes(existing, new string) string {
	if existing == "" {
		return new
	}
	if new == "" {
		return existing
	}
	return existing + " " + new
}
