package data

import "github.com/roidaradal/fn/number"

type Bins struct {
	NumBins  int
	Capacity float64
	Weight   []float64
}

// NewBins loads a Bins config
func NewBins(name string) *Bins {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Bins{
		NumBins:  number.ParseInt(data["numBins"]),
		Capacity: number.ParseFloat(data["capacity"]),
		Weight:   floatList(data["weight"]),
	}
}
