package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Note: Max Bipartite Match is a special case of Max Cardinality Match

// Create new Max Cardinality Matching problem
func MaxCardinalityMatching(n int) *discrete.Problem {
	name := newName(MAX_CARDINALITY_MATCH, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
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

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
