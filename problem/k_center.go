package problem

import (
	"math"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new K-Center problem
func KCenter(n int) *discrete.Problem {
	name := newName(K_CENTER, n)
	cfg := newKCenter(name)
	if cfg == nil {
		return nil
	}
	numCenters := len(cfg.centers)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.centers)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Get selected centers, ensure has correct count
		selected := fn.AsSubset(solution)
		return len(selected) == cfg.count
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Compute the maximum distance of any center to selected centers
		selected := fn.AsSubset(solution)
		var maxDistance float64 = 0
		for x := range numCenters {
			var minDistance float64 = math.Inf(1)
			for _, selectedCenter := range selected {
				minDistance = min(minDistance, cfg.distance[x][selectedCenter])
			}
			maxDistance = max(maxDistance, minDistance)
		}
		return discrete.Score(maxDistance)
	}

	p.SolutionStringFn = fn.String_Subset(cfg.centers)

	return p
}

type kCenterCfg struct {
	count    int
	centers  []string
	distance [][]float64
}

// Load k-center test case
func newKCenter(name string) *kCenterCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 3 {
		return nil
	}
	cfg := &kCenterCfg{
		count:    number.ParseInt(lines[0]),
		centers:  strings.Fields(lines[1]),
		distance: make([][]float64, 0),
	}
	idx := 2
	for range len(cfg.centers) {
		d := list.Map(strings.Fields(lines[idx]), number.ParseFloat)
		cfg.distance = append(cfg.distance, d)
		idx++
	}
	return cfg
}
