package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Partition creates a new Partition problem
func Partition(variant string, n int) *discrete.Problem {
	name := newName(PARTITION, variant, n)
	switch variant {
	case "graph":
		return partitionGraph(name)
	case "number":
		return partitionNumber(name)
	default:
		return nil
	}
}

// Graph Partition problem
func partitionGraph(name string) *discrete.Problem {
	graph, cfg := newGraphPartition(name)
	if graph == nil || cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.RangeDomain(1, cfg.numPartitions)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check all partition sizes are not less that minimum
		partitionSizes := dict.TallyValues(solution.Map, domain)
		return list.All(dict.Values(partitionSizes), func(size int) bool {
			return size >= cfg.minPartitionSize
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
				score += cfg.edgeWeight[i]
			}
		}
		return score
	}

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.StringPartition(domain, graph.Vertices)

	return p
}

// Number Partition problem
func partitionNumber(name string) *discrete.Problem {
	numbers := newSequence(name)
	if numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(numbers)
	domain := discrete.RangeDomain(1, 2)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		if p.IsOptimization() {
			return true // don't test if optimization problem
		}
		// Check if the 2 partition sums are the same
		sums := fn.PartitionSums(solution, domain, numbers)
		return list.AllSame(sums)
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Minimize difference between the 2 partition sums
		sums := fn.PartitionSums(solution, domain, numbers)
		return discrete.Score(number.Abs(sums[0] - sums[1]))
	}

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, numbers)
	p.SolutionStringFn = fn.StringPartition(domain, numbers)

	return p
}

// Load graph partition test case
func newGraphPartition(name string) (*ds.Graph, *graphPartitionCfg) {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 5 {
		return nil, nil
	}
	cfg := &graphPartitionCfg{
		numPartitions:    number.ParseInt(lines[0]),
		minPartitionSize: number.ParseInt(lines[1]),
		edgeWeight:       fn.FloatList(lines[4]),
	}
	graph := ds.GraphFrom(lines[2], lines[3])
	return graph, cfg
}
