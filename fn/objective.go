package fn

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/discrete"
)

// ObjectiveFn: count the size of the solution as subset
func Score_SubsetSize(solution *discrete.Solution) discrete.Score {
	return discrete.Score(SubsetSize(solution))
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

// ObjectiveFn: count unique values
func Score_CountUniqueValues(solution *discrete.Solution) discrete.Score {
	uniqueValues := ds.SetFrom(solution.Values())
	return discrete.Score(uniqueValues.Len())
}

// ObjectiveFn: schedule makespan (total length)
func Score_ScheduleMakespan(tasks []*a.Task) discrete.ObjectiveFn {
	return func(solution *discrete.Solution) discrete.Score {
		makespan := 0
		for x, start := range solution.Map {
			end := start + tasks[x].Duration
			makespan = max(makespan, end)
		}
		return discrete.Score(makespan)
	}
}
