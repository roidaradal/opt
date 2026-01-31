package fn

import (
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// CoreSortedPartition groups similar partitions by using their sorted versions
func CoreSortedPartition[T any](values []discrete.Value, items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		groups := PartitionStrings(solution, values, items)
		groups = list.Filter(groups, list.NotEmpty)
		partition := sortedPartitionGroups(groups)
		slices.Sort(partition)
		return strings.Join(partition, "/")
	}
}
