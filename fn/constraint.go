package fn

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// ConstraintAllUnique makes sure all solution values are unique
func ConstraintAllUnique(solution *discrete.Solution) bool {
	return list.AllUnique(solution.Values())
}
