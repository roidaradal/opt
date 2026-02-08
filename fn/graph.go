package fn

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
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

// ConnectedComponents performs BFS traversal multiple times to discover the connected components
// of the graph, considering the active edge set, and returns the list of components
func ConnectedComponents(graph *ds.Graph, activeEdges ds.EdgeSet) [][]ds.Vertex {
	covered := dict.Flags(graph.Vertices, false)
	components := make([][]ds.Vertex, 0)
	for _, vertex := range graph.Vertices {
		if covered[vertex] {
			continue // skip if already covered
		}
		component := graph.BFSTraversal(vertex, activeEdges)
		components = append(components, component)
		for _, v := range component {
			covered[v] = true
		}
	}
	return components
}

// IsEulerianPath checks if edge sequence is a valid Eulerian path (visit each edge exactly once)
// Also returns the head/tail of the sequence
func IsEulerianPath(graph *ds.Graph, edges []ds.Edge) (bool, [2]ds.Vertex) {
	var pair [2]ds.Vertex
	if len(edges) < 2 {
		return false, pair
	}
	visitCount := dict.NewCounter(graph.Edges)
	a1, b1 := edges[0].Tuple()
	a2, b2 := edges[1].Tuple()
	var head, tail ds.Vertex
	switch {
	case a1 == a2:
		head, tail = b1, b2
	case a1 == b2:
		head, tail = b1, a2
	case b1 == a2:
		head, tail = a1, b2
	case b1 == b2:
		head, tail = a1, a2
	default:
		return false, pair
	}
	visitCount[edges[0]] += 1
	visitCount[edges[1]] += 1
	for _, edge := range edges[2:] {
		visitCount[edge] += 1
		a, b := edge.Tuple()
		switch tail {
		case a:
			tail = b
		case b:
			tail = a
		default:
			return false, pair
		}
	}
	// Check that all edges visited exactly once
	return list.AllEqual(dict.Values(visitCount), 1), [2]ds.Vertex{head, tail}
}

// IsHamiltonianPath checks if vertex path is a valid Hamiltonian path (visit each vertex exactly once)
func IsHamiltonianPath(graph *ds.Graph, vertices []ds.Vertex) bool {
	numVertices := len(vertices)
	if numVertices == 0 {
		return false
	}
	visitCount := dict.NewCounter(graph.Vertices)
	for i, curr := range vertices {
		visitCount[curr] += 1
		if i == numVertices-1 {
			break
		}
		next := vertices[i+1]
		if graph.NeighborsOf[curr].HasNo(next) {
			return false // invalid path if no edge from curr -> next
		}
	}
	// Check all vertices visited exactly once
	return list.AllEqual(dict.Values(visitCount), 1)
}

// PathDistances returns the distance values of the edges in the path
func PathDistances(solution *discrete.Solution, cfg *data.GraphPath) []float64 {
	distances := make([]float64, 0)
	path := AsGraphPath(solution, cfg)
	prev := path[0]
	for _, curr := range path[1:] {
		distances = append(distances, cfg.Distance[prev][curr])
		prev = curr
	}
	return distances
}
