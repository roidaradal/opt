package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Magic Series problem
func MagicSeries(n int) *discrete.Problem {
	name := newName(MAGIC_SERIES, n)
	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy
	p.Type = discrete.Assignment

	p.Variables = discrete.RangeVariables(0, n)
	domain := discrete.RangeDomain(0, n)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Check if number assigned at index x is also
		// the number of times x appears in the solution
		value := solution.Map
		count := dict.TallyValues(solution.Map, domain)
		return list.All(p.Variables, func(x discrete.Variable) bool {
			return value[x] == count[x]
		})
	}
	p.AddUniversalConstraint(test)

	p.SolutionStringFn = fn.String_Values[int](p, nil)

	return p
}
