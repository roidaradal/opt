package problem

import (
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new K-MST problem
func KMinimumSpanningTree(n int) *discrete.Problem {
	name := newName(K_MST, n)
	graph, edgeWeight, k := newKMST(name)
	if graph == nil || edgeWeight == nil {
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
		// Ensure there are k-1 edges in the tree
		edges := list.MapList(fn.AsSubset(solution), graph.Edges)
		if len(edges) != k-1 {
			return false
		}
		// Check that the tree formed by the edges have k vertices
		activeEdges := ds.SetFrom(edges)
		start := edges[0][0]
		reachable := ds.SetFrom(graph.BFSTraversal(start, activeEdges))
		return reachable.Len() == k
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}

func newKMST(name string) (*ds.Graph, []float64, int) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 4 {
		return nil, nil, 0
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	edgeWeight := list.Map(strings.Fields(lines[2]), number.ParseFloat)
	k := number.ParseInt(lines[3])
	return graph, edgeWeight, k
}
