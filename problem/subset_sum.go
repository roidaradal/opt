package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// SubsetSum creates a new Subset Sum problem
func SubsetSum(variant string, n int) *discrete.Problem {
	name := newName(SUBSETSUM, variant, n)
	switch variant {
	case "basic":
		return subsetSumBasic(name)
	default:
		return nil
	}
}

// Basic Subset Sum problem
func subsetSumBasic(name string) *discrete.Problem {
	target, numbers := newBasicSubsetSum(name)
	if target == 0 || numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(numbers)
	p.AddVariableDomains(discrete.BooleanDomain())
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get solution subset sum
		total := list.Sum(list.MapList(fn.AsSubset(solution), numbers))
		if p.IsSatisfaction() {
			// Satisfaction: check if subset sum == target sum
			return total == target
		}
		// Optimization: check if subset sum does not exceed target
		return total <= target
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For optimization version, minimize difference between target and subset sum
		// If it exceeds target, invalid solution
		total := list.Sum(list.MapList(fn.AsSubset(solution), numbers))
		if total > target {
			return discrete.Inf
		}
		return discrete.Score(target - total)
	}

	p.SolutionStringFn = fn.StringSubset(numbers)
	return p
}

// Load 'subsetsum.basic' test case
func newBasicSubsetSum(name string) (int, []int) {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) < 2 {
		return 0, nil
	}
	return number.ParseInt(lines[0][0]), fn.IntList(lines[1][0])
}
