package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewJobShopScheduling creates a new Job Shop Scheduling problem
func NewJobShopScheduling(variant string, n int) *discrete.Problem {
	name := newName(JobShopScheduling, variant, n)
	switch variant {
	case "basic":
		return jobShopScheduling(name)
	default:
		return nil
	}
}

// Job Shop Scheduling
func jobShopScheduling(name string) *discrete.Problem {
	cfg := data.NewShopSchedule(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	variable := 0
	jobTasks := make(map[int][]discrete.Variable)
	tasks := cfg.GetTasks()
	for jobID, job := range cfg.Jobs {
		jobTasks[jobID] = make([]discrete.Variable, 0)
		for taskID, task := range cfg.JobTasks[job] {
			first := list.SumOf(cfg.JobTasks[job][:taskID], data.Task.GetDuration)
			after := list.SumOf(cfg.JobTasks[job][taskID+1:], data.Task.GetDuration)
			last := cfg.MaxMakespan - after - task.Duration

			p.Variables = append(p.Variables, variable)
			p.Domain[variable] = discrete.RangeDomain(first, last)
			jobTasks[jobID] = append(jobTasks[jobID], variable)
			variable += 1
		}
	}

	// Constraint: job tasks in order and no overlap
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		for _, variables := range jobTasks {
			for i := range len(variables) - 1 {
				curr, next := variables[i], variables[i+1]
				start1, start2 := solution.Map[curr], solution.Map[next]
				end1 := start1 + tasks[curr].Duration
				// Not in order or has overlap
				if start2 <= start1 || start2 < end1 {
					return false
				}
			}
		}
		return true
	})

	// Constraint: no machine overlap
	p.AddUniversalConstraint(fn.ConstraintNoMachineOverlap(cfg, tasks))

	p.Goal = discrete.Minimize
	p.ObjectiveFn = fn.ScoreScheduleMakespan(tasks)
	p.SolutionStringFn = fn.StringShopSchedule(tasks, cfg.Machines)

	return p
}
