package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewWarehouseLocation creates a new Warehouse Location problem
func NewWarehouseLocation(variant string, n int) *discrete.Problem {
	name := newName(WarehouseLocation, variant, n)
	switch variant {
	case "basic":
		return warehouseLocation(name)
	default:
		return nil
	}
}

// Warehouse Location
func warehouseLocation(name string) *discrete.Problem {
	cfg := data.NewWarehouse(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Stores)
	domain := discrete.Domain(cfg.Warehouses)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check warehouse usage doesn't exceed their capacity
		usage := dict.TallyValues(solution.Map, domain)
		return list.AllTrue(list.IndexedMap(domain, func(i, w int) bool {
			return usage[w] <= cfg.Capacity[i]
		}))
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var totalCost discrete.Score = 0
		// Warehouse Cost
		usage := dict.TallyValues(solution.Map, domain)
		for i, w := range domain {
			if usage[w] > 0 {
				totalCost += cfg.WarehouseCost[i]
			}
		}
		// Store Cost
		for x, w := range solution.Map {
			totalCost += cfg.StoreCost[x][w]
		}
		return totalCost
	}

	p.SolutionStringFn = fn.StringPartition(domain, cfg.Stores)
	return p
}
