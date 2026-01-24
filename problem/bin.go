package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Bin creates a new Bin problem
func Bin(variant string, n int) *discrete.Problem {
	name := newName(BIN, variant, n)
	switch variant {
	case "cover":
		return binCover(name)
	case "pack":
		return binPacking(name)
	default:
		return nil
	}
}

// Common steps of creating Bin problem
func binProblem(name string) (*discrete.Problem, *binCfg, []discrete.Value) {
	cfg := newBin(name)
	if cfg == nil {
		return nil, nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.weight)
	domain := discrete.RangeDomain(1, cfg.numBins)
	p.AddVariableDomains(domain)

	p.ObjectiveFn = fn.ScoreCountUniqueValues
	p.SolutionCoreFn = fn.CoreSortedPartition(domain, cfg.weight)
	p.SolutionStringFn = fn.StringPartition(domain, cfg.weight)

	return p, cfg, domain
}

// Bin Cover problem
func binCover(name string) *discrete.Problem {
	p, cfg, domain := binProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Maximize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get sum of item weights in each partition
		sums := fn.PartitionSums(solution, domain, cfg.weight)
		sums = list.Filter(sums, func(sum float64) bool {
			return sum > 0 // remove unused bins
		})
		// Check all partition sums are at least min capacity
		return list.All(sums, func(sum float64) bool {
			return sum >= cfg.capacity
		})
	})

	return p
}

// Bin Packing problem
func binPacking(name string) *discrete.Problem {
	p, cfg, domain := binProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Minimize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get sum of item weights in each partition
		sums := fn.PartitionSums(solution, domain, cfg.weight)
		// Check all partition sums do not exceed capacity
		return list.All(sums, func(sum float64) bool {
			return sum <= cfg.capacity
		})
	})

	return p
}

// Load bin test case
func newBin(name string) *binCfg {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &binCfg{
		numBins:  number.ParseInt(lines[0]),
		capacity: number.ParseFloat(lines[1]),
		weight:   fn.FloatList(lines[2]),
	}
}
