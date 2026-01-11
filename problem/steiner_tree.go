package problem

import (
	"strings"

	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/constraint"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Steiner Tree problem
func SteinerTree(n int) *discrete.Problem {
	name := newName(STEINER_TREE, n)
	graph, edgeWeight, terminals := newSteinerTree(name)
	if graph == nil || edgeWeight == nil || terminals == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Subset

	edgeNames := list.Map(graph.Edges, ds.Edge.String)
	p.Variables = discrete.Variables(edgeNames)
	domain := discrete.BooleanDomain()
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	// Constraint: all terminal vertices are spanned
	p.AddUniversalConstraint(constraint.AllVerticesSpanned(graph, terminals))

	// Constraint: solution forms a tree: all vertices reachable from tree traversal
	p.AddUniversalConstraint(constraint.SpanningTree(graph, terminals))

	p.ObjectiveFn = fn.Score_SumWeightedValues(p.Variables, edgeWeight)
	p.SolutionStringFn = fn.String_Subset(edgeNames)

	return p
}

// Load steiner tree test case
func newSteinerTree(name string) (*ds.Graph, []float64, []ds.Vertex) {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 4 {
		return nil, nil, nil
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	edgeWeight := list.Map(strings.Fields(lines[2]), number.ParseFloat)
	terminals := strings.Fields(lines[3])
	return graph, edgeWeight, terminals
}
