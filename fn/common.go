package fn

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// Convert partition groups into sorted partition groups, wraped in braces
func sortedPartitionGroups(groups [][]string) []string {
	return list.Map(groups, func(group []string) string {
		slices.Sort(group)
		return str.WrapBraces(group)
	})
}

// Convert solution to sequence of item strings
func sequenceStrings[T any](solution *discrete.Solution, items []T) []string {
	return list.Map(AsSequence(solution), func(x discrete.Variable) string {
		return str.Any(items[x])
	})
}

// Convert solution to list of value strings
func valueStrings[T any](p *discrete.Problem, solution *discrete.Solution, items []T) []string {
	return list.Map(p.Variables, func(x discrete.Variable) string {
		value := solution.Map[x]
		if items == nil {
			return str.Int(value)
		}
		return str.Any(items[value])
	})
}
