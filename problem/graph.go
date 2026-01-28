package problem

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/opt/fn"
)

// Load new unweighted graph and return offset lines
func newUnweightedGraph(name string) (*ds.Graph, []string) {
	lines, err := fn.LoadLines(name)
	numLines := len(lines)
	if err != nil || numLines < 2 {
		return nil, nil
	}
	graph := ds.GraphFrom(lines[0], lines[1])
	var extra []string
	if numLines > 2 {
		extra = lines[2:]
	}
	return graph, extra
}
