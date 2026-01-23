package fn

import (
	"cmp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// CoreMirroredSequence groups the normal and the reversed sequence as one
func CoreMirroredSequence[T any](items []T) discrete.SolutionCoreFn {
	return func(solution *discrete.Solution) string {
		sequence := list.Map(AsSequence(solution), func(x discrete.Variable) string {
			return str.Any(items[x])
		})
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
		output := list.Map(p.Variables, func(x discrete.Variable) string {
			value := solution.Map[x]
			if items == nil {
				return str.Int(value)
			}
			return str.Any(items[value])
		})
		first, last := output[0], list.Last(output, 1)
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(output)
		}
		return strings.Join(output, " ")
	}
}
