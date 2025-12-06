package a

import (
	"strings"

	"github.com/roidaradal/fn/str"
)

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
