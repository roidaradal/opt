package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewFlowShopScheduling creates a new Flow Shop Scheduling problem
func NewFlowShopScheduling(variant string, n int) *discrete.Problem {
	name := newName(FlowShopScheduling, variant, n)
	switch variant {
	case "basic":
		return flowShopScheduling(name)
	default:
		return nil
	}
}

// Flow Shop Scheduling
func flowShopScheduling(name string) *discrete.Problem {
	cfg := data.NewShopSchedule(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(cfg.Jobs)
	p.AddVariableDomains(discrete.IndexDomain(len(cfg.Jobs)))
	p.AddUniversalConstraint(fn.ConstraintAllUnique)

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		sequence := fn.AsSequence(solution)
		end := make(map[ds.Coords]int)
		for m := range len(cfg.Machines) {
			for i, x := range sequence {
				above := end[ds.Coords{m - 1, i}]
				prev := end[ds.Coords{m, i - 1}]
				// Pick later ending: above or prev
				// above: previous task on same job: can only process one job task at a time
				// prev: previous task on same machine: machine can only process one task at time
				start := max(above, prev)
				end[ds.Coords{m, i}] = start + cfg.TaskTimes[cfg.Jobs[x]][m]
			}
		}
		lastRow, lastCol := len(cfg.Machines)-1, len(cfg.Jobs)-1
		return discrete.Score(end[ds.Coords{lastRow, lastCol}])
	}

	p.SolutionStringFn = fn.StringSequence(cfg.Jobs)
	return p
}
