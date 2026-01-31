package data

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/io"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
)

// Keeps the read test cases in cache
var cacheData = make(map[[2]string]dict.StringMap)

const (
	modeOutside = iota
	modeInside
	modeList
	modeMap
)

const (
	commentStart   string = "#"
	entrySeparator string = ":"
	listSeparator  string = "|"
)

// Load the problem test case
func load(name string) (dict.StringMap, error) {
	parts := strings.SplitN(name, ".", 2)
	problemName, mainTestCase := parts[0], parts[1]
	mainKey := [2]string{problemName, mainTestCase}

	// Check if already in cache
	if data, ok := cacheData[mainKey]; ok {
		return data, nil
	}

	// Load test case file
	path := filepath.Join("data", problemName+".txt")
	lines, err := io.ReadNonEmptyLines(path)
	if err != nil {
		return nil, err
	}
	lines = list.Filter(lines, func(line string) bool {
		return !strings.HasPrefix(line, commentStart) // remove comments
	})

	// Read lines to load data
	testCase, currentKey := "", ""
	group := make([]string, 0)
	data := make(dict.StringMap)
	mode := modeOutside
	for _, line := range lines {
		braceEnd := strings.HasSuffix(line, "{")
		bracketEnd := strings.HasSuffix(line, "[")
		isEntry := strings.Contains(line, entrySeparator)
		if mode == modeOutside && braceEnd {
			// Start of test case block: read testCase name and change to modeInside
			testCase = strings.TrimSpace(strings.TrimSuffix(line, "{"))
			mode = modeInside
		} else if mode == modeInside && line == "}" {
			// End of test case block: save data to cache, reset data map, and change to modeOutside
			cacheData[[2]string{problemName, testCase}] = data
			data = make(dict.StringMap)
			mode = modeOutside
		} else if mode == modeInside && isEntry && braceEnd {
			// Inside test case block, start of map value
			currentKey = str.CleanSplit(line, entrySeparator)[0]
			mode = modeMap
		} else if mode == modeInside && isEntry && bracketEnd {
			// Inside test case block, start of list value
			currentKey = str.CleanSplit(line, entrySeparator)[0]
			mode = modeList
		} else if mode == modeInside && isEntry {
			// Inside test case block, normal key-value pair
			parts := str.CleanSplitN(line, ":", 2)
			data[parts[0]] = parts[1]
		} else if (mode == modeList && line == "]") || (mode == modeMap && line == "}") {
			// End of list/map value: add group list to data, reset group, and change to modeInside
			data[currentKey] = strings.Join(group, listSeparator)
			group = make([]string, 0)
			mode = modeInside
		} else if mode == modeList || mode == modeMap {
			// Inside list/map value: add line to group
			group = append(group, line)
		}
	}

	// Find the given test case name from cache
	problemData, ok := cacheData[mainKey]
	if !ok {
		return nil, fmt.Errorf("unknown test case: %s", name)
	}
	return problemData, nil
}

// Transform line into list of strings, separated by space
func stringList(line string) []string {
	return strings.Fields(line)
}

// Transform line into list of ints, separated by space
func intList(line string) []int {
	return list.Map(strings.Fields(line), number.ParseInt)
}

// Transform line into list of floats, separated by space
func floatList(line string) []float64 {
	return list.Map(strings.Fields(line), number.ParseFloat)
}

// Parse the list value into list of strings
func parseList(value string) []string {
	return str.CleanSplit(value, listSeparator)
}

// Parse the map value into StringMap
func parseMap(value string) dict.StringMap {
	data := make(dict.StringMap)
	for _, entry := range parseList(value) {
		parts := str.CleanSplitN(entry, entrySeparator, 2)
		data[parts[0]] = parts[1]
	}
	return data
}
