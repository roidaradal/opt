package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Exact Cover problem
func ExactCover(n int) *discrete.Problem {
	name := newName(EXACT_COVER, n)
	cfg := fn.NewSubsets(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy

	p.Variables = discrete.Variables(cfg.Names)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		count := dict.NewCounter(cfg.Universal)
		// Check each selected subset
		for _, x := range fn.AsSubset(solution) {
			// Update the counter for each subset item
			dict.UpdateCounter(count, cfg.Subsets[x])
		}
		// Check all counts are 1 = each universal item is covered
		// by exactly once by the selected subsets
		return list.AllEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test)

	p.SolutionStringFn = fn.String_Subset(cfg.Names)

	return p
}
