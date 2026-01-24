// Package problem contains definitions of select discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	INTERVAL    = "interval"
	KNAPSACK    = "knapsack"
	SAT         = "sat"
	SUBSEQUENCE = "subseq"
	SUBSETSUM   = "subsetsum"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	INTERVAL:    Interval,
	KNAPSACK:    Knapsack,
	SAT:         Satisfaction,
	SUBSEQUENCE: Subsequence,
	SUBSETSUM:   SubsetSum,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
