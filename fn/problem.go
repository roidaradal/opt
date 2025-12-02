package fn

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
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
