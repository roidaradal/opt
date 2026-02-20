package data

import (
	"fmt"

	"github.com/roidaradal/fn/number"
)

type AssignmentCfg struct {
	Workers    []string
	Tasks      []string
	Cost       [][]float64
	Teams      [][]string
	MaxPerTeam int
}

type QuadraticAssignment struct {
	Count    int
	Distance [][]float64
	Flow     [][]float64
}

// NewAssignment creates a new AssignmentCfg
func NewAssignment(name string) *AssignmentCfg {
	data, err := load(name)
	if err != nil {
		return nil
	}
	workers := stringList(data["workers"])
	tasks := stringList(data["tasks"])
	numWorkers, numTasks := len(workers), len(tasks)
	if numTasks > numWorkers {
		fmt.Println("Invalid assignment problem: more tasks than workers")
		return nil
	}
	cost := make([][]float64, 0)
	for _, line := range parseList(data["cost"]) {
		// Ensure equal number of workers and tasks
		// Add 0-cost tasks to end of list if more workers than tasks
		row := make([]float64, numWorkers)
		copy(row, matrixRow(line, true))
		cost = append(cost, row)
	}
	teams := make([][]string, 0)
	for _, line := range parseList(data["teams"]) {
		teams = append(teams, stringList(line))
	}
	return &AssignmentCfg{
		Workers:    workers,
		Tasks:      tasks,
		Cost:       cost,
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
