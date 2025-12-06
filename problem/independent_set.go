package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Independent Set problem
func IndependentSet(n int) *discrete.Problem {
	name := newName(INDEPENDENT_SET, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that the subset of vertices formed by the solution
		// forms an independent set: none of the vertices are connected to each other
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return graph.IsIndependentSet(vertices)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(graph.Vertices)

	return p
}
