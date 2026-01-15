package problem

import (
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Minimum Spanning Tree problem
func MinimumSpanningTree(n int) *discrete.Problem {
	name := newName(MST, n)
	graph, edgeWeight := fn.NewWeightedGraph(name)
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

	// Constraint: all vertices are spanned
	p.AddUniversalConstraint(constraint.AllVerticesSpanned(graph, graph.Vertices))

	// Constraint: solution forms a tree: all vertices reachable from tree traversal
	p.AddUniversalConstraint(constraint.SpanningTree(graph, graph.Vertices))

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
