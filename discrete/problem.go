// Package discrete contains the common parts of discrete optimization problems
package discrete

import "github.com/roidaradal/fn/list"

// Goal types
const (
	Maximize Goal = "max"
	Minimize Goal = "min"
	Satisfy  Goal = "sat"
)

// Problem types
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
	Variables   []Variable
	Domain      map[Variable][]Value
	Constraints []Constraint
	Type        ProblemType
	Goal
	ObjectiveFn
	SolutionCoreFn
	SolutionStringFn
	SolutionDisplayFn
}

// NewProblem creates a new Problem
func NewProblem(name string) *Problem {
	return &Problem{
		Name:        name,
		Variables:   make([]Variable, 0),
		Domain:      make(map[Variable][]Value),
		Constraints: make([]Constraint, 0),
	}
}

// AddConstraint adds a new problem constraint
func (p *Problem) AddConstraint(constraint Constraint) {
	p.Constraints = append(p.Constraints, constraint)
}

// IsSatisfied checks if solution satisfies all problem constraints
func (p *Problem) IsSatisfied(solution *Solution) bool {
	return list.All(p.Constraints, func(constraint Constraint) bool {
		return constraint.IsSatisfied(solution)
	})
}

// IsSatisfaction checks if satisfaction problem
func (p *Problem) IsSatisfaction() bool {
	return p.Goal == Satisfy
}

// IsOptimization checks if optimization problem
func (p *Problem) IsOptimization() bool {
	return p.Goal == Minimize || p.Goal == Maximize
}

// ComputeScore computes and attach solution score by calling the ObjectiveFn
func (p *Problem) ComputeScore(solution *Solution) {
	if p.ObjectiveFn != nil {
		solution.Score = p.ObjectiveFn(solution)
	}
}
