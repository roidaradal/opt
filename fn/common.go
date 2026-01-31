package fn

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
)

// Convert partition groups into sorted partition groups, wraped in braces
func sortedPartitionGroups(groups [][]string) []string {
	return list.Map(groups, func(group []string) string {
		slices.Sort(group)
		return str.WrapBraces(group)
	})
}
