package data

type FlowShop struct {
	Machines  []string
	Jobs      []string
	TaskTimes map[string][]int
}

// NewFlowShop creates a new FlowShop config
func NewFlowShop(name string) *FlowShop {
	data, err := load(name)
	if err != nil {
		return nil
	}
	taskTimes := make(map[string][]int)
	for task, value := range parseMap(data["taskTimes"]) {
		taskTimes[task] = intList(value)
	}
	return &FlowShop{
		Machines:  stringList(data["machines"]),
		Jobs:      stringList(data["jobs"]),
		TaskTimes: taskTimes,
	}
}
