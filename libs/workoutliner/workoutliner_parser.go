// Code generated from Workoutliner.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workoutliner // Workoutliner
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

type WorkoutlinerParser struct {
	*antlr.BaseParser
}

var WorkoutlinerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func workoutlinerParserInit() {
	staticData := &WorkoutlinerParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'*'", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "NUMBER", "X", "BY", "STAR", "WORD", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"paragraph", "line", "continuationLine", "supersetLine", "supersetName",
		"exerciseLine", "proseLine", "numericToken", "token", "byExpr", "multiplier",
		"number", "note",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 126, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 4, 0, 28, 8, 0, 11, 0, 12, 0, 29,
		1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1,
		1, 1, 1, 3, 1, 44, 8, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 3, 1, 50, 8, 1, 1,
		2, 1, 2, 5, 2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 3, 1, 3, 1, 3, 4, 3,
		62, 8, 3, 11, 3, 12, 3, 63, 1, 4, 4, 4, 67, 8, 4, 11, 4, 12, 4, 68, 1,
		4, 1, 4, 1, 5, 4, 5, 74, 8, 5, 11, 5, 12, 5, 75, 1, 5, 1, 5, 5, 5, 80,
		8, 5, 10, 5, 12, 5, 83, 9, 5, 1, 6, 4, 6, 86, 8, 6, 11, 6, 12, 6, 87, 1,
		7, 1, 7, 1, 7, 3, 7, 93, 8, 7, 1, 8, 1, 8, 3, 8, 97, 8, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 107, 8, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 117, 8, 10, 1, 11, 1, 11, 1,
		12, 4, 12, 122, 8, 12, 11, 12, 12, 12, 123, 1, 12, 0, 0, 13, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 0, 132, 0, 27, 1, 0, 0, 0, 2, 49,
		1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 66, 1, 0, 0, 0, 10,
		73, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 92, 1, 0, 0, 0, 16, 96, 1, 0, 0,
		0, 18, 106, 1, 0, 0, 0, 20, 116, 1, 0, 0, 0, 22, 118, 1, 0, 0, 0, 24, 121,
		1, 0, 0, 0, 26, 28, 3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0,
		29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 5,
		0, 0, 1, 32, 1, 1, 0, 0, 0, 33, 35, 3, 6, 3, 0, 34, 36, 5, 6, 0, 0, 35,
		34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 50, 1, 0, 0, 0, 37, 39, 3, 10,
		5, 0, 38, 40, 5, 6, 0, 0, 39, 38, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 50,
		1, 0, 0, 0, 41, 43, 3, 4, 2, 0, 42, 44, 5, 6, 0, 0, 43, 42, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 50, 1, 0, 0, 0, 45, 47, 3, 12, 6, 0, 46, 48, 5,
		6, 0, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49,
		33, 1, 0, 0, 0, 49, 37, 1, 0, 0, 0, 49, 41, 1, 0, 0, 0, 49, 45, 1, 0, 0,
		0, 50, 3, 1, 0, 0, 0, 51, 55, 3, 14, 7, 0, 52, 54, 3, 16, 8, 0, 53, 52,
		1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 5, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 3, 8, 4, 0, 59, 61, 3, 8,
		4, 0, 60, 62, 3, 16, 8, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63,
		61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 7, 1, 0, 0, 0, 65, 67, 5, 5, 0,
		0, 66, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 5, 4, 0, 0, 71, 9, 1, 0, 0, 0,
		72, 74, 5, 5, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1,
		0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 81, 3, 14, 7, 0, 78,
		80, 3, 16, 8, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0,
		0, 0, 81, 82, 1, 0, 0, 0, 82, 11, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86,
		5, 5, 0, 0, 85, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 13, 1, 0, 0, 0, 89, 93, 3, 18, 9, 0, 90, 93, 3,
		20, 10, 0, 91, 93, 3, 22, 11, 0, 92, 89, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 91, 1, 0, 0, 0, 93, 15, 1, 0, 0, 0, 94, 97, 3, 14, 7, 0, 95, 97, 3,
		24, 12, 0, 96, 94, 1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 17, 1, 0, 0, 0,
		98, 99, 5, 1, 0, 0, 99, 100, 5, 3, 0, 0, 100, 107, 5, 1, 0, 0, 101, 102,
		5, 1, 0, 0, 102, 103, 5, 3, 0, 0, 103, 104, 5, 1, 0, 0, 104, 105, 5, 3,
		0, 0, 105, 107, 5, 1, 0, 0, 106, 98, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0,
		107, 19, 1, 0, 0, 0, 108, 109, 5, 1, 0, 0, 109, 110, 5, 2, 0, 0, 110, 117,
		5, 1, 0, 0, 111, 112, 5, 1, 0, 0, 112, 113, 5, 2, 0, 0, 113, 114, 5, 1,
		0, 0, 114, 115, 5, 2, 0, 0, 115, 117, 5, 1, 0, 0, 116, 108, 1, 0, 0, 0,
		116, 111, 1, 0, 0, 0, 117, 21, 1, 0, 0, 0, 118, 119, 5, 1, 0, 0, 119, 23,
		1, 0, 0, 0, 120, 122, 5, 5, 0, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 25, 1, 0, 0, 0,
		17, 29, 35, 39, 43, 47, 49, 55, 63, 68, 75, 81, 87, 92, 96, 106, 116, 123,
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

// WorkoutlinerParserInit initializes any static state used to implement WorkoutlinerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWorkoutlinerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WorkoutlinerParserInit() {
	staticData := &WorkoutlinerParserStaticData
	staticData.once.Do(workoutlinerParserInit)
}

// NewWorkoutlinerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWorkoutlinerParser(input antlr.TokenStream) *WorkoutlinerParser {
	WorkoutlinerParserInit()
	this := new(WorkoutlinerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &WorkoutlinerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Workoutliner.g4"

	return this
}

// WorkoutlinerParser tokens.
const (
	WorkoutlinerParserEOF     = antlr.TokenEOF
	WorkoutlinerParserNUMBER  = 1
	WorkoutlinerParserX       = 2
	WorkoutlinerParserBY      = 3
	WorkoutlinerParserSTAR    = 4
	WorkoutlinerParserWORD    = 5
	WorkoutlinerParserNEWLINE = 6
	WorkoutlinerParserWS      = 7
)

// WorkoutlinerParser rules.
const (
	WorkoutlinerParserRULE_paragraph        = 0
	WorkoutlinerParserRULE_line             = 1
	WorkoutlinerParserRULE_continuationLine = 2
	WorkoutlinerParserRULE_supersetLine     = 3
	WorkoutlinerParserRULE_supersetName     = 4
	WorkoutlinerParserRULE_exerciseLine     = 5
	WorkoutlinerParserRULE_proseLine        = 6
	WorkoutlinerParserRULE_numericToken     = 7
	WorkoutlinerParserRULE_token            = 8
	WorkoutlinerParserRULE_byExpr           = 9
	WorkoutlinerParserRULE_multiplier       = 10
	WorkoutlinerParserRULE_number           = 11
	WorkoutlinerParserRULE_note             = 12
)

// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext

	// IsParagraphContext differentiates from other interfaces.
	IsParagraphContext()
}

type ParagraphContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParagraphContext() *ParagraphContext {
	var p = new(ParagraphContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_paragraph
	return p
}

func InitEmptyParagraphContext(p *ParagraphContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_paragraph
}

func (*ParagraphContext) IsParagraphContext() {}

func NewParagraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParagraphContext {
	var p = new(ParagraphContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_paragraph

	return p
}

func (s *ParagraphContext) GetParser() antlr.Parser { return s.parser }

func (s *ParagraphContext) EOF() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserEOF, 0)
}

func (s *ParagraphContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ParagraphContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
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

	return t.(ILineContext)
}

func (s *ParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParagraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterParagraph(s)
	}
}

func (s *ParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitParagraph(s)
	}
}

