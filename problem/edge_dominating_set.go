package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Edge Dominating Set problem
func EdgeDominatingSet(n int) *discrete.Problem {
	name := newName(EDGE_DOMINATING_SET, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	edgeNames := list.Map(graph.Edges, ds.Edge.String)
	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that the subset of edges formed by the solution forms an
		// edge dominating set: all edges have at least one endpoint covered by an edge in the set
		edges := list.MapList(fn.AsSubset(solution), graph.Edges)
		return graph.IsEdgeDominatingSet(edges)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
