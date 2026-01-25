// Package problem contains definitions of select discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	BIN         = "bin"
	COVER       = "cover"
	INTERVAL    = "interval"
	KNAPSACK    = "knapsack"
	SAT         = "sat"
	SET         = "set"
	SUBSEQUENCE = "subseq"
	SUBSETSUM   = "subsetsum"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	BIN:         Bin,
	COVER:       Cover,
	INTERVAL:    Interval,
	KNAPSACK:    Knapsack,
	SAT:         Satisfaction,
	SET:         Set,
	SUBSEQUENCE: Subsequence,
	SUBSETSUM:   SubsetSum,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
