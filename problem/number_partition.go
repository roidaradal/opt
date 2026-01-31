package problem

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/data"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// NewNumberPartition creates a new Number Partition problem
func NewNumberPartition(variant string, n int) *discrete.Problem {
	name := newName(NumberPartition, variant, n)
	switch variant {
	case "basic":
		return numberPartition(name)
	default:
		return nil
	}
}

// Number Partition
func numberPartition(name string) *discrete.Problem {
	cfg := data.NewNumbers(name)
	if cfg == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Type = discrete.Partition

	domain := discrete.RangeDomain(1, 2)
	p.Variables = discrete.Variables(cfg.Numbers)
	p.AddVariableDomains(domain)

	p.AddUniversalConstraint(func(solution *discrete.Solution) bool {
		if p.IsOptimization() {
			return true // don't test if optimization problem
		}
		// Check if the 2 partition sums are the same
		sums := fn.PartitionSums(solution, domain, cfg.Numbers)
		return list.AllSame(sums)
	})

	p.Goal = discrete.Minimize
	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Minimize difference between the 2 partition sums
		sums := fn.PartitionSums(solution, domain, cfg.Numbers)
		return discrete.Score(number.Abs(sums[0] - sums[1]))
	}

	p.SolutionCoreFn = fn.CoreSortedPartition(domain, cfg.Numbers)
	p.SolutionStringFn = fn.StringPartition(domain, cfg.Numbers)
	return p
}
