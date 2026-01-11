package problem

import (
	"slices"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Min-Degree Spanning Tree problem
func MinDegreeSpanningTree(n int) *discrete.Problem {
	name := newName(MDST, n)
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

	// Constraint: all vertices are spanned
	p.AddUniversalConstraint(constraint.AllVerticesSpanned(graph, graph.Vertices))

	// Constraint: solution forms a tree: all vertices reachable from tree traversal
	p.AddUniversalConstraint(constraint.SpanningTree(graph, graph.Vertices))

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count the degree of each vertex from the spanning tree
		count := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			count[v1] += 1
			count[v2] += 1
		}
		maxDegree := slices.Max(dict.Values(count))
		return discrete.Score(maxDegree)
	}

	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}
