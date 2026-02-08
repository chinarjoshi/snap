// Code generated from ./Workout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // Workout
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by WorkoutParser.
type WorkoutVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by WorkoutParser#workoutLog.
	VisitWorkoutLog(ctx *WorkoutLogContext) interface{}

	// Visit a parse tree produced by WorkoutParser#exercise.
	VisitExercise(ctx *ExerciseContext) interface{}

	// Visit a parse tree produced by WorkoutParser#exerciseName.
	VisitExerciseName(ctx *ExerciseNameContext) interface{}

	// Visit a parse tree produced by WorkoutParser#token.
	VisitToken(ctx *TokenContext) interface{}

	// Visit a parse tree produced by WorkoutParser#byExpr.
	VisitByExpr(ctx *ByExprContext) interface{}

	// Visit a parse tree produced by WorkoutParser#multiplier.
	VisitMultiplier(ctx *MultiplierContext) interface{}

	// Visit a parse tree produced by WorkoutParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by WorkoutParser#note.
	VisitNote(ctx *NoteContext) interface{}
}
