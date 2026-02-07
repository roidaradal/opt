package fn

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/lang"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
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

// CoreMirroredSequence groups the normal and the reversed sequence as one
func CoreMirroredSequence[T any](items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		sequence := sequenceStrings(solution, items)
		return MirroredSequence(sequence)
	}
}

// MirroredSequence checks the first and last items if in order; if not, mirrors the sequence
func MirroredSequence(sequence []string) string {
	return MirroredList(sequence, " ")
}

// MirroredList checks the first and last items if in order; if not, mirrors the list
func MirroredList(sequence []string, separator string) string {
	if len(sequence) == 0 {
		return ""
	}
	first, last := sequence[0], list.Last(sequence, 1)
	if cmp.Compare(first, last) == 1 {
		slices.Reverse(sequence)
	}
	return strings.Join(sequence, separator)
}

// CoreMirroredValues groups the normal and reversed list of values as one
func CoreMirroredValues[T any](p *discrete.Problem, items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		output := valueStrings(p, solution, items)
		return MirroredList(output, " ")
	}
}

// CoreSortedCycle groups the normal and sorted cycles as one
func CoreSortedCycle[T any](items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		sequence := sequenceStrings(solution, items)
		return SortedCycle(sequence, false)
	}
}

// SortedCycle finds the first element in sorted order,
// Rearrange the sequence so it becomes the first element
func SortedCycle(sequence []string, removeTail bool) string {
	limit := lang.Ternary(removeTail, len(sequence)-1, len(sequence))
	index := slices.Index(sequence, slices.Min(sequence))
	sequence2 := append([]string{}, sequence[index:limit]...)
	sequence2 = append(sequence2, sequence[:index]...)
	return strings.Join(sequence2, " ")
}

// CoreLookupValueOrder groups solutions based on the relative order of the solution values
func CoreLookupValueOrder(problem *discrete.Problem) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		values := solution.Tuple(problem)
		core := make([]string, len(values))
		lookup := make(map[discrete.Value]string)
		order := 0
		for i, value := range values {
			if dict.NoKey(lookup, value) {
				lookup[value] = str.Int(order)
				order += 1
			}
			core[i] = lookup[value]
		}
		return strings.Join(core, "")
	}
}
