package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Edge Cover problem
func EdgeCover(n int) *discrete.Problem {
	name := newName(EDGE_COVER, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}
	edgeNames := graph.EdgeNames()

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(graph.Edges)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check for all vertices, covered by at least one edge endpoint in solution subset
		count := dict.NewCounter(graph.Vertices)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			count[v1] += 1
			count[v2] += 1
		}
		return list.AllGreater(dict.Values(count), 0)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
