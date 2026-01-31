package data

import "github.com/roidaradal/fn/number"

type Subsets struct {
	Universal []string
	Names     []string
	Subsets   [][]string
	Limit     int
}

// NewSubsets loads a Subsets config
func NewSubsets(name string) *Subsets {
	data, err := load(name)
	if err != nil {
		return nil
	}
	cfg := &Subsets{
		Universal: stringList(data["universal"]),
		Names:     make([]string, 0),
		Subsets:   make([][]string, 0),
		Limit:     number.ParseInt(data["limit"]),
	}
	for key, value := range parseMap(data["subsets"]) {
		cfg.Names = append(cfg.Names, key)
		cfg.Subsets = append(cfg.Subsets, stringList(value))
	}
	return cfg
}
