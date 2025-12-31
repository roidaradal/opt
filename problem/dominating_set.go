package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Dominating Set problem
func DominatingSet(n int) *discrete.Problem {
	name := newName(DOMINATING_SET, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that the subset of vertices formed by the solution
		// forms a dominating set: all vertices are either in the set or has a neighbor in the set
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return graph.IsDominatingSet(vertices)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(graph.Vertices)

	return p
}
