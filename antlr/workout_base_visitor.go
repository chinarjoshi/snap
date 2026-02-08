// Code generated from ./Workout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // Workout
import "github.com/antlr4-go/antlr/v4"

type BaseWorkoutVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseWorkoutVisitor) VisitWorkoutLog(ctx *WorkoutLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitExercise(ctx *ExerciseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitExerciseName(ctx *ExerciseNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitToken(ctx *TokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitByExpr(ctx *ByExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitMultiplier(ctx *MultiplierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutVisitor) VisitNote(ctx *NoteContext) interface{} {
	return v.VisitChildren(ctx)
}
