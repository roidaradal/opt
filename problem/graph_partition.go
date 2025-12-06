package problem

import (
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Graph Partition problem
func GraphPartition(n int) *discrete.Problem {
	name := newName(GRAPH_PARTITION, n)
	cfg, graph := newGraphPartition(name)
	if cfg == nil || graph == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.RangeDomain(1, cfg.numPartitions)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check that all partition sizes are not less than minimum
		partitionSizes := dict.TallyValues(solution.Map, domain)
		return list.All(dict.Values(partitionSizes), func(size int) bool {
			return size >= cfg.minPartitionSize
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find edges that cross partitions, i.e. the partition of v1 and v2 are different
		// Sum up the weights of the crossing edges
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

	p.SolutionCoreFn = fn.Core_SortedPartition(domain, graph.Vertices)
	p.SolutionStringFn = fn.String_Partitions(domain, graph.Vertices)

	return p
}

type graphPartitionCfg struct {
	numPartitions    int
	minPartitionSize int
	edgeWeight       []float64
}

// Load graph partition test case
func newGraphPartition(name string) (*graphPartitionCfg, *ds.Graph) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 5 {
		return nil, nil
	}
	cfg := &graphPartitionCfg{
		numPartitions:    number.ParseInt(lines[0]),
		minPartitionSize: number.ParseInt(lines[1]),
		edgeWeight:       list.Map(strings.Fields(lines[4]), number.ParseFloat),
	}
	graph := ds.GraphFrom(lines[2], lines[3])
	return cfg, graph
}
