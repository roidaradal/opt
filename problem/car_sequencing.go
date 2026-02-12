package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewCarSequencing creates a new Car Sequencing problem
func NewCarSequencing(variant string, n int) *discrete.Problem {
	name := newName(CarSequencing, variant, n)
	switch variant {
	case "basic":
		return carSequencing(name)
	default:
		return nil
	}
}

// Car Sequencing
func carSequencing(name string) *discrete.Problem {
	cfg := data.NewCarSequence(name)
	if cfg == nil {
		return nil
	}
	if len(cfg.OptionMax) != len(cfg.OptionWindow) {
		return nil
	}
	numCars, numOptions := len(cfg.Cars), len(cfg.OptionMax)

	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Satisfy

	p.Variables = discrete.Variables(cfg.Cars)
	p.AddVariableDomains(discrete.IndexDomain(numCars))

	p.AddUniversalConstraint(fn.ConstraintAllUnique)
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Car sequencing constraint
		optionSequence := make([][]bool, numOptions)
		for i := range numOptions {
			optionSequence[i] = make([]bool, numCars)
		}
		for seqIdx, x := range fn.AsSequence(solution) {
			for optionIdx, flag := range cfg.CarOptions[cfg.Cars[x]] {
				optionSequence[optionIdx][seqIdx] = flag
			}
		}
		// Check each option's window and count number of cars
		// Ensure each window doesn't exceed max capacity for option
		for optionIdx, maxCount := range cfg.OptionMax {
			windowSize := cfg.OptionWindow[optionIdx]
			for i := range numCars {
				limit := min(numCars, i+windowSize) // clip window to end of cars list
				window := optionSequence[optionIdx][i:limit]
				if list.Count(window, true) > maxCount {
					return false
				}
			}
		}
		return true
	})

	p.SolutionCoreFn = fn.CoreMirroredSequence(cfg.Cars)
	p.SolutionStringFn = fn.StringSequence(cfg.Cars)
	return p
}
