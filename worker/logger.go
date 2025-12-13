package worker

import (
	"fmt"
	"time"
)

type Logger interface {
	Output(args ...any)
	Steps(args ...any)
	Clear(n int)
}

// All other loggers embed the QuietLogger, so that the default
// behavior for undefined Logger methods is to do nothing
type QuietLogger struct{}

type BatchLogger struct {
	QuietLogger
}

type StepsLogger struct {
	QuietLogger
	DelayNanosecond int
}

func (l QuietLogger) Output(args ...any) {
	// do nothing
}

func (l QuietLogger) Steps(args ...any) {
	// do nothing
}

func (l QuietLogger) Clear(n int) {
	// do nothing
}

func (l BatchLogger) Output(args ...any) {
	fmt.Println(args...)
}

func (l StepsLogger) Steps(args ...any) {
	fmt.Println(args...)
	if l.DelayNanosecond > 0 {
		time.Sleep(time.Duration(l.DelayNanosecond) * time.Nanosecond)
	}
}

func (l StepsLogger) Clear(n int) {
	for range n {
		fmt.Print("\033[1A\033[K")
	}
}
