package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewBinCover creates a new Bin Cover problem
func NewBinCover(variant string, n int) *discrete.Problem {
	name := newName(BinCover, variant, n)
	switch variant {
	case "basic":
		return binCover(name)
	default:
		return nil
	}
}

// Bin Cover
func binCover(name string) *discrete.Problem {
	p, cfg := newBinPartitionProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Maximize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get sum of item weights in each partition
		sums := fn.PartitionSums(solution, cfg.Bins, cfg.Weight)
		sums = list.Filter(sums, func(sum float64) bool {
			return sum > 0 // remove unused bins
		})
		// Check all partition sums are at least min capacity
		return list.All(sums, func(sum float64) bool {
			return sum >= cfg.Capacity
		})
	})
	return p
}
