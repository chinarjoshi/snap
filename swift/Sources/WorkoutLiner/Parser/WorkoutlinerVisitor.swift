// Generated from Workoutliner.g4 by ANTLR 4.13.2
import Antlr4

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link WorkoutlinerParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
open class WorkoutlinerVisitor<T>: ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#paragraph}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitParagraph(_ ctx: WorkoutlinerParser.ParagraphContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by the {@code ExerciseLineAlt}
	 * labeled alternative in {@link WorkoutlinerParser#line}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitExerciseLineAlt(_ ctx: WorkoutlinerParser.ExerciseLineAltContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by the {@code ProseLineAlt}
	 * labeled alternative in {@link WorkoutlinerParser#line}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitProseLineAlt(_ ctx: WorkoutlinerParser.ProseLineAltContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#exerciseLine}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitExerciseLine(_ ctx: WorkoutlinerParser.ExerciseLineContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#proseLine}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitProseLine(_ ctx: WorkoutlinerParser.ProseLineContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#numericToken}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitNumericToken(_ ctx: WorkoutlinerParser.NumericTokenContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#token}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitToken(_ ctx: WorkoutlinerParser.TokenContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by the {@code TwoPartBy}
	 * labeled alternative in {@link WorkoutlinerParser#byExpr}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitTwoPartBy(_ ctx: WorkoutlinerParser.TwoPartByContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by the {@code ThreePartBy}
	 * labeled alternative in {@link WorkoutlinerParser#byExpr}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitThreePartBy(_ ctx: WorkoutlinerParser.ThreePartByContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by the {@code PartialMultiplier}
	 * labeled alternative in {@link WorkoutlinerParser#multiplier}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitPartialMultiplier(_ ctx: WorkoutlinerParser.PartialMultiplierContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by the {@code FullMultiplier}
	 * labeled alternative in {@link WorkoutlinerParser#multiplier}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitFullMultiplier(_ ctx: WorkoutlinerParser.FullMultiplierContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#number}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitNumber(_ ctx: WorkoutlinerParser.NumberContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

	/**
	 * Visit a parse tree produced by {@link WorkoutlinerParser#note}.
	- Parameters:
	  - ctx: the parse tree
	- returns: the visitor result
	 */
	open func visitNote(_ ctx: WorkoutlinerParser.NoteContext) -> T {
	 	fatalError(#function + " must be overridden")
	}

}