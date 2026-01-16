package problem

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Generalized Assignment problem
func GeneralizedAssignment(n int) *discrete.Problem {
	name := newName(GEN_ASSIGNMENT, n)
	cfg := loadGeneralizedAssignment(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.tasks)
	domain := discrete.MapDomain(cfg.workers)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Compute total weight of each worker
		used := make(map[int]float64)
		for task, worker := range solution.Map {
			used[worker] += cfg.weight[worker][task]
		}
		// Check that each worker weight does not exceed capacity
		for worker, limit := range cfg.capacity {
			if used[worker] > limit {
				return false
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Total profit of assigning task to worker
		var totalProfit discrete.Score = 0
		for task, worker := range solution.Map {
			totalProfit += cfg.profit[worker][task]
		}
		return totalProfit
	}

	p.SolutionStringFn = func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(task discrete.Variable) string {
			worker := solution.Map[task]
			return fmt.Sprintf("t%s = w%s", cfg.tasks[task], cfg.workers[worker])
		})
		return str.WrapBraces(output)
	}

	return p
}

type generalizedAssignmentCfg struct {
	tasks    []string
	workers  []string
	profit   [][]float64
	weight   [][]float64
	capacity []float64
}

// Load generalized assignment problem test case
func loadGeneralizedAssignment(name string) *generalizedAssignmentCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	counts := list.Map(strings.Fields(lines[0]), number.ParseInt)
	numWorkers, numTasks := counts[0], counts[1]
	cfg := &generalizedAssignmentCfg{
		tasks:    make([]string, numTasks),
		workers:  make([]string, numWorkers),
		profit:   make([][]float64, numWorkers),
		weight:   make([][]float64, numWorkers),
		capacity: make([]float64, numWorkers),
	}
	cfg.tasks = strings.Fields(lines[1])
	idx := 2
	for i := range numWorkers {
		parts := strings.Fields(lines[idx])
		idx++
		cfg.workers[i] = parts[0]
		cfg.profit[i] = fn.ParseFloatRow(parts, true)
	}
	for i := range numWorkers {
		parts := strings.Fields(lines[idx])
		idx++
		cfg.weight[i] = fn.ParseFloatRow(parts, true)
	}
	for i := range numWorkers {
		parts := strings.Fields(lines[idx])
		idx++
		cfg.capacity[i] = number.ParseFloat(parts[1])
	}
	return cfg
}
