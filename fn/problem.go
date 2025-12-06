package fn

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/opt/a"
)

// Load problem test case
func LoadProblem(name string) ([]string, error) {
	path := fmt.Sprintf("data/%s.txt", name)
	lines, err := io.ReadLines(path)
	if err != nil {
		return nil, err
	}
	lines = list.Filter(lines, func(line string) bool {
		return !strings.HasPrefix(line, "#") && line != ""
	})
	return lines, nil
}

// Load new test case containing subsets data
func NewSubsets(name string) *a.Subsets {
	lines, err := LoadProblem(name)
	if err != nil || len(lines) < 2 {
		return nil
	}
	return a.NewSubsets(lines[0], lines[1:])
}
