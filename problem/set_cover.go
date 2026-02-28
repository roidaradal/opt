package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSetCover creates a new Set Cover problem
func NewSetCover(variant string, n int) *discrete.Problem {
	name := newName(SetCover, variant, n)
	switch variant {
	case "basic":
		return setCover(name)
	case "weighted":
		return weightedSetCover(name)
	default:
		return nil
	}
}

// Create new Set Cover problem
func newSetCoverProblem(name string) (*discrete.Problem, *data.Subsets) {
	p, cfg := newSubsetsProblem(name)
	if p == nil || cfg == nil {
		return nil, nil
	}

	p.Goal = discrete.Minimize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each selected subset
		covered := dict.Flags(cfg.Universal, false)
		for _, x := range fn.AsSubset(solution) {
			// Each subset item is covered
			for _, item := range cfg.Subsets[x] {
				covered[item] = true
			}
		}
		return list.AllTrue(dict.Values(covered))
	})
	return p, cfg
}

// Set Cover
func setCover(name string) *discrete.Problem {
	p, _ := newSetCoverProblem(name)
	return p
}

// Weighted Set Cover
func weightedSetCover(name string) *discrete.Problem {
	p, cfg := newSetCoverProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	if len(cfg.Weight) != len(cfg.Names) {
		return nil
	}

	p.ObjectiveFn = fn.ScoreSumWeightedSubset(cfg.Names, cfg.Weight)
	return p
}
