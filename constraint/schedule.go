package constraint

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/discrete"
)

// Constraint: No Machine Overlap
func NoMachineOverlap(cfg *a.ShopSchedCfg) discrete.ConstraintFn {
	return noOverlap(cfg, cfg.Machines, func(task *a.Task) string {
		return task.Machine
	})
}

// Constraint: No Job Task overlap
func NoJobTaskOverlap(cfg *a.ShopSchedCfg) discrete.ConstraintFn {
	keys := list.Map(cfg.Jobs, func(job *a.Job) string {
		return job.Name
	})
	return noOverlap(cfg, keys, func(task *a.Task) string {
		return task.JobName
	})
}

// Utility: check if no overlap
func noOverlap(cfg *a.ShopSchedCfg, keys []string, keyFn func(*a.Task) string) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Initialize all keys' schedules
		groupSched := make(map[string][]a.TimeRange)
		for _, key := range keys {
			groupSched[key] = make([]a.TimeRange, 0)
		}
		// For each task in the solution, add the schedule TimeRange to its group schedule
		for x, start := range solution.Map {
			task := cfg.Tasks[x]
			sched := a.TimeRange{start, start + task.Duration}
			key := keyFn(task)
			groupSched[key] = append(groupSched[key], sched)
		}
		// Sort the schedules for each group, and check if there is overlap
		for _, scheds := range groupSched {
			slices.SortFunc(scheds, a.SortByStartTime)
			for i := range len(scheds) - 1 {
				curr, next := scheds[i], scheds[i+1]
				start1, end1 := curr.Tuple()
				start2 := next[0]
				if start2 <= start1 || start2 < end1 {
					return false
				}
			}
		}
		return true
	}
}
