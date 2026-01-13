// Package fn contains various functions used in discrete optimization problems
package fn

import (
	"fmt"
	"slices"
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
		return a.Inf
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

// Load new test case containing directed graph
func NewDirectedGraph(name string) *ds.Graph {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return ds.DirectedGraphFrom(lines[0], lines[1])
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

// Load new test case for vertex coloring
func NewVertexColoring(name string) (*ds.Graph, int) {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil, 0
	}
	numColors := number.ParseInt(lines[0])
	graph := ds.GraphFrom(lines[1], lines[2])
	return graph, numColors
}

// Load new test case for path problem
func NewPathProblem(name string) *a.PathCfg {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) < 3 {
		return nil
	}

	cfg := &a.PathCfg{}
	parts := strings.Fields(lines[0])
	if len(parts) != 2 {
		return nil
	}
	start, end := parts[0], parts[1]

	cfg.Vertices = strings.Fields(lines[1])
	if !slices.Contains(cfg.Vertices, start) || !slices.Contains(cfg.Vertices, end) {
		return nil
	}
	cfg.Start = slices.Index(cfg.Vertices, start)
	cfg.End = slices.Index(cfg.Vertices, end)

	numVertices := len(cfg.Vertices)
	cfg.Distance = make([][]float64, numVertices)
	for i := range numVertices {
		cfg.Distance[i] = list.Map(strings.Fields(lines[2+i])[1:], ParseFloatInf)
	}

	cfg.IndexOf = make(map[int]int)
	cfg.Between = make([]ds.Vertex, 0, numVertices)
	betweenIdx := 0
	for i, vertex := range cfg.Vertices {
		if i == cfg.Start || i == cfg.End {
			continue
		}
		cfg.Between = append(cfg.Between, vertex)
		cfg.IndexOf[betweenIdx] = i
		betweenIdx += 1
	}
	return cfg
}

// Load new test case for traveling salesman problems
func NewTravelingSalesman(name string) *a.TSPCfg {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	cfg := &a.TSPCfg{
		Cities:   strings.Fields(lines[0]),
		Distance: make([][]float64, 0),
	}
	for _, line := range lines[1:] {
		d := list.Map(strings.Fields(line), ParseFloatInf)
		cfg.Distance = append(cfg.Distance, d)
	}
	return cfg
}
