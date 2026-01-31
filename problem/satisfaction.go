package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewSatisfaction creates a new Satisfaction problem
func NewSatisfaction(variant string, n int) *discrete.Problem {
	name := newName(Satisfaction, variant, n)
	switch variant {
	case "langford":
		return langfordPair(name, n)
	case "magic_series":
		return magicSeries(name, n)
	case "n_queens":
		return nQueens(name, n)
	default:
		return nil
	}
}

// Langford Pair
func langfordPair(name string, n int) *discrete.Problem {
	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Satisfy

	numPositions := n * 2
	numbers := make([]int, 0, numPositions)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i, i)
	}

	p.Variables = discrete.Variables(numbers)
	p.AddVariableDomains(discrete.IndexDomain(numPositions))

	p.AddUniversalConstraint(fn.ConstraintAllUnique)
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Distance constraint
		index := solution.Map
		for x := 0; x < numPositions; x += 2 {
			// Check that gap between number pair == number
			n := (x / 2) + 1
			gap := number.Abs(index[x+1]-index[x]) - 1
			if gap != n {
				return false
			}
		}
		return true
	})

	p.SolutionCoreFn = fn.CoreMirroredSequence(numbers)
	p.SolutionStringFn = fn.StringSequence(numbers)
	return p
}

// Magic Series
func magicSeries(name string, n int) *discrete.Problem {
	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment
	p.Goal = discrete.Satisfy

	domain := discrete.RangeDomain(0, n)
	p.Variables = discrete.RangeVariables(0, n)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check if number assigned at index x is also
		// number of times x appears in solution
		value := solution.Map
		count := dict.TallyValues(solution.Map, domain)
		return list.All(p.Variables, func(x discrete.Variable) bool {
			return value[x] == count[x]
		})
	})

	p.SolutionStringFn = fn.StringValues[int](p, nil)
	return p
}

// N-Queens
func nQueens(name string, n int) *discrete.Problem {
	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Satisfy

	p.Variables = discrete.RangeVariables(1, n)
	p.AddVariableDomains(discrete.RangeDomain(1, n))

	// No row conflict
	p.AddUniversalConstraint(fn.ConstraintAllUnique)
	// Check no diagonal conflict
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Gather coords occupied by queens
		row := solution.Map
		occupied := ds.NewSet[ds.Coords]()
		for _, x := range p.Variables {
			occupied.Add(ds.Coords{row[x], x})
		}
		// Check each queen for diagonal conflicts
		for _, x := range p.Variables {
			coords := ds.Coords{row[x], x}
			if hasDiagonalConflict(coords, occupied, n) {
				return false
			}
		}
		return true
	})

	// TODO: Update SolutionCoreFn
	p.SolutionCoreFn = fn.CoreMirroredValues[int](p, nil)
	p.SolutionStringFn = fn.StringValues[int](p, nil)
	return p
}

// Check if N-Queens solution has diagonal conflict
func hasDiagonalConflict(coords ds.Coords, occupied *ds.Set[ds.Coords], n int) bool {
	row, col := coords.Tuple()
	// Upper Left
	for y, x := row-1, col-1; y >= 1 && x >= 1; {
		if occupied.Has(ds.Coords{y, x}) {
			return true
		}
		y, x = y-1, x-1
	}
	// Upper Right
	for y, x := row-1, col+1; y >= 1 && x <= n; {
		if occupied.Has(ds.Coords{y, x}) {
			return true
		}
		y, x = y-1, x+1
	}
	// Bottom Left
	for y, x := row+1, col-1; y <= n && x >= 1; {
		if occupied.Has(ds.Coords{y, x}) {
			return true
		}
		y, x = y+1, x-1
	}
	// Bottom Right
	for y, x := row+1, col+1; y <= n && x <= n; {
		if occupied.Has(ds.Coords{y, x}) {
			return true
		}
		y, x = y+1, x+1
	}
	return false
}
