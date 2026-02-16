package problem

import (
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
