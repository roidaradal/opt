package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSetPacking creates a new Set Cover problem
func NewSetPacking(variant string, n int) *discrete.Problem {
	name := newName(SetPacking, variant, n)
	switch variant {
	case "basic":
		return setPacking(name)
	default:
		return nil
	}
}

// Set Packing
func setPacking(name string) *discrete.Problem {
	p, cfg := newSubsetsProblem(name)
	if p == nil || cfg == nil {
		return nil
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
	return p
}
