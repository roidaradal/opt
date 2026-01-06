package problem

import (
	"slices"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Harmonious Coloring problem
func HarmoniousColoring(n int) *discrete.Problem {
	name := newName(HARMONIOUS_COLOR, n)
	graph, numColors := fn.NewVertexColoring(name)
	if graph == nil || numColors == 0 {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.IndexDomain(numColors)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		count := make(map[[2]int]int) // ColorPair => Count
		color := solution.Map
		for _, edge := range graph.Edges {
			v1, v2 := edge.Tuple()
			x1, x2 := graph.IndexOf[v1], graph.IndexOf[v2]
			c1, c2 := color[x1], color[x2]
			if c1 == c2 {
				// If edge endpoints have same color,
				// it is not a proper vertex coloring = invalid
				return false
			}
			colors := []int{c1, c2}
			slices.Sort(colors)
			key := [2]int{colors[0], colors[1]}
			count[key] += 1
		}
		// Check all color pairs appears at most once
		return list.AllLessEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_LookupValueOrder(p)
	p.SolutionStringFn = fn.String_Values(p, list.NumRange(0, numColors))

	return p
}
