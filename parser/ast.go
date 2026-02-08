package parser

// Set represents a single set within an exercise
type Set struct {
	Reps   int
	Weight int
	Note   string
}

// Exercise represents a parsed exercise with all its sets
type Exercise struct {
	Name string
	Sets []Set
}

// WorkoutLog represents the complete parsed workout
type WorkoutLog struct {
	Exercises []Exercise
}

// MaxSets returns the maximum number of sets across all exercises
func (w *WorkoutLog) MaxSets() int {
	max := 0
	for _, e := range w.Exercises {
		if len(e.Sets) > max {
			max = len(e.Sets)
		}
	}
	return max
}
