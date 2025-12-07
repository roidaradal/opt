package base

import (
	"math"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
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

// Get solver's result
func (s *Solver) GetResult() *worker.Result {
	return s.Result
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
