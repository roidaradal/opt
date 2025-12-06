// Package constraint contains commonly used constraint functions
package constraint

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// All unique constraint
func AllUnique(solution *discrete.Solution) bool {
	return list.AllUnique(solution.Values())
}
