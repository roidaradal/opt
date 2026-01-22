// Package problem contains definitions of select discrete optimization problems
package problem

import (
	"fmt"
)

const (
	SUBSET = "subset"
)

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
