package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/solver/brute"
	"github.com/roidaradal/opt/worker"
)

/*
opt <task>

<task>
- ? | help 	Helper
- run		RunReporter
- run+sol	RunReporter{WithSolutions}
- sol.save 	SolutionReporter
- sol.read 	SolutionReader
- test		Tester
- multi		MultiTasker
*/

var usage string = strings.Join([]string{
	"Usage:",
	"opt w=workerName:params p=problemName:n s=solverName:params l=loggerName:params",
	"opt worker=workerName:params problem=problemName:n solver=solverName:params logger=loggerName:params",
	"opt f=pathToJSON",
	"opt file=pathToJSON",
}, "\n")

var (
	defaultWorker        worker.Worker        = worker.RunReporter{}
	defaultLogger        worker.Logger        = worker.BatchLogger{}
	defaultSolverCreator worker.SolverCreator = brute.NewLinearSolver
)

func main() {
	var p *discrete.Problem
	newSolver := defaultSolverCreator
	logger := defaultLogger
	mainWorker := defaultWorker

	for _, pair := range os.Args[1:] {
		parts := str.CleanSplit(pair, "=")
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "p", "problem":
			p = newProblem(value)
		case "s", "solver":
			newSolver = newSolverCreator(value)
		case "l", "logger":
			logger = newLogger(value)
		case "w", "worker":
			mainWorker = newWorker(value)
		}
	}

	if p == nil {
		fmt.Println("Error: Undefined problem")
		fmt.Println(usage)
		return
	}

	solver := newSolver(p)
	mainWorker.Run(solver, logger)
}
