package fn

import (
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
