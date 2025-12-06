package problem

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Number Partition problem
func NumberPartition(n int) *discrete.Problem {
	name := newName(NUMBER_PARTITION, n)
	numbers := newNumberPartition(name)
	if numbers == nil {
		return nil
	}

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize

	p.Variables = discrete.Variables(numbers)
	domain := discrete.RangeDomain(1, 2)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	test := func(solution *discrete.Solution) bool {
		if p.IsOptimization() {
			return true // don't test if optimization problem
		}
		// Check if the 2 partition sums are the same
		sums := fn.PartitionSums(solution, domain, numbers)
		return list.AllSame(sums)
	}
	p.AddUniversalConstraint(test)

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Minimize the difference between the 2 partition sums
		sums := fn.PartitionSums(solution, domain, numbers)
		return discrete.Score(number.Abs(sums[0] - sums[1]))
	}
	p.SolutionCoreFn = fn.Core_SortedPartition(domain, numbers)
	p.SolutionStringFn = fn.String_Partitions(domain, numbers)

	return p
}

// Load number partition test case
func newNumberPartition(name string) []int {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) != 1 {
		return nil
	}
	return list.Map(strings.Fields(lines[0]), number.ParseInt)
}
