package fn

import (
	"cmp"
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

// StringFn: display solution as subset
func String_Subset[T cmp.Ordered](items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		subset := list.MapList(AsSubset(solution), items)
		slices.Sort(subset)
		return str.WrapBraces(list.Map(subset, str.Any))
	}
}
