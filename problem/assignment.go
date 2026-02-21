package problem

import (
	"fmt"
	"slices"

	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
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
	case "general":
		return generalizedAssignment(name)
	case "quadratic":
		return quadraticAssignment(name)
	case "quadratic_bottleneck":
		return quadraticBottleneckAssignment(name)
	case "weapon":
		return weaponTargetAssignment(name)
	default:
		return nil
	}
}

// Common steps for creating Assignment problem
func newAssignmentProblem(name string) (*discrete.Problem, *data.AssignmentCfg) {
	cfg := data.NewAssignment(name, true)
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

// Generalized Assignment
func generalizedAssignment(name string) *discrete.Problem {
	cfg := data.NewAssignment(name, false)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Tasks)
	p.AddVariableDomains(discrete.Domain(cfg.Workers))

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Compute total cost of assigning tasks to each worker
		total := make(map[int]float64)
		for task, worker := range solution.Map {
			total[worker] += cfg.Cost[worker][task]
		}
		// Check that worker totals don't exceed their capacity
		for worker, limit := range cfg.Capacity {
			if total[worker] > limit {
				return false
			}
		}
		return true
	})

	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Total value of assigning tasks to workers
		var totalValue discrete.Score = 0
		for task, worker := range solution.Map {
			totalValue += cfg.Value[worker][task]
		}
		return totalValue
	}

	p.SolutionStringFn = func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(task discrete.Variable) string {
			worker := solution.Map[task]
			return fmt.Sprintf("t%s = w%s", cfg.Tasks[task], cfg.Workers[worker])
		})
		return str.WrapBraces(output)
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

// Weapon Target Assignment
func weaponTargetAssignment(name string) *discrete.Problem {
	cfg := data.NewWeapons(name)
	if cfg == nil {
		return nil
	}
	numWeapons, numTargets := len(cfg.Weapons), len(cfg.Targets)

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	// Expand weapons and count into weapons list (uses weapon index)
	weapons := make([]int, 0)
	for i := range numWeapons {
		weapons = append(weapons, slices.Repeat([]int{i}, cfg.Count[i])...)
	}

	p.Variables = discrete.Variables(weapons)
	p.AddVariableDomains(discrete.Domain(cfg.Targets))

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Compute survival rate of each target, with weapons assigned to attack it
		survival := list.Copy(cfg.Value)
		for w, target := range solution.Map {
			weapon := weapons[w]
			survival[target] *= 1 - cfg.Chance[weapon][target]
			// survival = 1 - weaponOnTargetEffectiveness
		}
		total := fmt.Sprintf("%.4f", list.Sum(survival))
		return number.ParseFloat(total)
	}

	weaponTargets := func(solution *discrete.Solution) string {
		// Group count of weapon => target assignments
		matrix := make([][]int, numWeapons)
		for i := range numWeapons {
			matrix[i] = make([]int, numTargets)
		}
		for w, target := range solution.Map {
			weapon := weapons[w]
			matrix[weapon][target] += 1
		}
		output := make([]string, 0)
		for i, weapon := range cfg.Weapons {
			for j, target := range cfg.Targets {
				if matrix[i][j] == 0 {
					continue // skip empty count
				}
				line := fmt.Sprintf("%d*%s = %s", matrix[i][j], weapon, target)
				output = append(output, line)
			}
		}
		return str.WrapBraces(output)
	}

	p.SolutionStringFn = weaponTargets
	p.SolutionCoreFn = weaponTargets

	return p
}
