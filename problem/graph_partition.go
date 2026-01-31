package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewGraphPartition creates a new Graph Partition problem
func NewGraphPartition(variant string, n int) *discrete.Problem {
	name := newName(GraphPartition, variant, n)
	switch variant {
	case "basic":
		return graphPartition(name)
	default:
		return nil
	}
}

// Graph Partition
func graphPartition(name string) *discrete.Problem {
	cfg := data.NewGraphPartition(name)
	if cfg == nil {
		return nil
	}
	graph := cfg.Graph
	if graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	domain := discrete.RangeDomain(1, cfg.NumPartitions)
	p.Variables = discrete.Variables(graph.Vertices)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check all partition sizes are not less that minimum
		partitionSizes := dict.TallyValues(solution.Map, domain)
		return list.All(dict.Values(partitionSizes), func(size int) bool {
			return size >= cfg.MinSize
		})
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find edges that cross partitions (v1 and v2 groups are different)
		// Sum up weight of crossing edges
		var score discrete.Score = 0
		group := solution.Map
		for i, edge := range graph.Edges {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			if group[x1] != group[x2] {
				score += cfg.EdgeWeight[i]
			}
		}
		return score
	}

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.StringPartition(domain, graph.Vertices)
	return p
}
