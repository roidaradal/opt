package fn

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// AsSubset assumes BooleanDomain {0,1} and returns list of variables with value=1 in solution
func AsSubset(solution *discrete.Solution) []discrete.Variable {
	subset := make([]discrete.Variable, 0, solution.Length())
	for variable, value := range solution.Map {
		if value == 1 {
			subset = append(subset, variable)
		}
	}
	return subset
}

// AsSequence returns list of variables sequenced by solution values
func AsSequence(solution *discrete.Solution) []discrete.Variable {
	sequence := make([]discrete.Variable, solution.Length())
	for variable, idx := range solution.Map {
		sequence[idx] = variable
	}
	return sequence
}

// AsPartition returns the partition from the solution
func AsPartition(solution *discrete.Solution, values []discrete.Value) [][]discrete.Variable {
	groups := make(map[discrete.Value][]discrete.Variable)
	for _, value := range values {
		groups[value] = make([]discrete.Variable, 0)
	}
	for variable, value := range solution.Map {
		groups[value] = append(groups[value], variable)
	}
	partition := make([][]discrete.Variable, len(values))
	for i, value := range values {
		partition[i] = groups[value]
	}
	return partition
}

// PartitionSums returns the sums of each partition group
func PartitionSums[T number.Number](solution *discrete.Solution, values []discrete.Value, items []T) []T {
	return list.Map(AsPartition(solution, values), func(group []discrete.Variable) T {
		return list.Sum(list.MapList(group, items))
	})
}

// PartitionStrings returns the partition as a list of strings
func PartitionStrings[T any](solution *discrete.Solution, values []discrete.Value, items []T) [][]string {
	return list.Map(AsPartition(solution, values), func(group []discrete.Variable) []string {
		return list.Map(list.MapList(group, items), str.Any)
	})
}
