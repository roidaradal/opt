package problem

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Subsequence creates a new Subsequence problem
func Subsequence(variant string, n int) *discrete.Problem {
	name := newName(SUBSEQUENCE, variant, n)
	switch variant {
	case "inc":
		return subsequenceIncreasing(name)
	case "alt":
		return subsequenceAlternating(name)
	default:
		return nil
	}
}

// Common create steps of Longest Subsequence problem
func subsequenceProblem(name string) (*discrete.Problem, []int) {
	sequence := newSequence(name)
	if sequence == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(sequence)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.Goal = discrete.Maximize
	p.ObjectiveFn = fn.ScoreSubsetSize
	p.SolutionStringFn = fn.StringSubset(sequence)
	return p, sequence
}

// Longest Increasing Subsequence problem
func subsequenceIncreasing(name string) *discrete.Problem {
	p, sequence := subsequenceProblem(name)
	if p == nil || sequence == nil {
		return nil
	}

	// Increasing Subsequence constraint
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		subset := fn.AsSubset(solution)
		numSelected := len(subset)
		if numSelected <= 1 {
			return true // no need to check if 0 or 1 item in sequence
		}

		slices.Sort(subset) // sort indexes
		subsequence := list.MapList(subset, sequence)
		for i := range numSelected - 1 {
			if subsequence[i] >= subsequence[i+1] {
				return false // invalid if current not less than next
			}
		}
		return true
	})

	return p
}

// Longest Alternating Subsequence problem
func subsequenceAlternating(name string) *discrete.Problem {
	p, sequence := subsequenceProblem(name)
	if p == nil || sequence == nil {
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
		subsequence := list.MapList(subset, sequence)
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

// Load int sequence test case
func newSequence(name string) []int {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 1 {
		return nil
	}
	return fn.IntList(lines[0])
}
