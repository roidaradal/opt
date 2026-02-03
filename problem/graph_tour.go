package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewGraphTour creates a new Graph Tour problem
func NewGraphTour(variant string, n int) *discrete.Problem {
	name := newName(GraphTour, variant, n)
	switch variant {
	case "euler_path":
		return eulerianPath(name)
	case "euler_cycle":
		return eulerianCycle(name)
	case "hamilton_path":
		return hamiltonianPath(name)
	case "hamilton_cycle":
		return hamiltonianCycle(name)
	default:
		return nil
	}
}

// Common steps for creating Graph Tour problem
func newGraphTourProblem(name string, variablesFn data.GraphVariablesFn) (*discrete.Problem, *data.Graph) {
	graph := data.NewUndirectedGraph(name)
	if graph == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Satisfy

	variables := variablesFn(graph.Graph)
	p.Variables = discrete.Variables(variables)
	p.AddVariableDomains(discrete.IndexDomain(len(variables)))
	return p, graph
}

// Eulerian Path
func eulerianPath(name string) *discrete.Problem {
	p, cfg := newGraphTourProblem(name, data.GraphEdges)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that path of edge sequence from solution forms Eulerian path:
		// visits each edge exactly once
		edgeSequence := list.MapList(fn.AsSequence(solution), graph.Edges)
		isEulerianPath, _ := fn.IsEulerianPath(graph, edgeSequence)
		return isEulerianPath
	})

	toEulerianPath := fn.StringEulerianPath(graph)
	p.SolutionStringFn = toEulerianPath
	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		path := toEulerianPath(solution)
		if path == fn.InvalidSolution {
			return path
		}
		// Mirrored sequence of Eulerian path
		return fn.MirroredSequence(str.SpaceSplit(path))
	}
	return p
}

// Eulerian Cycle
func eulerianCycle(name string) *discrete.Problem {
	p, cfg := newGraphTourProblem(name, data.GraphEdges)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that path of edge sequence from solution forms Eulerian cycle:
		// visits each edge exactly once, and ends at vertex where it started
		edgeSequence := list.MapList(fn.AsSequence(solution), graph.Edges)
		// Check if edges form Eulerian path
		isEulerianPath, pair := fn.IsEulerianPath(graph, edgeSequence)
		head, tail := pair[0], pair[1]
		return isEulerianPath && head == tail
	})

	toEulerianPath := fn.StringEulerianPath(graph)
	p.SolutionStringFn = toEulerianPath
	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		path := toEulerianPath(solution)
		if path == fn.InvalidSolution {
			return path
		}
		// Sorted cycle of Eulerian path, remove duplicated tail
		return fn.SortedCycle(str.SpaceSplit(path), true)
	}
	return p
}

// Hamiltonian Path
func hamiltonianPath(name string) *discrete.Problem {
	p, cfg := newGraphTourProblem(name, data.GraphVertices)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that path of vertex sequence from solution forms Hamiltonian path:
		// visit each vertex exactly once
		vertices := list.MapList(fn.AsSequence(solution), graph.Vertices)
		return fn.IsHamiltonianPath(graph, vertices)
	})

	p.SolutionCoreFn = fn.CoreMirroredSequence(graph.Vertices)
	p.SolutionStringFn = fn.StringSequence(graph.Vertices)
	return p
}

// Hamiltonian Cycle
func hamiltonianCycle(name string) *discrete.Problem {
	p, cfg := newGraphTourProblem(name, data.GraphVertices)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that path of vertex sequence from solution forms Hamiltonian cycle:
		// visit each vertex exactly once, and return to starting point
		vertices := list.MapList(fn.AsSequence(solution), graph.Vertices)
		// Check if vertices form Hamiltonian path
		if !fn.IsHamiltonianPath(graph, vertices) {
			return false
		}
		// Check if there is edge to connect last vertex and first vertex
		first, last := vertices[0], list.Last(vertices, 1)
		return graph.NeighborsOf[last].Has(first)
	})

	p.SolutionCoreFn = fn.CoreSortedCycle(graph.Vertices)
	p.SolutionStringFn = fn.StringSequence(graph.Vertices)
	return p
}
