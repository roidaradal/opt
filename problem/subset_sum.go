package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSubsetSum creates a new Subset Sum problem
func NewSubsetSum(variant string, n int) *discrete.Problem {
	name := newName(SubsetSum, variant, n)
	switch variant {
	case "basic":
		return subsetSum(name)
	default:
		return nil
	}
}

// Subset Sum
func subsetSum(name string) *discrete.Problem {
	p, cfg := newNumbersSubsetProblem(name)
	if p == nil || cfg == nil || cfg.Target == 0 {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get solution subset sum
		total := list.Sum(list.MapList(fn.AsSubset(solution), cfg.Numbers))
		if p.IsSatisfaction() {
			// Satisfaction: check if subset sum == target sum
			return total == cfg.Target
		}
		// Optimization: check if subset sum does not exceed target
		return total <= cfg.Target
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For optimization version, minimize difference between target and subset sum
		// If it exceeds target, invalid solution
		total := list.Sum(list.MapList(fn.AsSubset(solution), cfg.Numbers))
		if total > cfg.Target {
			return discrete.Inf
		}
		return discrete.Score(cfg.Target - total)
	}
	return p
}
