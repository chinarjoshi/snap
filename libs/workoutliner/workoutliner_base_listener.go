// Code generated from Workoutliner.g4 by ANTLR 4.13.2. DO NOT EDIT.

package workoutliner // Workoutliner
import "github.com/antlr4-go/antlr/v4"

// BaseWorkoutlinerListener is a complete listener for a parse tree produced by WorkoutlinerParser.
type BaseWorkoutlinerListener struct{}

var _ WorkoutlinerListener = &BaseWorkoutlinerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWorkoutlinerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWorkoutlinerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWorkoutlinerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWorkoutlinerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseWorkoutlinerListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseWorkoutlinerListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterExerciseLineAlt is called when production ExerciseLineAlt is entered.
func (s *BaseWorkoutlinerListener) EnterExerciseLineAlt(ctx *ExerciseLineAltContext) {}

// ExitExerciseLineAlt is called when production ExerciseLineAlt is exited.
func (s *BaseWorkoutlinerListener) ExitExerciseLineAlt(ctx *ExerciseLineAltContext) {}

// EnterProseLineAlt is called when production ProseLineAlt is entered.
func (s *BaseWorkoutlinerListener) EnterProseLineAlt(ctx *ProseLineAltContext) {}

// ExitProseLineAlt is called when production ProseLineAlt is exited.
func (s *BaseWorkoutlinerListener) ExitProseLineAlt(ctx *ProseLineAltContext) {}

// EnterExerciseLine is called when production exerciseLine is entered.
func (s *BaseWorkoutlinerListener) EnterExerciseLine(ctx *ExerciseLineContext) {}

// ExitExerciseLine is called when production exerciseLine is exited.
func (s *BaseWorkoutlinerListener) ExitExerciseLine(ctx *ExerciseLineContext) {}

// EnterProseLine is called when production proseLine is entered.
func (s *BaseWorkoutlinerListener) EnterProseLine(ctx *ProseLineContext) {}

// ExitProseLine is called when production proseLine is exited.
func (s *BaseWorkoutlinerListener) ExitProseLine(ctx *ProseLineContext) {}

// EnterNumericToken is called when production numericToken is entered.
func (s *BaseWorkoutlinerListener) EnterNumericToken(ctx *NumericTokenContext) {}

// ExitNumericToken is called when production numericToken is exited.
func (s *BaseWorkoutlinerListener) ExitNumericToken(ctx *NumericTokenContext) {}

// EnterToken is called when production token is entered.
func (s *BaseWorkoutlinerListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseWorkoutlinerListener) ExitToken(ctx *TokenContext) {}

// EnterTwoPartBy is called when production TwoPartBy is entered.
func (s *BaseWorkoutlinerListener) EnterTwoPartBy(ctx *TwoPartByContext) {}

// ExitTwoPartBy is called when production TwoPartBy is exited.
func (s *BaseWorkoutlinerListener) ExitTwoPartBy(ctx *TwoPartByContext) {}

// EnterThreePartBy is called when production ThreePartBy is entered.
func (s *BaseWorkoutlinerListener) EnterThreePartBy(ctx *ThreePartByContext) {}

// ExitThreePartBy is called when production ThreePartBy is exited.
func (s *BaseWorkoutlinerListener) ExitThreePartBy(ctx *ThreePartByContext) {}

// EnterPartialMultiplier is called when production PartialMultiplier is entered.
func (s *BaseWorkoutlinerListener) EnterPartialMultiplier(ctx *PartialMultiplierContext) {}

// ExitPartialMultiplier is called when production PartialMultiplier is exited.
func (s *BaseWorkoutlinerListener) ExitPartialMultiplier(ctx *PartialMultiplierContext) {}

// EnterFullMultiplier is called when production FullMultiplier is entered.
func (s *BaseWorkoutlinerListener) EnterFullMultiplier(ctx *FullMultiplierContext) {}

// ExitFullMultiplier is called when production FullMultiplier is exited.
func (s *BaseWorkoutlinerListener) ExitFullMultiplier(ctx *FullMultiplierContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseWorkoutlinerListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseWorkoutlinerListener) ExitNumber(ctx *NumberContext) {}

// EnterNote is called when production note is entered.
func (s *BaseWorkoutlinerListener) EnterNote(ctx *NoteContext) {}

// ExitNote is called when production note is exited.
func (s *BaseWorkoutlinerListener) ExitNote(ctx *NoteContext) {}
