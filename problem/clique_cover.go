package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewCliqueCover creates a new Clique Cover problem
func NewCliqueCover(variant string, n int) *discrete.Problem {
	name := newName(CliqueCover, variant, n)
	switch variant {
	case "basic":
		return cliqueCover(name)
	default:
		return nil
	}
}

// Clique Cover
func cliqueCover(name string) *discrete.Problem {
	graph := data.NewUndirectedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	domain := discrete.RangeDomain(1, len(graph.Vertices))
	p.Variables = discrete.Variables(graph.Vertices)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each partition group of vertices is a clique
		return list.All(fn.AsPartition(solution, domain), func(group []discrete.Variable) bool {
			vertices := list.MapList(group, graph.Vertices)
			return fn.IsClique(graph.Graph, vertices)
		})
	})

	// Minimize number of cliques used
	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScoreCountUniqueValues

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.StringPartition(domain, graph.Vertices)
	return p
}
