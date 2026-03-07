package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSteinerTree creates a new Steiner Tree problem
func NewSteinerTree(variant string, n int) *discrete.Problem {
	name := newName(SteinerTree, variant, n)
	switch variant {
	case "basic":
		return steinerTree(name)
	case "degree":
		return degreeConstrainedSteinerTree(name)
	case "group":
		return groupSteinerTree(name)
	default:
		return nil
	}
}

// Common steps for creating a Steiner Tree problem
func newSteinerTreeProblem(name string) (*discrete.Problem, *data.Graph) {
	p, graph := newSpanningTreeProblem(name, data.SpanTerminals)
	p = edgeWeightedProblem(p, graph)
	if p == nil || graph == nil || len(graph.Terminals) == 0 {
		return nil, nil
	}
	return p, graph
}

// Steiner Tree
func steinerTree(name string) *discrete.Problem {
	p, _ := newSteinerTreeProblem(name)
	return p
}

// Degree-Constrained Steiner Tree
func degreeConstrainedSteinerTree(name string) *discrete.Problem {
	p, cfg := newSteinerTreeProblem(name)
	if p == nil || cfg == nil || cfg.K == 0 {
		return nil
	}
	graph := cfg.Graph

	// Ensure that all vertex has degree less than or equal to K
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		activeEdges := ds.SetFrom(list.MapList(fn.AsSubset(solution), graph.Edges))
		for _, vertex := range cfg.Vertices {
			neighbors := graph.ActiveNeighbors(vertex, activeEdges)
			if len(neighbors) > cfg.K {
				return false
			}
		}
		return true
	})

	return p
}

// Group Steiner Tree
func groupSteinerTree(name string) *discrete.Problem {
	p, cfg := newGraphSubsetProblem(name, data.GraphEdges)
	if p == nil || cfg == nil || len(cfg.Groups) == 0 {
		return nil
	}
	p = edgeWeightedProblem(p, cfg)
	p.Goal = discrete.Minimize

	// Constraint: solution forms a tree where at least one vertex in each group is spanned
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check each group has at least one vertex covered
		spannedVertices := fn.SpannedVertices(solution, cfg.Graph)
		for _, group := range cfg.Groups {
			groupSet := ds.SetFrom(group)
			if groupSet.Intersection(spannedVertices).IsEmpty() {
				return false
			}
		}
		return true
	})
	return p
}
