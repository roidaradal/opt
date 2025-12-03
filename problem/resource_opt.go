package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Resource Optimization problem
func ResourceOptimization(n int) *discrete.Problem {
	name := newName(RESOURCE_OPT, n)
	cfg := newResourceOptimization(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(cfg.resources)
	for i, variable := range p.Variables {
		p.Domain[variable] = discrete.RangeDomain(0, cfg.supply[i])
	}

	test := func(solution *discrete.Solution) bool {
		// Check sum of weighted costs does not exceed limit
		count, cost := solution.Map, cfg.cost
		costs := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * cost[x]
		})
		return list.Sum(costs) <= cfg.limit
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, cfg.value)
	p.SolutionStringFn = fn.String_Values[int](p, nil)

	return p
}

type resourceOptCfg struct {
	limit     float64
	resources []string
	supply    []int
	cost      []float64
	value     []float64
}

// Load resource optimization test case
func newResourceOptimization(name string) *resourceOptCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 5 {
		return nil
	}
	return &resourceOptCfg{
		limit:     number.ParseFloat(lines[0]),
		resources: strings.Fields(lines[1]),
		supply:    list.Map(strings.Fields(lines[2]), number.ParseInt),
		cost:      list.Map(strings.Fields(lines[3]), number.ParseFloat),
		value:     list.Map(strings.Fields(lines[4]), number.ParseFloat),
	}
}
