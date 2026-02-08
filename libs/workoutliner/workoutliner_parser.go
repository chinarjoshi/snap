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
		"", "", "", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "NUMBER", "X", "BY", "WORD", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"paragraph", "line", "exerciseLine", "proseLine", "numericToken", "token",
		"byExpr", "multiplier", "number", "note",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 91, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4, 0,
		22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 30, 8, 1, 1,
		1, 1, 1, 3, 1, 34, 8, 1, 3, 1, 36, 8, 1, 1, 2, 4, 2, 39, 8, 2, 11, 2, 12,
		2, 40, 1, 2, 1, 2, 5, 2, 45, 8, 2, 10, 2, 12, 2, 48, 9, 2, 1, 3, 4, 3,
		51, 8, 3, 11, 3, 12, 3, 52, 1, 4, 1, 4, 1, 4, 3, 4, 58, 8, 4, 1, 5, 1,
		5, 3, 5, 62, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		72, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 82, 8,
		7, 1, 8, 1, 8, 1, 9, 4, 9, 87, 8, 9, 11, 9, 12, 9, 88, 1, 9, 0, 0, 10,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 0, 93, 0, 21, 1, 0, 0, 0, 2, 35,
		1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 50, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10,
		61, 1, 0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 81, 1, 0, 0, 0, 16, 83, 1, 0, 0,
		0, 18, 86, 1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 22, 23,
		1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0,
		25, 26, 5, 0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 29, 3, 4, 2, 0, 28, 30, 5, 5,
		0, 0, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 36, 1, 0, 0, 0, 31, 33,
		3, 6, 3, 0, 32, 34, 5, 5, 0, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0,
		34, 36, 1, 0, 0, 0, 35, 27, 1, 0, 0, 0, 35, 31, 1, 0, 0, 0, 36, 3, 1, 0,
		0, 0, 37, 39, 5, 4, 0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38,
		1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 46, 3, 8, 4, 0,
		43, 45, 3, 10, 5, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1,
		0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 5, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49,
		51, 5, 4, 0, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0,
		0, 52, 53, 1, 0, 0, 0, 53, 7, 1, 0, 0, 0, 54, 58, 3, 12, 6, 0, 55, 58,
		3, 14, 7, 0, 56, 58, 3, 16, 8, 0, 57, 54, 1, 0, 0, 0, 57, 55, 1, 0, 0,
		0, 57, 56, 1, 0, 0, 0, 58, 9, 1, 0, 0, 0, 59, 62, 3, 8, 4, 0, 60, 62, 3,
		18, 9, 0, 61, 59, 1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 11, 1, 0, 0, 0, 63,
		64, 5, 1, 0, 0, 64, 65, 5, 3, 0, 0, 65, 72, 5, 1, 0, 0, 66, 67, 5, 1, 0,
		0, 67, 68, 5, 3, 0, 0, 68, 69, 5, 1, 0, 0, 69, 70, 5, 3, 0, 0, 70, 72,
		5, 1, 0, 0, 71, 63, 1, 0, 0, 0, 71, 66, 1, 0, 0, 0, 72, 13, 1, 0, 0, 0,
		73, 74, 5, 1, 0, 0, 74, 75, 5, 2, 0, 0, 75, 82, 5, 1, 0, 0, 76, 77, 5,
		1, 0, 0, 77, 78, 5, 2, 0, 0, 78, 79, 5, 1, 0, 0, 79, 80, 5, 2, 0, 0, 80,
		82, 5, 1, 0, 0, 81, 73, 1, 0, 0, 0, 81, 76, 1, 0, 0, 0, 82, 15, 1, 0, 0,
		0, 83, 84, 5, 1, 0, 0, 84, 17, 1, 0, 0, 0, 85, 87, 5, 4, 0, 0, 86, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 19, 1, 0, 0, 0, 12, 23, 29, 33, 35, 40, 46, 52, 57, 61, 71, 81, 88,
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
	WorkoutlinerParserWORD    = 4
	WorkoutlinerParserNEWLINE = 5
	WorkoutlinerParserWS      = 6
)

// WorkoutlinerParser rules.
const (
	WorkoutlinerParserRULE_paragraph    = 0
	WorkoutlinerParserRULE_line         = 1
	WorkoutlinerParserRULE_exerciseLine = 2
	WorkoutlinerParserRULE_proseLine    = 3
	WorkoutlinerParserRULE_numericToken = 4
	WorkoutlinerParserRULE_token        = 5
	WorkoutlinerParserRULE_byExpr       = 6
	WorkoutlinerParserRULE_multiplier   = 7
	WorkoutlinerParserRULE_number       = 8
	WorkoutlinerParserRULE_note         = 9
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WorkoutlinerParserWORD {
		{
			p.SetState(20)
			p.Line()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
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

	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExerciseLineAltContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.ExerciseLine()
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkoutlinerParserNEWLINE {
			{
				p.SetState(28)
				p.Match(WorkoutlinerParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewProseLineAltContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.ProseLine()
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WorkoutlinerParserNEWLINE {
			{
				p.SetState(32)
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
	p.EnterRule(localctx, 4, WorkoutlinerParserRULE_exerciseLine)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WorkoutlinerParserWORD {
		{
			p.SetState(37)
			p.Match(WorkoutlinerParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.NumericToken()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Token()
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, WorkoutlinerParserRULE_proseLine)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(49)
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

		p.SetState(52)
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
	p.EnterRule(localctx, 8, WorkoutlinerParserRULE_numericToken)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.ByExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Multiplier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
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
	p.EnterRule(localctx, 10, WorkoutlinerParserRULE_token)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkoutlinerParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.NumericToken()
		}

	case WorkoutlinerParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
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
	p.EnterRule(localctx, 12, WorkoutlinerParserRULE_byExpr)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTwoPartByContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*TwoPartByContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(WorkoutlinerParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)

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
			p.SetState(66)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*ThreePartByContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(WorkoutlinerParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*ThreePartByContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(WorkoutlinerParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)

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
	p.EnterRule(localctx, 14, WorkoutlinerParserRULE_multiplier)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPartialMultiplierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*PartialMultiplierContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(WorkoutlinerParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)

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
			p.SetState(76)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*FullMultiplierContext).sets = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(WorkoutlinerParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)

			var _m = p.Match(WorkoutlinerParserNUMBER)

			localctx.(*FullMultiplierContext).reps = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(WorkoutlinerParserX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)

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
	p.EnterRule(localctx, 16, WorkoutlinerParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
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
	p.EnterRule(localctx, 18, WorkoutlinerParserRULE_note)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(85)
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

		p.SetState(88)
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
