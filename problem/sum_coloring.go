package problem

import (
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Sum Coloring problem
func SumColoring(n int) *discrete.Problem {
	name := newName(SUM_COLOR, n)
	graph, numbers := newSumColoring(name)
	if graph == nil || numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.MapDomain(numbers)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		color := solution.Map
		// For all graph edges, check that color of 2 vertices are different
		return list.All(graph.Edges, func(edge ds.Edge) bool {
			x1, x2 := graph.IndexOf[edge[0]], graph.IndexOf[edge[1]]
			return color[x1] != color[x2]
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		total := list.Sum(list.MapList(dict.Values(solution.Map), numbers))
		return discrete.Score(total)
	}
	p.SolutionStringFn = fn.String_Values(p, numbers)

	return p
}

// Load sum coloring test case
func newSumColoring(name string) (*ds.Graph, []int) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil, nil
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	numbers := list.Map(strings.Fields(lines[2]), number.ParseInt)
	return graph, numbers
}
