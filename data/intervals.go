package data

type Intervals struct {
	Activities []string
	Start      []float64
	End        []float64
	Weight     []float64
}

// NewIntervals loads an Intervals config
func NewIntervals(name string) *Intervals {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Intervals{
		Activities: stringList(data["activities"]),
		Start:      floatList(data["start"]),
		End:        floatList(data["end"]),
		Weight:     floatList(data["weight"]),
	}
}
