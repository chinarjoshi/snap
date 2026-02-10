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

	// State for handling continuation lines
	var lastExercise *Exercise  // Currently being built
	var pendingProse string     // Prose that might become exercise name

	flush := func() {
		if lastExercise != nil && len(lastExercise.Sets) > 0 {
			result.Exercises = append(result.Exercises, *lastExercise)
		} else if pendingProse != "" {
			result.ProseLines = append(result.ProseLines, pendingProse)
		}
		lastExercise = nil
		pendingProse = ""
	}

	for _, lineCtx := range ctx.AllLine() {
		lineResult := v.Visit(lineCtx)
		if lineResult == nil {
			continue
		}
		switch lr := lineResult.(type) {
		case string: // proseLine
			flush()
			pendingProse = lr

		case Exercise: // exerciseLine
			flush()
			if len(lr.Sets) > 0 {
				lastExercise = &lr
			} else {
				pendingProse = lr.Name
			}

		case []Exercise: // supersetLine
			flush()
			for _, ex := range lr {
				if len(ex.Sets) > 0 {
					result.Exercises = append(result.Exercises, ex)
				}
			}

		case []tokenResult: // continuationLine
			if lastExercise != nil {
				// Extend existing exercise
				lastWeight := 0
				if len(lastExercise.Sets) > 0 {
					lastWeight = lastExercise.Sets[len(lastExercise.Sets)-1].Weight
				}
				newSets := v.buildSetsFromResultsWithInitialWeight(lr, lastWeight)
				lastExercise.Sets = append(lastExercise.Sets, newSets...)
			} else if pendingProse != "" {
				// Start new exercise from pending prose
				lastExercise = &Exercise{
					Name: titleCase(pendingProse),
				}
				sets := v.buildSetsFromResults(lr)
				lastExercise.Sets = sets
				pendingProse = ""
			}
			// If no pending prose or exercise, ignore the continuation line
		}
	}

	flush()
	return result
}

func (v *WorkoutlinerASTVisitor) VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{} {
	return v.Visit(ctx.ExerciseLine())
}

func (v *WorkoutlinerASTVisitor) VisitProseLineAlt(ctx *ProseLineAltContext) interface{} {
	return v.Visit(ctx.ProseLine())
}

func (v *WorkoutlinerASTVisitor) VisitContinuationLineAlt(ctx *ContinuationLineAltContext) interface{} {
	return v.Visit(ctx.ContinuationLine())
}

func (v *WorkoutlinerASTVisitor) VisitContinuationLine(ctx *ContinuationLineContext) interface{} {
	var results []tokenResult

	numericToken := ctx.NumericToken()
	if numResult := v.Visit(numericToken); numResult != nil {
		results = append(results, numResult.(tokenResult))
	}

	for _, tokenCtx := range ctx.AllToken() {
		if result := v.Visit(tokenCtx); result != nil {
			results = append(results, result.(tokenResult))
		}
	}

	return results
}

func (v *WorkoutlinerASTVisitor) VisitSupersetLineAlt(ctx *SupersetLineAltContext) interface{} {
	return v.Visit(ctx.SupersetLine())
}

func (v *WorkoutlinerASTVisitor) VisitSupersetLine(ctx *SupersetLineContext) interface{} {
	supersetNames := ctx.AllSupersetName()
	if len(supersetNames) != 2 {
		return nil
	}

	// Extract exercise names (keep * suffix)
	name1 := v.extractSupersetName(supersetNames[0])
	name2 := v.extractSupersetName(supersetNames[1])

	// Build all sets from tokens
	tokens := ctx.AllToken()
	allSets := v.buildSetsFromTokens(tokens)

	// Alternate sets between two exercises
	var sets1, sets2 []Set
	for i, set := range allSets {
		if i%2 == 0 {
			sets1 = append(sets1, set)
		} else {
			sets2 = append(sets2, set)
		}
	}

	return []Exercise{
		{Name: name1, Sets: sets1},
		{Name: name2, Sets: sets2},
	}
}

