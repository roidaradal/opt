package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewKnapsack creates a new Knapsack problem
func NewKnapsack(variant string, n int) *discrete.Problem {
	name := newName(Knapsack, variant, n)
	switch variant {
	case "basic":
		return knapsack(name)
	case "quadratic":
		return quadraticKnapsack(name)
	default:
		return nil
	}
}

// Common steps for creating Knapsack problem
func newKnapsackProblem(name string) (*discrete.Problem, *data.Knapsack) {
	cfg := data.NewKnapsack(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(cfg.Items)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check sum of item weights don't exceed capacity
		count := solution.Map
		weights := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * cfg.Weight[x]
		})
		return list.Sum(weights) <= cfg.Capacity
	})

	p.SolutionStringFn = fn.StringSubset(cfg.Items)
	return p, cfg
}

// Knapsack
func knapsack(name string) *discrete.Problem {
	p, cfg := newKnapsackProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.Value)
	return p
}

// Quadratic Knapsack
func quadraticKnapsack(name string) *discrete.Problem {
	p, cfg := newKnapsackProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Base knapsack value
		baseValue := fn.ScoreSumWeightedValues(p.Variables, cfg.Value)(solution)
		// Compute bonus value of item pairs
		selected := ds.SetFrom(list.MapList(fn.AsSubset(solution), cfg.Items))
		var bonusValue discrete.Score = 0
		for pair, bonus := range cfg.PairBonus {
			if selected.Has(pair[0]) && selected.Has(pair[1]) {
				bonusValue += bonus
			}
		}
		return baseValue + bonusValue
	}
	return p
}
