package paragraph

import (
	"github.com/antlr4-go/antlr/v4"
)

func Parse(input string) (*ParseResult, error) {
	lexer := NewParagraphLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewParagraphParser(stream)

	p.RemoveErrorListeners()
	errorListener := &errorCollector{}
	p.AddErrorListener(errorListener)

	tree := p.Paragraph()

	if errorListener.hasErrors {
		return nil, errorListener.errors[0]
	}

	visitor := NewParagraphASTVisitor()
	result := visitor.Visit(tree)

	return result.(*ParseResult), nil
}

type errorCollector struct {
	*antlr.DefaultErrorListener
	hasErrors bool
	errors    []error
}

func (e *errorCollector) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	e.hasErrors = true
	e.errors = append(e.errors, &ParseError{Line: line, Column: column, Message: msg})
}

type ParseError struct {
	Line    int
	Column  int
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}
