package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
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
	case "weighted":
		return weightedIndependentSet(name)
	default:
		return nil
	}
}

// Common steps for creating Independent Set problem
func newIndependentSetProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphSubsetProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil, nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an independent set:
		// none of the vertices are connected to each other
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		// IsIndependentSet check
		vertexSet := ds.SetFrom(vertices)
		for _, vertex := range vertices {
			adjacent := ds.SetFrom(graph.Neighbors(vertex))
			if vertexSet.Intersection(adjacent).NotEmpty() {
				return false
			}
		}
		return true
	})

	p.Goal = discrete.Maximize
	return p, graph
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

	// Check vertices have different colors
	p.AddUniversalConstraint(fn.ConstraintRainbowColoring(graph.VertexColor))
	return p
}

// Weighted Independent Set
func weightedIndependentSet(name string) *discrete.Problem {
	p, graph := newIndependentSetProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	if len(graph.Vertices) != len(graph.VertexWeight) {
		return nil
	}

	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, graph.VertexWeight)
	return p
}
