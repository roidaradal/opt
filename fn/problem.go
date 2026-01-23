// Package fn contains common functions used in discrete optimization problems
package fn

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
)

const separator string = "-----"

var cacheLines = make(map[[2]string][]string)

// LoadLines loads a problem test case
func LoadLines(name string) ([]string, error) {
	parts := strings.SplitN(name, ".", 2)
	problem, testcase := parts[0], parts[1]
	mainKey := [2]string{problem, testcase}

	// Check if already in cache
	if lines, ok := cacheLines[mainKey]; ok {
		return lines, nil
	}

	// Load full dataset
	path := fmt.Sprintf("data/%s.txt", problem)
	lines, err := io.ReadNonEmptyLines(path)
	if err != nil {
		return nil, err
	}
	lines = list.Filter(lines, func(line string) bool {
		return !strings.HasPrefix(line, "#") // remove comments
	})

	// Group test cases within dataset
	group := make([]string, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, separator) {
			key := [2]string{problem, group[0]}
			cacheLines[key] = group[1:]
			group = []string{}
		} else {
			group = append(group, line)
		}
	}

	lines, ok := cacheLines[mainKey]
	if !ok {
		return nil, fmt.Errorf("unknown testcase: %s", name)
	}

	return lines, nil
}
