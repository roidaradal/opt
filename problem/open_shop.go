package problem

import (
	"strings"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Open Shop Schedule problem
func OpenShopSchedule(n int) *discrete.Problem {
	name := newName(OPENSHOP_SCHED, n)
	cfg := newOpenShop(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Tasks)
	domain := discrete.RangeDomain(0, cfg.MaxMakespan)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// Constraint: no job task overlap
	p.AddUniversalConstraint(constraint.NoJobTaskOverlap(cfg))

	// Constraint: no machine overlap
	p.AddUniversalConstraint(constraint.NoMachineOverlap(cfg))

	p.ObjectiveFn = fn.Score_ScheduleMakespan(cfg.Tasks)
	p.SolutionStringFn = fn.String_ShopSchedule(cfg.Tasks, cfg.Machines)

	return p
}

// Load open shop schedule test case
func newOpenShop(name string) *a.ShopSchedCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	cfg := &a.ShopSchedCfg{
		Machines: strings.Fields(lines[0]),
		Jobs:     make([]*a.Job, 0),
		Tasks:    make([]*a.Task, 0),
	}
	totalDuration := 0
	for _, line := range lines[1:] {
		parts := str.CleanSplit(line, "=")
		if len(parts) != 2 {
			continue
		}
		jobName := parts[0]
		job := &a.Job{
			Name:  jobName,
			Tasks: make([]*a.Task, 0),
		}
		for taskIndex, duration := range strings.Fields(parts[1]) {
			taskName := a.TaskString(cfg.Machines[taskIndex], duration)
			job.Tasks = append(job.Tasks, a.NewTask(taskName, jobName, taskIndex))
		}
		cfg.Jobs = append(cfg.Jobs, job)
		cfg.Tasks = append(cfg.Tasks, job.Tasks...)
		totalDuration += job.TotalDuration()
	}
	cfg.MaxMakespan = totalDuration
	return cfg
}
