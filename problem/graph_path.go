package problem

import (
	"slices"

	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewGraphPath creates a new Graph Path problem
func NewGraphPath(variant string, n int) *discrete.Problem {
	name := newName(GraphPath, variant, n)
	switch variant {
	case "longest":
		return longestPath(name)
	case "minimax":
		return minimaxPath(name)
	case "shortest":
		return shortestPath(name)
	case "widest":
		return widestPath(name)
	default:
		return nil
	}
}

// Common steps for creating a Graph Path problem
func newGraphPathProblem(name string) (*discrete.Problem, *data.GraphPath) {
	cfg := data.NewGraphPath(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Path

	p.Variables = discrete.Variables(cfg.Between)
	p.AddVariableDomains(discrete.PathDomain(len(cfg.Between)))

	p.AddUniversalConstraint(fn.ConstraintSimplePath(cfg))
	p.SolutionStringFn = fn.StringGraphPath(cfg)
	return p, cfg
}

// Longest Path
func longestPath(name string) *discrete.Problem {
	p, cfg := newGraphPathProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.Goal = discrete.Maximize
	p.ObjectiveFn = fn.ScorePathCost(cfg)
	return p
}

// Shortest Path
func shortestPath(name string) *discrete.Problem {
	p, cfg := newGraphPathProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScorePathCost(cfg)
	return p
}

// Minimax Path
func minimaxPath(name string) *discrete.Problem {
	p, cfg := newGraphPathProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find max-weight edge of path
		return slices.Max(fn.PathDistances(solution, cfg))
	}
	return p
}

// Widest Path
func widestPath(name string) *discrete.Problem {
	p, cfg := newGraphPathProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.Goal = discrete.Maximize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find min-weight edge of path
		return slices.Min(fn.PathDistances(solution, cfg))
	}
	return p
}
