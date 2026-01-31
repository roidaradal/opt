package problem

import (
	"fmt"

	"github.com/roidaradal/opt/discrete"
)

const (
	BinCover   = "bin_cover"
	BinPacking = "bin_packing"
)

var Creator = map[string]func(string, int) *discrete.Problem{
	BinCover:   NewBinCover,
	BinPacking: NewBinPacking,
}

// Create problem test case name
func newName(problem, variant string, n int) string {
	return fmt.Sprintf("%s.%s.%d", problem, variant, n)
}
