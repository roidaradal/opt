package data

import "github.com/roidaradal/fn/number"

type Cars struct {
	Cars      []string
	Sequence  []string
	CarColors []string
	MaxShift  int
}

// NewCars loads a Cars config
func NewCars(name string) *Cars {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Cars{
		Cars:      stringList(data["cars"]),
		Sequence:  stringList(data["sequence"]),
		CarColors: stringList(data["carColors"]),
		MaxShift:  number.ParseInt(data["maxshift"]),
	}
}
