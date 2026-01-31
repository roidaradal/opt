package problem

import (
	"cmp"
	"slices"

	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewInterval creates a new Interval problem
func NewInterval(variant string, n int) *discrete.Problem {
	name := newName(Interval, variant, n)
	switch variant {
	case "basic":
		return activitySelection(name)
	case "weighted":
		return weightedActivitySelection(name)
	default:
		return nil
	}
}

// Common steps for creating Activity Selection problem
func newActivitySelectionProblem(name string) (*discrete.Problem, *data.Intervals) {
	cfg := data.NewIntervals(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.Activities)
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
		start, end := cfg.Start, cfg.End
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
	p.SolutionStringFn = fn.StringSubset(cfg.Activities)
	return p, cfg
}

// Activity Selection
func activitySelection(name string) *discrete.Problem {
	p, _ := newActivitySelectionProblem(name)
	if p == nil {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSubsetSize
	return p
}

// Weighted Activity Selection
func weightedActivitySelection(name string) *discrete.Problem {
	p, cfg := newActivitySelectionProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	if len(cfg.Activities) != len(cfg.Weight) {
		return nil
	}
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.Weight)
	return p
}
