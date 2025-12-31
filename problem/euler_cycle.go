package problem

import (
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Eulerian Cycle problem
func EulerCycle(n int) *discrete.Problem {
	name := newName(EULER_CYCLE, n)
	graph := fn.NewUnweightedGraph(name)
	if graph == nil {
		return nil
	}
	numEdges := len(graph.Edges)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(graph.Edges)
	domain := discrete.IndexDomain(numEdges)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that the path of edge sequence formed by solution
		// forms an Eulerian cycle: visits each edge exactly once
		// and ends at the vertex where it started
		edgeSequence := list.MapList(fn.AsSequence(solution), graph.Edges)
		// TODO: Replace with graph.IsEulerianCycle
		// return graph.IsEulerianCycle(edgeSequence)
		ok, pair := graph.IsEulerianPath(edgeSequence)
		if !ok {
			return false
		}
		return pair[0] == pair[1] // check if head == tail
	}
	p.AddUniversalConstraint(test)

	eulerianPath := fn.String_EulerianPath(graph)
	p.SolutionStringFn = eulerianPath
	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		path := eulerianPath(solution)
		if path == fn.InvalidSolution {
			return path
		}
		// Sorted cycle of Eulerian path
		sequence := strings.Fields(path)
		index := slices.Index(sequence, slices.Min(sequence))
		sequence2 := append([]string{}, sequence[index:len(sequence)-1]...) // remove duplicate tail
		sequence2 = append(sequence2, sequence[:index]...)
		return strings.Join(sequence2, " ")
	}

	return p
}
