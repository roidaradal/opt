package fn

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// Assumes BooleanDomain {0, 1}, returns list of variables
// that has value = 1 in the solution
func AsSubset(solution *discrete.Solution) []discrete.Variable {
	subset := make([]discrete.Variable, 0, len(solution.Map))
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
