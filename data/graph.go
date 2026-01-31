package data

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/number"
)

type Graph struct {
	*ds.Graph
	VertexColor []string
}

type GraphPartition struct {
	*ds.Graph
	EdgeWeight    []float64
	NumPartitions int
	MinSize       int
}

// NewUndirectedGraph loads a Graph config
func NewUndirectedGraph(name string) *Graph {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Graph{
		Graph:       ds.GraphFrom(data["vertices"], data["edges"]),
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

// GraphVertices returns graph vertices
func GraphVertices(graph *Graph) []string {
	return graph.Vertices
}

// GraphEdges returns graph edge names
func GraphEdges(graph *Graph) []string {
	return graph.EdgeNames()
}
