package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSetPacking creates a new Set Cover problem
func NewSetPacking(variant string, n int) *discrete.Problem {
	name := newName(SetPacking, variant, n)
	switch variant {
	case "basic":
		return setPacking(name)
	case "weighted":
		return weightedSetPacking(name)
	default:
		return nil
	}
}

// Create new Set Packing problem
func newSetPackingProblem(name string) (*discrete.Problem, *data.Subsets) {
	p, cfg := newSubsetsProblem(name)
	if p == nil || cfg == nil {
		return nil, nil
	}

	p.Goal = discrete.Maximize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each selected subset
		covered := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			// Increment counter for each item in selected subset
			dict.UpdateCounter(covered, cfg.Subsets[x])
		}
		// Make sure all covered items are only covered once (no overlap)
		return list.AllEqual(dict.Values(covered), 1)
	})
	return p, cfg
}

// Set Packing
func setPacking(name string) *discrete.Problem {
	p, _ := newSetPackingProblem(name)
	return p
}

// Weighted Set Packing
func weightedSetPacking(name string) *discrete.Problem {
	p, cfg := newSetPackingProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	if len(cfg.Weight) != len(cfg.Names) {
		return nil
	}

	p.ObjectiveFn = fn.ScoreSumWeightedSubset(cfg.Names, cfg.Weight)
	return p
}
