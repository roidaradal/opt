package problem

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Linear Bottleneck Assignment problem
func LinearBottleneckAssignment(n int) *discrete.Problem {
	name := newName(LBAP, n)
	cfg := newLBAP(name)
	if cfg == nil {
		return nil
	}
	numWorkers := len(cfg.Cost)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.IndexVariables(numWorkers)
	domain := discrete.IndexDomain(numWorkers)
	indexOf := make(dict.IntMap)
	for i, variable := range p.Variables {
		p.Domain[variable] = domain[:]
		indexOf[cfg.Workers[i]] = i
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Total cost of assigning worker to task
		var maxCost discrete.Score = 0
		for worker, task := range solution.Map {
			maxCost = max(maxCost, cfg.Cost[worker][task])
		}
		return maxCost
	}

	p.SolutionStringFn = fn.String_Assignment(p, cfg)
	p.SolutionCoreFn = fn.String_Assignment(p, cfg)

	return p
}

// Load lbap test case
func newLBAP(name string) *a.AssignmentCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	counts := list.Map(strings.Fields(lines[0]), number.ParseInt)
	numWorkers, numTasks := counts[0], counts[1]
	if numTasks > numWorkers {
		fmt.Println("Invalid Assignment problem: more tasks than workers")
		return nil
	}
	cfg := &a.AssignmentCfg{
		Tasks:   make([]string, numTasks),
		Workers: make([]string, numWorkers),
		Cost:    make([][]float64, numWorkers),
	}
	copy(cfg.Tasks, strings.Fields(lines[1]))
	idx := 2
	for i := range numWorkers {
		// Ensure equal number of workers and tasks
		// Adds 0-cost tasks to end of list if more workers than tasks
		line := lines[idx]
		idx++
		parts := strings.Fields(line)
		name := parts[0]
		costs := fn.ParseFloatInfRow(parts, true)
		workerCost := make([]float64, numWorkers)
		copy(workerCost, costs)
		cfg.Workers[i] = name
		cfg.Cost[i] = workerCost
	}
	return cfg
}
