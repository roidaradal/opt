package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewGraphMatching creates a new Graph Matching problem
func NewGraphMatching(variant string, n int) *discrete.Problem {
	name := newName(GraphMatching, variant, n)
	switch variant {
	case "cardinal":
		return maxCardinalityMatching(name)
	case "weighted":
		return maxWeightMatching(name)
	case "rainbow":
		return rainbowMatching(name)
	default:
		return nil
	}
}

// Common steps to creating a Graph Matching problem
func newGraphMatchingProblem(name string) (*discrete.Problem, *data.Graph) {
	graph := data.NewUndirectedGraph(name)
	if graph == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(graph.Edges)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Graph matching constraint
		count := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			count[v1] += 1
			count[v2] += 1
		}
		// Check all vertices covered by matching are only covered once
		return list.AllEqual(dict.Values(count), 1)
	})

	p.SolutionStringFn = fn.StringSubset(graph.EdgeNames())
	return p, graph
}

// Max Cardinality Matching
// Note: Max Bipartite Matching is a special case of Max Cardinality Matching
func maxCardinalityMatching(name string) *discrete.Problem {
	p, _ := newGraphMatchingProblem(name)
	if p == nil {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSubsetSize
	return p
}

// Max Weight Matching
func maxWeightMatching(name string) *discrete.Problem {
	p, graph := newGraphMatchingProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	if len(graph.Edges) != len(graph.EdgeWeight) {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, graph.EdgeWeight)
	return p
}

// Rainbow Matching
func rainbowMatching(name string) *discrete.Problem {
	p, graph := newGraphMatchingProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	if len(graph.Edges) != len(graph.EdgeColor) {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that selected edges have different colors
		return list.AllUnique(list.MapList(fn.AsSubset(solution), graph.EdgeColor))
	})
	p.ObjectiveFn = fn.ScoreSubsetSize
	return p
}
