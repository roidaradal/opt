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

type NoLogger struct{}
type BatchLogger struct {
	NoLogger
}
type StepsLogger struct {
	NoLogger
	DelayNanosecond int
}

func (l NoLogger) Output(args ...any) {
	// do nothing
}

func (l NoLogger) Steps(args ...any) {
	// do nothing
}

func (l NoLogger) Clear(n int) {
	// do nothing
}

func (l BatchLogger) Output(args ...any) {
	fmt.Println(args...)
}

func (l StepsLogger) Steps(args ...any) {
	fmt.Println(args...)
	if l.DelayNanosecond > 0 {
		time.Sleep(time.Duration(l.DelayNanosecond) * time.Millisecond)
	}
}

func (l StepsLogger) Clear(n int) {
	for range n {
		fmt.Print("\033[1A\033[K")
	}
}
