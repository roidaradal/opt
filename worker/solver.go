// Pacakge worker contains common discrete optimization workers
package worker

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/opt/discrete"
)

const IterationBatch int = 1_000_000

type SolverCreator = func(*discrete.Problem) Solver

type Solver interface {
	Solve(LogLevel)
	GetResult() *Result
}

type Result struct {
	NumSteps          int
	BestScore         discrete.Score
	BestSolutions     []*discrete.Solution
	CoreSolutions     map[string][]*discrete.Solution
	FeasibleSolutions dict.Counter[discrete.Score]
}
