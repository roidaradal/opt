package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSetCover creates a new Set Cover problem
func NewSetCover(variant string, n int) *discrete.Problem {
	name := newName(SetCover, variant, n)
	switch variant {
	case "basic":
		return setCover(name)
	default:
		return nil
	}
}

// Set Cover
func setCover(name string) *discrete.Problem {
	p, cfg := newSubsetsProblem(name)
	if p == nil || cfg == nil {
		return nil
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
	return p
}
