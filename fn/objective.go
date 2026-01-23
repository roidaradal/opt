package fn

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// ObjectiveFn: count size of solution as subset
func Score_SubsetSize(solution *discrete.Solution) discrete.Score {
	return discrete.Score(list.Sum(solution.Values()))
}

// ObjectiveFn: sum of weighted values
func Score_SumWeightedValues(variables []discrete.Variable, weight []float64) discrete.ObjectiveFn {
	return func(solution *discrete.Solution) discrete.Score {
		count := solution.Map
		return list.Sum(list.Map(variables, func(x discrete.Variable) discrete.Score {
			return float64(count[x]) * weight[x]
		}))
	}
}
