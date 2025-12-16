package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

type Manager interface {
	Run([]*discrete.Problem, *Config)
}

type Solo struct{}

type Pool struct {
	NumWorkers int
}

// Run cfg.Worker on problems
func (m Solo) Run(problems []*discrete.Problem, cfg *Config) {
	runStart := time.Now()
	numProblems := len(problems)
	for i, problem := range problems {
		fmt.Printf("[%3d / %3d] ", i+1, numProblems)
		start := time.Now()
		cfg.Problem = problem
		output := cfg.Worker.Run(cfg)
		output = workerFormat(output, cfg.Worker)
		duration := getDuration(start)
		fmt.Printf("%-15s | %s\n", duration, output)
	}
	fmt.Println("\nTime:", time.Since(runStart).Round(time.Millisecond))
}

// Run cfg.Worker on problems using a pool of workers
func (m Pool) Run(problems []*discrete.Problem, cfg *Config) {
	runStart := time.Now()
	problemCh := make(chan *discrete.Problem)
	outputCh := make(chan string)

	// Worker function
	runWorker := func(problemCh <-chan *discrete.Problem, outputCh chan<- string) {
		workerCfg := cfg.Copy()
		worker := workerCfg.Worker
		for problem := range problemCh {
			start := time.Now()
			workerCfg.Problem = problem
			output := worker.Run(workerCfg)
			output = workerFormat(output, worker)
			outputCh <- fmt.Sprintf("%-15s | %s", getDuration(start), output)
		}
	}

	var wg sync.WaitGroup
	for range m.NumWorkers {
		wg.Go(func() {
			runWorker(problemCh, outputCh)
		})
	}

	// Feed problems to problem channel
	go func() {
		for _, problem := range problems {
			problemCh <- problem
		}
		close(problemCh)
	}()

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	progress, numProblems := 0, len(problems)
	for output := range outputCh {
		progress += 1
		fmt.Printf("[%3d / %3d] %s\n", progress, numProblems, output)
	}

	fmt.Println("\nTime:", time.Since(runStart).Round(time.Millisecond))
}

// Compute duration since startTime and wrap in violet color
func getDuration(start time.Time) string {
	duration := time.Since(start)
	factor, unit := 1, time.Millisecond
	if duration > 1*time.Second {
		factor = 100
	}
	if duration > 1*time.Minute {
		factor, unit = 1, time.Second
	}
	duration = duration.Round(time.Duration(factor) * unit)
	return str.Violet(fmt.Sprintf("%v", duration))
}

// Use the Worker format if can be split by |
func workerFormat(output string, worker Worker) string {
	parts := str.CleanSplit(output, "|")
	if len(parts) > 1 {
		output = fmt.Sprintf(worker.Columns(), list.ToAny(parts)...)
	}
	return output
}
