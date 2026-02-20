package problem

import (
	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewAssignment creates a new Assignment problem
func NewAssignment(variant string, n int) *discrete.Problem {
	name := newName(Assignment, variant, n)
	switch variant {
	case "basic":
		return assignment(name)
	case "bottleneck":
		return bottleneckAssignment(name)
	case "quadratic":
		return quadraticAssignment(name)
	case "quadratic_bottleneck":
		return quadraticBottleneckAssignment(name)
	default:
		return nil
	}
}

// Common steps for creating Assignment problem
func newAssignmentProblem(name string) (*discrete.Problem, *data.AssignmentCfg) {
	cfg := data.NewAssignment(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(cfg.Workers)
	p.AddVariableDomains(discrete.IndexDomain(len(cfg.Workers)))
	p.AddUniversalConstraint(fn.ConstraintAllUnique)

	p.SolutionStringFn = fn.StringAssignment(p, cfg)
	p.SolutionCoreFn = fn.StringAssignment(p, cfg)
	return p, cfg
}

// Assignment
func assignment(name string) *discrete.Problem {
	p, cfg := newAssignmentProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	// Add team constraint if it has more than 1 team
	if len(cfg.Teams) > 1 {
		indexOf := list.IndexMap(cfg.Workers)
		p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
			for _, team := range cfg.Teams {
				// Count team members with tasks in the solution
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
		})
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Total cost of assigning workers to tasks
		var totalCost discrete.Score = 0
		for worker, task := range solution.Map {
			totalCost += cfg.Cost[worker][task]
		}
		return totalCost
	}

	return p
}

// Bottleneck Assignment
func bottleneckAssignment(name string) *discrete.Problem {
	p, cfg := newAssignmentProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Max cost of assigning worker to task
		var maxCost discrete.Score = 0
		for worker, task := range solution.Map {
			maxCost = max(maxCost, cfg.Cost[worker][task])
		}
		return maxCost
	}
	return p
}

// Common steps for creating Quadratic Assignment problem
func newQuadraticAssignmentProblem(name string) (*discrete.Problem, *data.QuadraticAssignment) {
	cfg := data.NewQuadraticAssignment(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Minimize

	p.Variables = discrete.IndexVariables(cfg.Count)
	p.AddVariableDomains(discrete.IndexDomain(cfg.Count))

	p.AddUniversalConstraint(fn.ConstraintAllUnique)
	p.SolutionStringFn = fn.StringSequence(list.NumRange(1, cfg.Count+1))
	return p, cfg
}

// Quadratic Assignment
func quadraticAssignment(name string) *discrete.Problem {
	p, cfg := newQuadraticAssignmentProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For each pair, sum up cost of flow from facility1 => facility2
		// multiplied by distance of traveling from location1 (facility1's value) => location2 (facility2's value)
		var totalCost discrete.Score = 0
		for _, pair := range comb.Combinations(p.Variables, 2) {
			facility1, facility2 := pair[0], pair[1]
			location1, location2 := solution.Map[facility1], solution.Map[facility2]
			totalCost += cfg.Flow[facility1][facility2] * cfg.Distance[location1][location2] // 1 => 2
			totalCost += cfg.Flow[facility2][facility1] * cfg.Distance[location2][location1] // 2 => 1
		}
		return totalCost
	}
	return p
}

// Quadratic Bottleneck Assignment
func quadraticBottleneckAssignment(name string) *discrete.Problem {
	p, cfg := newQuadraticAssignmentProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For each pair, get the cost of flow from facility1 => facility2
		// multiplied by distance of traveling from location1 (facility1's value) => location2 (facility2's value)
		// Find maximum cost from pairs
		var maxCost discrete.Score = 0
		for _, pair := range comb.Combinations(p.Variables, 2) {
			facility1, facility2 := pair[0], pair[1]
			location1, location2 := solution.Map[facility1], solution.Map[facility2]
			cost1 := cfg.Flow[facility1][facility2] * cfg.Distance[location1][location2] // 1 => 2
			cost2 := cfg.Flow[facility2][facility1] * cfg.Distance[location2][location1] // 2 => 1
			maxCost = max(maxCost, cost1, cost2)
		}
		return maxCost
	}
	return p
}
