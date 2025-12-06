// Package discrete contains the common parts to discrete optimization problems
package discrete

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
)

const (
	Maximize Goal = "max"
	Minimize Goal = "min"
	Satisfy  Goal = "sat"
)

type Goal string

type Problem struct {
	Name        string
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
	return list.Product(list.Map(dict.Values(p.Domain), list.Length))
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
