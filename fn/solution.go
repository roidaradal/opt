package fn

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// Assumes BooleanDomain {0, 1}, returns list of variables
// that has value = 1 in the solution
func AsSubset(solution *discrete.Solution) []discrete.Variable {
	subset := make([]discrete.Variable, 0, solution.Length())
	for variable, value := range solution.Map {
		if value == 1 {
			subset = append(subset, variable)
		}
	}
	return subset
}

// Assumes BooleanDomain {0, 1}, returns the number of
// selected items (value = 1) in the solution
func SubsetSize(solution *discrete.Solution) int {
	return list.Sum(dict.Values(solution.Map))
}

// Return list of variables sequenced by the solution values
func AsSequence(solution *discrete.Solution) []discrete.Variable {
	sequence := make([]discrete.Variable, solution.Length())
	for variable, idx := range solution.Map {
		sequence[idx] = variable
	}
	return sequence
}
