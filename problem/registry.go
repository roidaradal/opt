package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	BIN_PACKING        = "binpacking"
	EDGE_COLOR         = "edgecolor"
	EXACT_COVER        = "exactcover"
	KNAPSACK           = "knapsack"
	LANGFORD_PAIR      = "langford"
	LIS                = "lis"
	MAGIC_SERIES       = "magicseries"
	NQUEENS            = "nqueens"
	NUMBER_PARTITION   = "numberpartition"
	RESOURCE_OPT       = "resource"
	SET_COVER          = "setcover"
	SUBSET_SUM         = "subsetsum"
	WAREHOUSE          = "warehouse"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	BIN_PACKING:        BinPacking,
	EDGE_COLOR:         EdgeColoring,
	EXACT_COVER:        ExactCover,
	KNAPSACK:           Knapsack,
	LANGFORD_PAIR:      LangfordPair,
	LIS:                LongestIncreasingSubsequence,
	MAGIC_SERIES:       MagicSeries,
	NQUEENS:            NQueens,
	NUMBER_PARTITION:   NumberPartition,
	RESOURCE_OPT:       ResourceOptimization,
	SET_COVER:          SetCover,
	SUBSET_SUM:         SubsetSum,
	WAREHOUSE:          WarehouseLocation,
}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
