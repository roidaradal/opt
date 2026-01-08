package problem

import (
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Shortest Path problem
func ShortestPath(n int) *discrete.Problem {
	name := newName(SHORTEST_PATH, n)
	cfg := fn.NewPathProblem(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Path

	p.Variables = discrete.Variables(cfg.Between)
	domain := discrete.PathDomain(len(cfg.Between))
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	p.AddUniversalConstraint(constraint.SimplePath(cfg))

	p.ObjectiveFn = fn.Score_PathCost(cfg)
	p.SolutionStringFn = fn.String_Path(cfg)

	return p
}
