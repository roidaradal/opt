package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	Allocation         = "allocation"
	Assignment         = "assignment"
	BinCover           = "bin_cover"
	BinPacking         = "bin_packing"
	CarPainting        = "car_painting"
	CarSequencing      = "car_sequencing"
	Clique             = "clique"
	CliqueCover        = "clique_cover"
	DominatingSet      = "dominating_set"
	EdgeColoring       = "edge_coloring"
	EdgeCover          = "edge_cover"
	FlowShopScheduling = "flow_shop"
	GraphMatching      = "graph_matching"
	GraphPartition     = "graph_partition"
	GraphPath          = "graph_path"
	GraphTour          = "graph_tour"
	IndependentSet     = "independent_set"
	InducedPath        = "induced_path"
	Interval           = "interval"
	KCenter            = "k_center"
	KCut               = "k_cut"
	Knapsack           = "knapsack"
	MaxCoverage        = "max_coverage"
	NumberColoring     = "number_coloring"
	NumberPartition    = "number_partition"
	NurseScheduling    = "nurse_scheduling"
	Satisfaction       = "satisfaction"
	SetCover           = "set_cover"
	SetPacking         = "set_packing"
	SetSplitting       = "set_splitting"
	SpanningTree       = "spanning_tree"
	SteinerTree        = "steiner_tree"
	Subsequence        = "subsequence"
	SubsetSum          = "subset_sum"
	TravelingPurchaser = "traveling_purchaser"
	TravelingSalesman  = "traveling_salesman"
	VertexColoring     = "verttex_coloring"
	VertexCover        = "vertex_cover"
	WarehouseLocation  = "warehouse_location"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	Allocation:         NewAllocation,
	Assignment:         NewAssignment,
	BinCover:           NewBinCover,
	BinPacking:         NewBinPacking,
	CarPainting:        NewCarPainting,
	CarSequencing:      NewCarSequencing,
	Clique:             NewClique,
	CliqueCover:        NewCliqueCover,
	DominatingSet:      NewDominatingSet,
	EdgeColoring:       NewEdgeColoring,
	EdgeCover:          NewEdgeCover,
	FlowShopScheduling: NewFlowShopScheduling,
	GraphMatching:      NewGraphMatching,
	GraphPartition:     NewGraphPartition,
	GraphPath:          NewGraphPath,
	GraphTour:          NewGraphTour,
	IndependentSet:     NewIndependentSet,
	InducedPath:        NewInducedPath,
	Interval:           NewInterval,
	KCenter:            NewKCenter,
	KCut:               NewKCut,
	Knapsack:           NewKnapsack,
	MaxCoverage:        NewMaxCoverage,
	NumberColoring:     NewNumberColoring,
	NumberPartition:    NewNumberPartition,
	NurseScheduling:    NewNurseScheduling,
	Satisfaction:       NewSatisfaction,
	SetCover:           NewSetCover,
	SetPacking:         NewSetPacking,
	SetSplitting:       NewSetSplitting,
	SpanningTree:       NewSpanningTree,
	SteinerTree:        NewSteinerTree,
	Subsequence:        NewSubsequence,
	SubsetSum:          NewSubsetSum,
	TravelingPurchaser: NewTravelingPurchaser,
	TravelingSalesman:  NewTravelingSalesman,
	VertexColoring:     NewVertexColoring,
	VertexCover:        NewVertexCover,
	WarehouseLocation:  NewWarehouseLocation,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