func (s *ParagraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitParagraph(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) Paragraph() (localctx IParagraphContext) {
	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WorkoutlinerParserRULE_paragraph)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WorkoutlinerParserNUMBER || _la == WorkoutlinerParserWORD {
		{
			p.SetState(26)
			p.Line()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(WorkoutlinerParserEOF)
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

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) CopyAll(ctx *LineContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExerciseLineAltContext struct {
	LineContext
}

func NewExerciseLineAltContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExerciseLineAltContext {
	var p = new(ExerciseLineAltContext)

	InitEmptyLineContext(&p.LineContext)
	p.parser = parser
	p.CopyAll(ctx.(*LineContext))

	return p
}

func (s *ExerciseLineAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExerciseLineAltContext) ExerciseLine() IExerciseLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExerciseLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExerciseLineContext)
}

func (s *ExerciseLineAltContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNEWLINE, 0)
}

func (s *ExerciseLineAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterExerciseLineAlt(s)
	}
}

func (s *ExerciseLineAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitExerciseLineAlt(s)
	}
}

func (s *ExerciseLineAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitExerciseLineAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinuationLineAltContext struct {
	LineContext
}

func NewContinuationLineAltContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinuationLineAltContext {
	var p = new(ContinuationLineAltContext)

	InitEmptyLineContext(&p.LineContext)
	p.parser = parser
	p.CopyAll(ctx.(*LineContext))

	return p
}

func (s *ContinuationLineAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuationLineAltContext) ContinuationLine() IContinuationLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuationLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuationLineContext)
}

func (s *ContinuationLineAltContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNEWLINE, 0)
}

func (s *ContinuationLineAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterContinuationLineAlt(s)
	}
}

func (s *ContinuationLineAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitContinuationLineAlt(s)
	}
}

func (s *ContinuationLineAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitContinuationLineAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SupersetLineAltContext struct {
	LineContext
}

func NewSupersetLineAltContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SupersetLineAltContext {
	var p = new(SupersetLineAltContext)

	InitEmptyLineContext(&p.LineContext)
	p.parser = parser
	p.CopyAll(ctx.(*LineContext))

	return p
}

func (s *SupersetLineAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SupersetLineAltContext) SupersetLine() ISupersetLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISupersetLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISupersetLineContext)
}

func (s *SupersetLineAltContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNEWLINE, 0)
}

func (s *SupersetLineAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterSupersetLineAlt(s)
	}
}

func (s *SupersetLineAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitSupersetLineAlt(s)
	}
}

func (s *SupersetLineAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitSupersetLineAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ProseLineAltContext struct {
	LineContext
}

func NewProseLineAltContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProseLineAltContext {
	var p = new(ProseLineAltContext)

	InitEmptyLineContext(&p.LineContext)
	p.parser = parser
	p.CopyAll(ctx.(*LineContext))

	return p
}

func (s *ProseLineAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProseLineAltContext) ProseLine() IProseLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProseLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProseLineContext)
}

func (s *ProseLineAltContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNEWLINE, 0)
}

func (s *ProseLineAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterProseLineAlt(s)
	}
}

func (s *ProseLineAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitProseLineAlt(s)
	}
}

