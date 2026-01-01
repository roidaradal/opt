// Package fn contains various functions used in discrete optimization problems
package fn

import (
	"fmt"
	"math"
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/a"
)

// Load problem test case
func LoadProblem(name string) ([]string, error) {
	path := fmt.Sprintf("data/%s.txt", name)
	lines, err := io.ReadLines(path)
	if err != nil {
		return nil, err
	}
	lines = list.Filter(lines, func(line string) bool {
		return !strings.HasPrefix(line, "#") && line != ""
	})
	return lines, nil
}

// Parse float or inf if "x"
func ParseFloatInf(x string) float64 {
	if x == "x" {
		return math.Inf(1)
	} else {
		return number.ParseFloat(x)
	}
}

// Load new test case containing subsets data
func NewSubsets(name string) *a.Subsets {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	return a.NewSubsets(lines[0], lines[1:])
}

// Load new test case containing unweighted graph
func NewUnweightedGraph(name string) *ds.Graph {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return ds.GraphFrom(lines[0], lines[1])
}

// Load new test case containing weighted graph
func NewWeightedGraph(name string) (*ds.Graph, []float64) {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil, nil
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	edgeWeight := list.Map(strings.Fields(lines[2]), number.ParseFloat)
	return graph, edgeWeight
}

// Load new test case for bin problem
func NewBinProblem(name string) *a.BinProblemCfg {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &a.BinProblemCfg{
		NumBins:  number.ParseInt(lines[0]),
		Capacity: number.ParseFloat(lines[1]),
		Weight:   list.Map(strings.Fields(lines[2]), number.ParseFloat),
	}
}
