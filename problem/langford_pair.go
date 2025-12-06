package problem

import (
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Langford Pair problem
func LangfordPair(n int) *discrete.Problem {
	name := newName(LANGFORD_PAIR, n)
	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy

	numPositions := n * 2
	numbers := make([]int, 0, numPositions)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i, i)
	}

	p.Variables = discrete.Variables(numbers)
	domain := discrete.IndexDomain(numPositions)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// AllUnique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	// Distance constraint
	test := func(solution *discrete.Solution) bool {
		index := solution.Map
		for x := 0; x < numPositions; x += 2 {
			// Check that the gap between the number pair == number
			n := (x / 2) + 1
			gap := number.Abs(index[x+1]-index[x]) - 1
			if gap != n {
				return false
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.SolutionCoreFn = fn.Core_MirroredSequence(numbers)
	p.SolutionStringFn = fn.String_Sequence(numbers)

	return p
}
