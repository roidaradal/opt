package fn

import (
	"slices"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/a"
	"github.com/roidaradal/opt/discrete"
)

// Assumes BooleanDomain {0, 1}, returns list of variables
// that has value = 1 in the solution
func AsSubset(solution *discrete.Solution) []discrete.Variable {
	subset := make([]discrete.Variable, 0, solution.Length())
	for variable, value := range solution.Map {
		if value == 1 {
			subset = append(subset, variable)
		}
	}
	return subset
}

// Assumes BooleanDomain {0, 1}, returns the number of
// selected items (value = 1) in the solution
func SubsetSize(solution *discrete.Solution) int {
	return list.Sum(dict.Values(solution.Map))
}

// Return list of variables sequenced by the solution values
func AsSequence(solution *discrete.Solution) []discrete.Variable {
	sequence := make([]discrete.Variable, solution.Length())
	for variable, idx := range solution.Map {
		sequence[idx] = variable
	}
	return sequence
}

// Return path formed by  solution values
func AsPath(solution *discrete.Solution, cfg *a.PathCfg) []int {
	length := slices.Max(solution.Values()) + 1
	path := make([]int, length+1)
	for variable, idx := range solution.Map {
		if idx < 0 {
			continue
		}
		path[idx] = cfg.IndexOf[variable] // convert to original index
	}
	path[length] = cfg.End
	path = append([]int{cfg.Start}, path...)
	return path
}

// Return list of partitions from the solution
func AsPartitions(solution *discrete.Solution, values []discrete.Value) [][]discrete.Variable {
	groups := make(map[discrete.Value][]discrete.Variable)
	for _, value := range values {
		groups[value] = make([]discrete.Variable, 0)
	}
	for variable, value := range solution.Map {
		groups[value] = append(groups[value], variable)
	}
	partitions := make([][]discrete.Variable, len(values))
	for i, value := range values {
		partitions[i] = groups[value]
	}
	return partitions
}

// Return sums of partition lists
func PartitionSums[T number.Number](solution *discrete.Solution, values []discrete.Value, variables []T) []T {
	return list.Map(AsPartitions(solution, values), func(partition []discrete.Variable) T {
		return list.Sum(list.MapList(partition, variables))
	})
}

// Return partitions as list of strings
func PartitionStrings[T any](solution *discrete.Solution, values []discrete.Value, variables []T) [][]string {
	return list.Map(AsPartitions(solution, values), func(partition []discrete.Variable) []string {
		return list.Map(list.MapList(partition, variables), str.Any)
	})
}

// Count color changes in the sequence
func CountColorChanges[T comparable](colorSequence []T) int {
	var prevColor T
	changes := 0
	for i, currColor := range colorSequence {
		if i > 0 && prevColor != currColor {
			changes += 1
		}
		prevColor = currColor
	}
	return changes
}
