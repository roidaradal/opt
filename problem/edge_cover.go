package problem

import (
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

	p.AddUniversalConstraint(fn.ConstraintAllVerticesCovered(graph.Graph, graph.Vertices))
	return p
}
