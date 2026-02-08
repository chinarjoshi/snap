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

	// Visit a parse tree produced by WorkoutParser#TwoPartBy.
	VisitTwoPartBy(ctx *TwoPartByContext) interface{}

	// Visit a parse tree produced by WorkoutParser#ThreePartBy.
	VisitThreePartBy(ctx *ThreePartByContext) interface{}

	// Visit a parse tree produced by WorkoutParser#PartialMultiplier.
	VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{}

	// Visit a parse tree produced by WorkoutParser#FullMultiplier.
	VisitFullMultiplier(ctx *FullMultiplierContext) interface{}

	// Visit a parse tree produced by WorkoutParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by WorkoutParser#note.
	VisitNote(ctx *NoteContext) interface{}
}
