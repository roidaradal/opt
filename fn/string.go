package fn

import (
	"cmp"
	"slices"
	"strings"

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

// SolutionStringFn: display solution mapped to values
func String_Values[T any](p *discrete.Problem, values []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(x discrete.Variable) string {
			value := solution.Map[x]
			if values == nil {
				return str.Int(value)
			} else {
				return str.Any(values[value])
			}
		})
		return strings.Join(output, " ")
	}
}
