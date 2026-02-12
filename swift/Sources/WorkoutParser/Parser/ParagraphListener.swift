// Generated from ./Paragraph.g4 by ANTLR 4.13.2
import Antlr4

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ParagraphParser}.
 */
public protocol ParagraphListener: ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#paragraph}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterParagraph(_ ctx: ParagraphParser.ParagraphContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#paragraph}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitParagraph(_ ctx: ParagraphParser.ParagraphContext)
	/**
	 * Enter a parse tree produced by the {@code ExerciseLineAlt}
	 * labeled alternative in {@link ParagraphParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterExerciseLineAlt(_ ctx: ParagraphParser.ExerciseLineAltContext)
	/**
	 * Exit a parse tree produced by the {@code ExerciseLineAlt}
	 * labeled alternative in {@link ParagraphParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitExerciseLineAlt(_ ctx: ParagraphParser.ExerciseLineAltContext)
	/**
	 * Enter a parse tree produced by the {@code ProseLineAlt}
	 * labeled alternative in {@link ParagraphParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterProseLineAlt(_ ctx: ParagraphParser.ProseLineAltContext)
	/**
	 * Exit a parse tree produced by the {@code ProseLineAlt}
	 * labeled alternative in {@link ParagraphParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitProseLineAlt(_ ctx: ParagraphParser.ProseLineAltContext)
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#exerciseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterExerciseLine(_ ctx: ParagraphParser.ExerciseLineContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#exerciseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitExerciseLine(_ ctx: ParagraphParser.ExerciseLineContext)
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#proseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterProseLine(_ ctx: ParagraphParser.ProseLineContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#proseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitProseLine(_ ctx: ParagraphParser.ProseLineContext)
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#numericToken}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterNumericToken(_ ctx: ParagraphParser.NumericTokenContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#numericToken}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitNumericToken(_ ctx: ParagraphParser.NumericTokenContext)
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#token}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterToken(_ ctx: ParagraphParser.TokenContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#token}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitToken(_ ctx: ParagraphParser.TokenContext)
	/**
	 * Enter a parse tree produced by the {@code TwoPartBy}
	 * labeled alternative in {@link ParagraphParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterTwoPartBy(_ ctx: ParagraphParser.TwoPartByContext)
	/**
	 * Exit a parse tree produced by the {@code TwoPartBy}
	 * labeled alternative in {@link ParagraphParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitTwoPartBy(_ ctx: ParagraphParser.TwoPartByContext)
	/**
	 * Enter a parse tree produced by the {@code ThreePartBy}
	 * labeled alternative in {@link ParagraphParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterThreePartBy(_ ctx: ParagraphParser.ThreePartByContext)
	/**
	 * Exit a parse tree produced by the {@code ThreePartBy}
	 * labeled alternative in {@link ParagraphParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitThreePartBy(_ ctx: ParagraphParser.ThreePartByContext)
	/**
	 * Enter a parse tree produced by the {@code PartialMultiplier}
	 * labeled alternative in {@link ParagraphParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterPartialMultiplier(_ ctx: ParagraphParser.PartialMultiplierContext)
	/**
	 * Exit a parse tree produced by the {@code PartialMultiplier}
	 * labeled alternative in {@link ParagraphParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitPartialMultiplier(_ ctx: ParagraphParser.PartialMultiplierContext)
	/**
	 * Enter a parse tree produced by the {@code FullMultiplier}
	 * labeled alternative in {@link ParagraphParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterFullMultiplier(_ ctx: ParagraphParser.FullMultiplierContext)
	/**
	 * Exit a parse tree produced by the {@code FullMultiplier}
	 * labeled alternative in {@link ParagraphParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitFullMultiplier(_ ctx: ParagraphParser.FullMultiplierContext)
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#number}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterNumber(_ ctx: ParagraphParser.NumberContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#number}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitNumber(_ ctx: ParagraphParser.NumberContext)
	/**
	 * Enter a parse tree produced by {@link ParagraphParser#note}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterNote(_ ctx: ParagraphParser.NoteContext)
	/**
	 * Exit a parse tree produced by {@link ParagraphParser#note}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitNote(_ ctx: ParagraphParser.NoteContext)
}