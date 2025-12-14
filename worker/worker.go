// Package worker contains common discrete optimization workers
package worker

import (
	"errors"
	"fmt"

	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
)

var (
	errMissingProblem = errors.New("Undefined problem")
)

type Worker interface {
	Run(*Config) string
}

type Config struct {
	Problem   *discrete.Problem
	NewSolver SolverCreator
	Logger    Logger
	Worker    Worker
}

func (c Config) Copy() *Config {
	return &Config{
		Problem:   c.Problem,
		NewSolver: c.NewSolver,
		Logger:    c.Logger,
	}
}

// Create error message string
func errMessage(err error) string {
	return fmt.Sprintf("%s %s", str.Red("Error:"), err)
}
