package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Minimum Spanning Tree problem
func MinimumSpanningTree(n int) *discrete.Problem {
	name := newName(MST, n)
	graph, edgeWeight := fn.NewWeightedGraph(name)
	if graph == nil || edgeWeight == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	edgeNames := list.Map(graph.Edges, ds.Edge.String)
	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// Constraint: all vertices are spanned
	test1 := func(solution *discrete.Solution) bool {
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
	p.AddUniversalConstraint(test1)

	// Constraint: solution forms a tree: all vertices are reachable
	// by doing a tree traversal
	test2 := func(solution *discrete.Solution) bool {
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
	p.AddUniversalConstraint(test2)

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
