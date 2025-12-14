package worker

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
)

type SolverRunner struct {
	DisplaySolutions bool
}

type SolutionSaver struct{}

// Runs Solver.Solve() using Logger
func (r SolverRunner) Run(cfg *Config) string {
	problem := cfg.Problem
	if problem == nil {
		return errMessage(errMissingProblem)
	}

	solver := cfg.NewSolver(problem)
	stringFn := problem.SolutionStringFn
	coreFn := problem.SolutionCoreFn

	start := time.Now()
	solver.Solve(cfg.Logger)
	result := solver.GetResult()
	out := make([]string, 0)

	if r.DisplaySolutions {
		if coreFn == nil {
			for i, solution := range result.BestSolutions {
				prefix := fmt.Sprintf("S%-3d :", i+1)
				if stringFn == nil {
					out = append(out, fmt.Sprintf("%s %v", prefix, solution))
				} else {
					out = append(out, fmt.Sprintf("%s %s", prefix, stringFn(solution)))
				}
			}
		} else {
			coreKeys := dict.Keys(result.CoreSolutions)
			slices.Sort(coreKeys)
			for i, key := range coreKeys {
				// Get first solution as representative
				solution := result.CoreSolutions[key][0]
				count := len(result.CoreSolutions[key])
				prefix := fmt.Sprintf("S%-3d : %s | %3d |", i+1, key, count)
				if stringFn == nil {
					out = append(out, fmt.Sprintf("%s %v", prefix, solution))
				} else {
					out = append(out, fmt.Sprintf("%s %s", prefix, stringFn(solution)))
				}
			}
		}
	}

	items := [][2]string{
		{"Problem", problem.Name},
		{"Solver", solver.GetName()},
		{"Iterations", number.Comma(result.NumSteps)},
		{"Feasible Solutions", number.Comma(list.Sum(dict.Values(result.FeasibleSolutions)))},
		{"Best Solutions", number.Comma(len(result.BestSolutions))},
		{"Core Solutions", number.Comma(len(result.CoreSolutions))},
		{"Best Score", fmt.Sprintf("%.2f", result.BestScore)},
		{"Time", str.Any(time.Since(start).Round(time.Millisecond))},
	}
	lengths := list.Map(items, func(pair [2]string) int {
		return len(pair[0])
	})
	template := fmt.Sprintf("%%-%ds : %%s", slices.Max(lengths))

	for _, pair := range items {
		key, value := pair[0], pair[1]
		out = append(out, fmt.Sprintf(template, key, value))
	}
	return strings.Join(out, "\n")
}

// Runs Solver.Solve() and saves solutions to solution/<problemname>.txt
func (r SolutionSaver) Run(cfg *Config) string {
	problem := cfg.Problem
	if problem == nil {
		return errMessage(errMissingProblem)
	}

	solver := cfg.NewSolver(problem)
	stringFn := problem.SolutionStringFn
	coreFn := problem.SolutionCoreFn

	solver.Solve(cfg.Logger)
	result := solver.GetResult()

	out := make([]string, 0)
	out = append(out, fmt.Sprintf("%.2f", result.BestScore))
	if coreFn == nil {
		for _, solution := range result.BestSolutions {
			if stringFn == nil {
				out = append(out, fmt.Sprintf("+ %v", solution))
			} else {
				out = append(out, fmt.Sprintf("+ %s", stringFn(solution)))
			}
		}
	} else {
		coreKeys := dict.Keys(result.CoreSolutions)
		slices.Sort(coreKeys)
		for _, key := range coreKeys {
			out = append(out, fmt.Sprintf("+ %s", key))
			solutions := ds.NewSet[string]()
			for _, solution := range result.CoreSolutions[key] {
				if stringFn == nil {
					solutions.Add(fmt.Sprintf("\t- %v", solution))
				} else {
					solutions.Add(fmt.Sprintf("\t- %s", stringFn(solution)))
				}
			}
			solutionStrings := solutions.Items()
			slices.Sort(solutionStrings)
			for _, solutionString := range solutionStrings {
				out = append(out, solutionString)
			}
		}
	}

	err := io.EnsurePathExists("solution/")
	if err != nil {
		return errMessage(err)
	}

	path := fmt.Sprintf("solution/%s.txt", problem.Name)
	err = io.SaveString(strings.Join(out, "\n"), path)
	if err != nil {
		return errMessage(err)
	}

	return fmt.Sprintf("%s %s", str.Green("Saved:"), path)
}
