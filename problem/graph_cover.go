package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// GraphCover creates a new Graph Cover problem
func GraphCover(variant string, n int) *discrete.Problem {
	name := newName(GRAPHCOVER, variant, n)
	switch variant {
	case "vertex":
		return graphCoverVertex(name)
	case "edge":
		return graphCoverEdge(name)
	case "clique":
		return graphCoverClique(name)
	default:
		return nil
	}
}

// Common steps for creating graph cover problem
func graphCoverProblem(name string, getVariables func(*graphCfg) []string) (*discrete.Problem, *graphCfg) {
	graph := newUndirectedGraph(name)
	if graph == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	variables := getVariables(graph)
	p.Variables = discrete.Variables(variables)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(variables)

	return p, graph
}

// Vertex Cover
func graphCoverVertex(name string) *discrete.Problem {
	p, graph := graphCoverProblem(name, graphVertices)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check for all edges, at least one vertex is covered by the solution subset
		used := solution.Map
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return used[x1]+used[x2] > 0 // at least 1 is covered
		})
	})
	return p
}

// Edge Cover
func graphCoverEdge(name string) *discrete.Problem {
	p, graph := graphCoverProblem(name, graphEdges)
	if p == nil || graph == nil {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check for all vertices, covered by at least one edge endpoint in solution subset
		count := dict.NewCounter(graph.Vertices)
		for _, x := range fn.AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			count[v1] += 1
			count[v2] += 1
		}
		return list.AllGreater(dict.Values(count), 0)
	})
	return p
}

// Clique Cover
func graphCoverClique(name string) *discrete.Problem {
	graph := newUndirectedGraph(name)
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.RangeDomain(1, len(graph.Vertices))
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each partition group of vertices is a clique
		return list.All(fn.AsPartition(solution, domain), func(group []discrete.Variable) bool {
			vertices := list.MapList(group, graph.Vertices)
			return graph.IsClique(vertices)
		})
	})

	// Minimize the number of cliques used
	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScoreCountUniqueValues

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.StringPartition(domain, graph.Vertices)

	return p
}
