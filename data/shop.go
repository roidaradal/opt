package data

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
)

type TimeRange [2]int

type ShopSchedule struct {
	Machines    []string
	Jobs        []string
	TaskTimes   map[string][]int
	JobTasks    map[string][]Task
	MaxMakespan int
}

type Task struct {
	Job      string
	Name     string
	Machine  string
	Duration int
}

type SlotSched struct {
	Start int
	End   int
	Name  string
}

func (t TimeRange) Tuple() (int, int) {
	return t[0], t[1]
}

func (t Task) GetDuration() int {
	return t.Duration
}

func SortBySchedStart(a, b SlotSched) int {
	return cmp.Compare(a.Start, b.Start)
}

func SortByStartTime(a, b TimeRange) int {
	return cmp.Compare(a[0], b[0])
}

// NewShopSchedule creates a new ShopSchedule config
func NewShopSchedule(name string) *ShopSchedule {
	data, err := load(name)
	if err != nil {
		return nil
	}
	taskTimes := make(map[string][]int)
	for job, value := range parseMap(data["taskTimes"]) {
		taskTimes[job] = intList(value)
	}
	jobTasks := make(map[string][]Task)
	totalDuration := 0
	for job, value := range parseMap(data["tasks"]) {
		jobTasks[job] = list.IndexedMap(stringList(value), func(i int, text string) Task {
			parts := strings.Split(text, ":")
			return Task{
				Job:      job,
				Name:     fmt.Sprintf("J%s_T%d", job, i),
				Machine:  parts[0],
				Duration: number.ParseInt(parts[1]),
			}
		})
		totalDuration += list.SumOf(jobTasks[job], func(task Task) int {
			return task.Duration
		})
	}
	return &ShopSchedule{
		Machines:    stringList(data["machines"]),
		Jobs:        stringList(data["jobs"]),
		TaskTimes:   taskTimes,
		JobTasks:    jobTasks,
		MaxMakespan: totalDuration,
	}
}

func (s *ShopSchedule) GetTasks() []Task {
	tasks := make([]Task, 0)
	for _, job := range s.Jobs {
		for _, task := range s.JobTasks[job] {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
