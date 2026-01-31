package fn

import (
	"strings"

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
