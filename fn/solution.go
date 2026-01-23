package fn

import (
	"github.com/roidaradal/opt/discrete"
)

// AsSubset assumes BooleanDomain {0,1} and returns list of variables with value=1 in solution
func AsSubset(solution *discrete.Solution) []discrete.Variable {
	subset := make([]discrete.Variable, 0, solution.Length())
	for variable, value := range solution.Map {
		if value == 1 {
			subset = append(subset, variable)
		}
	}
	return subset
}

// AsSequence returns list of variables sequenced by solution values
func AsSequence(solution *discrete.Solution) []discrete.Variable {
	sequence := make([]discrete.Variable, solution.Length())
	for variable, idx := range solution.Map {
		sequence[idx] = variable
	}
	return sequence
}
