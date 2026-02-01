package problem

import (
	"github.com/roidaradal/fn/list"
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
	p, graph := newGraphPartitionProblem(name)
	if p == nil || graph == nil {
		return nil
	}

	// Minimize number of cliques used
	p.Goal = discrete.Minimize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each partition group of vertices is a clique
		return list.All(fn.AsPartition(solution, p.UniformDomain()), func(group []discrete.Variable) bool {
			vertices := list.MapList(group, graph.Vertices)
			return fn.IsClique(graph.Graph, vertices)
		})
	})
	return p
}
