package problem

import (
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
