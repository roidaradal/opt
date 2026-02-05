package data

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/number"
)

type Graph struct {
	*ds.Graph
	K           int
	EdgeWeight  []float64
	EdgeColor   []string
	VertexColor []string
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

// NewUndirectedGraph loads a Graph config
func NewUndirectedGraph(name string) *Graph {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Graph{
		Graph:       ds.GraphFrom(data["vertices"], data["edges"]),
		K:           number.ParseInt(data["k"]),
		EdgeWeight:  floatList(data["edgeWeight"]),
		EdgeColor:   stringList(data["edgeColor"]),
		VertexColor: stringList(data["vertexColor"]),
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

type GraphVariablesFn = func(*ds.Graph) []string

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
