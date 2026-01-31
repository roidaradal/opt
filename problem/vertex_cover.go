package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
)

// NewVertexCover creates a new Vertex Cover problem
func NewVertexCover(variant string, n int) *discrete.Problem {
	name := newName(VertexCover, variant, n)
	switch variant {
	case "basic":
		return vertexCover(name)
	default:
		return nil
	}
}

// Vertex Cover
func vertexCover(name string) *discrete.Problem {
	p, graph := newGraphCoverProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check for all edges, at least one vertex is covered by the solution subset
		used := solution.Map
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return used[x1]+used[x2] > 0 // at least 1 is covered
		})
	})
	return p
}
