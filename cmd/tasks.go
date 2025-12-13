package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/worker"
)

// Run Reporter task
func runReporterTask(task worker.Reporter, args []string) {
	var err error
	var p *discrete.Problem
	newSolver := defaultSolverCreator
	logger := defaultLogger

	for _, pair := range args[1:] {
		parts := str.CleanSplit(pair, "=")
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "p", "problem":
			p, err = newProblem(value)
			if err != nil {
				fmt.Println(str.Red("Error:"), err)
				return
			}
		case "s", "solver":
			newSolver = newSolverCreator(value)
		case "l", "logger":
			logger = newLogger(value)
		}
	}

	if p == nil {
		fmt.Println(str.Red("Error:"), "Undefined problem")
		displayUsage(strings.ToLower(args[0]))
		return
	}

	solver := newSolver(p)
	task.Run(solver, logger)
}
