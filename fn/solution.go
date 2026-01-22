package fn

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// Assumes BooleanDomain {0,1}, return list of variables with value=1 in solution
func AsSubset(solution *discrete.Solution) []discrete.Variable {
	subset := make([]discrete.Variable, 0, solution.Length())
	for variable, value := range solution.Map {
		if value == 1 {
			subset = append(subset, variable)
		}
	}
	return subset
}

// Assumes BooleanDomain {0,1}, return number of selected values in solution
func SubsetSize(solution *discrete.Solution) int {
	return list.Sum(solution.Values())
}
