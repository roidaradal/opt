package problem

import (
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewTravelingSalesman creates a new Traveling Salesman problem
func NewTravelingSalesman(variant string, n int) *discrete.Problem {
	name := newName(TravelingSalesman, variant, n)
	switch variant {
	case "basic":
		return travelingSalesman(name)
	case "bottleneck":
		return bottleneckTravelingSalesman(name)
	default:
		return nil
	}
}

// Common steps for creating a Traveling Salesman problem
func newTravelingSalesmanProblem(name string) (*discrete.Problem, *data.GraphPath) {
	cfg := data.NewGraphTour(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(cfg.Vertices)
	p.AddVariableDomains(discrete.IndexDomain(len(cfg.Vertices)))

	p.AddUniversalConstraint(fn.ConstraintAllUnique)

	p.SolutionCoreFn = fn.CoreSortedCycle(cfg.Vertices)
	p.SolutionStringFn = fn.StringSequence(cfg.Vertices)
	return p, cfg
}

// Traveling Salesman
func travelingSalesman(name string) *discrete.Problem {
	p, cfg := newTravelingSalesmanProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Treat solution as sequence, add first variable to end to form loop
		sequence := fn.AsSequence(solution)
		sequence = append(sequence, sequence[0])
		// Compute total distance between succeeding cities in the sequence
		var totalDistance discrete.Score = 0
		for i := range len(cfg.Vertices) {
			curr, next := sequence[i], sequence[i+1]
			totalDistance += cfg.Distance[curr][next]
		}
		return totalDistance
	}
	return p
}

// Bottleneck Traveling Salesman
func bottleneckTravelingSalesman(name string) *discrete.Problem {
	p, cfg := newTravelingSalesmanProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Treat solution as sequence, add first variable to end to form loop
		sequence := fn.AsSequence(solution)
		sequence = append(sequence, sequence[0])
		// Compute max distance in the path (bottleneck)
		var maxDistance = -discrete.Inf
		for i := range len(cfg.Vertices) {
			curr, next := sequence[i], sequence[i+1]
			maxDistance = max(maxDistance, cfg.Distance[curr][next])
		}
		return maxDistance
	}
	return p
}
