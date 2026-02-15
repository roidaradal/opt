package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewAllocation creates a new Allocation problem
func NewAllocation(variant string, n int) *discrete.Problem {
	name := newName(Allocation, variant, n)
	switch variant {
	case "resource":
		return resourceAllocation(name)
	default:
		return nil
	}
}

// Resource Allocation
func resourceAllocation(name string) *discrete.Problem {
	cfg := data.NewResource(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Items)
	for i, variable := range p.Variables {
		p.Domain[variable] = discrete.RangeDomain(0, cfg.Count[i])
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check sum of weighted costs don't exceed budget
		count := solution.Map
		costs := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * cfg.Cost[x]
		})
		return list.Sum(costs) <= cfg.Budget
	})

	p.Goal = discrete.Maximize
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.Value)
	p.SolutionStringFn = fn.StringValues[int](p, nil)
	return p
}
