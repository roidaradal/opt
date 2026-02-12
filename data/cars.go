package data

import (
	"slices"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
)

type Cars struct {
	Cars      []string
	Sequence  []string
	CarColors []string
	MaxShift  int
}

type CarSequence struct {
	OptionMax    []int
	OptionWindow []int
	Cars         []string
	CarOptions   map[string][]bool
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

// NewCarSequence loads a CarSequence config
func NewCarSequence(name string) *CarSequence {
	data, err := load(name)
	if err != nil {
		return nil
	}
	carTypes := stringList(data["cars"])
	carCounts := intList(data["carCount"])
	if len(carTypes) != len(carCounts) {
		return nil
	}
	allCars := make([]string, 0)
	for i, car := range carTypes {
		allCars = append(allCars, slices.Repeat([]string{car}, carCounts[i])...)
	}
	options := stringList(data["options"])
	optionIdx := list.IndexMap(options)
	carOptions := make(map[string][]bool)
	for car, v := range parseMap(data["carOptions"]) {
		flags := make([]bool, len(options))
		for _, option := range stringList(v) {
			flags[optionIdx[option]] = true
		}
		carOptions[car] = flags
	}
	return &CarSequence{
		OptionMax:    intList(data["optionMax"]),
		OptionWindow: intList(data["optionWindow"]),
		Cars:         allCars,
		CarOptions:   carOptions,
	}
}
