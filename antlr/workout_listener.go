// Code generated from ./Workout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // Workout
import "github.com/antlr4-go/antlr/v4"

// WorkoutListener is a complete listener for a parse tree produced by WorkoutParser.
type WorkoutListener interface {
	antlr.ParseTreeListener

	// EnterWorkoutLog is called when entering the workoutLog production.
	EnterWorkoutLog(c *WorkoutLogContext)

	// EnterExercise is called when entering the exercise production.
	EnterExercise(c *ExerciseContext)

	// EnterExerciseName is called when entering the exerciseName production.
	EnterExerciseName(c *ExerciseNameContext)

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

	// ExitWorkoutLog is called when exiting the workoutLog production.
	ExitWorkoutLog(c *WorkoutLogContext)

	// ExitExercise is called when exiting the exercise production.
	ExitExercise(c *ExerciseContext)

	// ExitExerciseName is called when exiting the exerciseName production.
	ExitExerciseName(c *ExerciseNameContext)

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
