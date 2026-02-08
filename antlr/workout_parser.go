// Code generated from ./Workout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // Workout
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type WorkoutParser struct {
	*antlr.BaseParser
}

var WorkoutParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func workoutParserInit() {
	staticData := &WorkoutParserStaticData
	staticData.SymbolicNames = []string{
		"", "NUMBER", "X", "BY", "WORD", "DELIMITER", "WS",
	}
	staticData.RuleNames = []string{
		"workoutLog", "exercise", "exerciseName", "token", "byExpr", "multiplier",
		"number", "note",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 72, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 5, 0, 20, 8, 0, 10,
		0, 12, 0, 23, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 29, 8, 1, 10, 1, 12,
		1, 32, 9, 1, 1, 2, 4, 2, 35, 8, 2, 11, 2, 12, 2, 36, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 43, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 53, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 63,
		8, 5, 1, 6, 1, 6, 1, 7, 4, 7, 68, 8, 7, 11, 7, 12, 7, 69, 1, 7, 0, 0, 8,
		0, 2, 4, 6, 8, 10, 12, 14, 0, 0, 72, 0, 16, 1, 0, 0, 0, 2, 26, 1, 0, 0,
		0, 4, 34, 1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8, 52, 1, 0, 0, 0, 10, 62, 1,
		0, 0, 0, 12, 64, 1, 0, 0, 0, 14, 67, 1, 0, 0, 0, 16, 21, 3, 2, 1, 0, 17,
		18, 5, 5, 0, 0, 18, 20, 3, 2, 1, 0, 19, 17, 1, 0, 0, 0, 20, 23, 1, 0, 0,
		0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 24, 1, 0, 0, 0, 23, 21,
		1, 0, 0, 0, 24, 25, 5, 0, 0, 1, 25, 1, 1, 0, 0, 0, 26, 30, 3, 4, 2, 0,
		27, 29, 3, 6, 3, 0, 28, 27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33,
		35, 5, 4, 0, 0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0,
		0, 36, 37, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0, 38, 43, 3, 8, 4, 0, 39, 43, 3,
		10, 5, 0, 40, 43, 3, 12, 6, 0, 41, 43, 3, 14, 7, 0, 42, 38, 1, 0, 0, 0,
		42, 39, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 7, 1, 0,
		0, 0, 44, 45, 5, 1, 0, 0, 45, 46, 5, 3, 0, 0, 46, 53, 5, 1, 0, 0, 47, 48,
		5, 1, 0, 0, 48, 49, 5, 3, 0, 0, 49, 50, 5, 1, 0, 0, 50, 51, 5, 3, 0, 0,
		51, 53, 5, 1, 0, 0, 52, 44, 1, 0, 0, 0, 52, 47, 1, 0, 0, 0, 53, 9, 1, 0,
		0, 0, 54, 55, 5, 1, 0, 0, 55, 56, 5, 2, 0, 0, 56, 63, 5, 1, 0, 0, 57, 58,
		5, 1, 0, 0, 58, 59, 5, 2, 0, 0, 59, 60, 5, 1, 0, 0, 60, 61, 5, 2, 0, 0,
		61, 63, 5, 1, 0, 0, 62, 54, 1, 0, 0, 0, 62, 57, 1, 0, 0, 0, 63, 11, 1,
		0, 0, 0, 64, 65, 5, 1, 0, 0, 65, 13, 1, 0, 0, 0, 66, 68, 5, 4, 0, 0, 67,
		66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0,
		0, 70, 15, 1, 0, 0, 0, 7, 21, 30, 36, 42, 52, 62, 69,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// WorkoutParserInit initializes any static state used to implement WorkoutParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWorkoutParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WorkoutParserInit() {
	staticData := &WorkoutParserStaticData
	staticData.once.Do(workoutParserInit)
}

// NewWorkoutParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWorkoutParser(input antlr.TokenStream) *WorkoutParser {
	WorkoutParserInit()
	this := new(WorkoutParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &WorkoutParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Workout.g4"

	return this
}

// WorkoutParser tokens.
const (
	WorkoutParserEOF       = antlr.TokenEOF
	WorkoutParserNUMBER    = 1
	WorkoutParserX         = 2
	WorkoutParserBY        = 3
	WorkoutParserWORD      = 4
	WorkoutParserDELIMITER = 5
	WorkoutParserWS        = 6
)

// WorkoutParser rules.
const (
	WorkoutParserRULE_workoutLog   = 0
	WorkoutParserRULE_exercise     = 1
	WorkoutParserRULE_exerciseName = 2
	WorkoutParserRULE_token        = 3
	WorkoutParserRULE_byExpr       = 4
	WorkoutParserRULE_multiplier   = 5
	WorkoutParserRULE_number       = 6
	WorkoutParserRULE_note         = 7
)

// IWorkoutLogContext is an interface to support dynamic dispatch.
type IWorkoutLogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExercise() []IExerciseContext
	Exercise(i int) IExerciseContext
	EOF() antlr.TerminalNode
	AllDELIMITER() []antlr.TerminalNode
	DELIMITER(i int) antlr.TerminalNode

	// IsWorkoutLogContext differentiates from other interfaces.
	IsWorkoutLogContext()
}

type WorkoutLogContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorkoutLogContext() *WorkoutLogContext {
	var p = new(WorkoutLogContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_workoutLog
	return p
}

func InitEmptyWorkoutLogContext(p *WorkoutLogContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_workoutLog
}

func (*WorkoutLogContext) IsWorkoutLogContext() {}

func NewWorkoutLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorkoutLogContext {
	var p = new(WorkoutLogContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_workoutLog

	return p
}

func (s *WorkoutLogContext) GetParser() antlr.Parser { return s.parser }

func (s *WorkoutLogContext) AllExercise() []IExerciseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExerciseContext); ok {
			len++
		}
	}

	tst := make([]IExerciseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExerciseContext); ok {
			tst[i] = t.(IExerciseContext)
			i++
		}
	}

	return tst
}

func (s *WorkoutLogContext) Exercise(i int) IExerciseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExerciseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExerciseContext)
}

func (s *WorkoutLogContext) EOF() antlr.TerminalNode {
	return s.GetToken(WorkoutParserEOF, 0)
}

func (s *WorkoutLogContext) AllDELIMITER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserDELIMITER)
}

func (s *WorkoutLogContext) DELIMITER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserDELIMITER, i)
}

func (s *WorkoutLogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkoutLogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorkoutLogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterWorkoutLog(s)
	}
}

func (s *WorkoutLogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitWorkoutLog(s)
	}
}

func (s *WorkoutLogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitWorkoutLog(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) WorkoutLog() (localctx IWorkoutLogContext) {
	localctx = NewWorkoutLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WorkoutParserRULE_workoutLog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Exercise()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WorkoutParserDELIMITER {
		{
			p.SetState(17)
			p.Match(WorkoutParserDELIMITER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)
			p.Exercise()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(24)
		p.Match(WorkoutParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExerciseContext is an interface to support dynamic dispatch.
type IExerciseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExerciseName() IExerciseNameContext
	AllToken() []ITokenContext
	Token(i int) ITokenContext

	// IsExerciseContext differentiates from other interfaces.
	IsExerciseContext()
}

type ExerciseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExerciseContext() *ExerciseContext {
	var p = new(ExerciseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_exercise
	return p
}

func InitEmptyExerciseContext(p *ExerciseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_exercise
}

func (*ExerciseContext) IsExerciseContext() {}

func NewExerciseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExerciseContext {
	var p = new(ExerciseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_exercise

	return p
}

func (s *ExerciseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExerciseContext) ExerciseName() IExerciseNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExerciseNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExerciseNameContext)
}

func (s *ExerciseContext) AllToken() []ITokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITokenContext); ok {
			len++
		}
	}

	tst := make([]ITokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITokenContext); ok {
			tst[i] = t.(ITokenContext)
			i++
		}
	}

	return tst
}

func (s *ExerciseContext) Token(i int) ITokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITokenContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *ExerciseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExerciseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExerciseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterExercise(s)
	}
}

func (s *ExerciseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitExercise(s)
	}
}

func (s *ExerciseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitExercise(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) Exercise() (localctx IExerciseContext) {
	localctx = NewExerciseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WorkoutParserRULE_exercise)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.ExerciseName()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WorkoutParserNUMBER || _la == WorkoutParserWORD {
		{
			p.SetState(27)
			p.Token()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExerciseNameContext is an interface to support dynamic dispatch.
type IExerciseNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode

	// IsExerciseNameContext differentiates from other interfaces.
	IsExerciseNameContext()
}

type ExerciseNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExerciseNameContext() *ExerciseNameContext {
	var p = new(ExerciseNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_exerciseName
	return p
}

func InitEmptyExerciseNameContext(p *ExerciseNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_exerciseName
}

func (*ExerciseNameContext) IsExerciseNameContext() {}

func NewExerciseNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExerciseNameContext {
	var p = new(ExerciseNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_exerciseName

	return p
}

func (s *ExerciseNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ExerciseNameContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserWORD)
}

func (s *ExerciseNameContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserWORD, i)
}

func (s *ExerciseNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExerciseNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExerciseNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterExerciseName(s)
	}
}

func (s *ExerciseNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitExerciseName(s)
	}
}

