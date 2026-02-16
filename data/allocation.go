package data

import (
	"slices"

	"github.com/roidaradal/fn/number"
)

type Resource struct {
	Budget float64
	Items  []string
	Count  []int
	Cost   []float64
	Value  []float64
}

type Scene struct {
	NumDays     int
	DayMin      []int
	DayMax      []int
	DailyCost   map[string]float64
	Scenes      []string
	SceneActors map[string][]string
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

// NewScene creates a new Scene config
func NewScene(name string) *Scene {
	data, err := load(name)
	if err != nil {
		return nil
	}
	dailyCost := make(map[string]float64)
	for actor, value := range parseMap(data["actors"]) {
		dailyCost[actor] = number.ParseFloat(value)
	}
	scenes := make([]string, 0)
	sceneActors := make(map[string][]string)
	for scene, value := range parseMap(data["scenes"]) {
		sceneActors[scene] = stringList(value)
		scenes = append(scenes, scene)
	}
	slices.Sort(scenes)
	return &Scene{
		NumDays:     number.ParseInt(data["numDays"]),
		DayMin:      intList(data["dayMin"]),
		DayMax:      intList(data["dayMax"]),
		DailyCost:   dailyCost,
		Scenes:      scenes,
		SceneActors: sceneActors,
	}
}
