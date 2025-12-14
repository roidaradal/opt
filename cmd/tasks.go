package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/worker"
)

// Run worker
func runWorker(task worker.Worker, args []string) {
	cfg := &worker.Config{
		NewSolver: defaultSolverCreator,
		Logger:    defaultLogger,
	}

	for _, pair := range args[1:] {
		parts := str.CleanSplit(pair, "=")
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "p", "problem":
			p, err := newProblem(value)
			if err != nil {
				fmt.Println(redError, err)
				return
			}
			cfg.Problem = p
		case "s", "solver":
			cfg.NewSolver = newSolverCreator(value)
		case "l", "logger":
			cfg.Logger = newLogger(value)
		}
	}

	output := task.Run(cfg)
	fmt.Println(output)
}

// Run manager
func runManager(manager worker.Manager, args []string) {
	cfg := &worker.Config{
		NewSolver: defaultSolverCreator,
		Logger:    defaultLogger,
	}

	var problems []*discrete.Problem = nil
	hasWorker := false
	for _, pair := range args[1:] {
		parts := str.CleanSplit(pair, "=")
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "w", "worker":
			if worker, ok := newWorker(value); ok {
				cfg.Worker = worker
				hasWorker = true
			}
		case "d", "data":
			problems = newDataset(value)
		case "s", "solver":
			cfg.NewSolver = newSolverCreator(value)
		case "l", "logger":
			cfg.Logger = newLogger(value)
		}
	}

	if problems == nil {
		fmt.Println(redError, "No problems from dataset")
		displayUsage(strings.ToLower(args[0]), false)
		return
	}

	if !hasWorker {
		fmt.Println(redError, "Undefined base worker")
		displayUsage(strings.ToLower(args[0]), false)
		return
	}

	manager.Run(problems, cfg)
}
