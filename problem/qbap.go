package problem

import (
	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Quadratic Bottleneck Assignment problem
func QuadraticBottleneckAssignment(n int) *discrete.Problem {
	name := newName(QBAP, n)
	cfg := newQuadraticAssignment(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.IndexVariables(cfg.size)
	domain := discrete.IndexDomain(cfg.size)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For each pair, get the cost of the flow from facility1 => facility2
		// multiplied by the cost of traveling from location1 (fac1's value) => location2 (fac2's value)
		// Find the maximum cost from the pairs
		var maxCost discrete.Score = 0
		for _, pair := range comb.Combinations(p.Variables, 2) {
			var1, var2 := pair[0], pair[1]
			val1, val2 := solution.Map[var1], solution.Map[var2]
			cost1 := cfg.flow[var1][var2] * cfg.cost[val1][val2] // 1 => 2
			cost2 := cfg.flow[var2][var1] * cfg.cost[val2][val1] // 2 => 1
			maxCost = max(maxCost, cost1+cost2)
		}
		return maxCost
	}

	p.SolutionStringFn = fn.String_Sequence(list.NumRange(1, cfg.size+1))

	return p
}
