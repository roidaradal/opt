package problem

import (
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Scene Allocation problem
func SceneAllocation(n int) *discrete.Problem {
	name := newName(SCENE_ALLOCATION, n)
	cfg := newSceneAllocation(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(cfg.scenes)
	domain := discrete.MapDomain(cfg.days)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		// Group the scene distributon to days as solution partitions
		return list.IndexedAll(fn.AsPartitions(solution, domain), func(day int, scenes []discrete.Variable) bool {
			// Check that for each day, the minimum and maximum number of scenes are followed
			limits := cfg.days[day]
			minScenes, maxScenes := limits[0], limits[1]
			numScenes := len(scenes)
			return minScenes <= numScenes && numScenes <= maxScenes
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		var score discrete.Score = 0
		// Group the scene distribution to days as solution partitions
		for _, scenes := range fn.AsPartitions(solution, domain) {
			// For each day, get the set of actors working by getting the union
			// of all scene actors for the scenes scheduled on that day
			actorsToday := ds.NewSet[string]()
			for _, x := range scenes {
				scene := cfg.scenes[x]
				actorsToday.AddItems(cfg.sceneActors[scene])
			}
			// Add to the total cost the sum of all actors today's daily cost
			score += list.Sum(list.Translate(actorsToday.Items(), cfg.dailyCost))
		}
		return score
	}

	// p.SolutionCoreFn = fn.Core_SortedPartition(domain, cfg.scenes)
	p.SolutionStringFn = fn.String_Partitions(domain, cfg.scenes)

	return p
}

type sceneCfg struct {
	days        [][2]int            // [MinScenes, MaxScenes] per day
	dailyCost   map[string]float64  // {Actor => DailyCost}
	scenes      []string            // List of scenes
	sceneActors map[string][]string // {Scene => []Actors}
}

// Load scene allocation test case
func newSceneAllocation(name string) *sceneCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 5 {
		return nil
	}
	cfg := &sceneCfg{
		days:        make([][2]int, 0),
		dailyCost:   make(map[string]float64),
		scenes:      make([]string, 0),
		sceneActors: make(map[string][]string),
	}
	counts := list.Map(strings.Fields(lines[0]), number.ParseInt)
	minScenes := list.Map(strings.Fields(lines[1]), number.ParseInt)
	maxScenes := list.Map(strings.Fields(lines[2]), number.ParseInt)
	numDays, numActors, numScenes := counts[0], counts[1], counts[2]
	for i := range numDays {
		cfg.days = append(cfg.days, [2]int{minScenes[i], maxScenes[i]})
	}
	idx := 3
	for range numActors {
		parts := strings.Fields(lines[idx])
		actor, cost := parts[0], number.ParseFloat(parts[1])
		cfg.dailyCost[actor] = cost
		idx++
	}
	for range numScenes {
		parts := strings.Fields(lines[idx])
		scene, actors := parts[0], parts[1:]
		cfg.scenes = append(cfg.scenes, scene)
		cfg.sceneActors[scene] = actors
		idx++
	}
	return cfg
}
