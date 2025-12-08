package base

import (
	"fmt"

	"github.com/roidaradal/opt/worker"
)

// Base Linear Solver
type LinearSolver struct {
	Solver
	IterationLimit int // if lte 0: no iteration bounds
	SolutionLimit  int // if lte 0: find all solutions
}

// LinearSolver's identifier
func (s LinearSolver) FullName() string {
	problemName := ""
	if s.Problem != nil {
		problemName = s.Problem.Name
	}
	return fmt.Sprintf("%s(%s)", s.Name, problemName)
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
	iterationBatch := worker.IterationBatch
	if s.NumSteps%iterationBatch == 0 {
		batch := s.NumSteps / iterationBatch
		if batch > 0 {
			var bestScore string
			if s.Problem.IsOptimization() {
				bestScore = fmt.Sprintf("BestScore: %.2f, %d solutions", s.BestScore, len(s.BestSolutions))
			}
			logger.Output("IterBatch:", batch, bestScore, s.FullName())
		}
	}
	return s.IterationLimit > 0 && s.NumSteps >= s.IterationLimit
}
