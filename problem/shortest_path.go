package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/a"
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
	domain := discrete.IndexDomain(len(cfg.Between))
	domain = append(domain, -1) // for not included
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		path := fn.AsPath(solution, cfg)
		prev := path[0]

		visited := ds.NewSet[int]()
		visited.Add(prev)
		for _, curr := range path[1:] {
			if visited.Has(curr) {
				return false // not a simple path (repeated vertex)
			}
			if cfg.Distance[prev][curr] == a.Inf {
				return false // no edge from prev => curr
			}
			visited.Add(curr)
			prev = curr
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_PathCost(cfg)
	p.SolutionStringFn = fn.String_Path(cfg)

	return p
}