func (s *ExerciseNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitExerciseName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) ExerciseName() (localctx IExerciseNameContext) {
	localctx = NewExerciseNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WorkoutParserRULE_exerciseName)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(33)
				p.Match(WorkoutParserWORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ByExpr() IByExprContext
	Multiplier() IMultiplierContext
	Number() INumberContext
	Note() INoteContext

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_token
	return p
}

func InitEmptyTokenContext(p *TokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_token
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) ByExpr() IByExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IByExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IByExprContext)
}

func (s *TokenContext) Multiplier() IMultiplierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplierContext)
}

func (s *TokenContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *TokenContext) Note() INoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitToken(s)
	}
}

func (s *TokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) Token() (localctx ITokenContext) {
	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WorkoutParserRULE_token)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.ByExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Multiplier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Number()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Note()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IByExprContext is an interface to support dynamic dispatch.
type IByExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsByExprContext differentiates from other interfaces.
	IsByExprContext()
}

type ByExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByExprContext() *ByExprContext {
	var p = new(ByExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_byExpr
	return p
}

func InitEmptyByExprContext(p *ByExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_byExpr
}

func (*ByExprContext) IsByExprContext() {}

func NewByExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByExprContext {
	var p = new(ByExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_byExpr

	return p
}

func (s *ByExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ByExprContext) CopyAll(ctx *ByExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ByExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TwoPartByContext struct {
	ByExprContext
	reps   antlr.Token
	weight antlr.Token
}

func NewTwoPartByContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TwoPartByContext {
	var p = new(TwoPartByContext)

	InitEmptyByExprContext(&p.ByExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ByExprContext))

	return p
}

func (s *TwoPartByContext) GetReps() antlr.Token { return s.reps }

func (s *TwoPartByContext) GetWeight() antlr.Token { return s.weight }

func (s *TwoPartByContext) SetReps(v antlr.Token) { s.reps = v }

func (s *TwoPartByContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *TwoPartByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TwoPartByContext) BY() antlr.TerminalNode {
	return s.GetToken(WorkoutParserBY, 0)
}

func (s *TwoPartByContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserNUMBER)
}

func (s *TwoPartByContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserNUMBER, i)
}

func (s *TwoPartByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterTwoPartBy(s)
	}
}

func (s *TwoPartByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitTwoPartBy(s)
	}
}

func (s *TwoPartByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitTwoPartBy(s)

	default:
		return t.VisitChildren(s)
	}
}

type ThreePartByContext struct {
	ByExprContext
	sets   antlr.Token
	reps   antlr.Token
	weight antlr.Token
}

func NewThreePartByContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThreePartByContext {
	var p = new(ThreePartByContext)

	InitEmptyByExprContext(&p.ByExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ByExprContext))

	return p
}

func (s *ThreePartByContext) GetSets() antlr.Token { return s.sets }

func (s *ThreePartByContext) GetReps() antlr.Token { return s.reps }

func (s *ThreePartByContext) GetWeight() antlr.Token { return s.weight }

func (s *ThreePartByContext) SetSets(v antlr.Token) { s.sets = v }

func (s *ThreePartByContext) SetReps(v antlr.Token) { s.reps = v }

func (s *ThreePartByContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *ThreePartByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreePartByContext) AllBY() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserBY)
}

func (s *ThreePartByContext) BY(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserBY, i)
}

func (s *ThreePartByContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserNUMBER)
}

func (s *ThreePartByContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserNUMBER, i)
}

func (s *ThreePartByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterThreePartBy(s)
	}
}

func (s *ThreePartByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitThreePartBy(s)
	}
}

func (s *ThreePartByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitThreePartBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) ByExpr() (localctx IByExprContext) {
	localctx = NewByExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WorkoutParserRULE_byExpr)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTwoPartByContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*TwoPartByContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(WorkoutParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*TwoPartByContext).weight = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewThreePartByContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*ThreePartByContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(WorkoutParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*ThreePartByContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(WorkoutParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*ThreePartByContext).weight = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplierContext is an interface to support dynamic dispatch.
type IMultiplierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMultiplierContext differentiates from other interfaces.
	IsMultiplierContext()
}

type MultiplierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplierContext() *MultiplierContext {
	var p = new(MultiplierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_multiplier
	return p
}

func InitEmptyMultiplierContext(p *MultiplierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_multiplier
}

func (*MultiplierContext) IsMultiplierContext() {}

func NewMultiplierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplierContext {
	var p = new(MultiplierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_multiplier

	return p
}

func (s *MultiplierContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplierContext) CopyAll(ctx *MultiplierContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MultiplierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PartialMultiplierContext struct {
	MultiplierContext
	sets antlr.Token
	reps antlr.Token
}

func NewPartialMultiplierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PartialMultiplierContext {
	var p = new(PartialMultiplierContext)

	InitEmptyMultiplierContext(&p.MultiplierContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplierContext))

	return p
}

func (s *PartialMultiplierContext) GetSets() antlr.Token { return s.sets }

func (s *PartialMultiplierContext) GetReps() antlr.Token { return s.reps }

func (s *PartialMultiplierContext) SetSets(v antlr.Token) { s.sets = v }

func (s *PartialMultiplierContext) SetReps(v antlr.Token) { s.reps = v }

func (s *PartialMultiplierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartialMultiplierContext) X() antlr.TerminalNode {
	return s.GetToken(WorkoutParserX, 0)
}

func (s *PartialMultiplierContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserNUMBER)
}

func (s *PartialMultiplierContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserNUMBER, i)
}

func (s *PartialMultiplierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterPartialMultiplier(s)
	}
}

func (s *PartialMultiplierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitPartialMultiplier(s)
	}
}

func (s *PartialMultiplierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitPartialMultiplier(s)

	default:
		return t.VisitChildren(s)
	}
}

type FullMultiplierContext struct {
	MultiplierContext
	sets   antlr.Token
	reps   antlr.Token
	weight antlr.Token
}

func NewFullMultiplierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullMultiplierContext {
	var p = new(FullMultiplierContext)

	InitEmptyMultiplierContext(&p.MultiplierContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplierContext))

	return p
}

func (s *FullMultiplierContext) GetSets() antlr.Token { return s.sets }

func (s *FullMultiplierContext) GetReps() antlr.Token { return s.reps }

func (s *FullMultiplierContext) GetWeight() antlr.Token { return s.weight }

func (s *FullMultiplierContext) SetSets(v antlr.Token) { s.sets = v }

func (s *FullMultiplierContext) SetReps(v antlr.Token) { s.reps = v }

func (s *FullMultiplierContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *FullMultiplierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullMultiplierContext) AllX() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserX)
}

func (s *FullMultiplierContext) X(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserX, i)
}

func (s *FullMultiplierContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserNUMBER)
}

func (s *FullMultiplierContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserNUMBER, i)
}

func (s *FullMultiplierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterFullMultiplier(s)
	}
}

func (s *FullMultiplierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitFullMultiplier(s)
	}
}

func (s *FullMultiplierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitFullMultiplier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) Multiplier() (localctx IMultiplierContext) {
	localctx = NewMultiplierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WorkoutParserRULE_multiplier)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPartialMultiplierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*PartialMultiplierContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(WorkoutParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*PartialMultiplierContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFullMultiplierContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*FullMultiplierContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(WorkoutParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*FullMultiplierContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(WorkoutParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)

			var _m = p.Match(WorkoutParserNUMBER)

			localctx.(*FullMultiplierContext).weight = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(WorkoutParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WorkoutParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(WorkoutParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoteContext is an interface to support dynamic dispatch.
type INoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode

	// IsNoteContext differentiates from other interfaces.
	IsNoteContext()
}

type NoteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteContext() *NoteContext {
	var p = new(NoteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_note
	return p
}

func InitEmptyNoteContext(p *NoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutParserRULE_note
}

func (*NoteContext) IsNoteContext() {}

func NewNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteContext {
	var p = new(NoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutParserRULE_note

	return p
}

func (s *NoteContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(WorkoutParserWORD)
}

func (s *NoteContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutParserWORD, i)
}

func (s *NoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.EnterNote(s)
	}
}

func (s *NoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutListener); ok {
		listenerT.ExitNote(s)
	}
}

func (s *NoteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutVisitor:
		return t.VisitNote(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutParser) Note() (localctx INoteContext) {
	localctx = NewNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WorkoutParserRULE_note)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(66)
				p.Match(WorkoutParserWORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
