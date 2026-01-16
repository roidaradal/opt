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

// Create new Assignment problem
func Assignment(n int) *discrete.Problem {
	name := newName(ASSIGNMENT, n)
	cfg := newAssignment(name)
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

	// Team constraint if has constraint
	if len(cfg.Teams) > 1 {
		test := func(solution *discrete.Solution) bool {
			for _, team := range cfg.Teams {
				count := 0
				for _, workerName := range team {
					worker := indexOf[workerName]
					task := solution.Map[worker]
					if cfg.Cost[worker][task] > 0 {
						count += 1
					}
				}
				if count > cfg.MaxPerTeam {
					return false
				}
			}
			return true
		}
		p.AddUniversalConstraint(test)
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Total cost of assigning worker to task
		var totalCost discrete.Score = 0
		for worker, task := range solution.Map {
			totalCost += cfg.Cost[worker][task]
		}
		return totalCost
	}

	p.SolutionStringFn = fn.String_Assignment(p, cfg)
	p.SolutionCoreFn = fn.String_Assignment(p, cfg)

	return p
}

// Load assignment problem test case
func newAssignment(name string) *a.AssignmentCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	counts := list.Map(strings.Fields(lines[0]), number.ParseInt)
	numWorkers, numTasks := counts[0], counts[1]
	numTeams, maxPerTeam := counts[2], counts[3]
	if numTasks > numWorkers {
		fmt.Println("Invalid Assignment problem: more tasks than workers")
		return nil
	}
	cfg := &a.AssignmentCfg{
		Tasks:      make([]string, numTasks),
		Workers:    make([]string, numWorkers),
		Cost:       make([][]float64, numWorkers),
		Teams:      make([][]string, numTeams),
		MaxPerTeam: maxPerTeam,
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
	for i := range numTeams {
		team := strings.Fields(lines[idx])
		idx++
		cfg.Teams[i] = team
	}
	return cfg
}
