package problem

import (
	"github.com/roidaradal/fn/lang"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Common steps for creating Bins partition problems
func newBinPartitionProblem(name string) (*discrete.Problem, *data.Bins) {
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

// Common steps for creating Graph Partition problems (Clique Cover, Graph Partition)
func newGraphPartitionProblem(name string) (*discrete.Problem, *data.GraphPartition) {
	cfg := data.NewGraphPartition(name)
	if cfg == nil || cfg.Graph == nil {
		return nil, nil
	}
	graph := cfg.Graph

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	// If numPartitions is not set, defaults to number of vertices
	numPartitions := lang.Ternary(cfg.NumPartitions == 0, len(graph.Vertices), cfg.NumPartitions)
	domain := discrete.RangeDomain(1, numPartitions)
	p.Variables = discrete.Variables(graph.Vertices)
	p.AddVariableDomains(domain)

	p.ObjectiveFn = fn.ScoreCountUniqueValues
	p.SolutionCoreFn = fn.CoreSortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.StringPartition(domain, graph.Vertices)
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

	variables := variablesFn(graph.Graph)
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

// Common steps for creating Graph Coloring problems
func newGraphColoringProblem[T any](name string, variablesFn data.GraphVariablesFn, domainFn data.GraphColorsFn[T]) (*discrete.Problem, *data.GraphColoring) {
	cfg := data.NewGraphColoring(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment
	p.Goal = discrete.Minimize

	domain := domainFn(cfg)
	p.Variables = discrete.Variables(variablesFn(cfg.Graph))
	p.AddVariableDomains(discrete.Domain(domain))

	p.SolutionStringFn = fn.StringValues(p, domain)
	return p, cfg
}

// Common steps for creating Subsets problem
func newSubsetsProblem(name string) (*discrete.Problem, *data.Subsets) {
	cfg := data.NewSubsets(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.Names)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(cfg.Names)
	return p, cfg
}

// Common steps for creating a Numbers subset problem
func newNumbersSubsetProblem(name string) (*discrete.Problem, *data.Numbers) {
	cfg := data.NewNumbers(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.Numbers)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.SolutionStringFn = fn.StringSubset(cfg.Numbers)
	return p, cfg
}

// Common steps for creating Spanning Tree problem
func newSpanningTreeProblem(name string, spanFn data.GraphSpanFn) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphSubsetProblem(name, data.GraphEdges)
	if p == nil || graph == nil {
		return nil, nil
	}

	vertices := spanFn(graph)
	// Constraint: all vertices are spanned
	p.AddUniversalConstraint(fn.ConstraintAllVerticesCovered(graph.Graph, vertices))
	// Constraint: solution forms tree and all vertices are reachable from tree traversal
	p.AddUniversalConstraint(fn.ConstraintSpanningTree(graph.Graph, vertices))

	p.Goal = discrete.Minimize
	return p, graph
}
