// Code generated from Workoutliner.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workoutliner

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type WorkoutlinerLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var WorkoutlinerLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func workoutlinerlexerLexerInit() {
	staticData := &WorkoutlinerLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "NUMBER", "X", "BY", "WORD", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"NUMBER", "X", "BY", "WORD", "NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 37, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 4, 0, 15, 8, 0, 11, 0, 12, 0, 16, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 25, 8, 3, 11, 3, 12, 3, 26, 1, 4, 1, 4, 1,
		5, 4, 5, 32, 8, 5, 11, 5, 12, 5, 33, 1, 5, 1, 5, 0, 0, 6, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 1, 0, 6, 1, 0, 48, 57, 2, 0, 88, 88, 120, 120, 2,
		0, 66, 66, 98, 98, 2, 0, 89, 89, 121, 121, 2, 0, 65, 90, 97, 122, 2, 0,
		9, 9, 32, 32, 39, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 1, 14, 1, 0, 0,
		0, 3, 18, 1, 0, 0, 0, 5, 20, 1, 0, 0, 0, 7, 24, 1, 0, 0, 0, 9, 28, 1, 0,
		0, 0, 11, 31, 1, 0, 0, 0, 13, 15, 7, 0, 0, 0, 14, 13, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 2, 1, 0, 0, 0,
		18, 19, 7, 1, 0, 0, 19, 4, 1, 0, 0, 0, 20, 21, 7, 2, 0, 0, 21, 22, 7, 3,
		0, 0, 22, 6, 1, 0, 0, 0, 23, 25, 7, 4, 0, 0, 24, 23, 1, 0, 0, 0, 25, 26,
		1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 8, 1, 0, 0, 0,
		28, 29, 5, 10, 0, 0, 29, 10, 1, 0, 0, 0, 30, 32, 7, 5, 0, 0, 31, 30, 1,
		0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34,
		35, 1, 0, 0, 0, 35, 36, 6, 5, 0, 0, 36, 12, 1, 0, 0, 0, 4, 0, 16, 26, 33,
		1, 6, 0, 0,
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

// WorkoutlinerLexerInit initializes any static state used to implement WorkoutlinerLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewWorkoutlinerLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func WorkoutlinerLexerInit() {
	staticData := &WorkoutlinerLexerLexerStaticData
	staticData.once.Do(workoutlinerlexerLexerInit)
}

// NewWorkoutlinerLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewWorkoutlinerLexer(input antlr.CharStream) *WorkoutlinerLexer {
	WorkoutlinerLexerInit()
	l := new(WorkoutlinerLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &WorkoutlinerLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Workoutliner.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// WorkoutlinerLexer tokens.
const (
	WorkoutlinerLexerNUMBER  = 1
	WorkoutlinerLexerX       = 2
	WorkoutlinerLexerBY      = 3
	WorkoutlinerLexerWORD    = 4
	WorkoutlinerLexerNEWLINE = 5
	WorkoutlinerLexerWS      = 6
)
