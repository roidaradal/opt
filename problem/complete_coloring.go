package problem

import (
	"github.com/roidaradal/fn/comb"
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Complete Coloring problem
func CompleteColoring(n int) *discrete.Problem {
	name := newName(COMPLETE_COLOR, n)
	graph, numColors := newCompleteColoring(name)
	if graph == nil || numColors == 0 {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Satisfy
	p.Type = discrete.Assignment

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.IndexDomain(numColors)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	colors := list.NumRange(0, numColors)
	test := func(solution *discrete.Solution) bool {
		count := make(map[[2]int]int) // ColorPair => Count
		for _, pair := range comb.Combinations(colors, 2) {
			key := [2]int{pair[0], pair[1]}
			count[key] = 0
		}
		color := solution.Map
		for _, edge := range graph.Edges {
			v1, v2 := edge.Tuple()
			x1, x2 := graph.IndexOf[v1], graph.IndexOf[v2]
			c1, c2 := color[x1], color[x2]
			if c1 == c2 {
				continue // skip same color
			}
			key := [2]int{c1, c2}
			if dict.NoKey(count, key) {
				key = [2]int{c2, c1} // flip if original order not found
			}
			count[key] += 1
		}
		// Check all color pairs appears at least once
		return list.AllGreaterEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test)

	p.SolutionCoreFn = fn.Core_LookupValueOrder(p)
	p.SolutionStringFn = fn.String_Values(p, colors)

	return p
}

// Load complete coloring test case
func newCompleteColoring(name string) (*ds.Graph, int) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil, 0
	}
	numColors := number.ParseInt(lines[0])
	graph := ds.GraphFrom(lines[1], lines[2])
	return graph, numColors
}
