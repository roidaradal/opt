package fn

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// StringPartition displays the solution as a partition
func StringPartition[T any](values []discrete.Value, items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		groups := PartitionStrings(solution, values, items)
		partition := sortedPartitionGroups(groups)
		return strings.Join(partition, " ")
	}
}

// StringSubset displays the solution as subset
func StringSubset[T cmp.Ordered](items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		subset := list.MapList(AsSubset(solution), items)
		slices.Sort(subset)
		return str.WrapBraces(list.Map(subset, str.Any))
	}
}