func (s *ProseLineAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitProseLineAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WorkoutlinerParserRULE_line)
	var _la int

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSupersetLineAltContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.SupersetLine()
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkoutlinerParserNEWLINE {
			{
				p.SetState(34)
				p.Match(WorkoutlinerParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewExerciseLineAltContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.ExerciseLine()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkoutlinerParserNEWLINE {
			{
				p.SetState(38)
				p.Match(WorkoutlinerParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewContinuationLineAltContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)
			p.ContinuationLine()
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkoutlinerParserNEWLINE {
			{
				p.SetState(42)
				p.Match(WorkoutlinerParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		localctx = NewProseLineAltContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.ProseLine()
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkoutlinerParserNEWLINE {
			{
				p.SetState(46)
				p.Match(WorkoutlinerParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// IContinuationLineContext is an interface to support dynamic dispatch.
type IContinuationLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericToken() INumericTokenContext
	AllToken() []ITokenContext
	Token(i int) ITokenContext

	// IsContinuationLineContext differentiates from other interfaces.
	IsContinuationLineContext()
}

type ContinuationLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinuationLineContext() *ContinuationLineContext {
	var p = new(ContinuationLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_continuationLine
	return p
}

func InitEmptyContinuationLineContext(p *ContinuationLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_continuationLine
}

func (*ContinuationLineContext) IsContinuationLineContext() {}

func NewContinuationLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuationLineContext {
	var p = new(ContinuationLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_continuationLine

	return p
}

func (s *ContinuationLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuationLineContext) NumericToken() INumericTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericTokenContext)
}

func (s *ContinuationLineContext) AllToken() []ITokenContext {
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

func (s *ContinuationLineContext) Token(i int) ITokenContext {
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

func (s *ContinuationLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuationLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuationLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterContinuationLine(s)
	}
}

func (s *ContinuationLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitContinuationLine(s)
	}
}

func (s *ContinuationLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitContinuationLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) ContinuationLine() (localctx IContinuationLineContext) {
	localctx = NewContinuationLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WorkoutlinerParserRULE_continuationLine)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.NumericToken()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(52)
				p.Token()
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
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

// ISupersetLineContext is an interface to support dynamic dispatch.
type ISupersetLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSupersetName() []ISupersetNameContext
	SupersetName(i int) ISupersetNameContext
	AllToken() []ITokenContext
	Token(i int) ITokenContext

	// IsSupersetLineContext differentiates from other interfaces.
	IsSupersetLineContext()
}

type SupersetLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySupersetLineContext() *SupersetLineContext {
	var p = new(SupersetLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_supersetLine
	return p
}

func InitEmptySupersetLineContext(p *SupersetLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_supersetLine
}

func (*SupersetLineContext) IsSupersetLineContext() {}

func NewSupersetLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SupersetLineContext {
	var p = new(SupersetLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_supersetLine

	return p
}

func (s *SupersetLineContext) GetParser() antlr.Parser { return s.parser }

func (s *SupersetLineContext) AllSupersetName() []ISupersetNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISupersetNameContext); ok {
			len++
		}
	}

	tst := make([]ISupersetNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISupersetNameContext); ok {
			tst[i] = t.(ISupersetNameContext)
			i++
		}
	}

	return tst
}

func (s *SupersetLineContext) SupersetName(i int) ISupersetNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISupersetNameContext); ok {
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

	return t.(ISupersetNameContext)
}

func (s *SupersetLineContext) AllToken() []ITokenContext {
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

func (s *SupersetLineContext) Token(i int) ITokenContext {
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

func (s *SupersetLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SupersetLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SupersetLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterSupersetLine(s)
	}
}

func (s *SupersetLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitSupersetLine(s)
	}
}

func (s *SupersetLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitSupersetLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) SupersetLine() (localctx ISupersetLineContext) {
	localctx = NewSupersetLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WorkoutlinerParserRULE_supersetLine)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.SupersetName()
	}
	{
		p.SetState(59)
		p.SupersetName()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(60)
				p.Token()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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

// ISupersetNameContext is an interface to support dynamic dispatch.
type ISupersetNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode

	// IsSupersetNameContext differentiates from other interfaces.
	IsSupersetNameContext()
}

type SupersetNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySupersetNameContext() *SupersetNameContext {
	var p = new(SupersetNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_supersetName
	return p
}

func InitEmptySupersetNameContext(p *SupersetNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_supersetName
}

func (*SupersetNameContext) IsSupersetNameContext() {}

func NewSupersetNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SupersetNameContext {
	var p = new(SupersetNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_supersetName

	return p
}

func (s *SupersetNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SupersetNameContext) STAR() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserSTAR, 0)
}

func (s *SupersetNameContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserWORD)
}

func (s *SupersetNameContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserWORD, i)
}

func (s *SupersetNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SupersetNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SupersetNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterSupersetName(s)
	}
}

func (s *SupersetNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitSupersetName(s)
	}
}

func (s *SupersetNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitSupersetName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) SupersetName() (localctx ISupersetNameContext) {
	localctx = NewSupersetNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WorkoutlinerParserRULE_supersetName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WorkoutlinerParserWORD {
		{
			p.SetState(65)
			p.Match(WorkoutlinerParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(WorkoutlinerParserSTAR)
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

// IExerciseLineContext is an interface to support dynamic dispatch.
type IExerciseLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericToken() INumericTokenContext
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllToken() []ITokenContext
	Token(i int) ITokenContext

	// IsExerciseLineContext differentiates from other interfaces.
	IsExerciseLineContext()
}

type ExerciseLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExerciseLineContext() *ExerciseLineContext {
	var p = new(ExerciseLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_exerciseLine
	return p
}

func InitEmptyExerciseLineContext(p *ExerciseLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_exerciseLine
}

func (*ExerciseLineContext) IsExerciseLineContext() {}

func NewExerciseLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExerciseLineContext {
	var p = new(ExerciseLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_exerciseLine

	return p
}

func (s *ExerciseLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ExerciseLineContext) NumericToken() INumericTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericTokenContext)
}

func (s *ExerciseLineContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserWORD)
}

func (s *ExerciseLineContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserWORD, i)
}

func (s *ExerciseLineContext) AllToken() []ITokenContext {
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

func (s *ExerciseLineContext) Token(i int) ITokenContext {
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

func (s *ExerciseLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExerciseLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExerciseLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterExerciseLine(s)
	}
}

func (s *ExerciseLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitExerciseLine(s)
	}
}

func (s *ExerciseLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitExerciseLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) ExerciseLine() (localctx IExerciseLineContext) {
	localctx = NewExerciseLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WorkoutlinerParserRULE_exerciseLine)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WorkoutlinerParserWORD {
		{
			p.SetState(72)
			p.Match(WorkoutlinerParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.NumericToken()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(78)
				p.Token()
			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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

// IProseLineContext is an interface to support dynamic dispatch.
type IProseLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode

	// IsProseLineContext differentiates from other interfaces.
	IsProseLineContext()
}

type ProseLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProseLineContext() *ProseLineContext {
	var p = new(ProseLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_proseLine
	return p
}

func InitEmptyProseLineContext(p *ProseLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_proseLine
}

func (*ProseLineContext) IsProseLineContext() {}

func NewProseLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProseLineContext {
	var p = new(ProseLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_proseLine

	return p
}

func (s *ProseLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ProseLineContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserWORD)
}

func (s *ProseLineContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserWORD, i)
}

func (s *ProseLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProseLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProseLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterProseLine(s)
	}
}

func (s *ProseLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitProseLine(s)
	}
}

func (s *ProseLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitProseLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) ProseLine() (localctx IProseLineContext) {
	localctx = NewProseLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WorkoutlinerParserRULE_proseLine)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(84)
				p.Match(WorkoutlinerParserWORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

// INumericTokenContext is an interface to support dynamic dispatch.
type INumericTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ByExpr() IByExprContext
	Multiplier() IMultiplierContext
	Number() INumberContext

	// IsNumericTokenContext differentiates from other interfaces.
	IsNumericTokenContext()
}

type NumericTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericTokenContext() *NumericTokenContext {
	var p = new(NumericTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_numericToken
	return p
}

func InitEmptyNumericTokenContext(p *NumericTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_numericToken
}

func (*NumericTokenContext) IsNumericTokenContext() {}

func NewNumericTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericTokenContext {
	var p = new(NumericTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_numericToken

	return p
}

func (s *NumericTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericTokenContext) ByExpr() IByExprContext {
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

func (s *NumericTokenContext) Multiplier() IMultiplierContext {
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

func (s *NumericTokenContext) Number() INumberContext {
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

func (s *NumericTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterNumericToken(s)
	}
}

func (s *NumericTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitNumericToken(s)
	}
}

func (s *NumericTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitNumericToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) NumericToken() (localctx INumericTokenContext) {
	localctx = NewNumericTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WorkoutlinerParserRULE_numericToken)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.ByExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Multiplier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Number()
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

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericToken() INumericTokenContext
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
	p.RuleIndex = WorkoutlinerParserRULE_token
	return p
}

func InitEmptyTokenContext(p *TokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_token
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) NumericToken() INumericTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericTokenContext)
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
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitToken(s)
	}
}

func (s *TokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) Token() (localctx ITokenContext) {
	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, WorkoutlinerParserRULE_token)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkoutlinerParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.NumericToken()
		}

	case WorkoutlinerParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Note()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = WorkoutlinerParserRULE_byExpr
	return p
}

