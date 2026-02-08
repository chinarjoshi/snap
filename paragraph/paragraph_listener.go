// Code generated from ./Paragraph.g4 by ANTLR 4.13.2. DO NOT EDIT.

package paragraph // Paragraph
import "github.com/antlr4-go/antlr/v4"

// ParagraphListener is a complete listener for a parse tree produced by ParagraphParser.
type ParagraphListener interface {
	antlr.ParseTreeListener

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterExerciseLineAlt is called when entering the ExerciseLineAlt production.
	EnterExerciseLineAlt(c *ExerciseLineAltContext)

	// EnterProseLineAlt is called when entering the ProseLineAlt production.
	EnterProseLineAlt(c *ProseLineAltContext)

	// EnterExerciseLine is called when entering the exerciseLine production.
	EnterExerciseLine(c *ExerciseLineContext)

	// EnterProseLine is called when entering the proseLine production.
	EnterProseLine(c *ProseLineContext)

	// EnterNumericToken is called when entering the numericToken production.
	EnterNumericToken(c *NumericTokenContext)

	// EnterToken is called when entering the token production.
	EnterToken(c *TokenContext)

	// EnterTwoPartBy is called when entering the TwoPartBy production.
	EnterTwoPartBy(c *TwoPartByContext)

	// EnterThreePartBy is called when entering the ThreePartBy production.
	EnterThreePartBy(c *ThreePartByContext)

	// EnterPartialMultiplier is called when entering the PartialMultiplier production.
	EnterPartialMultiplier(c *PartialMultiplierContext)

	// EnterFullMultiplier is called when entering the FullMultiplier production.
	EnterFullMultiplier(c *FullMultiplierContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterNote is called when entering the note production.
	EnterNote(c *NoteContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitExerciseLineAlt is called when exiting the ExerciseLineAlt production.
	ExitExerciseLineAlt(c *ExerciseLineAltContext)

	// ExitProseLineAlt is called when exiting the ProseLineAlt production.
	ExitProseLineAlt(c *ProseLineAltContext)

	// ExitExerciseLine is called when exiting the exerciseLine production.
	ExitExerciseLine(c *ExerciseLineContext)

	// ExitProseLine is called when exiting the proseLine production.
	ExitProseLine(c *ProseLineContext)

	// ExitNumericToken is called when exiting the numericToken production.
	ExitNumericToken(c *NumericTokenContext)

	// ExitToken is called when exiting the token production.
	ExitToken(c *TokenContext)

	// ExitTwoPartBy is called when exiting the TwoPartBy production.
	ExitTwoPartBy(c *TwoPartByContext)

	// ExitThreePartBy is called when exiting the ThreePartBy production.
	ExitThreePartBy(c *ThreePartByContext)

	// ExitPartialMultiplier is called when exiting the PartialMultiplier production.
	ExitPartialMultiplier(c *PartialMultiplierContext)

	// ExitFullMultiplier is called when exiting the FullMultiplier production.
	ExitFullMultiplier(c *FullMultiplierContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitNote is called when exiting the note production.
	ExitNote(c *NoteContext)
}
