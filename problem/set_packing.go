package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Set Packing problem
func SetPacking(n int) *discrete.Problem {
	name := newName(SET_PACKING, n)
	cfg := fn.NewSubsets(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.Names)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check each selected subset
		covered := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			// Increment the counter for each item in selected subset
			for _, item := range cfg.Subsets[x] {
				covered[item] += 1
			}
		}
		// Make sure all covered items are only covered once (no overlap)
		return list.AllEqual(dict.Values(covered), 1)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(cfg.Names)

	return p
}
