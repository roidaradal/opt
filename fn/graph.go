package fn

import "github.com/roidaradal/fn/ds"

// IsClique checks if a list of vertices forms a clique in the graph
func IsClique(graph *ds.Graph, vertices []ds.Vertex) bool {
	vertexSet := ds.SetFrom(vertices)
	for _, vertex := range vertices {
		adjacent := ds.SetFrom(graph.Neighbors(vertex))
		adjacent.Add(vertex)
		if vertexSet.Difference(adjacent).NotEmpty() {
			return false
		}
	}
	return true
}
