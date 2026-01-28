package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// TODO: Transfer definitions of IsIndependentSet, IsDominatingSet, etc from fn to opt

// DominatingSet creates a new Dominating Set problem
func DominatingSet(variant string, n int) *discrete.Problem {
	name := newName(DOMSET, variant, n)
	switch variant {
	case "basic":
		return dominatingSetBasic(name)
	case "edge":
		return dominatingSetEdge(name)
	case "eff":
		return dominatingSetEfficient(name)
	default:
		return nil
	}
}

// Common steps for creating Dominating Set problem
func dominatingSetProblem(name string, getVariables func(*ds.Graph) []string) (*discrete.Problem, *ds.Graph) {
	graph, _ := newUnweightedGraph(name)
	if graph == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	variables := getVariables(graph)
	p.Variables = discrete.Variables(variables)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(variables)

	return p, graph
}

// Basic Dominating Set
func dominatingSetBasic(name string) *discrete.Problem {
	p, graph := dominatingSetProblem(name, graphVertices)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms a dominating set:
		// all vertices are either in the set or has neighbor in the set
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return graph.IsDominatingSet(vertices)
	})
	return p
}

// Edge Dominating Set
func dominatingSetEdge(name string) *discrete.Problem {
	p, graph := dominatingSetProblem(name, (*ds.Graph).EdgeNames)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an edge dominating set:
		// all edges have at least one endpoint covered by an edge in the set
		edges := list.MapList(fn.AsSubset(solution), graph.Edges)
		return graph.IsEdgeDominatingSet(edges)
	})
	return p
}

// Efficient Dominating Set
func dominatingSetEfficient(name string) *discrete.Problem {
	p, graph := dominatingSetProblem(name, graphVertices)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an efficient dominating set:
		// all vertices are dominated (in the set or has neighbor) exactly once
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return graph.IsEfficientDominatingSet(vertices)
	})
	return p
}

// Temporary: get graph vertices
// TODO: transfer this to ds.Graph in fn package
func graphVertices(graph *ds.Graph) []string {
	return graph.Vertices
}
