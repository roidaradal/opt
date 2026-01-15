package problem

import (
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Edge Coloring problem
func EdgeColoring(n int) *discrete.Problem {
	name := newName(EDGE_COLOR, n)
	graph, colors := newEdgeColoring(name)
	if graph == nil || colors == nil {
		return nil
	}

	edgeNames := graph.EdgeNames()
	edgeIndex := make(dict.IntMap)
	for i, edgeName := range edgeNames {
		edgeIndex[edgeName] = i
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.MapDomain(colors)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		color := solution.Map
		return list.All(graph.Vertices, func(vertex ds.Vertex) bool {
			// For each vertex, check that all the edges connected to it have different colors
			edgeColors := list.Map(graph.EdgesOf[vertex], func(edge ds.Edge) discrete.Value {
				return color[edgeIndex[edge.String()]]
			})
			return list.AllUnique(edgeColors)
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_LookupValueOrder(p)
	p.SolutionStringFn = fn.String_Values(p, colors)

	return p
}

// Load edge coloring test case
func newEdgeColoring(name string) (*ds.Graph, []string) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil, nil
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	colors := strings.Fields(lines[2])
	return graph, colors
}
