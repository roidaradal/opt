package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Set Cover problem
func SetCover(n int) *discrete.Problem {
	name := newName(SET_COVER, n)
	cfg := fn.NewSubsets(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(cfg.Names)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check each selected subset
		covered := dict.Flags(cfg.Universal, false)
		for _, x := range fn.AsSubset(solution) {
			// Each subset item is covered
			for _, item := range cfg.Subsets[x] {
				covered[item] = true
			}
		}
		return list.AllTrue(dict.Values(covered))
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(cfg.Names)

	return p
}
