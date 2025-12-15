package worker

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
)

type SolutionReader struct {
	HorizontalOutput bool
}

// Read solution from file and display stats
func (r SolutionReader) Run(cfg *Config) string {
	problem := cfg.Problem
	if problem == nil {
		return errMessage(errMissingProblem)
	}

	items := [][2]string{
		{"Problem", problem.Name},
	}

	path := fmt.Sprintf("solution/%s.txt", problem.Name)
	lines, err := io.ReadLines(path)
	if err != nil {
		return errMessage(err)
	}
	items = append(items, [2]string{"Best Score", lines[0]})

	details := []string{""}
	if problem.SolutionCoreFn == nil {
		items = append(items, [2]string{"Best Solutions", number.Comma(len(lines[1:]))})
		items = append(items, [2]string{"Core Solutions", ""})
	} else {
		var key string
		var count int
		coreSolutions := make(dict.IntMap)
		for _, line := range lines[1:] {
			if strings.HasPrefix(line, "+ ") {
				if key != "" {
					coreSolutions[key] = count
					count = 0
				}
				key = strings.TrimPrefix(line, "+ ")
			} else if strings.HasPrefix(line, "- ") {
				count += 1
			}
		}
		if key != "" && count > 0 {
			coreSolutions[key] = count // last group
		}
		items = append(items, [2]string{"Best Solutions", number.Comma(list.Sum(dict.Values(coreSolutions)))})
		items = append(items, [2]string{"Core Solutions", number.Comma(len(coreSolutions))})
		entries := dict.Entries(coreSolutions)
		slices.SortFunc(entries, func(a, b dict.Entry[string, int]) int {
			// Sort by descending count
			return cmp.Compare(b.Value, a.Value)
		})
		for _, e := range entries {
			details = append(details, fmt.Sprintf("%7s : %s", number.Comma(e.Value), e.Key))
		}
	}

	if r.HorizontalOutput {
		out := list.Map(items, func(pair [2]string) string {
			return pair[1] // return right side of pair
		})
		return strings.Join(out, "|")
	}

	lengths := list.Map(items, func(pair [2]string) int {
		return len(pair[0])
	})
	template := fmt.Sprintf("%%-%ds : %%s", slices.Max(lengths))

	out := make([]string, 0)
	for _, pair := range items {
		key, value := pair[0], pair[1]
		if value == "" {
			continue // skip blank values
		}
		out = append(out, fmt.Sprintf(template, key, value))
	}
	out = append(out, details...)
	return strings.Join(out, "\n")
}

// SolutionReader columns
func (r SolutionReader) Columns() string {
	// ProblemName | BestScore | BestSolutions | CoreSolutions
	return "%-20s %7s %7s %7s"
}
