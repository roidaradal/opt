package problem

import (
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Graph Coloring problem
func GraphColoring(n int) *discrete.Problem {
	name := newName(GRAPH_COLOR, n)
	graph, colors, asMap := newGraphColoring(name)
	if graph == nil || colors == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.MapDomain(colors)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		color := solution.Map
		// For all graph edges, check that color of the 2 vertices are different
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return color[x1] != color[x2]
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_LookupValueOrder(p)
	if asMap {
		p.SolutionStringFn = fn.String_Map(p, graph.Vertices, colors)
	} else {
		p.SolutionStringFn = fn.String_Values(p, colors)
	}

	return p
}

// Load graph coloring test case
func newGraphColoring(name string) (*ds.Graph, []string, bool) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 4 {
		return nil, nil, false
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	colors := strings.Fields(lines[2])
	asMap := lines[3] == "map"
	return graph, colors, asMap
}
