package problem

import (
	"strings"

	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Quadratic Assignment problem
func QuadraticAssignment(n int) *discrete.Problem {
	name := newName(QUAD_ASSIGNMENT, n)
	cfg := newQuadraticAssignment(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.IndexVariables(cfg.size)
	domain := discrete.IndexDomain(cfg.size)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For each pair, sum up the cost of the flow from facility1 => facility2
		// multiplied by the cost of traveling from location1 (fac1's value) => location2 (fac2's value)
		var totalCost discrete.Score = 0
		for _, pair := range comb.Combinations(p.Variables, 2) {
			var1, var2 := pair[0], pair[1]
			val1, val2 := solution.Map[var1], solution.Map[var2]
			totalCost += cfg.flow[var1][var2] * cfg.cost[val1][val2] // 1 => 2
			totalCost += cfg.flow[var2][var1] * cfg.cost[val2][val1] // 2 => 1
		}
		return totalCost
	}

	p.SolutionStringFn = fn.String_Sequence(list.NumRange(1, cfg.size+1))

	return p
}

type quadraticAssignmentCfg struct {
	size int
	cost [][]float64
	flow [][]float64
}

// Load quadratic assignment test case
func newQuadraticAssignment(name string) *quadraticAssignmentCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	size := number.ParseInt(lines[0])
	cfg := &quadraticAssignmentCfg{
		size: size,
		cost: make([][]float64, size),
		flow: make([][]float64, size),
	}
	idx := 1
	for i := range size {
		cfg.cost[i] = list.Map(strings.Fields(lines[idx]), number.ParseFloat)
		idx++
	}
	for i := range size {
		cfg.flow[i] = list.Map(strings.Fields(lines[idx]), number.ParseFloat)
		idx++
	}
	return cfg
}
