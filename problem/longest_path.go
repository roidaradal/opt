package problem

import (
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Longest Path problem
func LongestPath(n int) *discrete.Problem {
	name := newName(LONGEST_PATH, n)
	cfg := fn.NewPathProblem(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Path

	p.Variables = discrete.Variables(cfg.Between)
	domain := discrete.IndexDomain(len(cfg.Between))
	domain = append(domain, -1) // for not included
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	p.AddUniversalConstraint(constraint.SimplePath(cfg))

	p.ObjectiveFn = fn.Score_PathCost(cfg)
	p.SolutionStringFn = fn.String_Path(cfg)

	return p
}
