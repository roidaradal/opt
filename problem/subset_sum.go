package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Subset Sum problem
func SubsetSum(n int) *discrete.Problem {
	name := newName(SUBSET_SUM, n)
	target, numbers := newSubsetSum(name)
	if target == 0 || numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(numbers)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Get the solution subset sum
		total := list.Sum(list.MapList(fn.AsSubset(solution), numbers))
		if p.IsSatisfaction() {
			//Check if subset sum == target sum
			return total == target
		} else {
			// Check if subset sum does not exceed target
			return total <= target
		}
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// For optimization version, minimize the difference between target and subset sum
		// If it exceeds target, invalid solution
		total := list.Sum(list.MapList(fn.AsSubset(solution), numbers))
		if total > target {
			return discrete.HardPenalty
		}
		return discrete.Score(target - total)
	}
	p.SolutionStringFn = fn.String_Subset(numbers)

	return p
}

// Load subset sum test case
func newSubsetSum(name string) (int, []int) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 2 {
		return 0, nil
	}
	target := number.ParseInt(lines[0])
	numbers := list.Map(strings.Fields(lines[1]), number.ParseInt)
	return target, numbers
}
