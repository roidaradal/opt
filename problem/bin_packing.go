package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Bin Packing problem
func BinPacking(n int) *discrete.Problem {
	name := newName(BIN_PACKING, n)
	cfg := newBinPacking(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(cfg.weight)
	domain := discrete.RangeDomain(1, cfg.numBins)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		sums := fn.PartitionSums(solution, domain, cfg.weight)
		return list.All(sums, func(sum float64) bool {
			return sum <= cfg.capacity
		})
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = fn.Score_CountUniqueValues
	p.SolutionCoreFn = fn.Core_SortedPartition(domain, cfg.weight)
	p.SolutionStringFn = fn.String_Partitions(domain, cfg.weight)

	return p
}

type binPackingCfg struct {
	numBins  int
	capacity float64
	weight   []float64
}

// Load bin packing test case
func newBinPacking(name string) *binPackingCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 3 {
		return nil
	}
	return &binPackingCfg{
		numBins:  number.ParseInt(lines[0]),
		capacity: number.ParseFloat(lines[1]),
		weight:   list.Map(strings.Fields(lines[2]), number.ParseFloat),
	}
}
