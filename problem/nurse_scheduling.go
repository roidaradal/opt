package problem

import (
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Nurse Schedule problem
func NurseSchedule(n int) *discrete.Problem {
	name := newName(NURSE_SCHED, n)
	cfg := newNurseSchedule(name)
	if cfg == nil {
		return nil
	}
	numDays, numShifts := len(cfg.days), len(cfg.shifts)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	slot := 0
	dayOf := make(map[int]int)   // slotIdx => dayIdx
	shiftOf := make(map[int]int) // slotIdx => shiftIdx
	allNurses := discrete.MapDomain(cfg.nurses)
	for day := range numDays {
		for shift := range numShifts {
			for range cfg.minPerShift[shift] {
				dayOf[slot] = day
				shiftOf[slot] = shift
				p.Variables = append(p.Variables, slot)
				p.Domain[slot] = allNurses[:]
				slot += 1
			}
		}
	}

	// Group assigned nurses per <day, shift>
	groupSlotSched := func(solution *discrete.Solution) map[[2]int]*ds.Set[int] {
		sched := make(map[[2]int]*ds.Set[int]) // <day,shift> = set of nurseIdx
		for day := range numDays {
			for shift := range numShifts {
				slotKey := [2]int{day, shift}
				sched[slotKey] = ds.NewSet[int]()
			}
		}
		for slot, nurse := range solution.Map {
			day, shift := dayOf[slot], shiftOf[slot]
			slotKey := [2]int{day, shift}
			sched[slotKey].Add(nurse)
		}
		return sched
	}

	// Hard constraint: Number of assigned nurses per shift (within min/max limit)
	test1 := func(solution *discrete.Solution) bool {
		sched := groupSlotSched(solution)
		for day := range numDays {
			for shift := range numShifts {
				slotKey := [2]int{day, shift}
				nurseCount := sched[slotKey].Len()
				if nurseCount < cfg.minPerShift[shift] || nurseCount > cfg.maxPerShift[shift] {
					return false
				}
			}
		}
		return true
	}
	p.AddUniversalConstraint(test1)

	// Hard constraint: Ensure all nurse shift count does not exceed MaxTotal limit
	test2 := func(solution *discrete.Solution) bool {
		shiftCount := make(dict.IntCounter) // nurse => count
		for _, nurse := range solution.Map {
			shiftCount[nurse] += 1
		}
		return list.AllLessEqual(dict.Values(shiftCount), cfg.maxTotal)
	}
	p.AddUniversalConstraint(test2)

	// Hard constraint: Ensure all <nurse,day> does not exceed MaxDaily limit
	test3 := func(solution *discrete.Solution) bool {
		dailySched := make(map[[2]int]int) // <nurse,day> => count
		for slot, nurse := range solution.Map {
			day := dayOf[slot]
			dailySched[[2]int{nurse, day}] += 1
		}
		return list.AllLessEqual(dict.Values(dailySched), cfg.maxDaily)
	}
	p.AddUniversalConstraint(test3)

	// Hard constraint: Ensure all nurse schedules follow MaxConsecutive limit
	test4 := func(solution *discrete.Solution) bool {
		nurseSlots := dict.GroupByValue(solution.Map) // nurse => list of slots
		return list.All(allNurses, func(nurse discrete.Value) bool {
			return constraint.MaxConsecutive(nurseSlots[nurse], cfg.maxConsecutive)
		})
	}
	p.AddUniversalConstraint(test4)

	// Soft constraint: balance scheds
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var penalty discrete.Score = 0
		for slot, nurse := range solution.Map {
			// Soft constraint: check nurse shift preference
			shift := cfg.shifts[shiftOf[slot]]
			if !slices.Contains(cfg.prefShifts[nurse], shift) {
				penalty += 1
			}

			// Soft constraint: check nurse day preference
			day := cfg.days[dayOf[slot]]
			if !slices.Contains(cfg.prefDays[nurse], day) {
				penalty += 1
			}
		}
		return penalty
	}

	nurseSched := func(solution *discrete.Solution) string {
		sched := groupSlotSched(solution)
		out := make([]string, 0)
		for day := range numDays {
			for shift := range numShifts {
				slotKey := [2]int{day, shift}
				nurses := list.MapList(sched[slotKey].Items(), cfg.nurses)
				slices.Sort(nurses)
				line := fmt.Sprintf("%s_%s = %s", cfg.days[day], cfg.shifts[shift], strings.Join(nurses, ", "))
				out = append(out, line)
			}
		}
		return strings.Join(out, " | ")
	}

	p.SolutionCoreFn = nurseSched
	p.SolutionStringFn = nurseSched

	return p
}

type nurseSchedCfg struct {
	nurses         []string
	days           []string
	shifts         []string
	minPerShift    []int
	maxPerShift    []int
	maxConsecutive int
	maxTotal       int
	maxDaily       int
	prefShifts     [][]string
	prefDays       [][]string
}

// Load nurse schedule test case
func newNurseSchedule(name string) *nurseSchedCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 7 {
		return nil
	}
	limits := list.Map(strings.Fields(lines[0]), number.ParseInt)
	cfg := &nurseSchedCfg{
		nurses:         strings.Fields(lines[1]),
		days:           strings.Fields(lines[2]),
		shifts:         strings.Fields(lines[3]),
		minPerShift:    list.Map(strings.Fields(lines[4]), number.ParseInt),
		maxPerShift:    list.Map(strings.Fields(lines[5]), number.ParseInt),
		maxConsecutive: limits[0],
		maxTotal:       limits[1],
		maxDaily:       limits[2],
	}
	numNurses := len(cfg.nurses)
	indexOf := list.IndexMap(cfg.nurses)
	cfg.prefShifts = make([][]string, numNurses)
	cfg.prefDays = make([][]string, numNurses)
	idx := 6
	for range numNurses {
		parts := str.CleanSplit(lines[idx], ":")
		nurse, body := parts[0], parts[1]
		nurseIdx := indexOf[nurse]
		prefs := str.CleanSplit(body, "|")
		cfg.prefShifts[nurseIdx] = strings.Fields(prefs[0])
		cfg.prefDays[nurseIdx] = strings.Fields(prefs[1])
		idx++
	}
	return cfg
}
