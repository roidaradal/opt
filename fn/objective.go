package fn

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// ScoreSubsetSize counts the size of solution as subset
func ScoreSubsetSize(solution *discrete.Solution) discrete.Score {
	return discrete.Score(list.Sum(solution.Values()))
}

// ScoreSumWeightedValues sums up the weighted count of each variable
func ScoreSumWeightedValues(variables []discrete.Variable, weight []float64) discrete.ObjectiveFn {
	return func(solution *discrete.Solution) discrete.Score {
		count := solution.Map
		return list.Sum(list.Map(variables, func(x discrete.Variable) discrete.Score {
			return float64(count[x]) * weight[x]
		}))
	}
}
