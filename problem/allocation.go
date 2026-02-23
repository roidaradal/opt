package problem

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewAllocation creates a new Allocation problem
func NewAllocation(variant string, n int) *discrete.Problem {
	name := newName(Allocation, variant, n)
	switch variant {
	case "resource":
		return resourceAllocation(name)
	case "scene":
		return sceneAllocation(name)
	case "fair_item":
		return fairItemAllocation(name)
	case "house":
		return houseAllocation(name)
	default:
		return nil
	}
}

// Resource Allocation
func resourceAllocation(name string) *discrete.Problem {
	cfg := data.NewResource(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Items)
	for i, variable := range p.Variables {
		p.Domain[variable] = discrete.RangeDomain(0, cfg.Count[i])
	}

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Check sum of weighted costs don't exceed budget
		count := solution.Map
		costs := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * cfg.Cost[x]
		})
		return list.Sum(costs) <= cfg.Budget
	})

	p.Goal = discrete.Maximize
	p.ObjectiveFn = fn.ScoreSumWeightedValues(p.Variables, cfg.Value)
	p.SolutionStringFn = fn.StringValues[int](p, nil)
	return p
}

// Scene Allocation
func sceneAllocation(name string) *discrete.Problem {
	cfg := data.NewScene(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.Scenes)
	domain := discrete.IndexDomain(cfg.NumDays)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		// Group scene distribution to days as solution partitions
		return list.IndexedAll(fn.AsPartition(solution, domain), func(day int, scenes []discrete.Variable) bool {
			// Check each day: min/max number of scenes are not violated
			minScenes, maxScenes := cfg.DayMin[day], cfg.DayMax[day]
			numScenes := len(scenes)
			return minScenes <= numScenes && numScenes <= maxScenes
		})
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var score discrete.Score = 0
		// Group scene distribution to days as solution partitions
		for _, scenes := range fn.AsPartition(solution, domain) {
			// Each day: get set of actors working by getting union
			// of all scene actors for scenes scheduled on that day
			actorsToday := ds.NewSet[string]()
			for _, x := range scenes {
				scene := cfg.Scenes[x]
				actorsToday.AddItems(cfg.SceneActors[scene])
			}
			// Add to total cost sum of all today actors' daily cost
			score += list.Sum(list.Translate(actorsToday.Items(), cfg.DailyCost))
		}
		return score
	}

	//p.SolutionCoreFn = fn.CoreSortedPartition(domain, cfg.Scenes)
	p.SolutionStringFn = fn.StringPartition(domain, cfg.Scenes)
	return p
}

// Common steps for creating Item Allocation problem
func newItemAllocationProblem(name string) (*discrete.Problem, *data.ItemAllocation) {
	cfg := data.NewItemAllocation(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Assignment
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(cfg.Items)
	p.AddVariableDomains(discrete.Domain(cfg.Persons))

	return p, cfg
}

// Fair Item Allocation
func fairItemAllocation(name string) *discrete.Problem {
	p, cfg := newItemAllocationProblem(name)
	if p == nil || cfg == nil {
		return nil
	}
	domain := p.UniformDomain()

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		partitions := fn.AsPartition(solution, domain)
		nonEmptyPartitions := list.Filter(partitions, func(partition []discrete.Value) bool {
			return len(partition) > 0
		})
		// If there are less items than persons, minimum number of person who received an item should be numItems
		// If there are less persons than items, minimum number of person who received an item should be numPersons
		// This ensures that one or few people are not hoarding the items and are being distributed evenly
		minCount := min(len(cfg.Items), len(cfg.Persons))
		return len(nonEmptyPartitions) >= minCount
	})

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var envy discrete.Score = 0
		partitions := fn.AsPartition(solution, domain)
		for x1, p1 := range cfg.Persons {
			value1 := list.Sum(list.Map(partitions[x1], func(item discrete.Variable) float64 {
				return cfg.Value[p1][item]
			}))
			for x2 := range cfg.Persons {
				if x1 == x2 {
					continue
				}
				value2 := list.Sum(list.Map(partitions[x2], func(item discrete.Variable) float64 {
					return cfg.Value[p1][item]
				}))
				if value2 > value1 {
					envy += value2 - value1
				}
			}
		}
		return envy
	}

	p.SolutionStringFn = fn.StringPartition(domain, cfg.Items)
	return p
}

// House Allocation
func houseAllocation(name string) *discrete.Problem {
	p, cfg := newItemAllocationProblem(name)
	if p == nil || cfg == nil {
		return nil
	}

	// Ensure each house is assigned to a different person
	p.AddUniversalConstraint(fn.ConstraintAllUnique)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Solution.Map contains House => Person, so we swap the dictionary
		// to form a lookup with Person => House
		houseOf := dict.Swap(solution.Map)
		// Sum up envy among pairs of people:
		// For pair A, B: if A's assigned house has less value to person A
		// than the the house assigned to B, then A is envious of B's house
		var envy discrete.Score = 0
		for x1, p1 := range cfg.Persons {
			value1 := cfg.Value[p1][houseOf[x1]]
			for x2 := range cfg.Persons {
				if x1 == x2 {
					continue
				}
				value2 := cfg.Value[p1][houseOf[x2]]
				if value2 > value1 {
					envy += value2 - value1
				}
			}
		}
		return envy
	}

	p.SolutionStringFn = fn.StringMap(p, cfg.Items, cfg.Persons)
	return p
}
