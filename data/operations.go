package data

type Warehouse struct {
	Stores        []string
	Warehouses    []string
	Capacity      []int
	WarehouseCost []float64
	StoreCost     [][]float64
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
	}
	cfg.StoreCost = make([][]float64, len(cfg.Stores))
	for i, line := range parseList(data["storeCost"]) {
		cfg.StoreCost[i] = matrixRow(line, true)
	}
	return cfg
}
