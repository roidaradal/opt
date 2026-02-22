package data

type ShopSchedule struct {
	Machines  []string
	Jobs      []string
	TaskTimes map[string][]int
}

// NewShopSchedule creates a new ShopSchedule config
func NewShopSchedule(name string) *ShopSchedule {
	data, err := load(name)
	if err != nil {
		return nil
	}
	taskTimes := make(map[string][]int)
	for task, value := range parseMap(data["taskTimes"]) {
		taskTimes[task] = intList(value)
	}
	return &ShopSchedule{
		Machines:  stringList(data["machines"]),
		Jobs:      stringList(data["jobs"]),
		TaskTimes: taskTimes,
	}
}
