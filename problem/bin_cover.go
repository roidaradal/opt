package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Bin Cover problem
func BinCover(n int) *discrete.Problem {
	name := newName(BIN_COVER, n)
	cfg := fn.NewBinProblem(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.Weight)
	domain := discrete.RangeDomain(1, cfg.NumBins)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Get sum of item weights in each partition
		sums := fn.PartitionSums(solution, domain, cfg.Weight)
		sums = list.Filter(sums, func(sum float64) bool {
			return sum > 0 // remove unused bins
		})
		// Check all partition sums are at least min capacity
		return list.All(sums, func(sum float64) bool {
			return sum >= cfg.Capacity
		})
	}
	p.AddUniversalConstraint(test)

	// Maximize number of covered bins used
	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_SortedPartition(domain, cfg.Weight)
	p.SolutionStringFn = fn.String_Partitions(domain, cfg.Weight)

	return p
}
