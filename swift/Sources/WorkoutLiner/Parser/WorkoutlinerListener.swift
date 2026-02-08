// Generated from Workoutliner.g4 by ANTLR 4.13.2
import Antlr4

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link WorkoutlinerParser}.
 */
public protocol WorkoutlinerListener: ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#paragraph}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterParagraph(_ ctx: WorkoutlinerParser.ParagraphContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#paragraph}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitParagraph(_ ctx: WorkoutlinerParser.ParagraphContext)
	/**
	 * Enter a parse tree produced by the {@code ExerciseLineAlt}
	 * labeled alternative in {@link WorkoutlinerParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterExerciseLineAlt(_ ctx: WorkoutlinerParser.ExerciseLineAltContext)
	/**
	 * Exit a parse tree produced by the {@code ExerciseLineAlt}
	 * labeled alternative in {@link WorkoutlinerParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitExerciseLineAlt(_ ctx: WorkoutlinerParser.ExerciseLineAltContext)
	/**
	 * Enter a parse tree produced by the {@code ProseLineAlt}
	 * labeled alternative in {@link WorkoutlinerParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterProseLineAlt(_ ctx: WorkoutlinerParser.ProseLineAltContext)
	/**
	 * Exit a parse tree produced by the {@code ProseLineAlt}
	 * labeled alternative in {@link WorkoutlinerParser#line}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitProseLineAlt(_ ctx: WorkoutlinerParser.ProseLineAltContext)
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#exerciseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterExerciseLine(_ ctx: WorkoutlinerParser.ExerciseLineContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#exerciseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitExerciseLine(_ ctx: WorkoutlinerParser.ExerciseLineContext)
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#proseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterProseLine(_ ctx: WorkoutlinerParser.ProseLineContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#proseLine}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitProseLine(_ ctx: WorkoutlinerParser.ProseLineContext)
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#numericToken}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterNumericToken(_ ctx: WorkoutlinerParser.NumericTokenContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#numericToken}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitNumericToken(_ ctx: WorkoutlinerParser.NumericTokenContext)
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#token}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterToken(_ ctx: WorkoutlinerParser.TokenContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#token}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitToken(_ ctx: WorkoutlinerParser.TokenContext)
	/**
	 * Enter a parse tree produced by the {@code TwoPartBy}
	 * labeled alternative in {@link WorkoutlinerParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterTwoPartBy(_ ctx: WorkoutlinerParser.TwoPartByContext)
	/**
	 * Exit a parse tree produced by the {@code TwoPartBy}
	 * labeled alternative in {@link WorkoutlinerParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitTwoPartBy(_ ctx: WorkoutlinerParser.TwoPartByContext)
	/**
	 * Enter a parse tree produced by the {@code ThreePartBy}
	 * labeled alternative in {@link WorkoutlinerParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterThreePartBy(_ ctx: WorkoutlinerParser.ThreePartByContext)
	/**
	 * Exit a parse tree produced by the {@code ThreePartBy}
	 * labeled alternative in {@link WorkoutlinerParser#byExpr}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitThreePartBy(_ ctx: WorkoutlinerParser.ThreePartByContext)
	/**
	 * Enter a parse tree produced by the {@code PartialMultiplier}
	 * labeled alternative in {@link WorkoutlinerParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterPartialMultiplier(_ ctx: WorkoutlinerParser.PartialMultiplierContext)
	/**
	 * Exit a parse tree produced by the {@code PartialMultiplier}
	 * labeled alternative in {@link WorkoutlinerParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitPartialMultiplier(_ ctx: WorkoutlinerParser.PartialMultiplierContext)
	/**
	 * Enter a parse tree produced by the {@code FullMultiplier}
	 * labeled alternative in {@link WorkoutlinerParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterFullMultiplier(_ ctx: WorkoutlinerParser.FullMultiplierContext)
	/**
	 * Exit a parse tree produced by the {@code FullMultiplier}
	 * labeled alternative in {@link WorkoutlinerParser#multiplier}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitFullMultiplier(_ ctx: WorkoutlinerParser.FullMultiplierContext)
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#number}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterNumber(_ ctx: WorkoutlinerParser.NumberContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#number}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitNumber(_ ctx: WorkoutlinerParser.NumberContext)
	/**
	 * Enter a parse tree produced by {@link WorkoutlinerParser#note}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func enterNote(_ ctx: WorkoutlinerParser.NoteContext)
	/**
	 * Exit a parse tree produced by {@link WorkoutlinerParser#note}.
	 - Parameters:
	   - ctx: the parse tree
	 */
	func exitNote(_ ctx: WorkoutlinerParser.NoteContext)
}