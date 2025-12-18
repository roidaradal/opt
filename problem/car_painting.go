package problem

import (
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Car Paint problem
func CarPainting(n int) *discrete.Problem {
	name := newName(CAR_PAINT, n)
	cfg := newCarPainting(name)
	if cfg == nil {
		return nil
	}
	numCars := len(cfg.cars)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(cfg.cars)
	minLimit, maxLimit := 0, numCars-1
	for _, variable := range p.Variables {
		first := max(minLimit, variable-cfg.maxShift)
		last := min(maxLimit, variable+cfg.maxShift)
		p.Domain[variable] = discrete.RangeDomain(first, last)
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// From the sequence formed by the solution, count the color changes
		colorSequence := make([]string, numCars)
		for i, x := range fn.AsSequence(solution) {
			colorSequence[i] = cfg.cars[x]
		}
		return discrete.Score(fn.CountColorChanges(colorSequence))
	}

	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		// From the sequence formed by the solution, get the color and the car pair
		sequence := make([]a.StrInt, numCars)
		for i, x := range fn.AsSequence(solution) {
			sequence[i] = a.NewStrInt(cfg.cars[x], x)
		}
		// Go through sequence of car-color pairs, and group same color subsequences
		var prevColor string
		groups := make([][]int, 0)
		group := make([]int, 0)
		for i, item := range sequence {
			currColor := item.Str
			if i > 0 && prevColor != currColor {
				// On color change, end current group and start new group
				groups = append(groups, group)
				group = []int{item.Int}
			} else {
				// Same color = add to current group
				group = append(group, item.Int)
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
		// From the sequence formed by the solution, get the color and the car pair
		sequence := make([]a.StrInt, numCars)
		for i, x := range fn.AsSequence(solution) {
			sequence[i] = a.NewStrInt(cfg.cars[x], x)
		}
		// Go through the sequence of car-color pairs, and display the car and their colors
		// Detect color changes from previous and current color and add separator |
		var prevColor string
		output := make([]string, 0)
		for i, item := range sequence {
			currColor := item.Str
			if i > 0 && prevColor != currColor {
				output = append(output, "|")
			}
			output = append(output, fmt.Sprintf("%d:%s", item.Int, currColor))
			prevColor = currColor
		}
		return strings.Join(output, " ")
	}

	return p
}

type carPaintCfg struct {
	maxShift int
	cars     []string
}

// Load car painting test case
func newCarPainting(name string) *carPaintCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return &carPaintCfg{
		maxShift: number.ParseInt(lines[0]),
		cars:     strings.Fields(lines[1]),
	}
}
