package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/fn"
)

// Load new unweighted graph and return offset lines
func newUnweightedGraph(name string, offset int) (*ds.Graph, []string) {
	lines, err := fn.LoadLines(name)
	if err != nil || len(lines) != offset+2 {
		return nil, nil
	}
	graph := ds.GraphFrom(lines[offset], lines[offset+1])
	return graph, lines[:offset]
}
