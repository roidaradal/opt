package data

import "github.com/roidaradal/fn/number"

type Numbers struct {
	Numbers []int
	Target  int
}

// NewNumbers creates a new Sequence config
func NewNumbers(name string) *Numbers {
	data, err := load(name)
	if err != nil {
		return nil
	}
	return &Numbers{
		Numbers: intList(data["numbers"]),
		Target:  number.ParseInt(data["target"]),
	}
}

// NewN loads the N value
func NewN(name string) int {
	data, err := load(name)
	if err != nil {
		return 0
	}
	return number.ParseInt(data["n"])
}
