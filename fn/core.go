package fn

import (
	"cmp"
	"slices"
	"strings"

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
		first, last := sequence[0], sequence[solution.Length()-1]
		if cmp.Compare(first, last) == 1 {
			slices.Reverse(sequence)
		}
		return strings.Join(sequence, " ")
	}
}
