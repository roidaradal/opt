package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Max Set Splitting problem
func SetSplitting(n int) *discrete.Problem {
	name := newName(SET_SPLITTING, n)
	cfg := fn.NewSubsets(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.Universal)
	domain := discrete.RangeDomain(1, 2)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count the number of subsets that are split by the partition
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
	p.SolutionCoreFn = fn.Core_SortedPartition(domain, cfg.Universal)
	p.SolutionStringFn = fn.String_Partitions(domain, cfg.Universal)

	return p
}
