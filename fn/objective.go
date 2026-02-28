package fn

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
)

// ScoreCountUniqueValues counts the number of unique values
func ScoreCountUniqueValues(solution *discrete.Solution) discrete.Score {
	uniqueValues := ds.SetFrom(solution.Values())
	return discrete.Score(uniqueValues.Len())
}

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

// ScoreSumWeightedSubset sums up the weight of the solution subset
func ScoreSumWeightedSubset(keys []string, weight map[string]float64) discrete.ObjectiveFn {
	return func(solution *discrete.Solution) discrete.Score {
		subsetKeys := list.MapList(AsSubset(solution), keys)
		return list.Sum(list.Translate(subsetKeys, weight))
	}
}

// ScorePathCost sums up the path cost
func ScorePathCost(cfg *data.GraphPath) discrete.ObjectiveFn {
	return func(solution *discrete.Solution) discrete.Score {
		return list.Sum(PathDistances(solution, cfg))
	}
}

// CountColorChanges counts the number of color changes in the sequence
func CountColorChanges[T comparable](colorSequence []T) int {
	var prevColor T
	changes := 0
	for i, currColor := range colorSequence {
		if i > 0 && prevColor != currColor {
			changes += 1
		}
		prevColor = currColor
	}
	return changes
}
