package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Minimum K-Cut problem
func MinimumKCut(n int) *discrete.Problem {
	name := newName(MIN_K_CUT, n)
	graph, edgeWeight, k := fn.NewKWeightedGraph(name)
	if graph == nil || edgeWeight == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	edgeNames := graph.EdgeNames()
	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Remove selected cut edges from active edges
		cutEdges := list.MapList(fn.AsSubset(solution), graph.Edges)
		activeEdges := ds.SetFrom(graph.Edges)
		for _, cutEdge := range cutEdges {
			activeEdges.Delete(cutEdge)
		}
		// Make sure that the cut produced at least k connected components
		components := graph.ConnectedComponents(activeEdges)
		return len(components) >= k
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
