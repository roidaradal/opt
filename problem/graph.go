package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/fn"
)

// Load new unweighted graph
func newUnweightedGraph(name string) *graphCfg {
	lines, err := fn.LoadLines(name)
	numLines := len(lines)
	if err != nil || numLines < 2 {
		return nil
	}
	cfg := &graphCfg{
		Graph: ds.GraphFrom(lines[0][0], lines[1][0]),
		extra: make([][]string, 0),
	}
	if numLines > 2 {
		cfg.extra = lines[2:]
	}
	return cfg
}

// Temporary: get graph vertices
// TODO: transfer this to ds.Graph in fn package
func graphVertices(graph *graphCfg) []string {
	return graph.Vertices
}

// Get edge names of graph
func graphEdges(graph *graphCfg) []string {
	return graph.EdgeNames()
}
