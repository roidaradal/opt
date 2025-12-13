package worker

import (
	"fmt"
	"log"
	"path/filepath"
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

type Reporter interface {
	Run(Solver, Logger)
}

type RunReporter struct {
	WithSolutions bool
}

type SolutionReporter struct{}

// Basic Reporter for running Solver
func (r RunReporter) Run(solver Solver, logger Logger) {
	problem := solver.GetProblem()
	stringFn := problem.SolutionStringFn
	coreFn := problem.SolutionCoreFn

	start := time.Now()
	solver.Solve(logger)
	result := solver.GetResult()

	if r.WithSolutions {
		if coreFn == nil {
			for i, solution := range result.BestSolutions {
				prefix := fmt.Sprintf("S%-3d : ", i+1)
				if stringFn == nil {
					fmt.Println(prefix, solution)
				} else {
					fmt.Println(prefix, stringFn(solution))
				}
			}
		} else {
			coreKeys := dict.Keys(result.CoreSolutions)
			slices.Sort(coreKeys)
			for i, key := range coreKeys {
				// Get first solution as representative
				solution := result.CoreSolutions[key][0]
				count := len(result.CoreSolutions[key])
				prefix := fmt.Sprintf("S%-3d : %s | %3d | ", i+1, key, count)
				if stringFn == nil {
					fmt.Println(prefix, solution)
				} else {
					fmt.Println(prefix, stringFn(solution))
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
	template := fmt.Sprintf("%%-%ds : %%s\n", slices.Max(lengths))

	for _, pair := range items {
		key, value := pair[0], pair[1]
		fmt.Printf(template, key, value)
	}
}

// Saves solutions to solution/<problemname>.txt
func (r SolutionReporter) Run(solver Solver, logger Logger) {
	problem := solver.GetProblem()
	stringFn := problem.SolutionStringFn
	coreFn := problem.SolutionCoreFn

	solver.Solve(logger)
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
		log.Fatal(err)
	}

	path := filepath.Join("solution", fmt.Sprintf("%s.txt", problem.Name))
	err = io.SaveString(strings.Join(out, "\n"), path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved: ", path)
}
