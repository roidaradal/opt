package problem

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSubsequence creates a new Subsequence problem
func NewSubsequence(variant string, n int) *discrete.Problem {
	name := newName(Subsequence, variant, n)
	switch variant {
	case "increasing":
		return longestIncreasingSubsequence(name)
	case "alternating":
		return longestAlternatingSubsequence(name)
	case "decreasing":
		return longestDecreasingSubsequence(name)
	case "max_sum_increasing":
		return maxSumIncreasingSubsequence(name)
	case "max_weight_increasing":
		return maxWeightIncreasingSubsequence(name)
	default:
		return nil
	}
}

// Common steps for creating Longest Subsequence problem
func newLongestSubsequenceProblem(name string) (*discrete.Problem, *data.Numbers) {
	p, cfg := newNumbersSubsetProblem(name)
	if p == nil || cfg == nil {
		return nil, nil
	}
	p.Goal = discrete.Maximize
	p.ObjectiveFn = fn.ScoreSubsetSize
	return p, cfg
}

// Longest Increasing Subsequence
func longestIncreasingSubsequence(name string) *discrete.Problem {
	p, cfg := newLongestSubsequenceProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	// Increasing Subsequence constraint
	p.AddUniversalConstraint(fn.ConstraintIncreasingSubsequence(cfg))
	return p
}

// Longest Alternating Subsequence
func longestAlternatingSubsequence(name string) *discrete.Problem {
	p, cfg := newLongestSubsequenceProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	// Alternating Subsequence constraint (down-up-down-up-....)
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		subset := fn.AsSubset(solution)
		numSelected := len(subset)
		if numSelected <= 1 {
			return true // no need to check if 0 or 1 item in sequence
		}

		slices.Sort(subset) // sort indexes
		subsequence := list.MapList(subset, cfg.Numbers)
		down := true
		for i := range numSelected - 1 {
			if down && subsequence[i] <= subsequence[i+1] {
				// invalid if going down, but current not greater than next
				return false
			} else if !down && subsequence[i] >= subsequence[i+1] {
				// invalid if going up, but current not less than next
				return false
			}
			down = !down // toggle to alternate
		}
		return true
	})
	return p
}

// Longest Decreasing Subsequence
func longestDecreasingSubsequence(name string) *discrete.Problem {
	p, cfg := newLongestSubsequenceProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	// Decreasing Subsequence constraint
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		subset := fn.AsSubset(solution)
		numSelected := len(subset)
		if numSelected <= 1 {
			return true // no need to check if 0 or 1 item in sequence
		}
		slices.Sort(subset) // sort indexes
		subsequence := list.MapList(subset, cfg.Numbers)
		for i := range numSelected - 1 {
			if subsequence[i] <= subsequence[i+1] {
				return false // invalid if current not greater than next
			}
		}
		return true
	})
	return p
}

// Max Sum Increasing Subsequence
func maxSumIncreasingSubsequence(name string) *discrete.Problem {
	p, cfg := newLongestSubsequenceProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.AddUniversalConstraint(fn.ConstraintIncreasingSubsequence(cfg))
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		subsequence := list.MapList(fn.AsSubset(solution), cfg.Numbers)
		return discrete.Score(list.Sum(subsequence))
	}
	return p
}

// Max Weight Increasing Subsequence
func maxWeightIncreasingSubsequence(name string) *discrete.Problem {
	p, cfg := newLongestSubsequenceProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	if len(cfg.Weight) != len(cfg.Numbers) {
		return nil
	}
	p.AddUniversalConstraint(fn.ConstraintIncreasingSubsequence(cfg))
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		subsequenceWeights := list.MapList(fn.AsSubset(solution), cfg.Weight)
		return list.Sum(subsequenceWeights)
	}
	return p
}
