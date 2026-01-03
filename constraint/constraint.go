// Package constraint contains commonly used constraint functions
package constraint

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// All unique constraint
func AllUnique(solution *discrete.Solution) bool {
	return list.AllUnique(solution.Values())
}

// Knapsack Constraint
func Knapsack(p *discrete.Problem, capacity float64, weight []float64) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Check sum of weighted items does not exceed capacity
		count := solution.Map
		weights := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * weight[x]
		})
		return list.Sum(weights) <= capacity
	}
}

// Graph Matching Constraint
func GraphMatching(graph *ds.Graph) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		count := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			count[v1] += 1
			count[v2] += 1
		}
		// Check all vertices covered by matching are only covered once
		return list.AllEqual(dict.Values(count), 1)
	}
}
