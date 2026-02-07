package problem

import (
	"slices"

	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewVertexColoring creates a new Vertex Coloring problem
func NewVertexColoring(variant string, n int) *discrete.Problem {
	name := newName(VertexColoring, variant, n)
	switch variant {
	case "basic":
		return vertexColoring(name)
	case "complete":
		return completeColoring(name)
	case "harmonious":
		return harmoniousColoring(name)
	default:
		return nil
	}
}

// Common steps for creating Vertex Coloring problem
func newVertexColoringProblem(name string) (*discrete.Problem, *data.GraphColoring) {
	p, cfg := newGraphColoringProblem(name, data.GraphVertices, data.GraphColors)
	if p == nil || cfg == nil || len(cfg.Colors) == 0 {
		return nil, nil
	}

	p.AddUniversalConstraint(fn.ConstraintProperVertexColoring(cfg.Graph))
	p.ObjectiveFn = fn.ScoreCountUniqueValues
	p.SolutionCoreFn = fn.CoreLookupValueOrder(p)
	return p, cfg
}

// Vertex Coloring
func vertexColoring(name string) *discrete.Problem {
	p, _ := newVertexColoringProblem(name)
	return p
}

// Complete Coloring
func completeColoring(name string) *discrete.Problem {
	p, cfg := newVertexColoringProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.Goal = discrete.Maximize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		color := solution.Map
		count := make(map[[2]int]int) // ColorPair => Count
		for _, pair := range comb.Combinations(p.UniformDomain(), 2) {
			key := [2]int{pair[0], pair[1]}
			count[key] = 0
		}
		for _, edge := range graph.Edges {
			v1, v2 := edge.Tuple()
			c1, c2 := color[graph.IndexOf[v1]], color[graph.IndexOf[v2]]
			key := [2]int{c1, c2}
			if dict.NoKey(count, key) {
				key = [2]int{c2, c1} // flip if original order not found
			}
			count[key] += 1
		}
		// Check all color pair appears at least once
		return list.AllGreaterEqual(dict.Values(count), 1)
	})
	return p
}

// Harmonious Coloring
func harmoniousColoring(name string) *discrete.Problem {
	p, cfg := newVertexColoringProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		color := solution.Map
		count := make(map[[2]int]int) // ColorPair => Count
		for _, edge := range graph.Edges {
			v1, v2 := edge.Tuple()
			c1, c2 := color[graph.IndexOf[v1]], color[graph.IndexOf[v2]]
			colors := []int{c1, c2}
			slices.Sort(colors)
			key := [2]int{colors[0], colors[1]}
			count[key] += 1
		}
		// Check all color pair appears at most once
		return list.AllLessEqual(dict.Values(count), 1)
	})
	return p
}
