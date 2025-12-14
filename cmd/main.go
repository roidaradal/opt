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
		displayUsage(allTasks, false)
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
		displayUsage(allTasks, true)
	case "run":
		task := worker.SolverRunner{}
		runWorker(task, args)
	case "run+sol":
		task := worker.SolverRunner{DisplaySolutions: true}
		runWorker(task, args)
	case "sol.save":
		task := worker.SolutionSaver{}
		runWorker(task, args)
	case "sol.read":
		task := worker.SolutionReader{}
		runWorker(task, args)
	case "solo":
		manager := worker.Solo{}
		runManager(manager, args)
	default:
		fmt.Println("Unknown command: ", option)
	}
}

func displayUsage(taskFocus string, detailed bool) {
	if taskFocus == allTasks {
		required, autocreate := str.Red("required"), str.Blue("autocreate")
		fmt.Println(str.Yellow("opt"), " - discrete optimization tool")
		folders := [][3]string{
			{"data/", required, "Contains test case text files"},
			{"test/", required, "Contains test suite configurations"},
			{"results/", autocreate, "Stores test suite result files"},
			{"solution/", autocreate, "Stores test case solutions"},
		}
		fmt.Println("\nCurrent Directory Assumptions:")
		for _, folder := range folders {
			name, tag, description := folder[0], folder[1], folder[2]
			fmt.Printf("%-10s %-10s\t%s\n", name, tag, description)
		}
	}

	fmt.Println("\nUsage:", str.Yellow("opt"), str.Cyan("<task>"), str.Green("(option=key(:param)*)*\n"))

	// Task Choices
	required, optional := str.Red("Required:"), str.Yellow("Option:")
	requiredProblem := fmt.Sprintf("%s %s", required, str.Green("problem"))
	solverOption := fmt.Sprintf("%s %s", optional, str.Green("solver, logger"))
	tasks := [][]string{
		{"help", "Display help"},
		{"load", "Load config from file",
			fmt.Sprintf("%s %s, defaults to %s", optional, str.Violet("<configPath>"), str.Violet("./config.json")),
		},
		{"run", "Run solver on problem",
			requiredProblem,
			solverOption,
		},
		{"run+sol", "Run solver on problem, display solutions",
			requiredProblem,
			solverOption,
		},
		{"sol.save", "Run solver on problem, save solution file",
			requiredProblem,
			solverOption,
		},
		{"sol.read", "Read solution from saved file",
			requiredProblem,
		},
		{"solo", "Run worker task on dataset (one worker)",
			fmt.Sprintf("%s %s", required, str.Green("worker, data")),
			solverOption,
		},
	}

	solverOptions := []string{"problem", "solver", "logger"}
	multiOptions := []string{"worker", "data", "solver", "logger"}
	taskOptions := dict.StringListMap{
		"run":      solverOptions,
		"run+sol":  solverOptions,
		"sol.save": solverOptions,
		"sol.read": []string{"problem"},
		"solo":     multiOptions,
	}
	// TODO:
	// test			Tester

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
		{"worker", "w", "name(:param)*", "Set base worker"},
		{"data", "d", "name", "Dataset name"},
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

		if detailed {
			switch key {
			case "problem":
				displayProblemOptions()
			case "solver":
				displaySolverOptions()
			case "logger":
				displayLoggerOptions()
			case "worker":
				displayWorkerOptions()
			case "data":
				displayDatasetOptions()
			}
		}

		fmt.Println()
	}

}
