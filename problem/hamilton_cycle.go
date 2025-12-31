package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Hamilton Cycle problem
func HamiltonCycle(n int) *discrete.Problem {
	name := newName(HAMILTON_CYCLE, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.IndexDomain(len(graph.Vertices))
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that the path of vertex sequence formed by solution
		// forms a Hamiltonian cycle: visits each vertex exactly once and returns to starting point
		vertices := list.MapList(fn.AsSequence(solution), graph.Vertices)
		return graph.IsHamiltonianCycle(vertices)
	}
	p.AddUniversalConstraint(test)

	p.SolutionCoreFn = fn.Core_SortedCycle(graph.Vertices)
	p.SolutionStringFn = fn.String_Sequence(graph.Vertices)

	return p
}
