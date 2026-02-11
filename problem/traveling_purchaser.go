package problem

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewTravelingPurchaser creates a new Traveling Purchaser problem
func NewTravelingPurchaser(variant string, n int) *discrete.Problem {
	name := newName(TravelingPurchaser, variant, n)
	switch variant {
	case "basic":
		return travelingPurchaser(name)
	default:
		return nil
	}
}

// Traveling Purchaser
func travelingPurchaser(name string) *discrete.Problem {
	cfg := data.NewGraphTour(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Path
	p.Goal = discrete.Minimize

	// Gather valid (item, market) pairs
	itemMarkets := make([][2]int, 0)
	names := make([]string, 0)
	for i, item := range cfg.Items {
		for m, market := range cfg.Vertices {
			if cfg.Cost[i][m] == discrete.Inf {
				continue // skip infinite price: not available
			}
			itemMarkets = append(itemMarkets, [2]int{i, m})
			names = append(names, fmt.Sprintf("%s@%s", item, market))
		}
	}

	p.Variables = discrete.Variables(itemMarkets)
	p.AddVariableDomains(discrete.PathDomain(len(itemMarkets)))

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each item is covered once
		covered := dict.NewCounter(cfg.Items)
		for idx, order := range solution.Map {
			if order < 0 {
				continue // not in path
			}
			itemIdx := itemMarkets[idx][0]
			covered[cfg.Items[itemIdx]] += 1
		}
		return list.AllEqual(dict.Values(covered), 1)
	})

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var totalCost discrete.Score = 0
		// Create path from solution
		path := fn.AsPathOrder(solution)
		if len(path) == 0 {
			return 0
		}
		// Compute item prices
		for _, idx := range path {
			itemMarket := itemMarkets[idx]
			item, market := itemMarket[0], itemMarket[1]
			totalCost += cfg.Cost[item][market]
		}
		// Compute path cost
		for i := range len(path) - 1 {
			curr, next := path[i], path[i+1]
			market1 := itemMarkets[curr][1]
			market2 := itemMarkets[next][1]
			totalCost += cfg.Distance[market1][market2]
		}
		// Add FromOrigin and ToOrigin
		start, end := path[0], list.Last(path, 1)
		marketStart := itemMarkets[start][1]
		marketEnd := itemMarkets[end][1]
		totalCost += cfg.FromOrigin[marketStart]
		totalCost += cfg.ToOrigin[marketEnd]
		return totalCost
	}

	// TODO: Add SolutionCoreFn
	// TODO: Improve SolutionStringFn
	p.SolutionStringFn = func(solution *discrete.Solution) string {
		path := list.MapList(fn.AsPathOrder(solution), names)
		return strings.Join(path, " -> ")
	}

	return p
}
