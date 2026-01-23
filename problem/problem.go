// Package problem contains definitions of select discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	INTERVAL    = "interval"
	SAT         = "sat"
	SUBSEQUENCE = "subsequence"
	SUBSETSUM   = "subsetsum"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	INTERVAL:    Interval,
	SAT:         Satisfaction,
	SUBSEQUENCE: Subsequence,
	SUBSETSUM:   SubsetSum,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
