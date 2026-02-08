package paragraph

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

type ParseResult struct {
	ProseLines []string
	Exercises  []parser.Exercise
}

func (r *ParseResult) HasExercises() bool {
	return len(r.Exercises) > 0
}

type ParagraphASTVisitor struct {
	BaseParagraphVisitor
}

func NewParagraphASTVisitor() *ParagraphASTVisitor {
	return &ParagraphASTVisitor{}
}

func (v *ParagraphASTVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ParagraphASTVisitor) VisitParagraph(ctx *ParagraphContext) interface{} {
	result := &ParseResult{}
	for _, lineCtx := range ctx.AllLine() {
		lineResult := v.Visit(lineCtx)
		if lineResult == nil {
			continue
		}
		switch lr := lineResult.(type) {
		case string:
			result.ProseLines = append(result.ProseLines, lr)
		case parser.Exercise:
			if len(lr.Sets) > 0 {
				result.Exercises = append(result.Exercises, lr)
			} else {
				result.ProseLines = append(result.ProseLines, lr.Name)
			}
		}
	}
	return result
}

func (v *ParagraphASTVisitor) VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{} {
	return v.Visit(ctx.ExerciseLine())
}

func (v *ParagraphASTVisitor) VisitProseLineAlt(ctx *ProseLineAltContext) interface{} {
	return v.Visit(ctx.ProseLine())
}

func (v *ParagraphASTVisitor) VisitExerciseLine(ctx *ExerciseLineContext) interface{} {
	words := ctx.AllWORD()
	numericToken := ctx.NumericToken()
	tokens := ctx.AllToken()

	name := v.extractName(words, numericToken)
	sets := v.buildSetsFromExercise(numericToken, tokens)

	return parser.Exercise{Name: name, Sets: sets}
}

func (v *ParagraphASTVisitor) VisitProseLine(ctx *ProseLineContext) interface{} {
	words := []string{}
	for _, word := range ctx.AllWORD() {
		words = append(words, word.GetText())
	}
	return strings.Join(words, " ")
}

func (v *ParagraphASTVisitor) extractName(words []antlr.TerminalNode, numericToken INumericTokenContext) string {
	firstTokenStart := numericToken.GetStart().GetStart()
	nameWords := []string{}
	for _, w := range words {
		if w.GetSymbol().GetStart() < firstTokenStart {
			nameWords = append(nameWords, w.GetText())
		}
	}
	return titleCase(strings.Join(nameWords, " "))
}

func (v *ParagraphASTVisitor) buildSetsFromExercise(numericToken INumericTokenContext, tokens []ITokenContext) []parser.Set {
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
	sets        []parser.Set
	value       int
	text        string
	pendingReps []int
}

func (v *ParagraphASTVisitor) VisitNumericToken(ctx *NumericTokenContext) interface{} {
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

func (v *ParagraphASTVisitor) VisitToken(ctx *TokenContext) interface{} {
	if numToken := ctx.NumericToken(); numToken != nil {
		return v.Visit(numToken)
	}
	if note := ctx.Note(); note != nil {
		return v.Visit(note)
	}
	return nil
}

func (v *ParagraphASTVisitor) VisitTwoPartBy(ctx *TwoPartByContext) interface{} {
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]parser.Set, defaultSets)
	for i := range sets {
		sets[i] = parser.Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *ParagraphASTVisitor) VisitThreePartBy(ctx *ThreePartByContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]parser.Set, numSets)
	for i := range sets {
		sets[i] = parser.Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *ParagraphASTVisitor) VisitFullMultiplier(ctx *FullMultiplierContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())
	weight := toInt(ctx.GetWeight().GetText())

	sets := make([]parser.Set, numSets)
	for i := range sets {
		sets[i] = parser.Set{Reps: reps, Weight: weight}
	}
	return tokenResult{kind: "sets", sets: sets}
}

func (v *ParagraphASTVisitor) VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{} {
	numSets := toInt(ctx.GetSets().GetText())
	reps := toInt(ctx.GetReps().GetText())

	pending := make([]int, numSets)
	for i := range pending {
		pending[i] = reps
	}
	return tokenResult{kind: "pending_reps", pendingReps: pending}
}

func (v *ParagraphASTVisitor) VisitNumber(ctx *NumberContext) interface{} {
	n := toInt(ctx.NUMBER().GetText())
	if n > repThreshold {
		return tokenResult{kind: "weight", value: n}
	}
	return tokenResult{kind: "reps", value: n}
}

func (v *ParagraphASTVisitor) VisitNote(ctx *NoteContext) interface{} {
	words := []string{}
	for _, word := range ctx.AllWORD() {
		words = append(words, word.GetText())
	}
	return tokenResult{kind: "note", text: strings.Join(words, " ")}
}

func (v *ParagraphASTVisitor) buildSetsFromResults(tokenResults []tokenResult) []parser.Set {
	var sets []parser.Set
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
				sets[len(sets)-1].Note = joinNotes(sets[len(sets)-1].Note, tok.text)
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

func joinNotes(existing, new string) string {
	if existing == "" {
		return new
	}
	if new == "" {
		return existing
	}
	return existing + " " + new
}
