// Package base contains base solver types
package base

import (
	"fmt"
	"math"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/worker"
)

// Base Solver
type Solver struct {
	Name    string
	Problem *discrete.Problem
	*worker.Result
}

// Initialize Solver
func (s *Solver) Initialize(name string, problem *discrete.Problem) {
	s.Name = name
	s.Problem = problem
	s.Result = &worker.Result{
		BestSolutions:     make([]*discrete.Solution, 0),
		CoreSolutions:     make(map[string][]*discrete.Solution),
		FeasibleSolutions: make(dict.Counter[discrete.Score]),
	}
	switch problem.Goal {
	case discrete.Maximize:
		s.BestScore = math.Inf(-1)
	case discrete.Minimize:
		s.BestScore = math.Inf(1)
	case discrete.Satisfy:
		s.BestScore = 0
	}
}

// Get solver's name
func (s Solver) GetName() string {
	return s.Name
}

// Get solver's identifier
func (s Solver) FullName() string {
	problemName := ""
	if s.Problem != nil {
		problemName = s.Problem.Name
	}
	return fmt.Sprintf("%s(%s)", s.Name, problemName)
}

// Get solver's result
func (s *Solver) GetResult() *worker.Result {
	return s.Result
}

// Get solver's problem
func (s *Solver) GetProblem() *discrete.Problem {
	return s.Problem
}

// Display and return solution space size
func (s Solver) Prelude(logger worker.Logger) int {
	solutionSpace := s.Problem.SolutionSpace()
	if solutionSpace > worker.IterationBatch {
		logger.Output("SolutionSpace:", number.Comma(solutionSpace))
	}
	return solutionSpace
}

// Add solution to solver results,
// Returns boolean indicating whether solution score is better than current
func (s *Solver) AddSolution(solution *discrete.Solution) bool {
	if solution == nil {
		return false
	}

	score := solution.Score
	s.FeasibleSolutions[score] += 1 // increment counter for feasible solutions with solution score

	isBetter := false
	if s.IsScoreBetter(score) {
		// Reset the best score, best solutions and core solutions if we find a better score
		s.BestScore = score
		s.BestSolutions = make([]*discrete.Solution, 0)
		s.CoreSolutions = make(map[string][]*discrete.Solution)
		isBetter = true
	} else if score != s.BestScore && s.Problem.IsOptimization() {
		// skip if optimization problem and score is not better than current best
		return false
	}

	coreFn := s.Problem.SolutionCoreFn
	if coreFn != nil {
		coreKey := coreFn(solution)
		s.CoreSolutions[coreKey] = append(s.CoreSolutions[coreKey], solution)
	}
	s.BestSolutions = append(s.BestSolutions, solution)
	return isBetter
}

// Display solver's progress
func (s Solver) DisplayProgress(logger worker.Logger) {
	iterationBatch := worker.IterationBatch
	if s.NumSteps%iterationBatch != 0 || s.NumSteps < iterationBatch {
		return
	}
	var bestScore string
	batch := s.NumSteps / iterationBatch
	if s.Problem.IsOptimization() {
		bestScore = fmt.Sprintf("BestScore: %.2f, %d solutions", s.BestScore, len(s.BestSolutions))
	}
	logger.Output("IterBatch:", batch, bestScore, s.FullName())
}

// Check if solution is complete
func (s Solver) IsComplete(solution *discrete.Solution) bool {
	return list.All(s.Problem.Variables, func(variable discrete.Variable) bool {
		return dict.HasKey(solution.Map, variable)
	})
}

// Check if new score is better than current best
func (s Solver) IsScoreBetter(score discrete.Score) bool {
	goal := s.Problem.Goal
	if goal == discrete.Maximize && score > s.BestScore {
		return true
	} else if goal == discrete.Minimize && score < s.BestScore {
		return true
	}
	return false
}
