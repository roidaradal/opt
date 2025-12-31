package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Traveling Salesman problem
func TravelingSalesman(n int) *discrete.Problem {
	name := newName(TSP, n)
	cfg := newTravelingSalesman(name)
	if cfg == nil {
		return nil
	}
	numCities := len(cfg.cities)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(cfg.cities)
	domain := discrete.IndexDomain(numCities)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// AllUnique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Treat the solution as a sequence
		// Add first variable to end of sequence (loop)
		sequence := fn.AsSequence(solution)
		sequence = append(sequence, sequence[0])
		// Compute total distance between succeeding cities in the sequence
		var totalDistance discrete.Score = 0
		for i := range numCities {
			curr, next := sequence[i], sequence[i+1]
			totalDistance += cfg.distance[curr][next]
		}
		return totalDistance
	}

	p.SolutionCoreFn = fn.Core_SortedCycle(cfg.cities)
	p.SolutionStringFn = fn.String_Sequence(cfg.cities)

	return p
}

type tspCfg struct {
	cities   []string
	distance [][]float64
}

// Load traveling salesman test case
func newTravelingSalesman(name string) *tspCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	cfg := &tspCfg{
		cities:   strings.Fields(lines[0]),
		distance: make([][]float64, 0),
	}
	for _, line := range lines[1:] {
		d := list.Map(strings.Fields(line), fn.ParseFloatInf)
		cfg.distance = append(cfg.distance, d)
	}
	return cfg
}
