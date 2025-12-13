package worker

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/opt/discrete"
)

const IterationBatch int = 1_000_000

type SolverCreator = func(*discrete.Problem) Solver

type Solver interface {
	GetName() string
	GetProblem() *discrete.Problem
	GetResult() *Result
	Solve(Logger)
}

type Result struct {
	NumSteps          int
	BestScore         discrete.Score
	BestSolutions     []*discrete.Solution
	CoreSolutions     map[string][]*discrete.Solution
	FeasibleSolutions dict.Counter[discrete.Score]
}

const (
	InfeasibleResult string = "FAIL"
	FeasibleResult   string = "PASS"
	BestResult       string = "BEST"
)
