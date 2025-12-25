package problem

import (
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create Flow Shop Schedule problem
func FlowShopSchedule(n int) *discrete.Problem {
	name := newName(FLOWSHOP_SCHED, n)
	cfg := newFlowShop(name)
	if cfg == nil {
		return nil
	}
	numJobs, numMachines := len(cfg.Jobs), len(cfg.Machines)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Sequence

	p.Variables = discrete.Variables(cfg.Jobs)
	domain := discrete.IndexDomain(numJobs)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// All Unique constraint
	p.AddUniversalConstraint(constraint.AllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		sequence := fn.AsSequence(solution)
		end := make(map[ds.Coords]int)
		for m := range numMachines {
			for i, x := range sequence {
				task := cfg.Jobs[x].Tasks[m]
				above := end[ds.Coords{m - 1, i}]
				prev := end[ds.Coords{m, i - 1}]
				// Pick the later ending: above or prev
				// above: previous task on same job: can only process one job task at a time
				// prev:  previous task on same machine: machine can only process one task at a time
				start := max(above, prev)
				end[ds.Coords{m, i}] = start + task.Duration
			}
		}
		return discrete.Score(end[ds.Coords{numMachines - 1, numJobs - 1}])
	}

	p.SolutionStringFn = fn.String_Sequence(cfg.Jobs)

	return p
}

// Load flow shop schedule test case
func newFlowShop(name string) *a.ShopSchedCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	cfg := &a.ShopSchedCfg{
		Machines: strings.Fields(lines[0]),
		Jobs:     make([]*a.Job, 0),
		Tasks:    make([]*a.Task, 0),
	}
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
	}
	return cfg
}
