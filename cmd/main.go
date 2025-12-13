package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/worker"
)

const allTasks string = "*"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		displayUsage(allTasks)
		return
	}

	// Load config from file
	if strings.ToLower(args[0]) == "load" {
		args = loadFileArgs(args)
		if args == nil {
			return
		}
	}

	option := strings.ToLower(args[0])
	switch option {
	case "?", "help":
		displayUsage(allTasks)
	case "run":
		task := worker.RunReporter{}
		runReporterTask(task, args)
	case "run+sol":
		task := worker.RunReporter{WithSolutions: true}
		runReporterTask(task, args)
	case "sol.save":
		task := worker.SolutionReporter{}
		runReporterTask(task, args)
	default:
		fmt.Println("Unknown command: ", option)
	}

}

func displayUsage(taskFocus string) {
	if taskFocus == allTasks {
		fmt.Println(str.Yellow("opt"), " - discrete optimization tool")
		folders := [][3]string{
			{"data/", str.Red("required"), "Contains test case text files"},
			{"test/", str.Red("required"), "Contains test suite configurations"},
			{"results/", str.Blue("autocreate"), "Stores test suite result files"},
			{"solution/", str.Blue("autocreate"), "Stores test case solutions"},
		}
		fmt.Println("\nCurrent Directory Assumptions:")
		for _, folder := range folders {
			name, tag, description := folder[0], folder[1], folder[2]
			fmt.Printf("%-10s %-10s\t%s\n", name, tag, description)
		}
	}

	fmt.Println("\nUsage:", str.Yellow("opt"), str.Cyan("<task>"), str.Green("(option=key(:param)*)*\n"))

	// Task Choices
	reporterRequired := fmt.Sprintf("%s %s", str.Red("Required:"), str.Green("problem"))
	reporterOption := fmt.Sprintf("%s %s", str.Yellow("Option:"), str.Green("solver, logger"))
	tasks := [][]string{
		{"help", "Display help"},
		{"load", "Load config from file",
			fmt.Sprintf("%s %s, defaults to %s", str.Yellow("Option:"), str.Violet("<configPath>"), str.Violet("./config.json")),
		},
		{"run", "Run solver on problem",
			reporterRequired,
			reporterOption,
		},
		{"run+sol", "Run solver on problem, display solutions",
			reporterRequired,
			reporterOption,
		},
		{"sol.save", "Run solver on problem, save solution file",
			reporterRequired,
			reporterOption,
		},
	}

	reporterOptions := []string{"problem", "solver", "logger"}
	taskOptions := dict.StringListMap{
		"run":      reporterOptions,
		"run+sol":  reporterOptions,
		"sol.save": reporterOptions,
	}
	// TODO:
	// sol.read 	SolutionReader
	// test			Tester
	// multi 		MultiTasker

	if taskFocus == allTasks {
		fmt.Println("Task Choices:", len(tasks))
	}
	for _, task := range tasks {
		key, description := task[0], task[1]
		if taskFocus != allTasks && taskFocus != key {
			continue
		}
		fmt.Printf("%-20s %s\n", str.Cyan(key), description)
		for _, detail := range task[2:] {
			fmt.Printf("%-13s - %s\n", "", detail)
		}
	}
	fmt.Println()

	// Options Choices
	options := [][4]string{
		{"problem", "p", "name:n", "Set problem test case"},
		{"solver", "s", "name(:param)*", "Set solver"},
		{"logger", "l", "name(:param)*", "Set logger"},
	}
	if taskFocus == allTasks {
		fmt.Println("Options:", len(options))
	}
	for _, option := range options {
		key, shortcut, value, description := option[0], option[1], option[2], option[3]
		if taskFocus != allTasks && !slices.Contains(taskOptions[taskFocus], key) {
			continue
		}
		command := fmt.Sprintf("%s=%s", key, value)
		command2 := fmt.Sprintf("%s=%s", shortcut, value)
		fmt.Printf("%-30s %s\n", str.Green(command), description)
		fmt.Println(str.Green(command2))

		switch key {
		case "problem":
			displayProblemOptions()
		case "solver":
			displaySolverOptions()
		case "logger":
			displayLoggerOptions()
		}

		fmt.Println()
	}

}
