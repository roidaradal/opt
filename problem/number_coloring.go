package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewNumberColoring creates a new Vertex Coloring problem that uses numbers as colors
func NewNumberColoring(variant string, n int) *discrete.Problem {
	name := newName(NumberColoring, variant, n)
	switch variant {
	case "sum":
		return sumColoring(name)
	default:
		return nil
	}
}

// Sum Coloring
func sumColoring(name string) *discrete.Problem {
	p, cfg := newGraphColoringProblem(name, data.GraphVertices, data.GraphNumbers)
	if p == nil || cfg == nil || len(cfg.Numbers) == 0 {
		return nil
	}
	graph := cfg.Graph

	p.AddUniversalConstraint(fn.ConstraintProperVertexColoring(graph))
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		total := list.Sum(list.MapList(dict.Values(solution.Map), cfg.Numbers))
		return discrete.Score(total)
	}
	return p
}
