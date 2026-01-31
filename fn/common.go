package fn

import (
	"slices"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
)

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

// Convert partition groups into sorted partition groups, wraped in braces
func sortedPartitionGroups(groups [][]string) []string {
	return list.Map(groups, func(group []string) string {
		slices.Sort(group)
		return str.WrapBraces(group)
	})
}
