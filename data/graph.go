package data

import (
	"slices"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/number"
)

type Graph struct {
	*ds.Graph
	K           int
	EdgeWeight  []float64
	EdgeColor   []string
	VertexColor []string
	Terminals   []string
}

type GraphPartition struct {
	*ds.Graph
	EdgeWeight    []float64
	NumPartitions int
	MinSize       int
}

type GraphColoring struct {
	*ds.Graph
	Colors  []string
	Numbers []int
}

type GraphPath struct {
	Start         int         // index of start vertex in vertices list
	End           int         // index of end vertex in vertices list
	OriginalIndex map[int]int // map VariableIndex => OriginalIndex
	Vertices      []ds.Vertex
	Between       []ds.Vertex // list of vertices that are not start, end
	Distance      [][]float64
}

// NewUndirectedGraph loads an undirected Graph config
func NewUndirectedGraph(name string) *Graph {
	return newGraph(name, false)
}

// NewDirectedGraph loads a directed Graph config
func NewDirectedGraph(name string) *Graph {
	return newGraph(name, true)
}

// Common steps for creating a Graph config
func newGraph(name string, isDirected bool) *Graph {
	data, err := load(name)
	if err != nil {
		return nil
	}
	var graph *ds.Graph
	if isDirected {
		graph = ds.DirectedGraphFrom(data["vertices"], data["edges"])
	} else {
		graph = ds.GraphFrom(data["vertices"], data["edges"])
	}
	return &Graph{
		Graph:       graph,
		K:           number.ParseInt(data["k"]),
		EdgeWeight:  floatList(data["edgeWeight"]),
		EdgeColor:   stringList(data["edgeColor"]),
		VertexColor: stringList(data["vertexColor"]),
		Terminals:   stringList(data["terminals"]),
	}
}

// NewGraphPartition laods a GraphPartition config
func NewGraphPartition(name string) *GraphPartition {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &GraphPartition{
		Graph:         ds.GraphFrom(data["vertices"], data["edges"]),
		EdgeWeight:    floatList(data["edgeWeight"]),
		NumPartitions: number.ParseInt(data["numPartitions"]),
		MinSize:       number.ParseInt(data["minSize"]),
	}
}

// NewGraphColoring loads a GraphColoring config
func NewGraphColoring(name string) *GraphColoring {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &GraphColoring{
		Graph:   ds.GraphFrom(data["vertices"], data["edges"]),
		Colors:  stringList(data["colors"]),
		Numbers: intList(data["numbers"]),
	}
}

// NewGraphPath loads a GraphPath config
func NewGraphPath(name string) *GraphPath {
	data, err := load(name)
	if err != nil {
		return nil
	}
	cfg := &GraphPath{
		Vertices:      stringList(data["vertices"]),
		OriginalIndex: make(map[int]int),
	}
	numVertices := len(cfg.Vertices)
	start, end := data["start"], data["end"]
	if !slices.Contains(cfg.Vertices, start) || !slices.Contains(cfg.Vertices, end) {
		return nil
	}
	cfg.Start = slices.Index(cfg.Vertices, start)
	cfg.End = slices.Index(cfg.Vertices, end)
	cfg.Distance = make([][]float64, numVertices)
	for i, line := range parseList(data["distance"]) {
		cfg.Distance[i] = matrixRow(line, true)
	}
	cfg.Between = make([]ds.Vertex, 0, numVertices-2)
	betweenIdx := 0
	for i, vertex := range cfg.Vertices {
		if i == cfg.Start || i == cfg.End {
			continue
		}
		cfg.Between = append(cfg.Between, vertex)
		cfg.OriginalIndex[betweenIdx] = i
		betweenIdx += 1
	}
	return cfg
}

// NewGraphTour loads a GraphPath config with only vertices and distance matrix
func NewGraphTour(name string) *GraphPath {
	data, err := load(name)
	if err != nil {
		return nil
	}
	cfg := &GraphPath{
		Vertices: stringList(data["vertices"]),
	}
	cfg.Distance = make([][]float64, len(cfg.Vertices))
	for i, line := range parseList(data["distance"]) {
		cfg.Distance[i] = matrixRow(line, true)
	}
	return cfg
}

type GraphVariablesFn = func(*ds.Graph) []string

type GraphSpanFn = func(*Graph) []string

type GraphColorsFn[T any] = func(*GraphColoring) []T

// GraphVertices returns graph vertices
func GraphVertices(graph *ds.Graph) []string {
	return graph.Vertices
}

// GraphEdges returns graph edge names
func GraphEdges(graph *ds.Graph) []string {
	return graph.EdgeNames()
}

// GraphColors returns graph colors
func GraphColors(cfg *GraphColoring) []string {
	return cfg.Colors
}

// GraphNumbers returns numbers as graph colors
func GraphNumbers(cfg *GraphColoring) []int {
	return cfg.Numbers
}

// SpanVertices returns graph vertices
func SpanVertices(graph *Graph) []string {
	return graph.Vertices
}

// SpanTerminals returns graph terminals
func SpanTerminals(graph *Graph) []string {
	return graph.Terminals
}
