package antlr

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/chijoshi/workoutliner/parser"
)

// Parse parses a workout log string and returns a WorkoutLog
func Parse(input string) (*parser.WorkoutLog, error) {
	// Create lexer and parser
	is := antlr.NewInputStream(input)
	lexer := NewWorkoutLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewWorkoutParser(stream)

	// Parse
	tree := p.WorkoutLog()

	// Visit to build AST
	visitor := NewWorkoutASTVisitor()
	result := visitor.Visit(tree)

	return result.(*parser.WorkoutLog), nil
}
