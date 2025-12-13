package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/worker"
)

type Config struct {
	problem   *discrete.Problem
	newSolver worker.SolverCreator
	logger    worker.Logger
}

// Build Config from args
func buildConfig(args []string) (*Config, error) {
	cfg := &Config{
		newSolver: defaultSolverCreator,
		logger:    defaultLogger,
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
				return nil, err
			}
			cfg.problem = p
		case "s", "solver":
			cfg.newSolver = newSolverCreator(value)
		case "l", "logger":
			cfg.logger = newLogger(value)
		}
	}

	return cfg, nil
}

// Execute Runner task
func runTask(task worker.Runner, args []string) {
	cfg, err := buildConfig(args)
	if err != nil {
		fmt.Println(redError, err)
		return
	}

	if cfg.problem == nil {
		fmt.Println(redError, "Undefined problem")
		displayUsage(strings.ToLower(args[0]), false)
		return
	}

	solver := cfg.newSolver(cfg.problem)
	task.Run(solver, cfg.logger)
}

// Read Solution
func readSolution(task worker.SolutionReader, args []string) {
	cfg, err := buildConfig(args)
	if err != nil {
		fmt.Println(redError, err)
		return
	}

	if cfg.problem == nil {
		fmt.Println(redError, "Undefined test case")
		displayUsage(strings.ToLower(args[0]), false)
		return
	}

	task.Read(cfg.problem)
}
