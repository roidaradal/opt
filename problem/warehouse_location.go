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
	case "minimax":
		return minimaxWarehouseLocation(name)
	case "maxmin":
		return maxminWarehouseLocation(name)
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

// Common steps for creating Warehouse Subset problems
func newWarehouseSubsetProblem(name string) (*discrete.Problem, *data.Warehouse) {
	cfg := data.NewWarehouse(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.Warehouses)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check subset size is equal to warehouse count
		return len(fn.AsSubset(solution)) == cfg.Count
	})

	p.SolutionStringFn = fn.StringSubset(cfg.Warehouses)
	return p, cfg
}

// Minimax Warehouse Location
func minimaxWarehouseLocation(name string) *discrete.Problem {
	p, cfg := newWarehouseSubsetProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find maximum distance of any store to any selected warehouse
		var maxDistance discrete.Score = 0
		warehouses := fn.AsSubset(solution)
		for store := range cfg.Stores {
			for _, warehouse := range warehouses {
				maxDistance = max(maxDistance, cfg.Distance[warehouse][store])
			}
		}
		return maxDistance
	}
	return p
}

// Maxmin Warehouse Location
func maxminWarehouseLocation(name string) *discrete.Problem {
	p, cfg := newWarehouseSubsetProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find minimum distance of any store to any selected warehouse
		var minDistance = discrete.Inf
		warehouses := fn.AsSubset(solution)
		for store := range cfg.Stores {
			for _, warehouse := range warehouses {
				minDistance = min(minDistance, cfg.Distance[warehouse][store])
			}
		}
		return minDistance
	}
	return p
}
