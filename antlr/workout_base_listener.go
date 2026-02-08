// Code generated from ./Workout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr // Workout
import "github.com/antlr4-go/antlr/v4"

// BaseWorkoutListener is a complete listener for a parse tree produced by WorkoutParser.
type BaseWorkoutListener struct{}

var _ WorkoutListener = &BaseWorkoutListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWorkoutListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWorkoutListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWorkoutListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWorkoutListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWorkoutLog is called when production workoutLog is entered.
func (s *BaseWorkoutListener) EnterWorkoutLog(ctx *WorkoutLogContext) {}

// ExitWorkoutLog is called when production workoutLog is exited.
func (s *BaseWorkoutListener) ExitWorkoutLog(ctx *WorkoutLogContext) {}

// EnterExercise is called when production exercise is entered.
func (s *BaseWorkoutListener) EnterExercise(ctx *ExerciseContext) {}

// ExitExercise is called when production exercise is exited.
func (s *BaseWorkoutListener) ExitExercise(ctx *ExerciseContext) {}

// EnterExerciseName is called when production exerciseName is entered.
func (s *BaseWorkoutListener) EnterExerciseName(ctx *ExerciseNameContext) {}

// ExitExerciseName is called when production exerciseName is exited.
func (s *BaseWorkoutListener) ExitExerciseName(ctx *ExerciseNameContext) {}

// EnterToken is called when production token is entered.
func (s *BaseWorkoutListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseWorkoutListener) ExitToken(ctx *TokenContext) {}

// EnterTwoPartBy is called when production TwoPartBy is entered.
func (s *BaseWorkoutListener) EnterTwoPartBy(ctx *TwoPartByContext) {}

// ExitTwoPartBy is called when production TwoPartBy is exited.
func (s *BaseWorkoutListener) ExitTwoPartBy(ctx *TwoPartByContext) {}

// EnterThreePartBy is called when production ThreePartBy is entered.
func (s *BaseWorkoutListener) EnterThreePartBy(ctx *ThreePartByContext) {}

// ExitThreePartBy is called when production ThreePartBy is exited.
func (s *BaseWorkoutListener) ExitThreePartBy(ctx *ThreePartByContext) {}

// EnterPartialMultiplier is called when production PartialMultiplier is entered.
func (s *BaseWorkoutListener) EnterPartialMultiplier(ctx *PartialMultiplierContext) {}

// ExitPartialMultiplier is called when production PartialMultiplier is exited.
func (s *BaseWorkoutListener) ExitPartialMultiplier(ctx *PartialMultiplierContext) {}

// EnterFullMultiplier is called when production FullMultiplier is entered.
func (s *BaseWorkoutListener) EnterFullMultiplier(ctx *FullMultiplierContext) {}

// ExitFullMultiplier is called when production FullMultiplier is exited.
func (s *BaseWorkoutListener) ExitFullMultiplier(ctx *FullMultiplierContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseWorkoutListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseWorkoutListener) ExitNumber(ctx *NumberContext) {}

// EnterNote is called when production note is entered.
func (s *BaseWorkoutListener) EnterNote(ctx *NoteContext) {}

// ExitNote is called when production note is exited.
func (s *BaseWorkoutListener) ExitNote(ctx *NoteContext) {}
