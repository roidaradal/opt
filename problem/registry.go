// Package problem contains definitions of some discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	BIN_PACKING        = "binpacking"
	BINARY_PAINTSHOP   = "binarypaint"
	CAR_PAINT          = "carpaint"
	CAR_SEQUENCE       = "carsequence"
	CLIQUE             = "clique"
	EDGE_COLOR         = "edgecolor"
	EXACT_COVER        = "exactcover"
	FLOWSHOP_SCHED     = "flowshop"
	GRAPH_COLOR        = "graphcolor"
	GRAPH_PARTITION    = "graphpartition"
	INDEPENDENT_SET    = "independentset"
	JOBSHOP_SCHED      = "jobshop"
	KNAPSACK           = "knapsack"
	LANGFORD_PAIR      = "langford"
	LIS                = "lis"
	MAGIC_SERIES       = "magicseries"
	MST                = "mst"
	NQUEENS            = "nqueens"
	NUMBER_PARTITION   = "numberpartition"
	OPENSHOP_SCHED     = "openshop"
	RESOURCE_OPT       = "resource"
	SCENE_ALLOCATION   = "scene"
	SET_COVER          = "setcover"
	SUBSET_SUM         = "subsetsum"
	VERTEX_COVER       = "vertexcover"
	WAREHOUSE          = "warehouse"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	BIN_PACKING:        BinPacking,
	BINARY_PAINTSHOP:   BinaryPaintShop,
	CAR_PAINT:          CarPainting,
	CAR_SEQUENCE:       CarSequencing,
	CLIQUE:             Clique,
	EDGE_COLOR:         EdgeColoring,
	EXACT_COVER:        ExactCover,
	FLOWSHOP_SCHED:     FlowShopSchedule,
	GRAPH_COLOR:        GraphColoring,
	GRAPH_PARTITION:    GraphPartition,
	INDEPENDENT_SET:    IndependentSet,
	JOBSHOP_SCHED:      JobShopSchedule,
	KNAPSACK:           Knapsack,
	LANGFORD_PAIR:      LangfordPair,
	LIS:                LongestIncreasingSubsequence,
	MAGIC_SERIES:       MagicSeries,
	MST:                MinimumSpanningTree,
	NQUEENS:            NQueens,
	NUMBER_PARTITION:   NumberPartition,
	OPENSHOP_SCHED:     OpenShopSchedule,
	RESOURCE_OPT:       ResourceOptimization,
	SCENE_ALLOCATION:   SceneAllocation,
	SET_COVER:          SetCover,
	SUBSET_SUM:         SubsetSum,
	VERTEX_COVER:       VertexCover,
	WAREHOUSE:          WarehouseLocation,
}

var NoFiles = []string{LANGFORD_PAIR, MAGIC_SERIES, NQUEENS}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
