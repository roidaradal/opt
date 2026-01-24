package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Knapsack creates a new Knapsack problem
func Knapsack(variant string, n int) *discrete.Problem {
	name := newName(KNAPSACK, variant, n)
	switch variant {
	case "basic":
		return knapsackBasic(name)
	case "quad":
		return knapsackQuadratic(name)
	default:
		return nil
	}
}

// Common create steps of Knapsack problem
func knapsackProblem(name string) (*discrete.Problem, *knapsackCfg) {
	cfg := newKnapsack(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(cfg.items)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check sum of item weights don't exceed capacity
		count := solution.Map
		weights := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * cfg.weight[x]
		})
		return list.Sum(weights) <= cfg.capacity
	})

	p.SolutionStringFn = fn.StringSubset(cfg.items)
	return p, cfg
}

// 0-1 Knapsack problem
func knapsackBasic(name string) *discrete.Problem {
	p, cfg := knapsackProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.value)
	return p
}

// Quadratic Knapsack problem
func knapsackQuadratic(name string) *discrete.Problem {
	p, cfg := knapsackProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Base knapsack value
		baseValue := fn.ScoreSumWeightedValues(p.Variables, cfg.value)(solution)
		// Compute bonus value of item pairs
		items := ds.SetFrom(fn.AsSubset(solution))
		var bonusValue discrete.Score = 0
		for pair, bonus := range cfg.pairBonus {
			item1, item2 := pair[0], pair[1]
			if items.Has(item1) && items.Has(item2) {
				bonusValue += bonus
			}
		}
		return baseValue + bonusValue
	}
	return p
}

type knapsackCfg struct {
	capacity  float64
	items     []string
	weight    []float64
	value     []float64
	pairBonus map[[2]int]float64
}

// Load knapsack test case (basic, quad)
func newKnapsack(name string) *knapsackCfg {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) < 4 {
		return nil
	}
	cfg := &knapsackCfg{
		capacity:  number.ParseFloat(lines[0]),
		items:     fn.StringList(lines[1]),
		weight:    fn.FloatList(lines[2]),
		value:     fn.FloatList(lines[3]),
		pairBonus: make(map[[2]int]float64),
	}
	if len(lines) > 4 {
		indexOf := list.IndexMap(cfg.items)
		for _, line := range lines[4:] {
			parts := fn.StringList(line)
			if len(parts) != 3 {
				continue
			}
			var1, ok1 := indexOf[parts[0]]
			var2, ok2 := indexOf[parts[1]]
			if !ok1 || !ok2 {
				continue
			}
			cfg.pairBonus[[2]int{var1, var2}] = number.ParseFloat(parts[2])
		}
	}
	return cfg
}
