package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
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
	case "weighted":
		return weightedSetSplitting(name)
	default:
		return nil
	}
}

// Create new Set Splitting problem
func newSetSplittingProblem(name string) (*discrete.Problem, *data.Subsets) {
	cfg := data.NewSubsets(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition
	p.Goal = discrete.Maximize

	domain := discrete.RangeDomain(1, 2)
	p.Variables = discrete.Variables(cfg.Universal)
	p.AddVariableDomains(domain)

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, cfg.Universal)
	p.SolutionStringFn = fn.StringPartition(domain, cfg.Universal)
	return p, cfg
}

// Set Splitting
func setSplitting(name string) *discrete.Problem {
	p, cfg := newSetSplittingProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count number of subsets that are split by the partition
		count := len(splitSubsets(solution, cfg, p.UniformDomain()))
		return discrete.Score(count)
	}

	return p
}

// Weighted Set Splitting
func weightedSetSplitting(name string) *discrete.Problem {
	p, cfg := newSetSplittingProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	if len(cfg.Weight) != len(cfg.Names) {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Sum up weight of subsets that are split by the partition
		subsets := splitSubsets(solution, cfg, p.UniformDomain())
		return list.Sum(list.Translate(subsets, cfg.Weight))
	}

	return p
}

// Get subsets that are split by the partition
func splitSubsets(solution *discrete.Solution, cfg *data.Subsets, domain []discrete.Value) []string {
	partitions := fn.PartitionStrings(solution, domain, cfg.Universal)
	part1 := ds.SetFrom(partitions[0])
	part2 := ds.SetFrom(partitions[1])
	splitSubsets := make([]string, 0)
	for i, s := range cfg.Subsets {
		subset := ds.SetFrom(s)
		diff1 := subset.Difference(part1).Len()
		diff2 := subset.Difference(part2).Len()
		if diff1 > 0 && diff2 > 0 {
			// If subset has at least 1 difference from both partitions, valid split
			splitSubsets = append(splitSubsets, cfg.Names[i])
		}
	}
	return splitSubsets
}
