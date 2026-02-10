package problem

import (
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
)

// NewSteinerTree creates a new Steiner Tree problem
func NewSteinerTree(variant string, n int) *discrete.Problem {
	name := newName(SteinerTree, variant, n)
	switch variant {
	case "basic":
		return steinerTree(name)
	default:
		return nil
	}
}

// Steiner Tree
func steinerTree(name string) *discrete.Problem {
	p, graph := newSpanningTreeProblem(name, data.SpanTerminals)
	p = edgeWeightedProblem(p, graph)
	if p == nil || graph == nil || len(graph.Terminals) == 0 {
		return nil
	}
	return p
}
