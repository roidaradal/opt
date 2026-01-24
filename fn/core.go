package fn

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// CoreMirroredSequence groups the normal and the reversed sequence as one
func CoreMirroredSequence[T any](items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		sequence := sequenceStrings(solution, items)
		first, last := sequence[0], list.Last(sequence, 1)
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(sequence)
		}
		return strings.Join(sequence, " ")
	}
}

// CoreMirroredValues groups the normal and reversed list of values as one
func CoreMirroredValues[T any](p *discrete.Problem, items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		output := valueStrings(p, solution, items)
		first, last := output[0], list.Last(output, 1)
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(output)
		}
		return strings.Join(output, " ")
	}
}

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
