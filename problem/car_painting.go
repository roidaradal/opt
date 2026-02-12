package problem

import (
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewCarPainting creates a new Car Paiting problem
func NewCarPainting(variant string, n int) *discrete.Problem {
	name := newName(CarPainting, variant, n)
	switch variant {
	case "basic":
		return carPainting(name)
	case "binary":
		return binaryPaintShop(name)
	default:
		return nil
	}
}

// Car Painting
func carPainting(name string) *discrete.Problem {
	cfg := data.NewCars(name)
	if cfg == nil {
		return nil
	}
	numCars := len(cfg.CarColors)

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.CarColors)
	minLimit, maxLimit := 0, numCars-1
	for idx, variable := range p.Variables {
		// Setup domain of each variable, max wiggle to left/right is MaxShift
		first := max(minLimit, idx-cfg.MaxShift)
		last := min(maxLimit, idx+cfg.MaxShift)
		p.Domain[variable] = discrete.RangeDomain(first, last)
	}

	// AllUnique constraint
	p.AddUniversalConstraint(fn.ConstraintAllUnique)

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Create the color sequence from the solution
		colorSequence := make([]string, numCars)
		for i, x := range fn.AsSequence(solution) {
			colorSequence[i] = cfg.CarColors[x]
		}
		// Count number of color changes
		return discrete.Score(fn.CountColorChanges(colorSequence))
	}

	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		// Go through sequence of colors, group same color subsequences
		var prevColor string
		groups := make([][]int, 0)
		group := make([]int, 0)
		for i, x := range fn.AsSequence(solution) {
			currColor := cfg.CarColors[x]
			if i > 0 && prevColor != currColor {
				// On color change, end curent group and start new group
				groups = append(groups, group)
				group = []int{x}
			} else {
				// Same color = add to current group
				group = append(group, x)
			}
			prevColor = currColor
		}
		groups = append(groups, group) // Add last group
		// For each group, sort the items
		output := list.Map(groups, func(group []int) string {
			slices.Sort(group)
			return strings.Join(list.Map(group, str.Int), " ")
		})
		return strings.Join(output, "|")
	}

	p.SolutionStringFn = func(solution *discrete.Solution) string {
		// Go through sequence of colors, display car and color
		// Detect color changes from previous to current color: add separator |
		var prevColor string
		output := make([]string, 0)
		for i, x := range fn.AsSequence(solution) {
			currColor := cfg.CarColors[x]
			if i > 0 && prevColor != currColor {
				output = append(output, "|")
			}
			output = append(output, fmt.Sprintf("%d:%s", x, currColor))
			prevColor = currColor
		}

		return strings.Join(output, " ")
	}

	return p
}

// Binary Paint Shop
func binaryPaintShop(name string) *discrete.Problem {
	cfg := data.NewCars(name)
	if cfg == nil {
		return nil
	}
	indexOf := list.IndexMap(cfg.Cars)

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Cars)
	p.AddVariableDomains(discrete.BooleanDomain())

	// Set first car to 0
	car0 := p.Variables[0]
	p.Domain[car0] = []discrete.Value{0}

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Initialize current colors of cars from solution
		color := make([]int, len(cfg.Cars))
		for x, c := range solution.Map {
			color[x] = c
		}
		colorSequence := make([]int, len(cfg.Sequence))
		for i, car := range cfg.Sequence {
			x := indexOf[car]
			colorSequence[i] = color[x]   // add car's current color to sequence
			color[x] = (color[x] + 1) % 2 // flip car's color, since two cars of same kind must have different colors
		}
		// Count color changes
		return discrete.Score(fn.CountColorChanges(colorSequence))
	}

	p.SolutionStringFn = fn.StringValues[int](p, nil)
	return p
}
