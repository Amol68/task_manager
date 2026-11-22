package main

import (
	"fmt"
	"os"
	"strings"
	"tskmanager/internal/storage"
	"tskmanager/internal/task"
)

func main() {
	args := os.Args

	if len(args) < 2 {

		fmt.Fprintln(os.Stderr, "No command Found")
		os.Exit(2)
	}
	command := strings.ToLower(args[1])

	if (command == "task-add" || command == "task-delete" || command == "task-complete") && len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Invalid command")
		os.Exit(2)
	}

	manager := task.TaskManager{}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load tasks")
		os.Exit(2)
	}
	manager.Tasks = tasks

	switch command {
	case "task-add":
		_, err := manager.AddTask(args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to add task.", err)
			os.Exit(1)
		}

		SaveErr := storage.SaveTasks(manager.Tasks)

		if SaveErr != nil {
			fmt.Fprintln(os.Stderr, "Failed to save new task", SaveErr)
			os.Exit(1)
		}

	case "task-list":
		manager.ListTasks()
	case "task-complete":
		err := manager.CompleteTask(args[2])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		saveErr := storage.SaveTasks(manager.Tasks)

		if saveErr != nil {
			fmt.Fprintln(os.Stderr, "Failed to save new task", saveErr)
			os.Exit(1)
		}

	case "task-delete":
		err := manager.DeleteTask(args[2])

		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to delete task", err)
			os.Exit(1)

		}

		SaveErr := storage.SaveTasks(manager.Tasks)

		if SaveErr != nil {
			fmt.Fprintln(os.Stderr, "Failed to update tasks", SaveErr)
			os.Exit(1)
		}

	case "task-help":
		fmt.Println("displaying help")
	default:
		fmt.Fprintln(os.Stderr, "Unknown Command", command)
		os.Exit(2)
	}
}
