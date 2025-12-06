package fn

import (
	"cmp"
	"fmt"
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

// SolutionStringFn: display solution as sequence of variables
func String_Sequence[T any](variables []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		sequence := list.Map(AsSequence(solution), func(x discrete.Variable) string {
			return str.Any(variables[x])
		})
		return strings.Join(sequence, " ")
	}
}

// SolutionStringFn: display solution as partitions
func String_Partitions[T any](values []discrete.Value, variables []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		groups := PartitionStrings(solution, values, variables)
		partitions := list.Map(groups, func(group []string) string {
			slices.Sort(group)
			return str.WrapBraces(group)
		})
		return strings.Join(partitions, " ")
	}
}

// SolutionStringFn: display solution as translated { variable = value }
func String_Map[T any, V any](p *discrete.Problem, variables []T, values []V) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(x discrete.Variable) string {
			value := solution.Map[x]
			text1, text2 := str.Int(x), str.Int(value)
			if variables != nil {
				text1 = str.Any(variables[x])
			}
			if values != nil {
				text2 = str.Any(values[value])
			}
			return fmt.Sprintf("%s = %s", text1, text2)
		})
		return str.WrapBraces(output)
	}
}
