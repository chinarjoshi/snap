// Code generated from ./Paragraph.g4 by ANTLR 4.13.2. DO NOT EDIT.

package paragraph // Paragraph
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ParagraphParser.
type ParagraphVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ParagraphParser#paragraph.
	VisitParagraph(ctx *ParagraphContext) interface{}

	// Visit a parse tree produced by ParagraphParser#ExerciseLineAlt.
	VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{}

	// Visit a parse tree produced by ParagraphParser#ProseLineAlt.
	VisitProseLineAlt(ctx *ProseLineAltContext) interface{}

	// Visit a parse tree produced by ParagraphParser#exerciseLine.
	VisitExerciseLine(ctx *ExerciseLineContext) interface{}

	// Visit a parse tree produced by ParagraphParser#proseLine.
	VisitProseLine(ctx *ProseLineContext) interface{}

	// Visit a parse tree produced by ParagraphParser#numericToken.
	VisitNumericToken(ctx *NumericTokenContext) interface{}

	// Visit a parse tree produced by ParagraphParser#token.
	VisitToken(ctx *TokenContext) interface{}

	// Visit a parse tree produced by ParagraphParser#TwoPartBy.
	VisitTwoPartBy(ctx *TwoPartByContext) interface{}

	// Visit a parse tree produced by ParagraphParser#ThreePartBy.
	VisitThreePartBy(ctx *ThreePartByContext) interface{}

	// Visit a parse tree produced by ParagraphParser#PartialMultiplier.
	VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{}

	// Visit a parse tree produced by ParagraphParser#FullMultiplier.
	VisitFullMultiplier(ctx *FullMultiplierContext) interface{}

	// Visit a parse tree produced by ParagraphParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ParagraphParser#note.
	VisitNote(ctx *NoteContext) interface{}
}
