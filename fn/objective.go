package fn

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
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
