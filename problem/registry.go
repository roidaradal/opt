package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	BIN_PACKING        = "binpacking"
	KNAPSACK           = "knapsack"
	LANGFORD_PAIR      = "langford"
	LIS                = "lis"
	MAGIC_SERIES       = "magicseries"
	NQUEENS            = "nqueens"
	RESOURCE_OPT       = "resource"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	BIN_PACKING:        BinPacking,
	KNAPSACK:           Knapsack,
	LANGFORD_PAIR:      LangfordPair,
	LIS:                LongestIncreasingSubsequence,
	MAGIC_SERIES:       MagicSeries,
	NQUEENS:            NQueens,
	RESOURCE_OPT:       ResourceOptimization,
}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
