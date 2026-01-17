// Package problem contains definitions of some discrete optimization problems
package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	ACTIVITY_SELECTION       = "activity"
	ASSIGNMENT               = "assignment"
	BIN_COVER                = "bincover"
	BIN_PACKING              = "binpacking"
	BINARY_PAINTSHOP         = "binarypaint"
	BOTTLENECK_TSP           = "btsp"
	CAR_PAINT                = "carpaint"
	CAR_SEQUENCE             = "carsequence"
	CLIQUE                   = "clique"
	CLIQUE_COVER             = "cliquecover"
	COMPLETE_COLOR           = "completecolor"
	DOMINATING_SET           = "dominatingset"
	EDGE_COLOR               = "edgecolor"
	EDGE_COVER               = "edgecover"
	EDGE_DOMINATING_SET      = "edgedominatingset"
	EFFICIENT_DOMINATING_SET = "efficientdominatingset"
	EULER_CYCLE              = "eulercycle"
	EULER_PATH               = "eulerpath"
	EXACT_COVER              = "exactcover"
	FLOWSHOP_SCHED           = "flowshop"
	GEN_ASSIGNMENT           = "genassignment"
	GRAPH_COLOR              = "graphcolor"
	GRAPH_PARTITION          = "graphpartition"
	HAMILTON_PATH            = "hamiltonpath"
	HAMILTON_CYCLE           = "hamiltoncycle"
	HARMONIOUS_COLOR         = "harmoniouscolor"
	INDEPENDENT_SET          = "independentset"
	JOBSHOP_SCHED            = "jobshop"
	K_CENTER                 = "kcenter"
	K_MST                    = "kmst"
	KNAPSACK                 = "knapsack"
	LANGFORD_PAIR            = "langford"
	LBAP                     = "lbap"
	LIS                      = "lis"
	LONGEST_PATH             = "longestpath"
	MAGIC_SERIES             = "magicseries"
	MAX_CARDINALITY_MATCH    = "maxcardinalitymatch"
	MAX_INDUCED_PATH         = "maxinducedpath"
	MAX_WEIGHT_MATCH         = "maxweightmatch"
	MINIMAX_PATH             = "minimaxpath"
	MIN_K_CUT                = "minkcut"
	MDST                     = "mdst"
	MST                      = "mst"
	NQUEENS                  = "nqueens"
	NUMBER_PARTITION         = "numberpartition"
	NURSE_SCHED              = "nurse"
	OPENSHOP_SCHED           = "openshop"
	QBAP                     = "qbap"
	QKP                      = "qkp"
	QUAD_ASSIGNMENT          = "quadassignment"
	RESOURCE_OPT             = "resource"
	SCENE_ALLOCATION         = "scene"
	SET_COVER                = "setcover"
	SET_PACKING              = "setpacking"
	SET_SPLITTING            = "setsplit"
	SHORTEST_PATH            = "shortestpath"
	STEINER_TREE             = "steiner"
	SUBSET_SUM               = "subsetsum"
	SUM_COLOR                = "sumcolor"
	TOPOLOGICAL              = "topological"
	TPP                      = "tpp"
	TSP                      = "tsp"
	VERTEX_COVER             = "vertexcover"
	WAREHOUSE                = "warehouse"
	WEAPON                   = "weapon"
	WIDEST_PATH              = "widestpath"
)

var Creator = map[string]func(int) *discrete.Problem{
	ACTIVITY_SELECTION:       ActivitySelection,
	ASSIGNMENT:               Assignment,
	BIN_COVER:                BinCover,
	BIN_PACKING:              BinPacking,
	BINARY_PAINTSHOP:         BinaryPaintShop,
	BOTTLENECK_TSP:           BottleneckTravelingSalesman,
	CAR_PAINT:                CarPainting,
	CAR_SEQUENCE:             CarSequencing,
	CLIQUE:                   Clique,
	CLIQUE_COVER:             CliqueCover,
	COMPLETE_COLOR:           CompleteColoring,
	DOMINATING_SET:           DominatingSet,
	EDGE_COLOR:               EdgeColoring,
	EDGE_COVER:               EdgeCover,
	EDGE_DOMINATING_SET:      EdgeDominatingSet,
	EFFICIENT_DOMINATING_SET: EfficientDominatingSet,
	EULER_CYCLE:              EulerCycle,
	EULER_PATH:               EulerPath,
	EXACT_COVER:              ExactCover,
	FLOWSHOP_SCHED:           FlowShopSchedule,
	GEN_ASSIGNMENT:           GeneralizedAssignment,
	GRAPH_COLOR:              GraphColoring,
	GRAPH_PARTITION:          GraphPartition,
	HAMILTON_PATH:            HamiltonPath,
	HAMILTON_CYCLE:           HamiltonCycle,
	HARMONIOUS_COLOR:         HarmoniousColoring,
	INDEPENDENT_SET:          IndependentSet,
	JOBSHOP_SCHED:            JobShopSchedule,
	K_CENTER:                 KCenter,
	K_MST:                    KMinimumSpanningTree,
	KNAPSACK:                 Knapsack,
	LANGFORD_PAIR:            LangfordPair,
	LBAP:                     LinearBottleneckAssignment,
	LIS:                      LongestIncreasingSubsequence,
	LONGEST_PATH:             LongestPath,
	MAGIC_SERIES:             MagicSeries,
	MAX_CARDINALITY_MATCH:    MaxCardinalityMatching,
	MAX_INDUCED_PATH:         MaxInducedPath,
	MAX_WEIGHT_MATCH:         MaxWeightMatching,
	MINIMAX_PATH:             MinimaxPath,
	MIN_K_CUT:                MinimumKCut,
	MDST:                     MinDegreeSpanningTree,
	MST:                      MinimumSpanningTree,
	NQUEENS:                  NQueens,
	NUMBER_PARTITION:         NumberPartition,
	NURSE_SCHED:              NurseSchedule,
	OPENSHOP_SCHED:           OpenShopSchedule,
	QBAP:                     QuadraticBottleneckAssignment,
	QKP:                      QuadraticKnapsack,
	QUAD_ASSIGNMENT:          QuadraticAssignment,
	RESOURCE_OPT:             ResourceOptimization,
	SCENE_ALLOCATION:         SceneAllocation,
	SET_COVER:                SetCover,
	SET_PACKING:              SetPacking,
	SET_SPLITTING:            SetSplitting,
	SHORTEST_PATH:            ShortestPath,
	STEINER_TREE:             SteinerTree,
	SUBSET_SUM:               SubsetSum,
	SUM_COLOR:                SumColoring,
	TOPOLOGICAL:              TopologicalSort,
	TPP:                      TravelingPurchaser,
	TSP:                      TravelingSalesman,
	VERTEX_COVER:             VertexCover,
	WAREHOUSE:                WarehouseLocation,
	WEAPON:                   WeaponTarget,
	WIDEST_PATH:              WidestPath,
}

var NoFiles = []string{LANGFORD_PAIR, MAGIC_SERIES, NQUEENS}

// Create problem test case name
func newName(problem string, n int) string {
	return fmt.Sprintf("%s%d", problem, n)
}
