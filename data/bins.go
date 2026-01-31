package data

import (
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
)

type Bins struct {
	Bins     []int
	Capacity float64
	Weight   []float64
}

// NewBins loads a Bins config
func NewBins(name string) *Bins {
	data, err := load(name)
	if err != nil {
		return nil
	}
	numBins := number.ParseInt(data["numBins"])
	return &Bins{
		Bins:     list.NumRange(1, numBins+1),
		Capacity: number.ParseFloat(data["capacity"]),
		Weight:   floatList(data["weight"]),
	}
}
