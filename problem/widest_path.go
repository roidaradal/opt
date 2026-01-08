package problem

import (
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Widest Path problem
func WidestPath(n int) *discrete.Problem {
	name := newName(WIDEST_PATH, n)
	cfg := fn.NewPathProblem(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Path

	p.Variables = discrete.Variables(cfg.Between)
	domain := discrete.PathDomain(len(cfg.Between))
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	p.AddUniversalConstraint(constraint.SimplePath(cfg))

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Find min-weight edge of path
		var minWeight discrete.Score = a.Inf
		path := fn.AsPath(solution, cfg)
		prev := path[0]
		for _, curr := range path[1:] {
			minWeight = min(minWeight, cfg.Distance[prev][curr])
			prev = curr
		}
		return minWeight
	}

	p.SolutionStringFn = fn.String_Path(cfg)

	return p
}
