package problem

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Max Induced Path problem
func MaxInducedPath(n int) *discrete.Problem {
	name := newName(MAX_INDUCED_PATH, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Path

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.PathDomain(len(graph.Vertices))
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		path := fn.AsPathOrder(solution)
		if len(path) == 0 {
			return false
		}
		for i := range len(path) - 1 {
			// Check that consecutive vertices in the path have edge between them
			vertex1 := graph.Vertices[path[i]]
			vertex2 := graph.Vertices[path[i+1]]
			neighbors := graph.Neighbors(vertex1)
			if !slices.Contains(neighbors, vertex2) {
				return false
			}
			// Check that the non-adjacent vertices to the current don't have edge
			for _, x := range path[i+2:] {
				if slices.Contains(neighbors, graph.Vertices[x]) {
					return false
				}
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Number of points in the path (those not assigned -1)
		points := list.Filter(solution.Values(), func(value discrete.Value) bool {
			return value >= 0
		})
		return discrete.Score(len(points))
	}

	p.SolutionStringFn = func(solution *discrete.Solution) string {
		path := fn.AsPathOrder(solution)
		out := list.MapList(path, graph.Vertices)
		return strings.Join(out, "-")
	}
	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		path := fn.AsPathOrder(solution)
		out := list.MapList(path, graph.Vertices)
		first, last := out[0], list.Last(out, 1)
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(out)
		}
		return strings.Join(out, "-")
	}

	return p
}
