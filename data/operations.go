package data

import "github.com/roidaradal/fn/number"

type Warehouse struct {
	Stores        []string
	Warehouses    []string
	Capacity      []int
	WarehouseCost []float64
	StoreCost     [][]float64
	Count         int
	Distance      [][]float64
}

type NurseSchedule struct {
	Nurses              []string
	Days                []string
	Shifts              []string
	ShiftMin            []int
	ShiftMax            []int
	MaxConsecutiveShift int
	MaxTotalShift       int
	DailyLimit          int
	PreferShifts        map[string][]string
	PreferDays          map[string][]string
}

// NewWarehouse creates a new Warehouse config
func NewWarehouse(name string) *Warehouse {
	data, err := load(name)
	if err != nil {
		return nil
	}
	cfg := &Warehouse{
		Stores:        stringList(data["stores"]),
		Warehouses:    stringList(data["warehouses"]),
		Capacity:      intList(data["capacity"]),
		WarehouseCost: floatList(data["warehouseCost"]),
		Count:         number.ParseInt(data["count"]),
	}
	cfg.StoreCost = make([][]float64, len(cfg.Stores))
	for i, line := range parseList(data["storeCost"]) {
		cfg.StoreCost[i] = matrixRow(line, true)
	}
	cfg.Distance = make([][]float64, len(cfg.Warehouses))
	for i, line := range parseList(data["distance"]) {
		cfg.Distance[i] = matrixRow(line, true)
	}
	return cfg
}

// NewNurseSchedule creates a new NurseSchedule config
func NewNurseSchedule(name string) *NurseSchedule {
	data, err := load(name)
	if err != nil {
		return nil
	}
	preferShifts := make(map[string][]string)
	for nurse, value := range parseMap(data["preferShifts"]) {
		preferShifts[nurse] = stringList(value)
	}
	preferDays := make(map[string][]string)
	for nurse, value := range parseMap(data["preferDays"]) {
		preferDays[nurse] = stringList(value)
	}
	return &NurseSchedule{
		Nurses:              stringList(data["nurses"]),
		Days:                stringList(data["days"]),
		Shifts:              stringList(data["shifts"]),
		ShiftMin:            intList(data["shiftMin"]),
		ShiftMax:            intList(data["shiftMax"]),
		MaxConsecutiveShift: number.ParseInt(data["maxConsecutiveShift"]),
		MaxTotalShift:       number.ParseInt(data["maxTotalShift"]),
		DailyLimit:          number.ParseInt(data["dailyLimit"]),
		PreferShifts:        preferShifts,
		PreferDays:          preferDays,
	}
}
