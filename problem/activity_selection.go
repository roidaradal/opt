package problem

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Activity Selection problem
func ActivitySelection(n int) *discrete.Problem {
	name := newName(ACTIVITY_SELECTION, n)
	cfg := newActivitySelection(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(cfg.activities)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
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
		for i := 1; i < numSelected; i++ {
			prev, curr := selected[i-1], selected[i]
			// Conflict: prev activity ends after start of curr activity
			if end[prev] > start[curr] {
				return false
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(cfg.activities)

	return p
}

type activitySelectionCfg struct {
	activities []string
	start      []float64
	end        []float64
}

// Load activity selection test case
func newActivitySelection(name string) *activitySelectionCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &activitySelectionCfg{
		activities: strings.Fields(lines[0]),
		start:      list.Map(strings.Fields(lines[1]), number.ParseFloat),
		end:        list.Map(strings.Fields(lines[2]), number.ParseFloat),
	}
}
