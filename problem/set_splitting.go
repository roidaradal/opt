package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSetSplitting creates a new Set Cover problem
func NewSetSplitting(variant string, n int) *discrete.Problem {
	name := newName(SetSplitting, variant, n)
	switch variant {
	case "basic":
		return setSplitting(name)
	default:
		return nil
	}
}

// Set Splitting
func setSplitting(name string) *discrete.Problem {
	cfg := data.NewSubsets(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	domain := discrete.RangeDomain(1, 2)
	p.Variables = discrete.Variables(cfg.Universal)
	p.AddVariableDomains(domain)

	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count number of subsets that are split by the partition
		partitions := fn.PartitionStrings(solution, domain, cfg.Universal)
		part1 := ds.SetFrom(partitions[0])
		part2 := ds.SetFrom(partitions[1])
		var count discrete.Score = 0
		for _, s := range cfg.Subsets {
			subset := ds.SetFrom(s)
			diff1 := subset.Difference(part1).Len()
			diff2 := subset.Difference(part2).Len()
			if diff1 > 0 && diff2 > 0 {
				// If subset has at least 1 difference from both partitions, valid split
				count += 1
			}
		}
		return count
	}

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, cfg.Universal)
	p.SolutionStringFn = fn.StringPartition(domain, cfg.Universal)
	return p
}
