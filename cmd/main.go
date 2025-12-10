package main

import (
	"github.com/roidaradal/opt/problem"
	"github.com/roidaradal/opt/solver/brute"
	"github.com/roidaradal/opt/worker"
)

func main() {
	p := problem.ActivitySelection(1)
	solver := brute.NewLinearSolver(p)
	logger := worker.StepsLogger{DelayMs: 1}
	// logger := worker.BatchLogger{}
	solver.Solve(logger)
}
