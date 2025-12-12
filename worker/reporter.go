package worker

import (
	"fmt"
	"slices"
	"time"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
)

type Worker interface {
	Run(Solver, Logger)
}

type RunReporter struct {
	WithSolutions bool
}

// Basic Reporter for running Solver
func (r RunReporter) Run(solver Solver, logger Logger) {
	problem := solver.GetProblem()
	stringFn := problem.SolutionStringFn
	coreFn := problem.SolutionCoreFn

	start := time.Now()
	solver.Solve(logger)
	result := solver.GetResult()
	items := [][2]string{
		{"Problem", problem.Name},
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

	if r.WithSolutions {
		if coreFn == nil {
			for i, solution := range result.BestSolutions {
				prefix := fmt.Sprintf("S%-3d : ", i+1)
				if stringFn == nil {
					logger.Output(prefix, solution)
				} else {
					logger.Output(prefix, stringFn(solution))
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
					logger.Output(prefix, solution)
				} else {
					logger.Output(prefix, stringFn(solution))
				}
			}
		}
	}

	for _, pair := range items {
		key, value := pair[0], pair[1]
		fmt.Printf(template, key, value)
	}
}
