// Package a contains types used in the discrete optimization problems
package a

import (
	"strings"

	"github.com/roidaradal/fn/str"
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
