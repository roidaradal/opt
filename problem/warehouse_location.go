package problem

import (
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Warehouse Location problem
func WarehouseLocation(n int) *discrete.Problem {
	name := newName(WAREHOUSE, n)
	cfg := newWarehouseLocation(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.stores)
	domain := discrete.MapDomain(cfg.warehouses)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Tally the number of time each warehouse is used
		usage := dict.TallyValues(solution.Map, domain)
		// Check all warehouse usage don't exceed their capacity
		return list.AllTrue(list.IndexedMap(domain, func(i, w int) bool {
			return usage[w] <= cfg.capacity[i]
		}))
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var totalCost discrete.Score = 0
		// Fixed cost
		usage := dict.TallyValues(solution.Map, domain)
		for i, w := range domain {
			if usage[w] > 0 {
				totalCost += cfg.fixedCost[i]
			}
		}
		// Supply cost
		for x, w := range solution.Map {
			totalCost += cfg.supplyCost[x][w]
		}
		return totalCost
	}

	p.SolutionStringFn = fn.String_Partitions(domain, cfg.stores)

	return p
}

type warehouseCfg struct {
	warehouses []string // size M
	stores     []string // size N
	capacity   []int
	fixedCost  []float64
	supplyCost [][]float64 // 1 row per store => vector of size M (cost per warehouse)
}

// Load warehouse location test case
func newWarehouseLocation(name string) *warehouseCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 4 {
		return nil
	}
	cfg := &warehouseCfg{
		warehouses: strings.Fields(lines[0]),
		capacity:   list.Map(strings.Fields(lines[1]), number.ParseInt),
		fixedCost:  list.Map(strings.Fields(lines[2]), number.ParseFloat),
		stores:     make([]string, 0),
		supplyCost: make([][]float64, 0),
	}
	for _, line := range lines[3:] {
		parts := str.CleanSplit(line, ":")
		cost := list.Map(strings.Fields(parts[1]), number.ParseFloat)
		cfg.stores = append(cfg.stores, parts[0])
		cfg.supplyCost = append(cfg.supplyCost, cost)
	}
	return cfg
}
