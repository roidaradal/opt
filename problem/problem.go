// Package problem contains definitions of select discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	INTERVAL    = "interval"
	SUBSEQUENCE = "subsequence"
	SUBSET_SUM  = "subsetsum"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	INTERVAL:    Interval,
	SUBSEQUENCE: Subsequence,
	SUBSET_SUM:  SubsetSum,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
