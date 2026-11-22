package task

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"slices"
	"strings"
	"time"
	"unicode"
)

type Task struct {
	ID    string
	Title string

	IsCompleted bool
	CreatedAt   time.Time
}

type TaskManager struct {
	Tasks []Task
}

func hasSpecial(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) && r != '-' {
			return true
		}
	}
	return false
}

var (
	ErrInvalidTask   = errors.New("invalid task")
	ErrTaskNotFound  = errors.New("task not found")
	ErrCompletedTask = errors.New("task already completed")
)

func (m *TaskManager) AddTask(title string) (Task, error) {

	if len(strings.TrimSpace(title)) == 0 || hasSpecial(strings.ToLower(title)) {
		return Task{}, ErrInvalidTask
	}

	id := uuid.NewString()

	newTask := Task{
		ID:        id,
		Title:     title,
		CreatedAt: time.Now(),
	}
	m.Tasks = append(m.Tasks, newTask)

	return newTask, nil

}

func (m *TaskManager) ListTasks() {

	if len(m.Tasks) < 1 {
		fmt.Fprintln(os.Stderr, "No tasks yet!")
		return

	}

	for index, t := range m.Tasks {
		status := "Pending"
		if t.IsCompleted {
			status = "Completed"
		}
		fmt.Printf("\nTask %d\n", index+1)
		fmt.Println("Task ID:", t.ID)
		fmt.Println("Title:", t.Title)

		fmt.Println("Completed:", status)

		fmt.Println("Created At:", t.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func (m *TaskManager) CompleteTask(id string) error {

	if len(strings.TrimSpace(id)) == 0 || hasSpecial(id) {
		return ErrInvalidTask
	}

	idx := slices.IndexFunc(m.Tasks, func(task Task) bool {
		return task.ID == id
	})

	if idx == -1 {
		return ErrTaskNotFound
	}

	if m.Tasks[idx].IsCompleted {
		return ErrCompletedTask
	}

	m.Tasks[idx].IsCompleted = true

	return nil

}

func (m *TaskManager) DeleteTask(id string) error {

	idx := slices.IndexFunc(m.Tasks, func(task Task) bool {
		return task.ID == id
	})

	if idx == -1 {
		return errors.New("task not found")
	}

	m.Tasks = slices.Delete(m.Tasks, idx, idx+1)

	return nil

}
