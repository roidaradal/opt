package constraint

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/discrete"
)

// Knapsack Constraint
func Knapsack(p *discrete.Problem, capacity float64, weight []float64) discrete.ConstraintFn {
	return func(solution *discrete.Solution) bool {
		// Check sum of weighted items does not exceed capacity
		count := solution.Map
		weights := list.Map(p.Variables, func(x discrete.Variable) float64 {
			return float64(count[x]) * weight[x]
		})
		return list.Sum(weights) <= capacity
	}
}
