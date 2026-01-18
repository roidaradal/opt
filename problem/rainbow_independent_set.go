package problem

import (
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Rainbow Independent Set problem
func RainbowIndependentSet(n int) *discrete.Problem {
	name := newName(RAINBOW_INDEPENDENT_SET, n)
	graph, colors, colorOf := newRainbowIndependentSet(name)
	if graph == nil || colors == nil || colorOf == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(graph.Vertices)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test1 := func(solution *discrete.Solution) bool {
		// Check independent set
		vertices := list.MapList(fn.AsSubset(solution), graph.Vertices)
		return graph.IsIndependentSet(vertices)
	}
	p.AddUniversalConstraint(test1)

	test2 := func(solution *discrete.Solution) bool {
		// Check that vertices have different colors
		count := dict.NewCounter(colors)
		for _, x := range fn.AsSubset(solution) {
			color := colorOf[graph.Vertices[x]]
			count[color] += 1
		}
		return list.AllLessEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test2)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(graph.Vertices)

	return p
}

// Load rainbow independent set test case
func newRainbowIndependentSet(name string) (*ds.Graph, []string, dict.StringMap) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 4 {
		return nil, nil, nil
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	numColors := number.ParseInt(lines[2])
	colorOf := make(dict.StringMap)
	colors := make([]string, numColors)
	for i := range numColors {
		parts := str.CleanSplit(lines[3+i], ":")
		if len(parts) != 2 {
			continue
		}
		color := parts[0]
		colors[i] = color
		for _, vertex := range strings.Fields(parts[1]) {
			colorOf[vertex] = color
		}
	}
	return graph, colors, colorOf
}
