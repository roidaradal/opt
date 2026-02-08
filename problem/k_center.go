package problem

import (
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewKCenter creates a new K-Center problem
func NewKCenter(variant string, n int) *discrete.Problem {
	name := newName(KCenter, variant, n)
	switch variant {
	case "basic":
		return kCenter(name)
	default:
		return nil
	}
}

// Common steps to creating a K-Center problem
func newKCenterProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newGraphSubsetProblem(name, data.GraphVertices)
	if p == nil || graph == nil || graph.K == 0 {
		return nil, nil
	}
	if len(graph.Edges) != len(graph.EdgeWeight) {
		return nil, nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check selected vertices has correct count
		return len(fn.AsSubset(solution)) == graph.K
	})

	p.Goal = discrete.Minimize
	return p, graph
}

// K-Center
func kCenter(name string) *discrete.Problem {
	p, graph := newKCenterProblem(name)
	if p == nil || graph == nil {
		return nil
	}
	pairEdgeIndex := make(map[[2]int]int)
	for i, edge := range graph.Edges {
		v1, v2 := edge.Tuple()
		x1, x2 := graph.IndexOf[v1], graph.IndexOf[v2]
		pairEdgeIndex[[2]int{x1, x2}] = i
		pairEdgeIndex[[2]int{x2, x1}] = i
	}
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Compute maximum shortest distance of any vertex to selected vertices
		selected := fn.AsSubset(solution)
		var maxDistance float64 = 0
		for i := range graph.Vertices {
			minDistance := discrete.Inf
			for _, j := range selected {
				edgeIndex := pairEdgeIndex[[2]int{i, j}]
				minDistance = min(minDistance, graph.EdgeWeight[edgeIndex])
			}
			maxDistance = max(maxDistance, minDistance)
		}
		return maxDistance
	}
	return p
}
