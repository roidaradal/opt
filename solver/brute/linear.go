// Package brute contains brute force solvers
package brute

import (
	"fmt"
	"iter"

	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/lang"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/solver/base"
	"github.com/roidaradal/opt/worker"
)

type LinearSolver struct {
	base.LinearSolver
}

// Create new Linear Brute-Force Solver
func NewLinearSolver(problem *discrete.Problem) worker.Solver {
	solver := &LinearSolver{}
	solver.Initialize("LinearBruteForce", problem)
	return solver
}

// Solve problem using the Linear Brute-Force solver
func (solver *LinearSolver) Solve(logger worker.Logger) {
	solutionSpace := solver.Prelude(logger)
	problem := solver.Problem
	domains := list.Translate(problem.Variables, problem.Domain)

	var valuesIterator iter.Seq2[int, []discrete.Value] = comb.Product(domains...)
	switch problem.Type {
	case discrete.Sequence:
		valuesIterator = comb.Permutations(domains[0], len(domains[0]))
	case discrete.Path:
		domain := list.Filter(domains[0], func(value discrete.Value) bool {
			return value >= 0 // remove -1
		})
		valuesIterator = comb.AllPermutationPositions(domain)
	}

	for _, values := range valuesIterator {
		solver.NumSteps += 1
		if solver.IsIterationLimitReached(logger) {
			break
		}
		var result string
		solution := discrete.ZipSolution(problem.Variables, values)
		if problem.IsSatisfied(solution) {
			problem.ComputeScore(solution)
			isBetter := solver.AddSolution(solution)
			result = lang.Ternary(isBetter, worker.BestResult, worker.FeasibleResult)
		} else {
			result = worker.InfeasibleResult
		}

		progress := (solver.NumSteps * 100) / solutionSpace
		logger.Clear(1)
		logger.Steps(fmt.Sprintf("%3d%% %v %s %.2f", progress, values, result, solver.BestScore))

		if solver.IsSolutionLimitReached() {
			break
		}
	}

}
