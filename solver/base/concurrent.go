package base

import (
	"fmt"
	"sync"

	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/worker"
)

// Task that takes in workerID, channel for incrementing steps,
// and channel for receiving solutions
type ConcurrentTaskFn = func(int, chan<- struct{}, chan<- *discrete.Solution)

// Base Concurrent Solver
type ConcurrentSolver struct {
	Solver
	NumWorkers  int
	StepsCh     chan struct{} // empty struct costs 0 bytes, only used to signal increment
	SolutionsCh chan *discrete.Solution
}

// Initialize base Concurrent solver
func (s *ConcurrentSolver) Initialize(name string, problem *discrete.Problem, numWorkers int) {
	name = fmt.Sprintf("%s%d", name, numWorkers)
	s.Solver.Initialize(name, problem)
	s.NumWorkers = numWorkers
	s.StepsCh = make(chan struct{}, numWorkers)
	s.SolutionsCh = make(chan *discrete.Solution, numWorkers)
}

// Run workers of Concurrent solver
func (s *ConcurrentSolver) RunWorkers(task ConcurrentTaskFn, logger worker.Logger) {
	var wg sync.WaitGroup
	for workerID := range s.NumWorkers {
		wg.Go(func() {
			task(workerID, s.StepsCh, s.SolutionsCh)
		})
	}

	go func() {
		wg.Wait()
		close(s.StepsCh)
		close(s.SolutionsCh)
	}()

	go func() {
		for range s.StepsCh {
			s.NumSteps += 1
			s.DisplayProgress(logger)
		}
	}()

	for solution := range s.SolutionsCh {
		s.AddSolution(solution)
	}
}
