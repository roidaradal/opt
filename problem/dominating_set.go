package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewDominatingSet creates a new Dominating Set problem
func NewDominatingSet(variant string, n int) *discrete.Problem {
	name := newName(DominatingSet, variant, n)
	switch variant {
	case "basic":
		return dominatingSet(name)
	case "edge":
		return edgeDominatingSet(name)
	case "efficient":
		return efficientDominatingSet(name)
	default:
		return nil
	}
}

// Dominating Set
func dominatingSet(name string) *discrete.Problem {
	p, graph := newGraphCoverProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms a dominating set:
		// all vertices are either in the set or has neighbor in the set
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		// IsDominatingSet check
		if len(vertices) == 0 {
			return false
		}
		vertexSet := ds.SetFrom(vertices)
		return list.All(graph.Vertices, func(vertex ds.Vertex) bool {
			adjacent := ds.SetFrom(graph.Neighbors(vertex))
			adjacent.Add(vertex)
			return vertexSet.Intersection(adjacent).NotEmpty()
		})
	})
	return p
}

// Edge Dominating Set
func edgeDominatingSet(name string) *discrete.Problem {
	p, graph := newGraphCoverProblem(name, data.GraphEdges)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an edge dominating set:
		// all edges have at least one endpoint covered by an edge in the set
		edges := list.MapList(fn.AsSubset(solution), graph.Edges)
		// IsEdgeDominatingSet check
		// Mark vertex endpoints of given edges as covered
		covered := ds.NewSet[ds.Vertex]()
		for _, edge := range edges {
			v1, v2 := edge.Tuple()
			covered.Add(v1)
			covered.Add(v2)
		}
		// Check that for all edges, at least one endpoint is in covered vertices
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			v1, v2 := edge.Tuple()
			return covered.Has(v1) || covered.Has(v2)
		})
	})
	return p
}

// Efficient Dominating Set
func efficientDominatingSet(name string) *discrete.Problem {
	p, graph := newGraphCoverProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an efficient dominating set:
		// all vertices are dominated (in the set or has neighbor) exactly once
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		// IsEfficientDominatingSet check
		if len(vertices) == 0 {
			return false
		}
		vertexSet := ds.SetFrom(vertices)
		return list.All(graph.Vertices, func(vertex ds.Vertex) bool {
			// Check that all vertices are only dominated by exactly one vertex in the set
			adjacent := ds.SetFrom(graph.Neighbors(vertex))
			adjacent.Add(vertex)
			return vertexSet.Intersection(adjacent).Len() == 1
		})
	})
	return p
}
