package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Max Weight Matching problem
func MaxWeightMatching(n int) *discrete.Problem {
	name := newName(MAX_WEIGHT_MATCH, n)
	graph, edgeWeight := fn.NewWeightedGraph(name)
	if graph == nil || edgeWeight == nil {
		return nil
	}
	edgeNames := list.Map(graph.Edges, ds.Edge.String)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(graph.Edges)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	p.AddUniversalConstraint(constraint.GraphMatching(graph))

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