func InitEmptyByExprContext(p *ByExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_byExpr
}

func (*ByExprContext) IsByExprContext() {}

func NewByExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByExprContext {
	var p = new(ByExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_byExpr

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
	return s.GetToken(WorkoutlinerParserBY, 0)
}

func (s *TwoPartByContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserNUMBER)
}

func (s *TwoPartByContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNUMBER, i)
}

func (s *TwoPartByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterTwoPartBy(s)
	}
}

func (s *TwoPartByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitTwoPartBy(s)
	}
}

func (s *TwoPartByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
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
	return s.GetTokens(WorkoutlinerParserBY)
}

func (s *ThreePartByContext) BY(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserBY, i)
}

func (s *ThreePartByContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserNUMBER)
}

func (s *ThreePartByContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNUMBER, i)
}

func (s *ThreePartByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterThreePartBy(s)
	}
}

func (s *ThreePartByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitThreePartBy(s)
	}
}

func (s *ThreePartByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitThreePartBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) ByExpr() (localctx IByExprContext) {
	localctx = NewByExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, WorkoutlinerParserRULE_byExpr)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTwoPartByContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*TwoPartByContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(WorkoutlinerParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)

			var _m = p.Match(WorkoutlinerParserNUMBER)

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
			p.SetState(101)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*ThreePartByContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(WorkoutlinerParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*ThreePartByContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(WorkoutlinerParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _m = p.Match(WorkoutlinerParserNUMBER)

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
	p.RuleIndex = WorkoutlinerParserRULE_multiplier
	return p
}

func InitEmptyMultiplierContext(p *MultiplierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_multiplier
}

func (*MultiplierContext) IsMultiplierContext() {}

func NewMultiplierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplierContext {
	var p = new(MultiplierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_multiplier

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
	return s.GetToken(WorkoutlinerParserX, 0)
}

func (s *PartialMultiplierContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserNUMBER)
}

func (s *PartialMultiplierContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNUMBER, i)
}

