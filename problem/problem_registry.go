package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	BinCover        = "bin_cover"
	BinPacking      = "bin_packing"
	CliqueCover     = "clique_cover"
	DominatingSet   = "dominating_set"
	EdgeColoring    = "edge_coloring"
	EdgeCover       = "edge_cover"
	GraphMatching   = "graph_matching"
	GraphPartition  = "graph_partition"
	GraphTour       = "graph_tour"
	IndependentSet  = "independent_set"
	Interval        = "interval"
	Knapsack        = "knapsack"
	MaxCoverage     = "max_coverage"
	NumberColoring  = "number_coloring"
	NumberPartition = "number_partition"
	Satisfaction    = "satisfaction"
	SetCover        = "set_cover"
	SetPacking      = "set_packing"
	SetSplitting    = "set_splitting"
	Subsequence     = "subsequence"
	SubsetSum       = "subset_sum"
	VertexCover     = "vertex_cover"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	BinCover:        NewBinCover,
	BinPacking:      NewBinPacking,
	CliqueCover:     NewCliqueCover,
	DominatingSet:   NewDominatingSet,
	EdgeColoring:    NewEdgeColoring,
	EdgeCover:       NewEdgeCover,
	GraphMatching:   NewGraphMatching,
	GraphPartition:  NewGraphPartition,
	GraphTour:       NewGraphTour,
	IndependentSet:  NewIndependentSet,
	Interval:        NewInterval,
	Knapsack:        NewKnapsack,
	MaxCoverage:     NewMaxCoverage,
	NumberColoring:  NewNumberColoring,
	NumberPartition: NewNumberPartition,
	Satisfaction:    NewSatisfaction,
	SetCover:        NewSetCover,
	SetPacking:      NewSetPacking,
	SetSplitting:    NewSetSplitting,
	Subsequence:     NewSubsequence,
	SubsetSum:       NewSubsetSum,
	VertexCover:     NewVertexCover,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
