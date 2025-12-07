package worker

import (
	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/opt/discrete"
)

const (
	LOG_NONE  LogLevel = iota // no logging
	LOG_BATCH                 // log iteration batch
	LOG_STEPS                 // log step-by-step
)

const IterationBatch int = 1_000_000

type LogLevel uint

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
