package fn

import (
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
)

// StringList transforms the line into list of strings, separated by space
func StringList(line string) []string {
	return strings.Fields(line)
}

// IntList transforms the line into list of ints, separated by space
func IntList(line string) []int {
	return list.Map(strings.Fields(line), number.ParseInt)
}

// FloatList transforms the line into list of floats, separated by space
func FloatList(line string) []float64 {
	return list.Map(strings.Fields(line), number.ParseFloat)
}
