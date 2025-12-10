package worker

import (
	"fmt"
	"time"
)

type Logger interface {
	Output(args ...any)
	Steps(args ...any)
}

type NoLogger struct{}
type BatchLogger struct{}
type StepsLogger struct {
	DelayMs int
}

func (l NoLogger) Output(args ...any) {
	// do nothing
}

func (l NoLogger) Steps(args ...any) {
	// do nothing
}

func (l BatchLogger) Output(args ...any) {
	fmt.Println(args...)
}

func (l BatchLogger) Steps(args ...any) {
	// do nothing
}

func (l StepsLogger) Output(args ...any) {
	// do nothing
}

func (l StepsLogger) Steps(args ...any) {
	fmt.Print("\033[1A\033[K")
	fmt.Println(args...)
	if l.DelayMs > 0 {
		time.Sleep(time.Duration(l.DelayMs) * time.Millisecond)
	}
}
