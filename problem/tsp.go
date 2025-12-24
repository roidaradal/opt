package problem

import (
	"math"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
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

	p.SolutionCoreFn = func(solution *discrete.Solution) string {
		sequence := list.Map(fn.AsSequence(solution), func(x discrete.Variable) string {
			return str.Any(cfg.cities[x])
		})
		// Find the first element alphabetically, rearrange the sequence so it is the first
		index := slices.Index(sequence, slices.Min(sequence))
		sequence2 := append([]string{}, sequence[index:]...)
		sequence2 = append(sequence2, sequence[:index]...)
		return strings.Join(sequence2, " ")
	}

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
		d := list.Map(strings.Fields(line), func(x string) float64 {
			if x == "x" {
				return math.Inf(1)
			} else {
				return number.ParseFloat(x)
			}
		})
		cfg.distance = append(cfg.distance, d)
	}
	return cfg
}