func (v *WorkoutlinerASTVisitor) extractSupersetName(ctx ISupersetNameContext) string {
	words := []string{}
	for _, w := range ctx.AllWORD() {
		words = append(words, w.GetText())
	}
	name := titleCase(strings.Join(words, " "))
	return name + "*"
}

func (v *WorkoutlinerASTVisitor) buildSetsFromTokens(tokens []ITokenContext) []Set {
	var sets []Set

	for _, tokenCtx := range tokens {
		result := v.Visit(tokenCtx)
		if result == nil {
			continue
		}
		tok := result.(tokenResult)
		switch tok.kind {
		case "sets":
			// Fully specified sets (e.g., 3x8x135) - add each as individual set
			sets = append(sets, tok.sets...)
		case "pending_reps":
			// Partial multiplier (e.g., 3x8) - add each as bodyweight
			for _, r := range tok.pendingReps {
				sets = append(sets, Set{Reps: r, Weight: 0})
			}
		case "weight":
			// Single weight - assume default reps
			sets = append(sets, Set{Reps: defaultReps, Weight: tok.value})
		case "reps":
			// Single rep count - bodyweight
			sets = append(sets, Set{Reps: tok.value, Weight: 0})
		}
	}

	return sets
}

func (v *WorkoutlinerASTVisitor) VisitExerciseLine(ctx *ExerciseLineContext) interface{} {
	words := ctx.AllWORD()
	numericToken := ctx.NumericToken()
	tokens := ctx.AllToken()

	// Collect first exercise name
	firstName := v.extractName(words, numericToken)

	// Collect all token results including notes
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

	// Split into exercises by detecting note+numericToken pattern
	exercises := v.splitIntoExercises(firstName, allTokenResults)
	if len(exercises) == 1 {
		return exercises[0]
	}
	return exercises
}

func (v *WorkoutlinerASTVisitor) splitIntoExercises(firstName string, tokens []tokenResult) []Exercise {
	var exercises []Exercise
	currentName := firstName
	var currentTokens []tokenResult

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		// Check if this note starts a new exercise (note followed by numeric token)
		if tok.kind == "note" && i+1 < len(tokens) && isNumericToken(tokens[i+1]) {
			// Flush the current exercise
			if currentName != "" {
				sets := v.buildSetsFromResults(currentTokens)
				exercises = append(exercises, Exercise{Name: currentName, Sets: sets})
			}
			// Start new exercise with this note as name
			currentName = titleCase(tok.text)
			currentTokens = nil
		} else {
			currentTokens = append(currentTokens, tok)
		}
	}

	// Flush last exercise
	if currentName != "" {
		sets := v.buildSetsFromResults(currentTokens)
		exercises = append(exercises, Exercise{Name: currentName, Sets: sets})
	}

	return exercises
}

func isNumericToken(tok tokenResult) bool {
	return tok.kind == "sets" || tok.kind == "pending_reps" || tok.kind == "weight" || tok.kind == "reps"
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

func (v *WorkoutlinerASTVisitor) buildSetsFromResultsWithInitialWeight(tokenResults []tokenResult, initialWeight int) []Set {
	return v.buildSetsFromResultsImpl(tokenResults, initialWeight)
}

func (v *WorkoutlinerASTVisitor) buildSetsFromResults(tokenResults []tokenResult) []Set {
	return v.buildSetsFromResultsImpl(tokenResults, 0)
}

func (v *WorkoutlinerASTVisitor) buildSetsFromResultsImpl(tokenResults []tokenResult, initialWeight int) []Set {
	var sets []Set
	var pendingReps []int
	lastWeight := initialWeight

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
			// Resolve pending reps before attaching note
			if len(pendingReps) > 0 {
				for _, r := range pendingReps {
					sets = append(sets, Set{Reps: r, Weight: lastWeight})
				}
				pendingReps = nil
			}
			if len(sets) > 0 {
				sets[len(sets)-1].Note = joinNotes(sets[len(sets)-1].Note, tok.text)
			}
		}
	}

	for _, r := range pendingReps {
		sets = append(sets, Set{Reps: r, Weight: lastWeight})
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
