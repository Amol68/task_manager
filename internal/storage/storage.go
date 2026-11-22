package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"tskmanager/internal/task"
)

type MyTask struct {
	*task.Task
}

func LoadTasks() ([]task.Task, error) {

	tasks := make([]task.Task, 0)

	data, err := os.ReadFile("tasks.json")

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
		return nil, err
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
	}

	return tasks, nil

}

func SaveTasks(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)

}
