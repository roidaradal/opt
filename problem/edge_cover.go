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
	case "weighted":
		return weightedEdgeCover(name)
	default:
		return nil
	}
}

// Create new Edge Cover problem
func newEdgeCoverProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphCoverProblem(name, data.GraphEdges)
	if p == nil || graph == nil {
		return nil, nil
	}

	p.AddUniversalConstraint(fn.ConstraintAllVerticesCovered(graph.Graph, graph.Vertices))
	return p, graph
}

// Edge Cover
func edgeCover(name string) *discrete.Problem {
	p, _ := newEdgeCoverProblem(name)
	return p
}

// Weighted Edge Cover
func weightedEdgeCover(name string) *discrete.Problem {
	p, graph := newEdgeCoverProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	if len(graph.EdgeWeight) != len(graph.Edges) {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, graph.EdgeWeight)
	return p
}
