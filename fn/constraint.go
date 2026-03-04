package fn

import (
	"slices"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
)

// ConstraintAllUnique makes sure all solution values are unique
func ConstraintAllUnique(solution *discrete.Solution) bool {
	return list.AllUnique(solution.Values())
}

// ConstraintProperVertexColoring makes sure all adjacent vertices in the graph have a different color
func ConstraintProperVertexColoring(graph *ds.Graph) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		color := solution.Map
		// For all graph edges, check that color of 2 vertices are different
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return color[x1] != color[x2]
		})
	}
}

// ConstraintAllVerticesCovered makes sure all vertices are covered at least once
func ConstraintAllVerticesCovered(graph *ds.Graph, vertices []ds.Vertex) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Go through all edges formed by the subset solution
		// Mark the 2 vertices as covered
		covered := dict.Flags(vertices, false)
		for _, x := range AsSubset(solution) {
			v1, v2 := graph.Edges[x].Tuple()
			covered[v1] = true
			covered[v2] = true
		}
		return list.AllTrue(dict.Values(covered))
	}
}

// ConstraintSpanningTree checks if the solution forms a tree, and all vertices are reachable from tree traversal
func ConstraintSpanningTree(graph *ds.Graph, vertices []ds.Vertex) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		reachable := SpannedVertices(solution, graph)
		if reachable == nil {
			return false
		}
		vertexSet := ds.SetFrom(vertices)
		// Check all vertices are reachable
		return vertexSet.Difference(reachable).IsEmpty()
	}
}

// ConstraintRainbowColoring makes sure all chosen items have different colors
func ConstraintRainbowColoring(colors []string) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		return list.AllUnique(list.MapList(AsSubset(solution), colors))
	}
}

// ConstraintSimplePath makes sure solution forms a valid simple path (no repeated vertices)
func ConstraintSimplePath(cfg *data.GraphPath) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		path := AsGraphPath(solution, cfg)
		prev := path[0]

		visited := ds.NewSet[int]()
		visited.Add(prev)
		for _, curr := range path[1:] {
			if visited.Has(curr) {
				return false // repeated vertex = not simple path
			}
			if cfg.Distance[prev][curr] == discrete.Inf {
				return false // no edge from prev -> curr
			}
			visited.Add(curr)
			prev = curr
		}
		return true
	}
}

// ConstraintIncreasingSubsequence makes sure that the subsequence is increasing
func ConstraintIncreasingSubsequence(cfg *data.Numbers) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		subset := AsSubset(solution)
		numSelected := len(subset)
		if numSelected <= 1 {
			return true // no need to check if 0 or 1 item in sequence
		}
		slices.Sort(subset) // sort indexes
		subsequence := list.MapList(subset, cfg.Numbers)
		for i := range numSelected - 1 {
			if subsequence[i] >= subsequence[i+1] {
				return false // invalid if current not less than next
			}
		}
		return true
	}
}

// ConstraintNoMachineOverlap checks that the schedule has no machine overlap
func ConstraintNoMachineOverlap(cfg *data.ShopSchedule, taskLookup map[discrete.Variable]data.Task) discrete.ConstraintFn {
	return noOverlap(taskLookup, cfg.Machines, func(task data.Task) string {
		return task.Machine
	})
}

// Common: check that schedule has no overlap
func noOverlap(taskLookup map[discrete.Variable]data.Task, keys []string, keyFn func(data.Task) string) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Initialize all key's schedules
		groupSched := make(map[string][]data.TimeRange)
		for _, key := range keys {
			groupSched[key] = make([]data.TimeRange, 0)
		}
		// For each task in solution, add schedule TimeRange to its group schedule
		for x, start := range solution.Map {
			task := taskLookup[x]
			sched := data.TimeRange{start, start + task.Duration}
			key := keyFn(task)
			groupSched[key] = append(groupSched[key], sched)
		}
		// Sort schedules for each group, check if there is overlap
		for _, scheds := range groupSched {
			slices.SortFunc(scheds, data.SortByStartTime)
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
