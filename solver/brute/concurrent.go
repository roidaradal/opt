package brute

import (
	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/solver/base"
	"github.com/roidaradal/opt/worker"
)

type ConcurrentSolver struct {
	base.ConcurrentSolver
	workerRange [][2]int
}

// Create new Concurrent Brute-Force solver
func NewConcurrentSolver(numWorkers int) worker.SolverCreator {
	return func(problem *discrete.Problem) worker.Solver {
		solver := &ConcurrentSolver{}
		solver.Initialize("ConcurrentBruteForce", problem, numWorkers)
		solver.workerRange = list.Divide(problem.SolutionSpace(), numWorkers)
		return solver
	}
}

// Solve problem using the Concurrent Brute-Force solver
func (solver *ConcurrentSolver) Solve(logger worker.Logger) {
	solver.Prelude(logger)
	problem := solver.Problem
	domains := list.Translate(problem.Variables, problem.Domain)
	task := func(workerID int, stepsCh chan<- struct{}, solutionsCh chan<- *discrete.Solution) {
		workerRange := solver.workerRange[workerID]
		start, end := workerRange[0], workerRange[1]
		for _, values := range comb.RangeProduct(start, end, domains...) {
			stepsCh <- struct{}{}
			solution := discrete.ZipSolution(problem.Variables, values)
			if !problem.IsSatisfied(solution) {
				continue
			}

			problem.ComputeScore(solution)
			solutionsCh <- solution
		}
	}
	solver.RunWorkers(task, logger)
}
