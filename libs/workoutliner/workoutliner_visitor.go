// Code generated from Workoutliner.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workoutliner // Workoutliner
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by WorkoutlinerParser.
type WorkoutlinerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by WorkoutlinerParser#paragraph.
	VisitParagraph(ctx *ParagraphContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#ExerciseLineAlt.
	VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#ProseLineAlt.
	VisitProseLineAlt(ctx *ProseLineAltContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#exerciseLine.
	VisitExerciseLine(ctx *ExerciseLineContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#proseLine.
	VisitProseLine(ctx *ProseLineContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#numericToken.
	VisitNumericToken(ctx *NumericTokenContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#token.
	VisitToken(ctx *TokenContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#TwoPartBy.
	VisitTwoPartBy(ctx *TwoPartByContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#ThreePartBy.
	VisitThreePartBy(ctx *ThreePartByContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#PartialMultiplier.
	VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#FullMultiplier.
	VisitFullMultiplier(ctx *FullMultiplierContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by WorkoutlinerParser#note.
	VisitNote(ctx *NoteContext) interface{}
}
