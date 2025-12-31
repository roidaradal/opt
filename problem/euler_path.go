package problem

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Eulerian Path problem
func EulerPath(n int) *discrete.Problem {
	name := newName(EULER_PATH, n)
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
		// forms an Eulerian path: visits each edge exactly once
		edgeSequence := list.MapList(fn.AsSequence(solution), graph.Edges)
		return graph.IsEulerianPath(edgeSequence)
	}
	p.AddUniversalConstraint(test)

	eulerianPath := func(solution *discrete.Solution) string {
		path := make([]string, 0, numEdges)
		edgeSequence := list.MapList(fn.AsSequence(solution), graph.Edges)
		a1, b1 := edgeSequence[0].Tuple()
		a2, b2 := edgeSequence[1].Tuple()
		var tail ds.Vertex
		if a1 == a2 {
			path = append(path, b1, a1)
			tail = b2
		} else if b1 == a2 {
			path = append(path, a1, b1)
			tail = b2
		} else if a1 == b2 {
			path = append(path, b1, a1)
			tail = a2
		} else if b1 == b2 {
			path = append(path, a1, b1)
			tail = a2
		} else {
			return "Invalid solution"
		}
		for _, edge := range edgeSequence[2:] {
			a, b := edge.Tuple()
			path = append(path, tail)
			switch tail {
			case a:
				tail = b
			case b:
				tail = a
			default:
				return "Invalid solution"
			}
		}
		path = append(path, tail)
		return strings.Join(path, " ")
	}

	p.SolutionStringFn = eulerianPath
	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		path := eulerianPath(solution)
		if path == "Invalid solution" {
			return path
		}
		// Mirrored sequence of Eulerian path
		sequence := strings.Fields(path)
		first, last := sequence[0], list.Last(sequence, 1)
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(sequence)
		}
		return strings.Join(sequence, " ")
	}

	return p
}
