package data

import "github.com/roidaradal/fn/number"

type Resource struct {
	Budget float64
	Items  []string
	Count  []int
	Cost   []float64
	Value  []float64
}

// NewResource creates a new Resource config
func NewResource(name string) *Resource {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Resource{
		Budget: number.ParseFloat(data["budget"]),
		Items:  stringList(data["items"]),
		Count:  intList(data["count"]),
		Cost:   floatList(data["cost"]),
		Value:  floatList(data["value"]),
	}
}
