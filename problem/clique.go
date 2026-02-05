package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewClique creates a new Clique problem
func NewClique(variant string, n int) *discrete.Problem {
	name := newName(Clique, variant, n)
	switch variant {
	case "basic":
		return clique(name)
	case "k":
		return kClique(name)
	default:
		return nil
	}
}

// Common steps for creating Clique problem
func newCliqueProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphSubsetProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil, nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms a clique:
		// all vertices are connected to each other
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return fn.IsClique(graph.Graph, vertices)
	})
	p.Goal = discrete.Maximize
	return p, graph
}

// Clique
func clique(name string) *discrete.Problem {
	p, _ := newCliqueProblem(name)
	return p
}

// K-Clique
func kClique(name string) *discrete.Problem {
	p, graph := newCliqueProblem(name)
	if p == nil || graph == nil || graph.K == 0 {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that clique size is K
		return len(fn.AsSubset(solution)) == graph.K
	})
	p.Goal = discrete.Satisfy
	p.ObjectiveFn = nil
	return p
}
