package problem

import (
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
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
	if p == nil || graph == nil || len(graph.Terminals) == 0 {
		return nil
	}
	if len(graph.Edges) != len(graph.EdgeWeight) {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, graph.EdgeWeight)
	return p
}
