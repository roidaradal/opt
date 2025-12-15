package worker

import (
	"fmt"
	"time"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

type Manager interface {
	Run([]*discrete.Problem, *Config)
}

type Solo struct{}

// Run cfg.Worker on problems
func (w Solo) Run(problems []*discrete.Problem, cfg *Config) {
	runStart := time.Now()
	numProblems := len(problems)
	for i, problem := range problems {
		fmt.Printf("[%3d / %3d] ", i+1, numProblems)
		start := time.Now()
		cfg.Problem = problem
		output := cfg.Worker.Run(cfg)
		duration := str.Violet(fmt.Sprintf("%v", time.Since(start).Round(time.Millisecond)))
		parts := str.CleanSplit(output, "|")
		if len(parts) > 1 {
			output = fmt.Sprintf(cfg.Worker.Columns(), list.ToAny(parts)...)
		}
		fmt.Printf("%-15s | %s\n", duration, output)
	}
	fmt.Println("\nTime:", time.Since(runStart).Round(time.Millisecond))
}
