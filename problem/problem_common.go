package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Common steps for creating Bin problems
func newBinProblem(name string) (*discrete.Problem, *data.Bins) {
	cfg := data.NewBins(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.Weight)
	p.AddVariableDomains(cfg.Bins)

	p.ObjectiveFn = fn.ScoreCountUniqueValues
	p.SolutionCoreFn = fn.CoreSortedPartition(cfg.Bins, cfg.Weight)
	p.SolutionStringFn = fn.StringPartition(cfg.Bins, cfg.Weight)
	return p, cfg
}

// Common steps for creating Graph Subset problems (Graph Cover, Independent Set)
func newGraphSubsetProblem(name string, variablesFn data.GraphVariablesFn) (*discrete.Problem, *data.Graph) {
	graph := data.NewUndirectedGraph(name)
	if graph == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	variables := variablesFn(graph)
	p.Variables = discrete.Variables(variables)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(variables)
	return p, graph
}

// Common steps for creating Graph Cover problems
func newGraphCoverProblem(name string, variablesFn data.GraphVariablesFn) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphSubsetProblem(name, variablesFn)
	if p == nil || graph == nil {
		return nil, nil
	}

	p.Goal = discrete.Minimize
	return p, graph
}

// Common steps for creating Independent Set problem
func newIndependentSetProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphSubsetProblem(name, data.GraphVertices)
	if p == nil || graph == nil {
		return nil, nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check that subset of vertices forms an independent set:
		// none of the vertices are connected to each other
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		// IsIndependentSet check
		vertexSet := ds.SetFrom(vertices)
		for _, vertex := range vertices {
			adjacent := ds.SetFrom(graph.Neighbors(vertex))
			if vertexSet.Intersection(adjacent).NotEmpty() {
				return false
			}
		}
		return true
	})

	p.Goal = discrete.Maximize
	return p, graph
}
