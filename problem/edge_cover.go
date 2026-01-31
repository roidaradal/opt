package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewEdgeCover creates a new Edge Cover problem
func NewEdgeCover(variant string, n int) *discrete.Problem {
	name := newName(EdgeCover, variant, n)
	switch variant {
	case "basic":
		return edgeCover(name)
	default:
		return nil
	}
}

// Edge Cover
func edgeCover(name string) *discrete.Problem {
	p, graph := newGraphCoverProblem(name, data.GraphEdges)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check for all vertices, covered by at least one edge endpoint in solution subset
		count := dict.NewCounter(graph.Vertices)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			count[v1] += 1
			count[v2] += 1
		}
		return list.AllGreater(dict.Values(count), 0)
	})
	return p
}
