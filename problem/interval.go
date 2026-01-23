package problem

import (
	"cmp"
	"slices"

	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Interval creates a new Interval problem
func Interval(variant string, n int) *discrete.Problem {
	name := newName(INTERVAL, variant, n)
	switch variant {
	case "unweighted":
		return intervalUnweighted(name)
	case "weighted":
		return intervalWeighted(name)
	default:
		return nil
	}
}

// Common create steps of Activity Selection problem
func intervalProblem(name string, cfg *intervalCfg) *discrete.Problem {
	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.activities)
	p.AddVariableDomains(discrete.BooleanDomain())
	p.AddUniversalConstraint(intervalTest(cfg))

	p.Goal = discrete.Maximize
	p.SolutionStringFn = fn.StringSubset(cfg.activities)
	return p
}

// Unweighted Activity Selection problem
func intervalUnweighted(name string) *discrete.Problem {
	cfg := newUnweightedInterval(name)
	if cfg == nil {
		return nil
	}

	p := intervalProblem(name, cfg)
	p.ObjectiveFn = fn.ScoreSubsetSize
	return p
}

// Weighted Activity Selection problem
func intervalWeighted(name string) *discrete.Problem {
	cfg := newWeightedInterval(name)
	if cfg == nil {
		return nil
	}

	p := intervalProblem(name, cfg)
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.weight)
	return p
}

// Interval (Activity Selection) constraint
func intervalTest(cfg *intervalCfg) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
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
	}
}

type intervalCfg struct {
	activities []string
	start      []float64
	end        []float64
	weight     []float64
}

// Load 'interval.unweighted' test case
func newUnweightedInterval(name string) *intervalCfg {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &intervalCfg{
		activities: fn.StringList(lines[0]),
		start:      fn.FloatList(lines[1]),
		end:        fn.FloatList(lines[2]),
	}
}

// Load 'interval.weighted' test case
func newWeightedInterval(name string) *intervalCfg {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != 4 {
		return nil
	}
	return &intervalCfg{
		activities: fn.StringList(lines[0]),
		start:      fn.FloatList(lines[1]),
		end:        fn.FloatList(lines[2]),
		weight:     fn.FloatList(lines[3]),
	}
}
