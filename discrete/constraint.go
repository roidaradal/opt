package discrete

import (
	"math"

	"github.com/roidaradal/fn/lang"
)

var HardPenalty Penalty = math.Inf(1)

type Penalty = float64
type ConstraintFn func(*Solution) bool

type Constraint interface {
	IsSatisfied(*Solution) bool
	ComputePenalty(*Solution) Penalty
}

// The base Constraint type
type BaseConstraint struct {
	Penalty
	Variables []Variable
	Test      ConstraintFn
	// TODO: Add PartialTest for solvers with PartialSolution
}

// Checks if solution satisfies the constraint test
func (c BaseConstraint) IsSatisfied(solution *Solution) bool {
	return c.Test(solution)
}

// Computes the penalty for the given solution
func (c BaseConstraint) ComputePenalty(solution *Solution) Penalty {
	return lang.Ternary(c.IsSatisfied(solution), 0, c.Penalty)
}

// Constraint with more than 2 variables
type GlobalConstraint struct {
	BaseConstraint
}

// Add Universal constraint to problem (all problem variables are involved)
func (p *Problem) AddUniversalConstraint(test ConstraintFn) {
	constraint := GlobalConstraint{}
	constraint.Variables = p.Variables
	constraint.Test = test
	constraint.Penalty = lang.Ternary(p.Goal == Maximize, -HardPenalty, HardPenalty)
	p.AddConstraint(constraint)
}

// Add Global constraint to problem, with given penalty and variables
func (p *Problem) AddGlobalConstraint(test ConstraintFn, penalty Penalty, variables ...Variable) {
	constraint := GlobalConstraint{}
	constraint.Variables = variables
	constraint.Test = test
	constraint.Penalty = penalty
	p.AddConstraint(constraint)
}
