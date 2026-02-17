package problem

import (
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewNurseScheduling creates a new Nurse Scheduling problem
func NewNurseScheduling(variant string, n int) *discrete.Problem {
	name := newName(NurseScheduling, variant, n)
	switch variant {
	case "basic":
		return nurseScheduling(name)
	default:
		return nil
	}
}

// Nurse Scheduling
func nurseScheduling(name string) *discrete.Problem {
	cfg := data.NewNurseSchedule(name)
	if cfg == nil {
		return nil
	}
	numDays, numShifts := len(cfg.Days), len(cfg.Shifts)

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	slotIdx := 0
	dayOf := make(map[int]int)   // slotIdx => dayIdx
	shiftOf := make(map[int]int) // slotIdx => shiftIdx
	for dayIdx := range numDays {
		for shiftIdx := range numShifts {
			for range cfg.ShiftMin[shiftIdx] {
				dayOf[slotIdx] = dayIdx
				shiftOf[slotIdx] = shiftIdx
				slotIdx += 1
			}
		}
	}
	p.Variables = discrete.IndexDomain(slotIdx)
	domain := discrete.Domain(cfg.Nurses)
	p.AddVariableDomains(domain)

	// Function to group assigned nurses per (day, shift)
	groupSlotSched := func(solution *discrete.Solution) map[[2]int]*ds.Set[int] {
		sched := make(map[[2]int]*ds.Set[int]) // (day, shift) => set of nurseIdx
		for dayIdx := range numDays {
			for shiftIdx := range numShifts {
				key := [2]int{dayIdx, shiftIdx}
				sched[key] = ds.NewSet[int]()
			}
		}
		for slot, nurse := range solution.Map {
			dayIdx, shiftIdx := dayOf[slot], shiftOf[slot]
			key := [2]int{dayIdx, shiftIdx}
			sched[key].Add(nurse)
		}
		return sched
	}

	// Hard constraint: number of assigned nurses per shift within min/max limit
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		sched := groupSlotSched(solution)
		for dayIdx := range numDays {
			for shiftIdx := range numShifts {
				key := [2]int{dayIdx, shiftIdx}
				nurseCount := sched[key].Len()
				if nurseCount < cfg.ShiftMin[shiftIdx] || nurseCount > cfg.ShiftMax[shiftIdx] {
					return false
				}
			}
		}
		return true
	})

	// Hard constraint: ensure all nurse shift count does not exceed MaxTotal limit
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		shiftCount := make(dict.IntCounter) // nurse => count
		for _, nurse := range solution.Map {
			shiftCount[nurse] += 1
		}
		return list.AllLessEqual(dict.Values(shiftCount), cfg.MaxTotalShift)
	})

	// Hard constraint: ensure all (nurse,day) does not exceed MaxDaily limit
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		dailySched := make(map[[2]int]int) // (nurse, day) => count
		for slot, nurse := range solution.Map {
			dayIdx := dayOf[slot]
			dailySched[[2]int{nurse, dayIdx}] += 1
		}
		return list.AllLessEqual(dict.Values(dailySched), cfg.DailyLimit)
	})

	// Hard constraint: ensure all nurse schedule follow MaxConsecutive limit
	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		nurseSlots := dict.GroupByValue(solution.Map)
		return list.All(domain, func(nurse discrete.Value) bool {
			return fn.MaxConsecutive(nurseSlots[nurse], cfg.MaxConsecutiveShift)
		})
	})

	// Soft constraint: minimize penalty by balancing schedules
	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var penalty discrete.Score = 0
		for slot, nurseIdx := range solution.Map {
			nurse := cfg.Nurses[nurseIdx]
			// Soft constraint: check nurse shift preference
			shift := cfg.Shifts[shiftOf[slot]]
			if !slices.Contains(cfg.PreferShifts[nurse], shift) {
				penalty += 1
			}
			// Soft constraint: check nurse day preference
			day := cfg.Days[dayOf[slot]]
			if !slices.Contains(cfg.PreferDays[nurse], day) {
				penalty += 1
			}
		}
		return penalty
	}

	// CoreFn and StringFn
	nurseSched := func(solution *discrete.Solution) string {
		sched := groupSlotSched(solution)
		out := str.NewBuilder()
		for dayIdx := range numDays {
			for shiftIdx := range numShifts {
				key := [2]int{dayIdx, shiftIdx}
				nurses := list.MapList(sched[key].Items(), cfg.Nurses)
				slices.Sort(nurses)
				line := fmt.Sprintf("%s_%s = %s", cfg.Days[dayIdx], cfg.Shifts[shiftIdx], strings.Join(nurses, ", "))
				out.Add(line)
			}
		}
		return out.Build(" | ")
	}
	p.SolutionCoreFn = nurseSched
	p.SolutionStringFn = nurseSched

	return p
}
