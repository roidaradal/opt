package discrete

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
)

type (
	Score             = float64
	ObjectiveFn       func(*Solution) Score
	SolutionCoreFn    func(*Solution) string
	SolutionStringFn  func(*Solution) string
	SolutionDisplayFn func(*Solution) string
)

type Solution struct {
	Score
	Map           map[Variable]Value
	VariableOrder []Variable
	IsPartial     bool
}

// NewSolution creates a new blank solution
func NewSolution() *Solution {
	return &Solution{
		IsPartial:     true,
		Map:           make(map[Variable]Value),
		VariableOrder: make([]Variable, 0),
	}
}

// ZipSolution creates a new solution by zipping variables, values
func ZipSolution(variables []Variable, values []Value) *Solution {
	return &Solution{
		IsPartial:     false,
		Map:           dict.Zip(variables, values),
		VariableOrder: variables,
	}
}

// Assign assigns variable=value in solution
func (s *Solution) Assign(variable Variable, value Value) {
	s.Map[variable] = value
	s.VariableOrder = append(s.VariableOrder, variable)
}

// Values gets values of solution in arbitrary order
func (s *Solution) Values() []Value {
	return dict.Values(s.Map)
}

// Tuple gets values of solution, ordered by problem variable order
func (s *Solution) Tuple(p *Problem) []Value {
	return list.Translate(p.Variables, s.Map)
}

// Length gets solution length
func (s *Solution) Length() int {
	return len(s.Map)
}
