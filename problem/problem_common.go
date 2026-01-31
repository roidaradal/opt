package problem

import (
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Common steps for creating Bin problems
func newBinProblem(name string) (*discrete.Problem, *data.Bins) {
	cfg := data.NewBins(name)
	if cfg == nil {
		return nil, nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	p.Variables = discrete.Variables(cfg.Weight)
	p.AddVariableDomains(cfg.Bins)

	p.ObjectiveFn = fn.ScoreCountUniqueValues
	p.SolutionCoreFn = fn.CoreSortedPartition(cfg.Bins, cfg.Weight)
	p.SolutionStringFn = fn.StringPartition(cfg.Bins, cfg.Weight)

	return p, cfg
}
