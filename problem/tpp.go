package problem

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Traveling Purchaser problem
func TravelingPurchaser(n int) *discrete.Problem {
	name := newName(TPP, n)
	cfg := newTravelingPurchaser(name)
	if cfg == nil {
		return nil
	}
	numItems, numMarkets := len(cfg.items), len(cfg.markets)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Path

	// Gather valid (item, market) pairs
	itemMarkets := make([][2]int, 0)
	names := make([]string, 0)
	for i := range numItems {
		for m := range numMarkets {
			if cfg.price[i][m] == a.Inf {
				continue // skip infinite price = not available
			}
			itemMarkets = append(itemMarkets, [2]int{i, m})
			names = append(names, fmt.Sprintf("%s@%s", cfg.items[i], cfg.markets[m]))
		}
	}

	p.Variables = discrete.Variables(itemMarkets)
	domain := discrete.PathDomain(len(itemMarkets))
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check each item is covered once
		count := dict.NewCounter(cfg.items)
		for idx, order := range solution.Map {
			if order < 0 {
				continue
			}
			itemIdx := itemMarkets[idx][0]
			count[cfg.items[itemIdx]] += 1
		}
		return list.AllEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var totalCost discrete.Score = 0
		// Build path
		path := fn.AsPathOrder(solution)
		// Compute item prices
		for _, idx := range path {
			itemMarket := itemMarkets[idx]
			item, market := itemMarket[0], itemMarket[1]
			totalCost += cfg.price[item][market]
		}
		// Compute path cost
		for i := range len(path) - 1 {
			curr, next := path[i], path[i+1]
			market1 := itemMarkets[curr][1]
			market2 := itemMarkets[next][1]
			totalCost += cfg.cost[market1][market2]
		}

		// Add fromOrigin and toOrigin
		start, end := path[0], list.Last(path, 1)
		marketStart := itemMarkets[start][1]
		marketEnd := itemMarkets[end][1]
		totalCost += cfg.fromOrigin[marketStart]
		totalCost += cfg.toOrigin[marketEnd]
		return totalCost
	}

	// TODO: Improve the SolutionStringFn
	// TODO: Add SolutionCoreFn
	p.SolutionStringFn = func(solution *discrete.Solution) string {
		// Build path
		path := fn.AsPathOrder(solution)
		out := list.MapList(path, names)
		return strings.Join(out, " -> ")
	}

	return p
}

type tppCfg struct {
	items      []string
	markets    []string
	price      [][]float64 // Price of (Item, Market)
	cost       [][]float64 // Cost of going from (Market1, Market2)
	fromOrigin []float64   // Cost of going from Origin to Market
	toOrigin   []float64   // Cost of going from Market to Origin
}

// Load traveling purchaser test case
func newTravelingPurchaser(name string) *tppCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 6 {
		return nil
	}
	cfg := &tppCfg{}
	cfg.items = strings.Fields(lines[0])
	cfg.markets = strings.Fields(lines[1])
	numItems, numMarkets := len(cfg.items), len(cfg.markets)
	cfg.price = make([][]float64, numItems)
	cfg.cost = make([][]float64, numMarkets)
	idx := 2
	for i := range numItems {
		cfg.price[i] = fn.ParseFloatInfRow(strings.Fields(lines[idx]), true)
		idx += 1
	}
	for i := range numMarkets {
		cfg.cost[i] = fn.ParseFloatInfRow(strings.Fields(lines[idx]), true)
		idx += 1
	}
	cfg.fromOrigin = fn.ParseFloatInfRow(strings.Fields(lines[idx]), true)
	cfg.toOrigin = fn.ParseFloatInfRow(strings.Fields(lines[idx+1]), true)
	return cfg
}
