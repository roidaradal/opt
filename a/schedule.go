package a

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
)

type TimeRange [2]int

type ShopSchedCfg struct {
	Machines    []string
	Jobs        []*Job
	Tasks       []*Task
	MaxMakespan int
}

type Job struct {
	Name  string
	Tasks []*Task
}

type Task struct {
	Name     string
	JobName  string
	Machine  string
	Duration int
}

type SlotSched struct {
	Start int
	End   int
	Name  string
}

// Create new Job
func NewJob(jobName string, line string) *Job {
	job := &Job{
		Name:  jobName,
		Tasks: make([]*Task, 0),
	}
	for taskIndex, taskName := range strings.Fields(line) {
		task := NewTask(taskName, jobName, taskIndex)
		if task != nil {
			job.Tasks = append(job.Tasks, task)
		}
	}
	return job
}

// Create new Task
func NewTask(taskName, jobName string, taskIndex int) *Task {
	parts := str.CleanSplit(taskName, ":")
	if len(parts) != 2 {
		return nil
	}
	return &Task{
		Name:     fmt.Sprintf("J%s_T%d", jobName, taskIndex),
		JobName:  jobName,
		Machine:  parts[0],
		Duration: number.ParseInt(parts[1]),
	}
}

// Unpack time range start, end
func (t TimeRange) Tuple() (int, int) {
	return t[0], t[1]
}

// Job string representation
func (j Job) String() string {
	return j.Name
}

// Return Job's total duration
func (j Job) TotalDuration() int {
	return list.Sum(list.Map(j.Tasks, taskDuration))
}

// Return Job's task margins before and after given task ID
func (j Job) TaskMargins(taskID int) (before int, after int) {
	before = list.Sum(list.Map(j.Tasks[:taskID], taskDuration))
	after = list.Sum(list.Map(j.Tasks[taskID+1:], taskDuration))
	return before, after
}

// Create machine:duration task string
func TaskString(machine, duration string) string {
	return fmt.Sprintf("%s:%s", machine, duration)
}

// Comparison function: sort time ranges by start time
func SortByStartTime(a, b TimeRange) int {
	return cmp.Compare(a[0], b[0])
}

// Comparison function: sort slot scheds by start time
func SortBySchedStart(a, b SlotSched) int {
	return cmp.Compare(a.Start, b.Start)
}

// Utility function for returning task duration
func taskDuration(task *Task) int {
	return task.Duration
}
