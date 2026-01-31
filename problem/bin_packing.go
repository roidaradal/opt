package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewBinPacking creates a new Bin Packing problem
func NewBinPacking(variant string, n int) *discrete.Problem {
	name := newName(BinPacking, variant, n)
	switch variant {
	case "basic":
		return binPacking(name)
	default:
		return nil
	}
}

// Bin Packing
func binPacking(name string) *discrete.Problem {
	p, cfg := newBinProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Minimize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get sum of item weights in each partition
		sums := fn.PartitionSums(solution, cfg.Bins, cfg.Weight)
		// Check all partition sums do not exceed capacity
		return list.All(sums, func(sum float64) bool {
			return sum <= cfg.Capacity
		})
	})
	return p
}
