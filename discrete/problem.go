// Package discrete contains the common parts to discrete optimization problems
package discrete

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
)

const (
	Maximize Goal = "max"
	Minimize Goal = "min"
	Satisfy  Goal = "sat"
)

const (
	Assignment ProblemType = "assignment"
	Partition  ProblemType = "partition"
	Sequence   ProblemType = "sequence"
	Subset     ProblemType = "subset"
	Path       ProblemType = "path"
)

type (
	Goal        string
	ProblemType string
)

type Problem struct {
	Name        string
	Type        ProblemType
	Variables   []Variable
	Domain      map[Variable][]Value
	Constraints []Constraint
	Goal
	ObjectiveFn
	SolutionCoreFn
	SolutionStringFn
}

// Create new Problem
func NewProblem(name string) *Problem {
	return &Problem{
		Name:        name,
		Variables:   make([]Variable, 0),
		Domain:      make(map[Variable][]Value),
		Constraints: make([]Constraint, 0),
	}
}

// Add new constraint to problem
func (p *Problem) AddConstraint(constraint Constraint) {
	p.Constraints = append(p.Constraints, constraint)
}

// Compute the total solution space of the problem
func (p Problem) SolutionSpace() int {
	switch p.Type {
	case Sequence:
		domain := p.Domain[p.Variables[0]]
		return comb.Factorial(len(domain))
	case Path:
		domainSize := len(p.Domain[p.Variables[0]]) - 1 // remove -1
		count := 0
		for take := range domainSize + 1 {
			count += comb.NumPermutations(domainSize, take)
		}
		return count
	default:
		return list.Product(list.Map(dict.Values(p.Domain), list.Length))
	}
}

// Create the solution space equation of the problem
func (p Problem) SolutionSpaceEquation() string {
	switch p.Type {
	case Sequence:
		domain := p.Domain[p.Variables[0]]
		return fmt.Sprintf("%d!", len(domain))
	default:
		entries := dict.Entries(dict.CounterFunc(dict.Values(p.Domain), list.Length))
		slices.SortFunc(entries, func(a, b dict.Entry[int, int]) int {
			// Sort by descending counts
			return cmp.Compare(b.Value, a.Value)
		})
		equation := make([]string, 0)
		for _, e := range entries {
			size, count := e.Tuple()
			equation = append(equation, fmt.Sprintf("%s^%s", number.Comma(size), number.Comma(count)))
		}
		if len(equation) == 1 {
			return equation[0]
		}
		equation = list.Map(equation, func(expr string) string {
			return fmt.Sprintf("(%s)", expr)
		})
		return strings.Join(equation, " * ")
	}
}

// Check if solution satisfies the problem constraints
func (p Problem) IsSatisfied(solution *Solution) bool {
	return list.All(p.Constraints, func(constraint Constraint) bool {
		return constraint.IsSatisfied(solution)
	})
}

// Check if satisfaction problem
func (p Problem) IsSatisfaction() bool {
	return p.Goal == Satisfy
}

// Check if optimization problem
func (p Problem) IsOptimization() bool {
	return p.Goal == Minimize || p.Goal == Maximize
}

// Computes and attaches solution score by calling the ObjectiveFn
func (p Problem) ComputeScore(solution *Solution) {
	if p.ObjectiveFn != nil {
		solution.Score = p.ObjectiveFn(solution)
	}
}
