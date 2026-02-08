// Code generated from ./Paragraph.g4 by ANTLR 4.13.2. DO NOT EDIT.

package paragraph // Paragraph
import "github.com/antlr4-go/antlr/v4"

type BaseParagraphVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseParagraphVisitor) VisitParagraph(ctx *ParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitExerciseLineAlt(ctx *ExerciseLineAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitProseLineAlt(ctx *ProseLineAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitExerciseLine(ctx *ExerciseLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitProseLine(ctx *ProseLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitNumericToken(ctx *NumericTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitToken(ctx *TokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitTwoPartBy(ctx *TwoPartByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitThreePartBy(ctx *ThreePartByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitPartialMultiplier(ctx *PartialMultiplierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitFullMultiplier(ctx *FullMultiplierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParagraphVisitor) VisitNote(ctx *NoteContext) interface{} {
	return v.VisitChildren(ctx)
}
