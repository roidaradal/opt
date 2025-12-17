package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Binary Paintshop problem
func BinaryPaintShop(n int) *discrete.Problem {
	name := newName(BINARY_PAINTSHOP, n)
	cfg := newBinaryPaintShop(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.RangeVariables(0, cfg.numCars-1)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// Set first car to 0
	car0 := p.Variables[0]
	p.Domain[car0] = []discrete.Value{0}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Initialize current colors of cars from solution
		color := make([]int, cfg.numCars)
		for x, c := range solution.Map {
			color[x] = c
		}
		colorSequence := make([]int, len(cfg.sequence))
		for i, x := range cfg.sequence {
			colorSequence[i] = color[x]   // add car x's current color to the sequence
			color[x] = (color[x] + 1) % 2 // flip car x's color to the other, since we have to paint the 2 car x's a different color
		}
		return discrete.Score(fn.CountColorChanges(colorSequence))
	}

	p.SolutionStringFn = fn.String_Values[int](p, nil)

	return p
}

type binaryPaintCfg struct {
	numCars  int
	sequence []int
}

// Load binary paintshop test case
func newBinaryPaintShop(name string) *binaryPaintCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 2 {
		return nil
	}
	return &binaryPaintCfg{
		numCars:  number.ParseInt(lines[0]),
		sequence: list.Map(strings.Fields(lines[1]), number.ParseInt),
	}
}
