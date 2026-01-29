package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// IndependentSet creates a new Independent Set problem
func IndependentSet(variant string, n int) *discrete.Problem {
	name := newName(INDEPSET, variant, n)
	switch variant {
	case "basic":
		return independentSetBasic(name)
	case "rainbow":
		return independentSetRainbow(name)
	default:
		return nil
	}
}

// Common steps of creating IndependentSet problem
func independentSetProblem(name string, graph *graphCfg) *discrete.Problem {
	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(graph.Vertices)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an independent set:
		// none of the vertices are connected to each other
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return graph.IsIndependentSet(vertices)
	})

	p.Goal = discrete.Maximize
	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(graph.Vertices)

	return p
}

// Independent Set problem
func independentSetBasic(name string) *discrete.Problem {
	graph := newUnweightedGraph(name)
	if graph == nil {
		return nil
	}
	return independentSetProblem(name, graph)
}

// Rainbow Independent Set problem
func independentSetRainbow(name string) *discrete.Problem {
	graph := newUnweightedGraph(name)
	if graph == nil || len(graph.extra) < 1 {
		return nil
	}
	colorOf := fn.StringList(graph.extra[0][0])

	p := independentSetProblem(name, graph)
	if p == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that vertices have different colors
		colors := list.MapList(fn.AsSubset(solution), colorOf)
		return list.AllUnique(colors)
	})

	return p
}
