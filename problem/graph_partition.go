package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
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
	p, cfg := newGraphPartitionProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check all partition sizes are not less that minimum
		partitionSizes := dict.TallyValues(solution.Map, p.UniformDomain())
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
	return p
}
