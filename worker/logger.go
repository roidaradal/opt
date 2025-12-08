package worker

import "fmt"

const (
	LOG_NONE  LogLevel = iota // no logging
	LOG_BATCH                 // log iteration batch
	LOG_STEPS                 // log step-by-step
)

type LogLevel uint

type Logger interface {
	Output(args ...any)
}

type CmdLogger struct{}
type NoLogger struct{}

func (l CmdLogger) Output(args ...any) {
	fmt.Println(args...)
}

func (l NoLogger) Output(args ...any) {
	// do nothing
}

// Create new Logger based on LogLevel
func NewLogger(level LogLevel) Logger {
	switch level {
	case LOG_NONE:
		return NoLogger{}
	default:
		return CmdLogger{}
	}
}
