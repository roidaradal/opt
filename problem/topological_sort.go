package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Topological Sort problem
func TopologicalSort(n int) *discrete.Problem {
	name := newName(TOPOLOGICAL, n)
	graph := fn.NewDirectedGraph(name)
	if graph == nil {
		return nil
	}
	numVertices := len(graph.Vertices)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.IndexDomain(numVertices)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		past := ds.NewSet[ds.Vertex]()
		for _, x := range fn.AsSequence(solution) {
			vertex := graph.Vertices[x]
			forward, hasNeighbors := graph.NeighborsOf[vertex]
			// Fail if vertex has a forward neighbor that has already been
			// encountered previously
			if hasNeighbors && forward.Intersection(past).NotEmpty() {
				return false
			}
			past.Add(vertex)
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.SolutionStringFn = fn.String_Sequence(graph.Vertices)

	return p
}
