package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	RESOURCE_OPT       = "resource"
	MAGIC_SERIES       = "magicseries"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	MAGIC_SERIES:       MagicSeries,
	RESOURCE_OPT:       ResourceOptimization,
}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
