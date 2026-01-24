package problem

import (
	"cmp"
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Interval creates a new Interval problem
func Interval(variant string, n int) *discrete.Problem {
	name := newName(INTERVAL, variant, n)
	switch variant {
	case "basic":
		return intervalBasic(name)
	case "weight":
		return intervalWeighted(name)
	default:
		return nil
	}
}

// Common create steps of Activity Selection problem
func intervalProblem(name string) (*discrete.Problem, *intervalCfg) {
	cfg := newInterval(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.activities)
	p.AddVariableDomains(discrete.BooleanDomain())

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Get selected activities
		selected := fn.AsSubset(solution)
		numSelected := len(selected)
		if numSelected <= 1 {
			// no conflict for 0 or 1 activities
			return true
		}
		// Sort selected activities by start time
		start, end := cfg.start, cfg.end
		slices.SortFunc(selected, func(x1, x2 int) int {
			return cmp.Compare(start[x1], start[x2])
		})
		// Check for activity overlaps
		for i := range numSelected - 1 {
			curr, next := selected[i], selected[i+1]
			// Conflict: curr activity ends after start of next activity
			if end[curr] > start[next] {
				return false
			}
		}
		return true
	})

	p.Goal = discrete.Maximize
	p.SolutionStringFn = fn.StringSubset(cfg.activities)
	return p, cfg
}

// Activity Selection problem
func intervalBasic(name string) *discrete.Problem {
	p, _ := intervalProblem(name)
	if p == nil {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSubsetSize
	return p
}

// Weighted Activity Selection problem
func intervalWeighted(name string) *discrete.Problem {
	p, cfg := intervalProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.weight)
	return p
}

type intervalCfg struct {
	activities []string
	start      []float64
	end        []float64
	weight     []float64
}

// Load interval test case (basic, weight)
func newInterval(name string) *intervalCfg {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) < 3 {
		return nil
	}
	cfg := &intervalCfg{
		activities: fn.StringList(lines[0]),
		start:      fn.FloatList(lines[1]),
		end:        fn.FloatList(lines[2]),
	}
	if len(lines) == 4 {
		cfg.weight = fn.FloatList(lines[3])
	} else {
		cfg.weight = list.Repeated(1.0, len(cfg.activities))
	}
	return cfg
}
