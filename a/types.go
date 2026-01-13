// Package a contains types used in the discrete optimization problems
package a

import (
	"math"
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/str"
)

var (
	Inf    = math.Inf(1)
	NegInf = math.Inf(-1)
)

type StrInt struct {
	Str string
	Int int
}

// Create new StrInt
func NewStrInt(str string, x int) StrInt {
	return StrInt{str, x}
}

// Destructure StrInt
func (s StrInt) Tuple() (string, int) {
	return s.Str, s.Int
}

// Config for Assignment problems
type AssignmentCfg struct {
	Tasks      []string
	Workers    []string
	Cost       [][]float64
	Teams      [][]string
	MaxPerTeam int
}

type Subsets struct {
	Universal []string
	Names     []string
	Subsets   [][]string
}

// Create new Subsets
func NewSubsets(universalLine string, subsetLines []string) *Subsets {
	numSubsets := len(subsetLines)
	names := make([]string, numSubsets)
	subsets := make([][]string, numSubsets)
	for i, line := range subsetLines {
		parts := str.CleanSplit(line, ":")
		names[i] = parts[0]
		subsets[i] = strings.Fields(parts[1])
	}
	return &Subsets{
		Universal: strings.Fields(strings.TrimSpace(universalLine)),
		Names:     names,
		Subsets:   subsets,
	}
}

// Config for Bin problems
type BinProblemCfg struct {
	NumBins  int
	Capacity float64 // Min / Max Capacity
	Weight   []float64
}

// Config for Path problems
type PathCfg struct {
	Start    int
	End      int
	IndexOf  map[int]int // VariableIndex => OriginalIndex
	Vertices []ds.Vertex
	Between  []ds.Vertex
	Distance [][]float64
}

// Config for Traveling Salesman problems
type TSPCfg struct {
	Cities   []string
	Distance [][]float64
}
