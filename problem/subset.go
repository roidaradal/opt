package problem

import (
	"cmp"
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Subset problem
func Subset(variant string, n int) *discrete.Problem {
	name := newName(SUBSET, variant, n)
	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	switch variant {
	case "activity":
		// Activity Selection
		cfg := newSubsetActivity(name)
		if cfg == nil {
			return nil
		}
		p.Variables = discrete.Variables(cfg.activities)
		p.AddVariableDomains(discrete.BooleanDomain())
		p.AddUniversalConstraint(activitySelectionTest(cfg))

		p.Goal = discrete.Maximize
		p.ObjectiveFn = fn.Score_SubsetSize
		p.SolutionStringFn = fn.String_Subset(cfg.activities)
	case "lis":
		// Longest Increasing Subsequence
		sequence := newSubsetLIS(name)
		if sequence == nil {
			return nil
		}
		p.Variables = discrete.Variables(sequence)
		p.AddVariableDomains(discrete.BooleanDomain())
		p.AddUniversalConstraint(lisTest(sequence))

		p.Goal = discrete.Maximize
		p.ObjectiveFn = fn.Score_SubsetSize
		p.SolutionStringFn = fn.String_Subset(sequence)
	case "sum":
		// Subset Sum
		target, numbers := newSubsetSum(name)
		if target == 0 || numbers == nil {
			return nil
		}
		p.Variables = discrete.Variables(numbers)
		p.AddVariableDomains(discrete.BooleanDomain())
		p.AddUniversalConstraint(subsetSumTest(p, target, numbers))

		p.Goal = discrete.Minimize
		p.ObjectiveFn = subsetSumScore(target, numbers)
		p.SolutionStringFn = fn.String_Subset(numbers)
	default:
		return nil
	}

	return p
}

type subsetActivityCfg struct {
	activities []string
	start      []float64
	end        []float64
}

// Load activity selection test case
func newSubsetActivity(name string) *subsetActivityCfg {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &subsetActivityCfg{
		activities: fn.StringList(lines[0]),
		start:      fn.FloatList(lines[1]),
		end:        fn.FloatList(lines[2]),
	}
}

// Activity Selection constraint
func activitySelectionTest(cfg *subsetActivityCfg) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Get selected activities
		selected := fn.AsSubset(solution)
		numSelected := len(selected)
		if numSelected <= 1 {
			// no conflict for 0 or 1 activities
			return true
		}
		// Sort selected activities by start time
		start, end := cfg.start, cfg.end
		slices.SortFunc(selected, func(x1, x2 int) int {
			return cmp.Compare(start[x1], start[x2])
		})
		// Check for activity overlaps
		for i := range numSelected - 1 {
			curr, next := selected[i], selected[i+1]
			// Conflict: curr activity ends after start of next activity
			if end[curr] > start[next] {
				return false
			}
		}
		return true
	}
}

// Load LIS test case
func newSubsetLIS(name string) []int {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 1 {
		return nil
	}
	return fn.IntList(lines[0])
}

// Longest Increasing Subsequence constraint
func lisTest(sequence []int) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Check if selected subsequence has increasing values
		subset := fn.AsSubset(solution)
		numSelected := len(subset)
		if numSelected <= 1 {
			return true // no need to check if 0 or 1 item in sequence
		}

		slices.Sort(subset) // sort the indexes
		subseq := list.MapList(subset, sequence)
		for i := range numSelected - 1 {
			if subseq[i] >= subseq[i+1] {
				return false // invalid if current not less than next
			}
		}
		return true
	}
}

// Load subset sum test case
func newSubsetSum(name string) (int, []int) {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 2 {
		return 0, nil
	}
	return number.ParseInt(lines[0]), fn.IntList(lines[1])
}

// Subset Sum constraint
func subsetSumTest(p *discrete.Problem, target int, numbers []int) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
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
}

// Subset Sum objective
func subsetSumScore(target int, numbers []int) discrete.ObjectiveFn {
	return func(solution *discrete.Solution) discrete.Score {
		// For optimization version, minimize the difference between target and subset sum
		// If it exceeds target, invalid solution
		total := list.Sum(list.MapList(fn.AsSubset(solution), numbers))
		if total > target {
			return discrete.Inf
		}
		return discrete.Score(target - total)
	}
}
