package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/problem"
	"github.com/roidaradal/opt/solver/brute"
	"github.com/roidaradal/opt/worker"
)

// Create new problem
func newProblem(value string) *discrete.Problem {
	parts := str.CleanSplit(value, ":")
	if len(parts) != 2 {
		return nil
	}
	name, n := parts[0], parts[1]
	if dict.NoKey(problem.Creator, name) {
		log.Fatal("Unknown problem: ", name)
	}
	p := problem.Creator[name](number.ParseInt(n))
	if p == nil {
		log.Fatal("Unknown test case: ", value)
	}
	return p
}

// Create new worker, defaults to RunReporter
func newWorker(value string) worker.Worker {
	var mainWorker worker.Worker = defaultWorker
	parts := str.CleanSplit(value, ":")
	name := parts[0]
	switch name {
	case "RunReporter":
		mainWorker = worker.RunReporter{}
	case "RunReporterWithSolutions":
		mainWorker = worker.RunReporter{WithSolutions: true}
	case "SolutionReporter":
		mainWorker = worker.SolutionReporter{}
	default:
		fmt.Printf("Unknown worker %q, using the default worker...\n", name)
	}
	return mainWorker
}

// Create new solver creator, defaults to LinearBruteForce
func newSolverCreator(value string) worker.SolverCreator {
	var solver worker.SolverCreator = defaultSolverCreator
	parts := str.CleanSplit(value, ":")
	name := parts[0]
	switch name {
	case "LinearBruteForce":
		solver = brute.NewLinearSolver
	case "ConcurrentBruteForce":
		numWorkers := 10
		if len(parts) > 1 {
			numWorkers = max(numWorkers, number.ParseInt(parts[1]))
		}
		solver = brute.NewConcurrentSolver(numWorkers)
	default:
		fmt.Printf("Unknown solver %q, using the default solver...\n", name)
	}
	return solver
}

// Create new logger, defaults to BatchLogger
func newLogger(value string) worker.Logger {
	var logger worker.Logger = defaultLogger
	parts := str.CleanSplit(value, ":")
	name := parts[0]
	switch strings.ToLower(name) {
	case "quiet":
		logger = worker.QuietLogger{}
	case "batch":
		logger = worker.BatchLogger{}
	case "steps":
		delay := 0
		if len(parts) > 1 {
			delay = number.ParseInt(parts[1])
		}
		logger = worker.StepsLogger{DelayNanosecond: delay}
	default:
		fmt.Printf("Unknown logger %q, using the default logger...\n", name)
	}
	return logger
}
