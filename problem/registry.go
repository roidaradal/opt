package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	KNAPSACK           = "knapsack"
	LIS                = "lis"
	MAGIC_SERIES       = "magicseries"
	RESOURCE_OPT       = "resource"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	KNAPSACK:           Knapsack,
	LIS:                LongestIncreasingSubsequence,
	MAGIC_SERIES:       MagicSeries,
	RESOURCE_OPT:       ResourceOptimization,
}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
