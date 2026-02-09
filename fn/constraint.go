package fn

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
)

// ConstraintAllUnique makes sure all solution values are unique
func ConstraintAllUnique(solution *discrete.Solution) bool {
	return list.AllUnique(solution.Values())
}

// ConstraintProperVertexColoring makes sure all adjacent vertices in the graph have a different color
func ConstraintProperVertexColoring(graph *ds.Graph) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		color := solution.Map
		// For all graph edges, check that color of 2 vertices are different
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return color[x1] != color[x2]
		})
	}
}

// ConstraintAllVerticesCovered makes sure all vertices are covered at least once
func ConstraintAllVerticesCovered(graph *ds.Graph, vertices []ds.Vertex) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Go through all edges formed by the subset solution
		// Mark the 2 vertices as covered
		covered := dict.Flags(vertices, false)
		for _, x := range AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			covered[v1] = true
			covered[v2] = true
		}
		return list.AllTrue(dict.Values(covered))
	}
}

// ConstraintSpanningTree checks if the solution forms a tree, and all vertices are reachable from tree traversal
func ConstraintSpanningTree(graph *ds.Graph, vertices []ds.Vertex) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		reachable := SpannedVertices(solution, graph)
		if reachable == nil {
			return false
		}
		vertexSet := ds.SetFrom(vertices)
		// Check all vertices are reachable
		return vertexSet.Difference(reachable).IsEmpty()
	}
}

// ConstraintRainbowColoring makes sure all chosen items have different colors
func ConstraintRainbowColoring(colors []string) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		return list.AllUnique(list.MapList(AsSubset(solution), colors))
	}
}

// ConstraintSimplePath makes sure solution forms a valid simple path (no repeated vertices)
func ConstraintSimplePath(cfg *data.GraphPath) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		path := AsGraphPath(solution, cfg)
		prev := path[0]

		visited := ds.NewSet[int]()
		visited.Add(prev)
		for _, curr := range path[1:] {
			if visited.Has(curr) {
				return false // repeated vertex = not simple path
			}
			if cfg.Distance[prev][curr] == discrete.Inf {
				return false // no edge from prev -> curr
			}
			visited.Add(curr)
			prev = curr
		}
		return true
	}
}
