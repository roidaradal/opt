package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Cover creates a new Cover problem
func Cover(variant string, n int) *discrete.Problem {
	name := newName(COVER, variant, n)
	switch variant {
	case "exact":
		return coverExact(name)
	case "max":
		return coverMax(name)
	default:
		return nil
	}
}

// Exact Cover problem
func coverExact(name string) *discrete.Problem {
	cfg, _ := newSubsets(name, 0)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset
	p.Goal = discrete.Satisfy

	p.Variables = discrete.Variables(cfg.names)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		count := dict.NewCounter(cfg.universal)
		// Check each seleected subset
		for _, x := range fn.AsSubset(solution) {
			// Update counter for each subset item
			dict.UpdateCounter(count, cfg.subsets[x])
		}
		// Check all counts are 1 = each universal item is
		// covered exactly once by selected subsets
		return list.AllEqual(dict.Values(count), 1)
	})

	p.SolutionStringFn = fn.StringSubset(cfg.names)

	return p
}

// Max Coverage problem
func coverMax(name string) *discrete.Problem {
	cfg, lines := newSubsets(name, 1)
	if cfg == nil || lines == nil {
		return nil
	}
	limit := number.ParseInt(lines[0])

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.names)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check number of selected subsets don't exceed limit
		return len(fn.AsSubset(solution)) <= limit
	})

	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count unique covered items
		covered := ds.NewSet[string]()
		for _, x := range fn.AsSubset(solution) {
			covered.AddItems(cfg.subsets[x])
		}
		return discrete.Score(covered.Len())
	}

	p.SolutionStringFn = fn.StringSubset(cfg.names)

	return p
}
