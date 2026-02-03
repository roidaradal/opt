package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewEdgeColoring creates a new Edge Coloring problem
func NewEdgeColoring(variant string, n int) *discrete.Problem {
	name := newName(EdgeColoring, variant, n)
	switch variant {
	case "basic":
		return edgeColoring(name)
	default:
		return nil
	}
}

// Edge Coloring
func edgeColoring(name string) *discrete.Problem {
	p, cfg := newGraphColoringProblem(name, data.GraphEdges, data.GraphColors)
	if p == nil || cfg == nil || len(cfg.Colors) == 0 {
		return nil
	}
	graph := cfg.Graph

	edgeIndex := make(dict.IntMap)
	for i, edge := range graph.Edges {
		edgeIndex[edge.String()] = i
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		color := solution.Map
		return list.All(graph.Vertices, func(vertex ds.Vertex) bool {
			// For each vertex, check that all edges connected to it have different colors
			edgeColors := list.Map(graph.EdgesOf[vertex], func(edge ds.Edge) discrete.Value {
				return color[edgeIndex[edge.String()]]
			})
			return list.AllUnique(edgeColors)
		})
	})

	p.ObjectiveFn = fn.ScoreCountUniqueValues
	p.SolutionCoreFn = fn.CoreLookupValueOrder(p)
	return p
}
