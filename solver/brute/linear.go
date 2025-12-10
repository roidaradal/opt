// Package brute contains brute force solvers
package brute

import (
	"fmt"

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
	for _, values := range comb.Product(domains...) {
		solver.NumSteps += 1
		if solver.IsIterationLimitReached(logger) {
			break
		}
		var result string
		solution := discrete.ZipSolution(problem.Variables, values)
		if problem.IsSatisfied(solution) {
			problem.ComputeScore(solution)
			isBetter := solver.AddSolution(solution)
			result = lang.Ternary(isBetter, worker.BestSolution, worker.FeasibleSolution)
		} else {
			result = worker.InfeasibleSolution
		}

		progress := (solver.NumSteps * 100) / solutionSpace
		logger.Steps(fmt.Sprintf("%3d%% %v %s %.2f", progress, values, result, solver.BestScore))

		if solver.IsSolutionLimitReached() {
			break
		}
	}

}
