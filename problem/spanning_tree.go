package problem

import (
	"slices"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSpanningTree creates a new Spanning Tree problem
func NewSpanningTree(variant string, n int) *discrete.Problem {
	name := newName(SpanningTree, variant, n)
	switch variant {
	case "mst":
		return minimumSpanningTree(name)
	case "mdst":
		return minDegreeSpanningTree(name)
	case "kmst":
		return kMinimumSpanningTree(name)
	default:
		return nil
	}
}

// Minimum Spanning Tree
func minimumSpanningTree(name string) *discrete.Problem {
	p, graph := newSpanningTreeProblem(name, data.SpanVertices)
	return edgeWeightedProblem(p, graph)
}

// Min-Degree Spanning Tree
func minDegreeSpanningTree(name string) *discrete.Problem {
	p, graph := newSpanningTreeProblem(name, data.SpanVertices)
	if p == nil || graph == nil {
		return nil
	}
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find the max-degree vertex from spanning tree
		degree := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			degree[v1] += 1
			degree[v2] += 1
		}
		maxDegree := slices.Max(dict.Values(degree))
		return discrete.Score(maxDegree)
	}
	return p
}

// K-Minimum Spanning Tree
func kMinimumSpanningTree(name string) *discrete.Problem {
	p, graph := newSpanningTreeProblem(name, data.SpanVertices)
	p = edgeWeightedProblem(p, graph)
	if p == nil || graph == nil || graph.K == 0 {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Ensure there are k-1 edges in the tree
		edges := list.MapList(fn.AsSubset(solution), graph.Edges)
		if len(edges) != graph.K-1 {
			return false
		}
		// Check that tree formed by edges have k vertices
		reachable := fn.SpannedVertices(solution, graph.Graph)
		return reachable.Len() == graph.K
	})
	return p
}
