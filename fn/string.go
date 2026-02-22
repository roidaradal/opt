package fn

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
)

// StringPartition displays the solution as a partition
func StringPartition[T any](values []discrete.Value, items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		groups := PartitionStrings(solution, values, items)
		partition := sortedPartitionGroups(groups)
		return strings.Join(partition, " ")
	}
}

// StringSubset displays the solution as subset
func StringSubset[T cmp.Ordered](items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		subset := list.MapList(AsSubset(solution), items)
		slices.Sort(subset)
		return str.WrapBraces(list.Map(subset, str.Any))
	}
}

// StringSequence displays the solution as sequence of variables
func StringSequence[T any](items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		sequence := sequenceStrings(solution, items)
		return strings.Join(sequence, " ")
	}
}

// StringValues displays the solution mapped to given values
func StringValues[T any](p *discrete.Problem, items []T) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		output := valueStrings(p, solution, items)
		return strings.Join(output, " ")
	}
}

// StringMap displays solution as {variable = value}
func StringMap[T any, V any](p *discrete.Problem, variables []T, values []V) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(x discrete.Variable) string {
			value := solution.Map[x]
			k, v := str.Int(x), str.Int(value)
			if variables != nil {
				k = str.Any(variables[x])
			}
			if values != nil {
				v = str.Any(values[value])
			}
			return fmt.Sprintf("%s = %s", k, v)
		})
		return str.WrapBraces(output)
	}
}

// StringAssignment displays assignment of {worker = task}
// Note: didn't explicitly return as SolutionStringFn so it can be reused as SolutionCoreFn
func StringAssignment(p *discrete.Problem, cfg *data.AssignmentCfg) func(*discrete.Solution) string {
	return func(solution *discrete.Solution) string {
		output := list.Map(p.Variables, func(worker discrete.Variable) string {
			task := solution.Map[worker]
			if cfg.Cost[worker][task] == 0 {
				return "" // skip dummy tasks
			}
			return fmt.Sprintf("w%s = t%s", cfg.Workers[worker], cfg.Tasks[task])
		})
		output = list.Filter(output, str.NotEmpty)
		return str.WrapBraces(output)
	}
}

// StringGraphPath displays the path sequence of vertices
func StringGraphPath(cfg *data.GraphPath) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		path := AsGraphPath(solution, cfg)
		return strings.Join(list.MapList(path, cfg.Vertices), "-")
	}
}

// StringEulerianPath displays the Eulerian path sequence of vertices
func StringEulerianPath(graph *ds.Graph) discrete.SolutionStringFn {
	return func(solution *discrete.Solution) string {
		path := make([]string, 0, len(graph.Edges))
		edgeSequence := list.MapList(AsSequence(solution), graph.Edges)
		a1, b1 := edgeSequence[0].Tuple()
		a2, b2 := edgeSequence[1].Tuple()
		var tail ds.Vertex
		switch {
		case a1 == a2:
			path = append(path, b1, a1)
			tail = b2
		case a1 == b2:
			path = append(path, b1, a1)
			tail = a2
		case b1 == a2:
			path = append(path, a1, b1)
			tail = b2
		case b1 == b2:
			path = append(path, a1, b1)
			tail = a2
		default:
			return InvalidSolution
		}
		for _, edge := range edgeSequence[2:] {
			a, b := edge.Tuple()
			path = append(path, tail)
			switch tail {
			case a:
				tail = b
			case b:
				tail = a
			default:
				return InvalidSolution
			}
		}
		path = append(path, tail)
		return strings.Join(path, " ")
	}
}
