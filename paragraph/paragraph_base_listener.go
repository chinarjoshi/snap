// Code generated from ./Paragraph.g4 by ANTLR 4.13.2. DO NOT EDIT.

package paragraph // Paragraph
import "github.com/antlr4-go/antlr/v4"

// BaseParagraphListener is a complete listener for a parse tree produced by ParagraphParser.
type BaseParagraphListener struct{}

var _ ParagraphListener = &BaseParagraphListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParagraphListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParagraphListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParagraphListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParagraphListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseParagraphListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseParagraphListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterExerciseLineAlt is called when production ExerciseLineAlt is entered.
func (s *BaseParagraphListener) EnterExerciseLineAlt(ctx *ExerciseLineAltContext) {}

// ExitExerciseLineAlt is called when production ExerciseLineAlt is exited.
func (s *BaseParagraphListener) ExitExerciseLineAlt(ctx *ExerciseLineAltContext) {}

// EnterProseLineAlt is called when production ProseLineAlt is entered.
func (s *BaseParagraphListener) EnterProseLineAlt(ctx *ProseLineAltContext) {}

// ExitProseLineAlt is called when production ProseLineAlt is exited.
func (s *BaseParagraphListener) ExitProseLineAlt(ctx *ProseLineAltContext) {}

// EnterExerciseLine is called when production exerciseLine is entered.
func (s *BaseParagraphListener) EnterExerciseLine(ctx *ExerciseLineContext) {}

// ExitExerciseLine is called when production exerciseLine is exited.
func (s *BaseParagraphListener) ExitExerciseLine(ctx *ExerciseLineContext) {}

// EnterProseLine is called when production proseLine is entered.
func (s *BaseParagraphListener) EnterProseLine(ctx *ProseLineContext) {}

// ExitProseLine is called when production proseLine is exited.
func (s *BaseParagraphListener) ExitProseLine(ctx *ProseLineContext) {}

// EnterNumericToken is called when production numericToken is entered.
func (s *BaseParagraphListener) EnterNumericToken(ctx *NumericTokenContext) {}

// ExitNumericToken is called when production numericToken is exited.
func (s *BaseParagraphListener) ExitNumericToken(ctx *NumericTokenContext) {}

// EnterToken is called when production token is entered.
func (s *BaseParagraphListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseParagraphListener) ExitToken(ctx *TokenContext) {}

// EnterTwoPartBy is called when production TwoPartBy is entered.
func (s *BaseParagraphListener) EnterTwoPartBy(ctx *TwoPartByContext) {}

// ExitTwoPartBy is called when production TwoPartBy is exited.
func (s *BaseParagraphListener) ExitTwoPartBy(ctx *TwoPartByContext) {}

// EnterThreePartBy is called when production ThreePartBy is entered.
func (s *BaseParagraphListener) EnterThreePartBy(ctx *ThreePartByContext) {}

// ExitThreePartBy is called when production ThreePartBy is exited.
func (s *BaseParagraphListener) ExitThreePartBy(ctx *ThreePartByContext) {}

// EnterPartialMultiplier is called when production PartialMultiplier is entered.
func (s *BaseParagraphListener) EnterPartialMultiplier(ctx *PartialMultiplierContext) {}

// ExitPartialMultiplier is called when production PartialMultiplier is exited.
func (s *BaseParagraphListener) ExitPartialMultiplier(ctx *PartialMultiplierContext) {}

// EnterFullMultiplier is called when production FullMultiplier is entered.
func (s *BaseParagraphListener) EnterFullMultiplier(ctx *FullMultiplierContext) {}

// ExitFullMultiplier is called when production FullMultiplier is exited.
func (s *BaseParagraphListener) ExitFullMultiplier(ctx *FullMultiplierContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseParagraphListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseParagraphListener) ExitNumber(ctx *NumberContext) {}

// EnterNote is called when production note is entered.
func (s *BaseParagraphListener) EnterNote(ctx *NoteContext) {}

// ExitNote is called when production note is exited.
func (s *BaseParagraphListener) ExitNote(ctx *NoteContext) {}
