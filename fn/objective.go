package fn

import "github.com/roidaradal/opt/discrete"

// ObjectiveFn: count size of solution as subset
func Score_SubsetSize(solution *discrete.Solution) discrete.Score {
	return discrete.Score(SubsetSize(solution))
}
