package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Clique Cover problem
func CliqueCover(n int) *discrete.Problem {
	name := newName(CLIQUE_COVER, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.RangeDomain(1, len(graph.Vertices))
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that each partition group of vertices is a clique
		return list.All(fn.AsPartitions(solution, domain), func(partition []discrete.Variable) bool {
			vertices := list.MapList(partition, graph.Vertices)
			return graph.IsClique(vertices)
		})
	}
	p.AddUniversalConstraint(test)

	// Minimize the number of cliques used
	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_SortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.String_Partitions(domain, graph.Vertices)

	return p
}
