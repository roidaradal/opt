// Package brute contains brute force solvers
package brute

import (
	"github.com/roidaradal/fn/comb"
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
func (solver *LinearSolver) Solve(logLevel worker.LogLevel) {
	problem := solver.Problem
	logger := solver.Prelude(logLevel)
	domains := list.Translate(problem.Variables, problem.Domain)
	for _, values := range comb.Product(domains...) {
		solver.NumSteps += 1
		if solver.IsIterationLimitReached(logger) {
			break
		}

		solution := discrete.ZipSolution(problem.Variables, values)
		if !problem.IsSatisfied(solution) {
			continue
		}

		problem.ComputeScore(solution)
		solver.AddSolution(solution)
		if solver.IsSolutionLimitReached() {
			break
		}
	}

}
