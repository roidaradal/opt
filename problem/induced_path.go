package problem

import (
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewInducedPath creates a new Induced Path problem
func NewInducedPath(variant string, n int) *discrete.Problem {
	name := newName(InducedPath, variant, n)
	switch variant {
	case "basic":
		return maxInducedPath(name)
	default:
		return nil
	}
}

// Common steps for creating Induced Path problems
func newInducedPathProblem(name string) (*discrete.Problem, *data.Graph) {
	graph := data.NewUndirectedGraph(name)
	if graph == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Path
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(graph.Vertices)
	p.AddVariableDomains(discrete.PathDomain(len(graph.Vertices)))

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		pathOrder := fn.AsPathOrder(solution)
		if len(pathOrder) == 0 {
			return false
		}
		for i := range len(pathOrder) - 1 {
			// Check that consecutive vertices in path have edge between them
			vertex1 := graph.Vertices[pathOrder[i]]
			vertex2 := graph.Vertices[pathOrder[i+1]]
			neighbors := graph.Neighbors(vertex1)
			if !slices.Contains(neighbors, vertex2) {
				return false
			}
			// Check that non-adjacent vertices to vertex1 don't have an edge
			for _, j := range pathOrder[i+2:] {
				if slices.Contains(neighbors, graph.Vertices[j]) {
					return false
				}
			}
		}
		return true
	})

	p.SolutionStringFn = func(solution *discrete.Solution) string {
		vertices := list.MapList(fn.AsPathOrder(solution), graph.Vertices)
		return strings.Join(vertices, "-")
	}
	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		vertices := list.MapList(fn.AsPathOrder(solution), graph.Vertices)
		return fn.MirroredList(vertices, "-")
	}
	return p, graph
}

// Max Induced Path
func maxInducedPath(name string) *discrete.Problem {
	p, _ := newInducedPathProblem(name)
	if p == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Number of points in path (not assigned -1)
		points := list.Filter(solution.Values(), func(value discrete.Value) bool {
			return value >= 0
		})
		return discrete.Score(len(points))
	}
	return p
}
