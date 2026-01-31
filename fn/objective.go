package fn

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/discrete"
)

// ScoreCountUniqueValues counts the number of unique values
func ScoreCountUniqueValues(solution *discrete.Solution) discrete.Score {
	uniqueValues := ds.SetFrom(solution.Values())
	return discrete.Score(uniqueValues.Len())
}
