// Package problem contains definitions of some discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	ASSIGNMENT         = "assignment"
	BIN_PACKING        = "binpacking"
	BINARY_PAINTSHOP   = "binarypaint"
	CAR_PAINT          = "carpaint"
	CAR_SEQUENCE       = "carsequence"
	CLIQUE             = "clique"
	EDGE_COLOR         = "edgecolor"
	EXACT_COVER        = "exactcover"
	FLOWSHOP_SCHED     = "flowshop"
	GEN_ASSIGNMENT     = "genassignment"
	GRAPH_COLOR        = "graphcolor"
	GRAPH_PARTITION    = "graphpartition"
	INDEPENDENT_SET    = "independentset"
	JOBSHOP_SCHED      = "jobshop"
	KNAPSACK           = "knapsack"
	LANGFORD_PAIR      = "langford"
	LBAP               = "lbap"
	LIS                = "lis"
	MAGIC_SERIES       = "magicseries"
	MST                = "mst"
	NQUEENS            = "nqueens"
	NUMBER_PARTITION   = "numberpartition"
	OPENSHOP_SCHED     = "openshop"
	QBAP               = "qbap"
	QUAD_ASSIGNMENT    = "quadassignment"
	RESOURCE_OPT       = "resource"
	SCENE_ALLOCATION   = "scene"
	SET_COVER          = "setcover"
	SUBSET_SUM         = "subsetsum"
	TSP                = "tsp"
	VERTEX_COVER       = "vertexcover"
	WAREHOUSE          = "warehouse"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	ASSIGNMENT:         Assignment,
	BIN_PACKING:        BinPacking,
	BINARY_PAINTSHOP:   BinaryPaintShop,
	CAR_PAINT:          CarPainting,
	CAR_SEQUENCE:       CarSequencing,
	CLIQUE:             Clique,
	EDGE_COLOR:         EdgeColoring,
	EXACT_COVER:        ExactCover,
	FLOWSHOP_SCHED:     FlowShopSchedule,
	GEN_ASSIGNMENT:     GeneralizedAssignment,
	GRAPH_COLOR:        GraphColoring,
	GRAPH_PARTITION:    GraphPartition,
	INDEPENDENT_SET:    IndependentSet,
	JOBSHOP_SCHED:      JobShopSchedule,
	KNAPSACK:           Knapsack,
	LANGFORD_PAIR:      LangfordPair,
	LBAP:               LinearBottleneckAssignment,
	LIS:                LongestIncreasingSubsequence,
	MAGIC_SERIES:       MagicSeries,
	MST:                MinimumSpanningTree,
	NQUEENS:            NQueens,
	NUMBER_PARTITION:   NumberPartition,
	OPENSHOP_SCHED:     OpenShopSchedule,
	QBAP:               QuadraticBottleneckAssignment,
	QUAD_ASSIGNMENT:    QuadraticAssignment,
	RESOURCE_OPT:       ResourceOptimization,
	SCENE_ALLOCATION:   SceneAllocation,
	SET_COVER:          SetCover,
	SUBSET_SUM:         SubsetSum,
	TSP:                TravelingSalesman,
	VERTEX_COVER:       VertexCover,
	WAREHOUSE:          WarehouseLocation,
}

var NoFiles = []string{LANGFORD_PAIR, MAGIC_SERIES, NQUEENS}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
