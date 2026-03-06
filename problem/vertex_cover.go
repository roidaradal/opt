package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewVertexCover creates a new Vertex Cover problem
func NewVertexCover(variant string, n int) *discrete.Problem {
	name := newName(VertexCover, variant, n)
	switch variant {
	case "basic":
		return vertexCover(name)
	case "weighted":
		return weightedVertexCover(name)
	default:
		return nil
	}
}

// Common steps for creating Vertex Cover problem
func newVertexCoverProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphCoverProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil, nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check for all edges, at least one vertex is covered by the solution subset
		used := solution.Map
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return used[x1]+used[x2] > 0 // at least 1 is covered
		})
	})
	return p, graph
}

// Vertex Cover
func vertexCover(name string) *discrete.Problem {
	p, _ := newVertexCoverProblem(name)
	return p
}

// Weighted Vertex Cover
func weightedVertexCover(name string) *discrete.Problem {
	p, graph := newVertexCoverProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	if len(graph.Vertices) != len(graph.VertexWeight) {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, graph.VertexWeight)
	return p
}
