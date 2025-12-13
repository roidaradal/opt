package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/problem"
	"github.com/roidaradal/opt/solver/brute"
	"github.com/roidaradal/opt/worker"
)

var (
	defaultLogger        worker.Logger        = worker.BatchLogger{}
	defaultSolverCreator worker.SolverCreator = brute.NewLinearSolver
)

// Load args from file, defaults to config.json
func loadFileArgs(args []string) []string {
	path := "config.json"
	if len(args) > 1 {
		path = args[1]
	}
	cfg, err := io.ReadJSONMap[string](path)
	if err != nil {
		fmt.Println(str.Red("Error:"), err)
		return nil
	}
	if dict.NoKey(cfg, "task") {
		fmt.Println(str.Red("Error:"), "Undefined task from config")
		return nil
	}

	args = make([]string, 0, len(cfg))
	args = append(args, cfg["task"])
	delete(cfg, "task")
	for key, value := range cfg {
		args = append(args, fmt.Sprintf("%s=%s", key, value))
	}
	return args
}

// Create new problem
func newProblem(value string) (*discrete.Problem, error) {
	parts := str.CleanSplit(value, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid problem string: %q", value)
	}

	name, n := parts[0], parts[1]
	if dict.NoKey(problem.Creator, name) {
		return nil, fmt.Errorf("Unknown problem: %q", name)
	}

	p := problem.Creator[name](number.ParseInt(n))
	if p == nil {
		return nil, fmt.Errorf("Unknown test case: %q", value)
	}

	return p, nil
}

// Display problem options
func displayProblemOptions() {
	entries, err := os.ReadDir("data/")
	if err != nil {
		fmt.Println(str.Red("Error:"), err)
		return
	}

	testCases := make(map[string][]int)
	letters := regexp.MustCompile("[a-zA-Z]+")
	for _, e := range entries {
		filename := e.Name()
		if !strings.HasSuffix(filename, ".txt") {
			continue // skip non-text files
		}
		filename = strings.Split(filename, ".")[0]
		name := letters.FindString(filename)
		if dict.NoKey(problem.Creator, name) {
			continue // skip unknown problem
		}
		n := number.ParseInt(strings.TrimPrefix(filename, name))
		testCases[name] = append(testCases[name], n)
	}
	names := dict.Keys(testCases)
	names = append(names, problem.NoFiles...)
	slices.Sort(names)
	for _, name := range names {
		if dict.HasKey(testCases, name) {
			first := slices.Min(testCases[name])
			last := slices.Max(testCases[name])
			fmt.Printf("%s:[%d,%d]\n", name, first, last)
		} else {
			fmt.Printf("%s:n\n", name)
		}
	}
}

// Create new SolverCreator, defaults to LinearBruteForce
func newSolverCreator(value string) worker.SolverCreator {
	newSolver := defaultSolverCreator
	parts := str.CleanSplit(value, ":")
	name := parts[0]
	switch name {
	case "LinearBruteForce":
		newSolver = brute.NewLinearSolver
	case "ConcurrentBruteForce":
		numWorkers := 10 // default
		if len(parts) > 1 {
			numWorkers = max(numWorkers, number.ParseInt(parts[1]))
		}
		newSolver = brute.NewConcurrentSolver(numWorkers)
	default:
		fmt.Printf("Unknown solver %q, using default...", name)
	}
	return newSolver
}

// Display solver options
func displaySolverOptions() {
	options := []string{
		"LinearBruteForce",
		"ConcurrentBruteForce:<numWorkers>",
	}
	for _, option := range options {
		fmt.Printf("%s\n", option)
	}
}

// Create new logger, defaults to BatchLogger
func newLogger(value string) worker.Logger {
	logger := defaultLogger
	parts := str.CleanSplit(value, ":")
	name := strings.ToLower(parts[0])
	switch name {
	case "quiet":
		logger = worker.QuietLogger{}
	case "batch":
		logger = worker.BatchLogger{}
	case "steps":
		delay := 0
		if len(parts) > 1 {
			delay = number.ParseInt(parts[1])
		}
		logger = worker.StepsLogger{DelayNanosecond: delay}
	default:
		fmt.Printf("Unknown logger %q, using default...", name)
	}
	return logger
}

func displayLoggerOptions() {
	options := []string{
		"quiet",
		"batch",
		"steps:<delay>",
	}
	for _, option := range options {
		fmt.Printf("%s\n", option)
	}
}
