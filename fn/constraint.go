package fn

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
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
