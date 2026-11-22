package task_test

import (
	"errors"
	"testing"
	"tskmanager/internal/task"
)

// func TestTaksToAdd(t *testing.T) {

// 	tests := []struct {
// 		name    string
// 		task    string
// 		wantErr error
// 	}{

// 		{name: "empty task", task: "", wantErr: task.ErrInvalidTask},
// 		{name: "invalid characters", task: "^%$&*(", wantErr: task.ErrInvalidTask},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			manager := task.TaskManager{}
// 			_, err := manager.AddTask(tt.task)

// 			if !errors.Is(err, tt.wantErr) {
// 				t.Fatalf("Addtask(%v) is invalid task.Cannot be added", tt.task)
// 			}

// 		})
// 	}

// }

func TestTaksToComplete(t *testing.T) {

	tests := []struct {
		name        string
		taskId      string
		wantErr     error
		isCompleted bool
	}{

		{name: "empty task", taskId: "  ", isCompleted: false, wantErr: task.ErrInvalidTask},
		{name: "invalid characters", taskId: "^%$&*(", isCompleted: false, wantErr: task.ErrInvalidTask},
		{name: "already completed task", taskId: "d6aec1e2-5046-4a4d-bff3-e9683da00bd", isCompleted: true, wantErr: task.ErrCompletedTask},
		{name: "complete a valid task", taskId: "a4642be7-88e0-4bb4-a60e-deea5f76ee57", isCompleted: true, wantErr: nil},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			manager := task.TaskManager{
				Tasks: []task.Task{
					{
						ID:          tt.taskId,
						Title:       "Learn Go",
						IsCompleted: tt.isCompleted,
					},
				},
			};

			err := manager.CompleteTask(tt.taskId)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("complete(%v) is invalid task.Cannot be marked as completed", tt.taskId)
			}

		})
	}

}
