package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new N-Queens problem
func NQueens(n int) *discrete.Problem {
	name := newName(NQUEENS, n)
	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy

	p.Variables = discrete.RangeVariables(1, n)
	domain := discrete.RangeDomain(1, n)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// No row conflict
	p.AddUniversalConstraint(constraint.AllUnique)

	// No diagonal conflict
	test := func(solution *discrete.Solution) bool {
		row := solution.Map
		occupied := ds.NewSet[ds.Coords]()
		for _, x := range p.Variables {
			occupied.Add(ds.Coords{row[x], x})
		}
		for _, x := range p.Variables {
			coords := ds.Coords{row[x], x}
			if hasDiagonalConflict(coords, occupied, n) {
				return false
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	// TODO: Update SolutionCoreFn
	p.SolutionCoreFn = fn.Core_MirroredValues[int](p, nil)
	p.SolutionStringFn = fn.String_Values[int](p, nil)

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
