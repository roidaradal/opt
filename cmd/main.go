package main

import (
	"github.com/roidaradal/opt/problem"
	"github.com/roidaradal/opt/solver/brute"
	"github.com/roidaradal/opt/worker"
)

func main() {
	p := problem.GraphColoring(7)
	// solver := brute.NewLinearSolver(p)
	solver := brute.NewConcurrentSolver(4)(p)
	// logger := worker.StepsLogger{DelayNanosecond: 1}
	logger := worker.BatchLogger{}
	reporter := worker.RunReporter{}
	reporter.Run(solver, logger)
}
