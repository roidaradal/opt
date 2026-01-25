package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Set creates a new Set problem
func Set(variant string, n int) *discrete.Problem {
	name := newName(SET, variant, n)
	switch variant {
	case "cover":
		return setCover(name)
	case "pack":
		return setPacking(name)
	case "split":
		return setSplitting(name)
	default:
		return nil
	}
}

// Common steps of creating Set problem
func setProblem(name string) (*discrete.Problem, *subsetsCfg) {
	cfg, _ := newSubsets(name, 0)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.names)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(cfg.names)

	return p, cfg
}

// Set Cover problem
func setCover(name string) *discrete.Problem {
	p, cfg := setProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Minimize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each selected subset
		covered := dict.Flags(cfg.universal, false)
		for _, x := range fn.AsSubset(solution) {
			// Each subset item is covered
			for _, item := range cfg.subsets[x] {
				covered[item] = true
			}
		}
		return list.AllTrue(dict.Values(covered))
	})

	return p
}

// Set Packing problem
func setPacking(name string) *discrete.Problem {
	p, cfg := setProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.Goal = discrete.Maximize
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each selected subset
		covered := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			// Increment counter for each item in selected subset
			dict.UpdateCounter(covered, cfg.subsets[x])
		}
		// Make sure all covered items are only covered once (no overlap)
		return list.AllEqual(dict.Values(covered), 1)
	})

	return p
}

// Set Splitting problem
func setSplitting(name string) *discrete.Problem {
	cfg, _ := newSubsets(name, 0)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.universal)
	domain := discrete.RangeDomain(1, 2)
	p.AddVariableDomains(domain)

	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Count number of subsets that are split by the partition
		partitions := fn.PartitionStrings(solution, domain, cfg.universal)
		part1 := ds.SetFrom(partitions[0])
		part2 := ds.SetFrom(partitions[1])
		var count discrete.Score = 0
		for _, s := range cfg.subsets {
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

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, cfg.universal)
	p.SolutionStringFn = fn.StringPartition(domain, cfg.universal)

	return p
}

// Load set test case, return subsetsCfg and extra lines (offeset)
func newSubsets(name string, offset int) (*subsetsCfg, []string) {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) < offset+2 {
		return nil, nil
	}
	numSubsets := len(lines[offset+1:])
	cfg := &subsetsCfg{
		universal: fn.StringList(lines[offset]),
		names:     make([]string, 0, numSubsets),
		subsets:   make([][]string, 0, numSubsets),
	}
	for _, line := range lines[offset+1:] {
		parts := str.CleanSplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		cfg.names = append(cfg.names, parts[0])
		cfg.subsets = append(cfg.subsets, fn.StringList(parts[1]))
	}
	return cfg, lines[:offset]
}
