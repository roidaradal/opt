package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new 0-1 Knapsack problem
func Knapsack(n int) *discrete.Problem {
	name := newName(KNAPSACK, n)
	cfg := newKnapsack(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.items)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check of sum of weighted items does not exceed capacity
		count, weight := solution.Map, cfg.weight
		weights := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * weight[x]
		})
		return list.Sum(weights) <= cfg.capacity
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, cfg.value)
	p.SolutionStringFn = fn.String_Subset(cfg.items)

	return p
}

type knapsackCfg struct {
	capacity float64
	items    []string
	weight   []float64
	value    []float64
}

// Load knapsack test case
func newKnapsack(name string) *knapsackCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 4 {
		return nil
	}
	return &knapsackCfg{
		capacity: number.ParseFloat(lines[0]),
		items:    strings.Fields(lines[1]),
		weight:   list.Map(strings.Fields(lines[2]), number.ParseFloat),
		value:    list.Map(strings.Fields(lines[3]), number.ParseFloat),
	}
}
