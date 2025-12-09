package base

import (
	"github.com/roidaradal/opt/worker"
)

// Base Linear Solver
type LinearSolver struct {
	Solver
	IterationLimit int // if lte 0: no iteration bounds
	SolutionLimit  int // if lte 0: find all solutions
}

// Check if solution limit is reached
func (s LinearSolver) IsSolutionLimitReached() bool {
	limit := s.SolutionLimit
	if limit <= 0 {
		return false
	}
	if s.Problem.SolutionCoreFn != nil {
		return len(s.CoreSolutions) >= limit
	}
	return len(s.BestSolutions) >= limit
}

// Display batch progress and check if iteration limit is reached
func (s LinearSolver) IsIterationLimitReached(logger worker.Logger) bool {
	s.DisplayProgress(logger)
	return s.IterationLimit > 0 && s.NumSteps >= s.IterationLimit
}
