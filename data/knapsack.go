package data

import "github.com/roidaradal/fn/number"

type Knapsack struct {
	Capacity  float64
	Items     []string
	Weight    []float64
	Value     []float64
	PairBonus map[[2]string]float64
}

// NewKnapsack loads a Knapsack config
func NewKnapsack(name string) *Knapsack {
	data, err := load(name)
	if err != nil {
		return nil
	}
	cfg := &Knapsack{
		Capacity:  number.ParseFloat(data["capacity"]),
		Items:     stringList(data["items"]),
		Weight:    floatList(data["weight"]),
		Value:     floatList(data["value"]),
		PairBonus: map[[2]string]float64{},
	}
	for _, line := range parseList(data["pairBonus"]) {
		parts := stringList(line)
		if len(parts) != 3 {
			continue
		}
		pair := [2]string{parts[0], parts[1]}
		cfg.PairBonus[pair] = number.ParseFloat(parts[2])
	}
	return cfg
}
