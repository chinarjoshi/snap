package paragraph

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
