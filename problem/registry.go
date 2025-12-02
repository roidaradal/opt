package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION string = "activity"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
