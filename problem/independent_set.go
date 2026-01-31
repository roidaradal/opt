package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewIndependentSet creates a new Independent Set problem
func NewIndependentSet(variant string, n int) *discrete.Problem {
	name := newName(IndependentSet, variant, n)
	switch variant {
	case "basic":
		return independentSet(name)
	case "rainbow":
		return rainbowIndependentSet(name)
	default:
		return nil
	}
}

// Independent Set
func independentSet(name string) *discrete.Problem {
	p, _ := newIndependentSetProblem(name)
	return p
}

// Rainbow Independent Set
func rainbowIndependentSet(name string) *discrete.Problem {
	p, graph := newIndependentSetProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	if len(graph.Vertices) != len(graph.VertexColor) {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that vertices have different colors
		colors := list.MapList(fn.AsSubset(solution), graph.VertexColor)
		return list.AllUnique(colors)
	})
	return p
}
