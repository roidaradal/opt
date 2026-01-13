package problem

import (
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Bottleneck Traveling Salesman problem
func BottleneckTravelingSalesman(n int) *discrete.Problem {
	name := newName(BOTTLENECK_TSP, n)
	cfg := fn.NewTravelingSalesman(name)
	if cfg == nil {
		return nil
	}
	numCities := len(cfg.Cities)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(cfg.Cities)
	domain := discrete.IndexDomain(numCities)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// AllUnique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Treat solution as sequence, add first variable to end of sequence (loop)
		sequence := fn.AsSequence(solution)
		sequence = append(sequence, sequence[0])
		// Compute the max distance in the path (bottleneck)
		var maxDistance discrete.Score = -a.Inf
		for i := range numCities {
			curr, next := sequence[i], sequence[i+1]
			maxDistance = max(maxDistance, cfg.Distance[curr][next])
		}
		return maxDistance
	}

	p.SolutionCoreFn = fn.Core_SortedCycle(cfg.Cities)
	p.SolutionStringFn = fn.String_Sequence(cfg.Cities)

	return p
}
