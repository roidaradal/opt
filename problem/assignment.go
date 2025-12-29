package problem

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
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
	numWorkers := len(cfg.cost)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.IndexVariables(numWorkers)
	domain := discrete.IndexDomain(numWorkers)
	indexOf := make(dict.IntMap)
	for i, variable := range p.Variables {
		p.Domain[variable] = domain[:]
		indexOf[cfg.workers[i]] = i
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	// Team constraint if has constraint
	if len(cfg.teams) > 1 {
		test := func(solution *discrete.Solution) bool {
			for _, team := range cfg.teams {
				count := 0
				for _, workerName := range team {
					worker := indexOf[workerName]
					task := solution.Map[worker]
					if cfg.cost[worker][task] > 0 {
						count += 1
					}
				}
				if count > cfg.maxPerTeam {
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
			totalCost += cfg.cost[worker][task]
		}
		return totalCost
	}

	assignmentString := func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(worker discrete.Variable) string {
			task := solution.Map[worker]
			if cfg.cost[worker][task] == 0 {
				return "" // skip dummy tasks
			}
			return fmt.Sprintf("w%s = t%s", cfg.workers[worker], cfg.tasks[task])
		})
		output = list.Filter(output, str.NotEmpty)
		return str.WrapBraces(output)
	}

	p.SolutionStringFn = assignmentString
	p.SolutionCoreFn = assignmentString

	return p
}

type assignmentCfg struct {
	tasks      []string
	workers    []string
	cost       [][]float64
	teams      [][]string
	maxPerTeam int
}

// Load assignment problem
func newAssignment(name string) *assignmentCfg {
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
	cfg := &assignmentCfg{
		tasks:      make([]string, numTasks),
		workers:    make([]string, numWorkers),
		cost:       make([][]float64, numWorkers),
		teams:      make([][]string, numTeams),
		maxPerTeam: maxPerTeam,
	}
	copy(cfg.tasks, strings.Fields(lines[1]))
	idx := 2
	for i := range numWorkers {
		// Ensure equal number of workers and tasks
		// Adds 0-cost tasks to end of list if more workers than tasks
		line := lines[idx]
		idx++
		parts := strings.Fields(line)
		name := parts[0]
		costs := list.Map(parts[1:], fn.ParseFloatInf)
		workerCost := make([]float64, numWorkers)
		copy(workerCost, costs)
		cfg.workers[i] = name
		cfg.cost[i] = workerCost
	}
	for i := range numTeams {
		team := strings.Fields(lines[idx])
		idx++
		cfg.teams[i] = team
	}
	return cfg
}
