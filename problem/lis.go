package problem

import (
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Longest Increasing Subsequence problem
func LongestIncreasingSubsequence(n int) *discrete.Problem {
	name := newName(LIS, n)
	sequence := newLIS(name)
	if sequence == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(sequence)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check if selected subsequence has increasing values
		subset := fn.AsSubset(solution)
		slices.Sort(subset) // sort the indexes
		subseq := list.MapList(subset, sequence)
		for i := 1; i < len(subseq); i++ {
			if subseq[i-1] >= subseq[i] {
				return false // invalid if prev not less than curr
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(sequence)

	return p
}

// Load LIS test case
func newLIS(name string) []int {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 1 {
		return nil
	}
	return list.Map(strings.Fields(lines[0]), number.ParseInt)
}
