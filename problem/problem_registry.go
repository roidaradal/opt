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
	EdgeCover       = "edge_cover"
	GraphPartition  = "graph_partition"
	IndependentSet  = "independent_set"
	Interval        = "interval"
	Knapsack        = "knapsack"
	MaxCoverage     = "max_coverage"
	NumberPartition = "number_partition"
	Satisfaction    = "satisfaction"
	SetCover        = "set_cover"
	SetPacking      = "set_packing"
	SetSplitting    = "set_splitting"
	VertexCover     = "vertex_cover"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	BinCover:        NewBinCover,
	BinPacking:      NewBinPacking,
	CliqueCover:     NewCliqueCover,
	DominatingSet:   NewDominatingSet,
	EdgeCover:       NewEdgeCover,
	GraphPartition:  NewGraphPartition,
	IndependentSet:  NewIndependentSet,
	Interval:        NewInterval,
	Knapsack:        NewKnapsack,
	MaxCoverage:     NewMaxCoverage,
	NumberPartition: NewNumberPartition,
	Satisfaction:    NewSatisfaction,
	SetCover:        NewSetCover,
	SetPacking:      NewSetPacking,
	SetSplitting:    NewSetSplitting,
	VertexCover:     NewVertexCover,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
