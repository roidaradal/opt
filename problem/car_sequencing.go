package problem

import (
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Car Sequencing problem
func CarSequencing(n int) *discrete.Problem {
	name := newName(CAR_SEQUENCE, n)
	cfg := newCarSequencing(name)
	if cfg == nil {
		return nil
	}
	numCars, numOptions := len(cfg.cars), len(cfg.options)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy

	p.Variables = discrete.Variables(cfg.cars)
	domain := discrete.IndexDomain(numCars)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	// Car sequencing constraint
	test := func(solution *discrete.Solution) bool {
		sequence := fn.AsSequence(solution)
		optionSequence := make([][]bool, numOptions)
		for i := range numOptions {
			optionSequence[i] = make([]bool, numCars)
		}
		for seqIdx, x := range sequence {
			for optionIdx, flag := range cfg.carOptions[cfg.cars[x]] {
				optionSequence[optionIdx][seqIdx] = flag
			}
		}
		// Check each option's window and count the number of cars
		// using that option for each window, ensure doesn't exceed maxCount
		for optionIdx, optionCfg := range cfg.options {
			maxCount, windowSize := optionCfg[0], optionCfg[1]
			limit := (numCars / windowSize) * windowSize
			for i := 0; i < limit; i += windowSize {
				window := optionSequence[optionIdx][i : i+windowSize]
				if list.Count(window, true) > maxCount {
					return false
				}
			}
			// Check last window
			window := optionSequence[optionIdx][limit:]
			if list.Count(window, true) > maxCount {
				return false
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.SolutionCoreFn = fn.Core_MirroredSequence(cfg.cars)
	p.SolutionStringFn = fn.String_Sequence(cfg.cars)

	return p
}

type carSequenceCfg struct {
	options    [][2]int
	cars       []string
	carOptions map[string][]bool
}

// Load car sequencing test case
func newCarSequencing(name string) *carSequenceCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 3 {
		return nil
	}
	cfg := &carSequenceCfg{
		options:    make([][2]int, 0),
		cars:       make([]string, 0),
		carOptions: make(map[string][]bool),
	}
	counts := list.Map(strings.Fields(lines[0]), number.ParseInt)
	numOptions, numCars := counts[0], counts[1]
	idx := 1
	optionIdx := make(map[string]int)
	for i := range numOptions {
		parts := strings.Fields(lines[idx])
		name, maxCount, windowSize := parts[0], number.ParseInt(parts[1]), number.ParseInt(parts[2])
		optionIdx[name] = i
		cfg.options = append(cfg.options, [2]int{maxCount, windowSize})
		idx++
	}
	for range numCars {
		parts := strings.Fields(lines[idx])
		car, count := parts[0], number.ParseInt(parts[1])
		cfg.cars = append(cfg.cars, slices.Repeat([]string{car}, count)...)
		flags := make([]bool, numOptions)
		for _, name := range parts[2:] {
			flags[optionIdx[name]] = true
		}
		cfg.carOptions[car] = flags
		idx++
	}
	return cfg
}
