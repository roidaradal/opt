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

type Reporter interface {
	Run(Solver, Logger)
}

type RunReporter struct{}

// Basic Reporter for running Solver
func (r RunReporter) Run(solver Solver, logger Logger) {
	start := time.Now()
	solver.Solve(logger)
	result := solver.GetResult()
	items := [][2]string{
		{"Problem", solver.GetProblem().Name},
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
