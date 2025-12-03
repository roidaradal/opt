package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION string = "activity"
	RESOURCE_OPT       string = "resource"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	RESOURCE_OPT:       ResourceOptimization,
}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
