package problem

import (
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Quadratic Knapsack problem
func QuadraticKnapsack(n int) *discrete.Problem {
	name := newName(QKP, n)
	cfg := newQuadraticKnapsack(name)
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

	p.AddUniversalConstraint(constraint.Knapsack(p, cfg.capacity, cfg.weight))

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		count := solution.Map
		items := ds.SetFrom(fn.AsSubset(solution))
		baseValue := list.Sum(list.Map(p.Variables, func(x discrete.Variable) discrete.Score {
			return float64(count[x]) * cfg.value[x]
		}))
		var bonusValue discrete.Score = 0
		for pair, bonus := range cfg.pairBonus {
			item1, item2 := pair[0], pair[1]
			if items.Has(item1) && items.Has(item2) {
				bonusValue += bonus
			}
		}
		return baseValue + bonusValue
	}

	p.SolutionStringFn = fn.String_Subset(cfg.items)

	return p
}

type quadraticKnapsackCfg struct {
	knapsackCfg
	pairBonus map[[2]int]float64
}

// Load quadratic knapsack test case
func newQuadraticKnapsack(name string) *quadraticKnapsackCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 5 {
		return nil
	}
	cfg := &quadraticKnapsackCfg{
		pairBonus: make(map[[2]int]float64),
	}
	cfg.capacity = number.ParseFloat(lines[0])
	cfg.items = strings.Fields(lines[1])
	cfg.weight = list.Map(strings.Fields(lines[2]), number.ParseFloat)
	cfg.value = list.Map(strings.Fields(lines[3]), number.ParseFloat)
	indexOf := list.IndexMap(cfg.items)
	for _, line := range lines[4:] {
		parts := strings.Fields(line)
		var1 := indexOf[parts[0]]
		var2 := indexOf[parts[1]]
		pair := [2]int{var1, var2}
		cfg.pairBonus[pair] = number.ParseFloat(parts[2])
	}
	return cfg
}
