package worker

import (
	"fmt"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/opt/discrete"
)

const IterationBatch int = 1_000_000

const (
	InfeasibleResult string = "FAIL"
	FeasibleResult   string = "PASS"
	BestResult       string = "BEST"
)

type SolverCreator = func(*discrete.Problem) Solver

type Solver interface {
	GetName() string
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

type SpaceSolver struct{}

// Return problem's solution space
func (s SpaceSolver) Run(cfg *Config) string {
	problem := cfg.Problem
	if problem == nil {
		return errMessage(errMissingProblem)
	}
	size := problem.SolutionSpace()
	space := number.Comma(size)
	if size == 0 {
		space = problem.SolutionSpaceEquation()
	}
	return fmt.Sprintf("%-20s: %15s", problem.Name, space)
}
