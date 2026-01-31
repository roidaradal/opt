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
	NumberPartition = "number_partition"
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
	NumberPartition: NewNumberPartition,
	VertexCover:     NewVertexCover,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
