package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Vertex Cover problem
func VertexCover(n int) *discrete.Problem {
	name := newName(VERTEX_COVER, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check for all edges, at least one vertex is covered by solution subset
		used := solution.Map
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return used[x1]+used[x2] > 0 // at least 1 is covered
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(graph.Vertices)

	return p
}
