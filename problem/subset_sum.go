package problem

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSubsetSum creates a new Subset Sum problem
func NewSubsetSum(variant string, n int) *discrete.Problem {
	name := newName(SubsetSum, variant, n)
	switch variant {
	case "basic":
		return subsetSum(name)
	case "max_sum":
		return maxSumMultipleSubsetSum(name)
	case "max_min":
		return maxMinMultipleSubsetSum(name)
	default:
		return nil
	}
}

// Subset Sum
func subsetSum(name string) *discrete.Problem {
	p, cfg := newNumbersSubsetProblem(name)
	if p == nil || cfg == nil || cfg.Target == 0 {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get solution subset sum
		total := list.Sum(list.MapList(fn.AsSubset(solution), cfg.Numbers))
		if p.IsSatisfaction() {
			// Satisfaction: check if subset sum == target sum
			return total == cfg.Target
		}
		// Optimization: check if subset sum does not exceed target
		return total <= cfg.Target
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For optimization version, minimize difference between target and subset sum
		// If it exceeds target, invalid solution
		total := list.Sum(list.MapList(fn.AsSubset(solution), cfg.Numbers))
		if total > cfg.Target {
			return discrete.Inf
		}
		return discrete.Score(cfg.Target - total)
	}
	return p
}

// Common steps for creating multiple subset sum problem
func newMultipleSubsetSumProblem(name string) (*discrete.Problem, *data.Numbers) {
	cfg := data.NewNumbers(name)
	if cfg == nil || cfg.Target == 0 || cfg.NumBins == 0 {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition
	p.Goal = discrete.Maximize

	p.Variables = discrete.Variables(cfg.Numbers)
	p.AddVariableDomains(discrete.RangeDomain(0, cfg.NumBins)) // 0 = not in bin

	validDomain := discrete.RangeDomain(1, cfg.NumBins)
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Total of numbers in each bin must not exceed capacity
		for _, partition := range fn.AsPartition(solution, validDomain) {
			total := list.Sum(list.MapList(partition, cfg.Numbers))
			if total > cfg.Target {
				return false
			}
		}
		return true
	})

	p.SolutionStringFn = fn.StringPartition(validDomain, cfg.Numbers)
	return p, cfg
}

// Max-Sum Multiple Subset Sum
func maxSumMultipleSubsetSum(name string) *discrete.Problem {
	p, cfg := newMultipleSubsetSumProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	validDomain := discrete.RangeDomain(1, cfg.NumBins)
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Sum up total numbers across bins
		total := 0
		for _, partition := range fn.AsPartition(solution, validDomain) {
			total += list.Sum(list.MapList(partition, cfg.Numbers))
		}
		return discrete.Score(total)
	}
	return p
}

// Max-Min Multiple Subset Sum
func maxMinMultipleSubsetSum(name string) *discrete.Problem {
	p, cfg := newMultipleSubsetSumProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	validDomain := discrete.RangeDomain(1, cfg.NumBins)
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find min total numbers across bins
		totals := list.Map(fn.AsPartition(solution, validDomain), func(partition []discrete.Variable) int {
			return list.Sum(list.MapList(partition, cfg.Numbers))
		})
		return discrete.Score(slices.Min(totals))
	}

	return p
}
