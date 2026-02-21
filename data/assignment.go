package data

import (
	"fmt"

	"github.com/roidaradal/fn/number"
)

type AssignmentCfg struct {
	Workers    []string
	Capacity   []float64
	Tasks      []string
	Cost       [][]float64
	Value      [][]float64
	Teams      [][]string
	MaxPerTeam int
}

type QuadraticAssignment struct {
	Count    int
	Distance [][]float64
	Flow     [][]float64
}

type Weapons struct {
	Weapons []string
	Count   []int
	Targets []string
	Value   []float64
	Chance  [][]float64
}

// NewAssignment creates a new AssignmentCfg
func NewAssignment(name string, mustBeEqual bool) *AssignmentCfg {
	data, err := load(name)
	if err != nil {
		return nil
	}
	workers := stringList(data["workers"])
	tasks := stringList(data["tasks"])
	numWorkers, numTasks := len(workers), len(tasks)
	if numTasks > numWorkers && mustBeEqual {
		fmt.Println("Invalid assignment problem: more tasks than workers")
		return nil
	}
	cost := make([][]float64, 0)
	for _, line := range parseList(data["cost"]) {
		costRow := matrixRow(line, true)
		if mustBeEqual {
			// Ensure equal number of workers and tasks
			// Add 0-cost tasks to end of list if more workers than tasks
			row := make([]float64, numWorkers)
			copy(row, costRow)
			cost = append(cost, row)
		} else {
			cost = append(cost, costRow)
		}
	}
	value := make([][]float64, 0)
	for _, line := range parseList(data["value"]) {
		value = append(value, matrixRow(line, true))
	}
	teams := make([][]string, 0)
	for _, line := range parseList(data["teams"]) {
		teams = append(teams, stringList(line))
	}
	return &AssignmentCfg{
		Workers:    workers,
		Capacity:   floatList(data["capacity"]),
		Tasks:      tasks,
		Cost:       cost,
		Value:      value,
		Teams:      teams,
		MaxPerTeam: number.ParseInt(data["maxPerTeam"]),
	}
}

// NewQuadraticAssignment creates a new Quadratic Assignment config
func NewQuadraticAssignment(name string) *QuadraticAssignment {
	data, err := load(name)
	if err != nil {
		return nil
	}
	distance := make([][]float64, 0)
	for _, line := range parseList(data["distance"]) {
		distance = append(distance, matrixRow(line, false))
	}
	flow := make([][]float64, 0)
	for _, line := range parseList(data["flow"]) {
		flow = append(flow, matrixRow(line, false))
	}
	return &QuadraticAssignment{
		Count:    number.ParseInt(data["count"]),
		Distance: distance,
		Flow:     flow,
	}
}

// NewWeapons creates a new Weapons config
func NewWeapons(name string) *Weapons {
	data, err := load(name)
	if err != nil {
		return nil
	}
	chance := make([][]float64, 0)
	for _, line := range parseList(data["chance"]) {
		chance = append(chance, matrixRow(line, true))
	}
	return &Weapons{
		Weapons: stringList(data["weapons"]),
		Targets: stringList(data["targets"]),
		Count:   intList(data["count"]),
		Value:   floatList(data["value"]),
		Chance:  chance,
	}
}
