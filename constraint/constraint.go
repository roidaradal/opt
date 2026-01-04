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

// Spanning tree: all vertices are spanned
func AllVerticesSpanned(graph *ds.Graph) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Go through all edges formed by the subset solution
		// 2 vertices of each edge are marked as spanned
		spanned := dict.Flags(graph.Vertices, false)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			spanned[v1] = true
			spanned[v2] = true
		}
		// Check that all vertices are spanned
		return list.AllTrue(dict.Values(spanned))
	}
}

// Spanning tree: solution forms a tree, all vertices reachable from tree traversal
func SpanningTree(graph *ds.Graph) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Get the edges from the subset solution
		edges := list.MapList(fn.AsSubset(solution), graph.Edges)
		if len(edges) == 0 {
			return false
		}
		activeEdges := ds.SetFrom(edges)
		start := edges[0][0] // first edge's first vertex, chosen arbitrarily
		// Perform a BFS traversal starting from the start vertex
		// using only edges from the spanning tree
		reachable := ds.SetFrom(graph.BFSTraversal(start, activeEdges))
		vertexSet := ds.SetFrom(graph.Vertices)
		// Check that all vertices are reachable
		return vertexSet.Difference(reachable).IsEmpty()
	}
}