func (s *PartialMultiplierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterPartialMultiplier(s)
	}
}

func (s *PartialMultiplierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitPartialMultiplier(s)
	}
}

func (s *PartialMultiplierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
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
	return s.GetTokens(WorkoutlinerParserX)
}

func (s *FullMultiplierContext) X(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserX, i)
}

func (s *FullMultiplierContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserNUMBER)
}

func (s *FullMultiplierContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNUMBER, i)
}

func (s *FullMultiplierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterFullMultiplier(s)
	}
}

func (s *FullMultiplierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitFullMultiplier(s)
	}
}

func (s *FullMultiplierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitFullMultiplier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) Multiplier() (localctx IMultiplierContext) {
	localctx = NewMultiplierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, WorkoutlinerParserRULE_multiplier)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPartialMultiplierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*PartialMultiplierContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(WorkoutlinerParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)

			var _m = p.Match(WorkoutlinerParserNUMBER)

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
			p.SetState(111)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*FullMultiplierContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(WorkoutlinerParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*FullMultiplierContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(WorkoutlinerParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)

			var _m = p.Match(WorkoutlinerParserNUMBER)

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
	p.RuleIndex = WorkoutlinerParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, WorkoutlinerParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(WorkoutlinerParserNUMBER)
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
	p.RuleIndex = WorkoutlinerParserRULE_note
	return p
}

func InitEmptyNoteContext(p *NoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkoutlinerParserRULE_note
}

func (*NoteContext) IsNoteContext() {}

func NewNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteContext {
	var p = new(NoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkoutlinerParserRULE_note

	return p
}

func (s *NoteContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(WorkoutlinerParserWORD)
}

func (s *NoteContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(WorkoutlinerParserWORD, i)
}

func (s *NoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.EnterNote(s)
	}
}

func (s *NoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkoutlinerListener); ok {
		listenerT.ExitNote(s)
	}
}

func (s *NoteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WorkoutlinerVisitor:
		return t.VisitNote(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WorkoutlinerParser) Note() (localctx INoteContext) {
	localctx = NewNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, WorkoutlinerParserRULE_note)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(120)
				p.Match(WorkoutlinerParserWORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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
