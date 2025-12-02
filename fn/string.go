package fn

import (
	"cmp"
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// SolutionStringFn: display solution as subset
func String_Subset[T cmp.Ordered](variables []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		subset := list.MapList(AsSubset(solution), variables)
		slices.Sort(subset)
		output := list.Map(subset, str.Any)
		return str.WrapBraces(output)
	}
}
