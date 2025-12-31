package fn

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// SolutionCoreFn: mirrored sequence
func Core_MirroredSequence[T any](variables []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		sequence := list.Map(AsSequence(solution), func(x discrete.Variable) string {
			return str.Any(variables[x])
		})
		first, last := sequence[0], sequence[len(sequence)-1]
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(sequence)
		}
		return strings.Join(sequence, " ")
	}
}

// SolutionCoreFn: mirrored values
func Core_MirroredValues[T any](p *discrete.Problem, values []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(x discrete.Variable) string {
			value := solution.Map[x]
			if values == nil {
				return str.Int(value)
			} else {
				return str.Any(values[value])
			}
		})
		first, last := output[0], output[len(output)-1]
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(output)
		}
		return strings.Join(output, " ")
	}
}

// SolutionCoreFn: sorted partitions
func Core_SortedPartition[T any](values []discrete.Value, variables []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		groups := PartitionStrings(solution, values, variables)
		groups = list.Filter(groups, list.NotEmpty)
		partitions := list.Map(groups, func(group []string) string {
			slices.Sort(group)
			return str.WrapBraces(group)
		})
		slices.Sort(partitions)
		return strings.Join(partitions, "/")
	}
}

// SolutionCoreFn: lookup value order
func Core_LookupValueOrder(problem *discrete.Problem) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		values := solution.Tuple(problem)
		core := make([]string, len(values))
		lookup := make(map[discrete.Value]string)
		order := 0
		for i, value := range values {
			if dict.NoKey(lookup, value) {
				lookup[value] = str.Int(order)
				order++
			}
			core[i] = lookup[value]
		}
		return strings.Join(core, "")
	}
}

// SolutionCoreFn: sorted cycle
func Core_SortedCycle[T any](variables []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		sequence := list.Map(AsSequence(solution), func(x discrete.Variable) string {
			return str.Any(variables[x])
		})
		// Find first element in sorted order
		// Rearrange sequence so it becomes first element
		index := slices.Index(sequence, slices.Min(sequence))
		sequence2 := append([]string{}, sequence[index:]...)
		sequence2 = append(sequence2, sequence[:index]...)
		return strings.Join(sequence2, " ")
	}
}
