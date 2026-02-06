package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewKCut creates a new K-Cut problem
func NewKCut(variant string, n int) *discrete.Problem {
	name := newName(KCut, variant, n)
	switch variant {
	case "min":
		return minKCut(name)
	case "max":
		return maxKCut(name)
	default:
		return nil
	}
}

// Common steps to creating a K-Cut problem
func newKCutProblem(name string, countTest func(int, int) bool) *discrete.Problem {
	p, graph := newGraphSubsetProblem(name, data.GraphEdges)
	if p == nil || graph == nil || graph.K == 0 {
		return nil
	}
	if len(graph.Edges) != len(graph.EdgeWeight) {
		return nil
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Remove selected cut edges from active edges
		cutEdges := ds.SetFrom(list.MapList(fn.AsSubset(solution), graph.Edges))
		activeEdges := ds.SetFrom(graph.Edges).Difference(cutEdges)
		// Make sure cut produced at least k connected components
		components := fn.ConnectedComponents(graph.Graph, activeEdges)
		return countTest(len(components), graph.K)
	})
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, graph.EdgeWeight)
	return p
}

// Min K-Cut
func minKCut(name string) *discrete.Problem {
	p := newKCutProblem(name, func(numComponents int, k int) bool {
		return numComponents >= k // produce at least k components
	})
	if p == nil {
		return nil
	}
	p.Goal = discrete.Minimize
	return p
}

// Max K-Cut
func maxKCut(name string) *discrete.Problem {
	p := newKCutProblem(name, func(numComponents int, k int) bool {
		return numComponents == k // produce exactly k components
	})
	if p == nil {
		return nil
	}
	p.Goal = discrete.Maximize
	return p
}
