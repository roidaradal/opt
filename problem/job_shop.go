package problem

import (
	"strings"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Job Shop Schedule problem
func JobShopSchedule(n int) *discrete.Problem {
	name := newName(JOBSHOP_SCHED, n)
	cfg := newJobShop(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	variableID := 0
	jobTasks := make(map[int][]discrete.Variable)
	for jobID, job := range cfg.Jobs {
		jobTasks[jobID] = make([]discrete.Variable, 0)
		for taskID, task := range job.Tasks {
			// Get before and after margins of task
			first, after := job.TaskMargins(taskID)
			last := cfg.MaxMakespan - after - task.Duration

			variable := discrete.Variable(variableID)
			p.Variables = append(p.Variables, variable)
			p.Domain[variable] = discrete.RangeDomain(first, last)
			jobTasks[jobID] = append(jobTasks[jobID], variable)
			variableID++
		}
	}

	// Constraint: job tasks in order and no overlap
	test := func(solution *discrete.Solution) bool {
		for _, variables := range jobTasks {
			for i := range len(variables) - 1 {
				curr, next := variables[i], variables[i+1]
				start1, start2 := solution.Map[curr], solution.Map[next]
				end1 := start1 + cfg.Tasks[curr].Duration
				// Not in order or has overlap
				if start2 <= start1 || start2 < end1 {
					return false
				}
			}
		}
		return true
	}
	p.AddUniversalConstraint(test)

	// Constraint: no machine overlap
	p.AddUniversalConstraint(constraint.NoMachineOverlap(cfg))

	p.ObjectiveFn = fn.Score_ScheduleMakespan(cfg.Tasks)
	p.SolutionStringFn = fn.String_ShopSchedule(cfg.Tasks, cfg.Machines)

	return p
}

// Load job shop schedule test case
func newJobShop(name string) *a.ShopSchedCfg {
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
		job := a.NewJob(parts[0], parts[1])
		cfg.Jobs = append(cfg.Jobs, job)
		cfg.Tasks = append(cfg.Tasks, job.Tasks...)
		totalDuration += job.TotalDuration()
	}
	cfg.MaxMakespan = totalDuration
	return cfg
}
