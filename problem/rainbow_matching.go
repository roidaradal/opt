package problem

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Rainbow Matching problem
func RainbowMatching(n int) *discrete.Problem {
	name := newName(RAINBOW_MATCH, n)
	colorOf := newRainbowMatching(name)
	if colorOf == nil {
		return nil
	}
	edges := dict.Keys(colorOf)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Maximize
	p.Type = discrete.Subset

	p.Variables = discrete.Variables(edges)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test1 := func(solution *discrete.Solution) bool {
		// Check if solution is a matching (edges have no common vertex)
		count := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			for _, vertex := range strings.Split(edges[x], "-") {
				count[vertex] += 1
			}
		}
		return list.AllEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test1)

	test2 := func(solution *discrete.Solution) bool {
		// Check that selected edges have different colors
		count := make(dict.StringCounter)
		for _, x := range fn.AsSubset(solution) {
			color := colorOf[edges[x]]
			count[color] += 1
		}
		return list.AllEqual(dict.Values(count), 1)
	}
	p.AddUniversalConstraint(test2)

	p.ObjectiveFn = fn.Score_SubsetSize
	p.SolutionStringFn = fn.String_Subset(edges)

	return p
}

// Load rainbow matching test case
func newRainbowMatching(name string) dict.StringMap {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 4 {
		return nil
	}
	numSide1 := number.ParseInt(lines[0])
	side2 := strings.Fields(lines[1])
	colorOf := make(dict.StringMap)
	for i := range numSide1 {
		parts := strings.Fields(lines[2+i])
		vertex1 := parts[0]
		for j, vertex2 := range side2 {
			edge := fmt.Sprintf("%s-%s", vertex1, vertex2)
			colorOf[edge] = parts[1+j]
		}
	}
	return colorOf
}
