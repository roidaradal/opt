package discrete

import (
	"math"

	"github.com/roidaradal/fn/lang"
)

var (
	Inf         = math.Inf(1)
	HardPenalty = Inf
)

type Penalty = float64
type ConstraintFn func(*Solution) bool

// Constraint interface
type Constraint interface {
	IsSatisfied(*Solution) bool
	ComputePenalty(*Solution) Penalty
}

// BaseConstraint type
type BaseConstraint struct {
	Penalty
	Variables []Variable
	Test      ConstraintFn
	// TODO: Add PartialTest for solvers with PartialSolution
}

// IsSatisfied checks if solution satisfies the constraint test
func (c BaseConstraint) IsSatisfied(solution *Solution) bool {
	return c.Test(solution)
}

// ComputePenalty computes the penalty of given solution
func (c BaseConstraint) ComputePenalty(solution *Solution) Penalty {
	return lang.Ternary(c.IsSatisfied(solution), 0, c.Penalty)
}

// GlobalConstraint is a constraint with more than 2 variables
type GlobalConstraint struct {
	BaseConstraint
}

// AddUniversalConstraint adds a GlobalConstraint to the problem (all problem variables are involved)
func (p *Problem) AddUniversalConstraint(test ConstraintFn) {
	constraint := GlobalConstraint{}
	constraint.Variables = p.Variables
	constraint.Test = test
	constraint.Penalty = lang.Ternary(p.Goal == Maximize, -HardPenalty, HardPenalty)
	p.AddConstraint(constraint)
}

// AddGlobalConstraint adds a GlobalConstraint to the problem, with given penalty and variables
func (p *Problem) AddGlobalConstraint(test ConstraintFn, penalty Penalty, variables ...Variable) {
	constraint := GlobalConstraint{}
	constraint.Variables = variables
	constraint.Test = test
	constraint.Penalty = penalty
	p.AddConstraint(constraint)
}
