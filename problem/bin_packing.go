package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Bin Packing problem
func BinPacking(n int) *discrete.Problem {
	name := newName(BIN_PACKING, n)
	cfg := fn.NewBinProblem(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.Weight)
	domain := discrete.RangeDomain(1, cfg.NumBins)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Get the sum of item weights in each partition
		sums := fn.PartitionSums(solution, domain, cfg.Weight)
		// Check that all partition sums do not execeed capacity
		return list.All(sums, func(sum float64) bool {
			return sum <= cfg.Capacity
		})
	}
	p.AddUniversalConstraint(test)

	// Minimize the number of bins used
	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_SortedPartition(domain, cfg.Weight)
	p.SolutionStringFn = fn.String_Partitions(domain, cfg.Weight)

	return p
}
