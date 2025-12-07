package worker

import "fmt"

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
