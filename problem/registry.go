// Package problem contains definitions of some discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION = "activity"
	ASSIGNMENT         = "assignment"
	BIN_COVER          = "bincover"
	BIN_PACKING        = "binpacking"
	BINARY_PAINTSHOP   = "binarypaint"
	CAR_PAINT          = "carpaint"
	CAR_SEQUENCE       = "carsequence"
	CLIQUE             = "clique"
	DOMINATING_SET     = "dominatingset"
	EDGE_COLOR         = "edgecolor"
	EULER_CYCLE        = "eulercycle"
	EULER_PATH         = "eulerpath"
	EXACT_COVER        = "exactcover"
	FLOWSHOP_SCHED     = "flowshop"
	GEN_ASSIGNMENT     = "genassignment"
	GRAPH_COLOR        = "graphcolor"
	GRAPH_PARTITION    = "graphpartition"
	HAMILTON_PATH      = "hamiltonpath"
	HAMILTON_CYCLE     = "hamiltoncycle"
	INDEPENDENT_SET    = "independentset"
	JOBSHOP_SCHED      = "jobshop"
	K_CENTER           = "kcenter"
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
	QKP                = "qkp"
	QUAD_ASSIGNMENT    = "quadassignment"
	RESOURCE_OPT       = "resource"
	SCENE_ALLOCATION   = "scene"
	SET_COVER          = "setcover"
	SET_PACKING        = "setpacking"
	SET_SPLITTING      = "setsplit"
	SUBSET_SUM         = "subsetsum"
	TSP                = "tsp"
	VERTEX_COVER       = "vertexcover"
	WAREHOUSE          = "warehouse"
	WEAPON             = "weapon"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION: ActivitySelection,
	ASSIGNMENT:         Assignment,
	BIN_COVER:          BinCover,
	BIN_PACKING:        BinPacking,
	BINARY_PAINTSHOP:   BinaryPaintShop,
	CAR_PAINT:          CarPainting,
	CAR_SEQUENCE:       CarSequencing,
	CLIQUE:             Clique,
	DOMINATING_SET:     DominatingSet,
	EDGE_COLOR:         EdgeColoring,
	EULER_CYCLE:        EulerCycle,
	EULER_PATH:         EulerPath,
	EXACT_COVER:        ExactCover,
	FLOWSHOP_SCHED:     FlowShopSchedule,
	GEN_ASSIGNMENT:     GeneralizedAssignment,
	GRAPH_COLOR:        GraphColoring,
	GRAPH_PARTITION:    GraphPartition,
	HAMILTON_PATH:      HamiltonPath,
	HAMILTON_CYCLE:     HamiltonCycle,
	INDEPENDENT_SET:    IndependentSet,
	JOBSHOP_SCHED:      JobShopSchedule,
	K_CENTER:           KCenter,
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
	QKP:                QuadraticKnapsack,
	QUAD_ASSIGNMENT:    QuadraticAssignment,
	RESOURCE_OPT:       ResourceOptimization,
	SCENE_ALLOCATION:   SceneAllocation,
	SET_COVER:          SetCover,
	SET_PACKING:        SetPacking,
	SET_SPLITTING:      SetSplitting,
	SUBSET_SUM:         SubsetSum,
	TSP:                TravelingSalesman,
	VERTEX_COVER:       VertexCover,
	WAREHOUSE:          WarehouseLocation,
	WEAPON:             WeaponTarget,
}

var NoFiles = []string{LANGFORD_PAIR, MAGIC_SERIES, NQUEENS}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
