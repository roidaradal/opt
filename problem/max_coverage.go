package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewMaxCoverage creates a new Max Coverage problem
func NewMaxCoverage(variant string, n int) *discrete.Problem {
	name := newName(MaxCoverage, variant, n)
	switch variant {
	case "basic":
		return maxCoverage(name)
	default:
		return nil
	}
}

// Max Coverage
func maxCoverage(name string) *discrete.Problem {
	p, cfg := newSubsetsProblem(name)
	if p == nil || cfg == nil || cfg.Limit == 0 {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check number of selected subsets don't exceed limit
		return len(fn.AsSubset(solution)) <= cfg.Limit
	})

	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count unique covered items
		covered := ds.NewSet[string]()
		for _, x := range fn.AsSubset(solution) {
			covered.AddItems(cfg.Subsets[x])
		}
		return discrete.Score(covered.Len())
	}
	return p
}
