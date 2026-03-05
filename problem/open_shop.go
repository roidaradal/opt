package problem

import (
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewOpenShopScheduling creates a new Open Shop Scheduling problem
func NewOpenShopScheduling(variant string, n int) *discrete.Problem {
	name := newName(OpenShopScheduling, variant, n)
	switch variant {
	case "basic":
		return openShopScheduling(name)
	default:
		return nil
	}
}

// Open Shop Scheduling
func openShopScheduling(name string) *discrete.Problem {
	cfg := data.NewShopSchedule(name)
	if cfg == nil {
		return nil
	}
	tasks := cfg.GetTasks()

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(tasks)
	p.AddVariableDomains(discrete.RangeDomain(0, cfg.MaxMakespan))

	// Constraint: No job task overlap
	p.AddUniversalConstraint(fn.NoOverlap(tasks, cfg.Jobs, func(task data.Task) string {
		return task.Job
	}))

	// Constraint: No machine overlap
	p.AddUniversalConstraint(fn.ConstraintNoMachineOverlap(cfg, tasks))

	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScoreScheduleMakespan(tasks)
	p.SolutionStringFn = fn.StringShopSchedule(tasks, cfg.Machines)

	return p
}
