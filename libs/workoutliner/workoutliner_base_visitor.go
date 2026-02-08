// Code generated from Workoutliner.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workoutliner // Workoutliner
import "github.com/antlr4-go/antlr/v4"

type BaseWorkoutlinerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseWorkoutlinerVisitor) VisitParagraph(ctx *ParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitProseLineAlt(ctx *ProseLineAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitExerciseLine(ctx *ExerciseLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitProseLine(ctx *ProseLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitNumericToken(ctx *NumericTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitToken(ctx *TokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitTwoPartBy(ctx *TwoPartByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitThreePartBy(ctx *ThreePartByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitFullMultiplier(ctx *FullMultiplierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWorkoutlinerVisitor) VisitNote(ctx *NoteContext) interface{} {
	return v.VisitChildren(ctx)
}
