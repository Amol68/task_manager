package task_test

import (
	"errors"
	"testing"
	"tskmanager/internal/task"
)

// func TestTaskEmpty(t *testing.T) {
// 	manager := task.TaskManager{}
// 	_, err := manager.AddTask("     ")
// 	if err != nil {
// 		t.Fatal("AddTask() expected a task but got an empty title, got nil")

// 	}
// }


func TestTaksToAdd(t *testing.T) {
	
	tests := []struct {
		name    string
		task    string
		wantErr error
		}{
			
			{name: "empty task", task: "", wantErr: task.ErrInvalidTask},
			{name: "invalid characters", task: "^%$&*(", wantErr: task.ErrInvalidTask},
		}
		
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
			manager := task.TaskManager{}
			_, err := manager.AddTask(tt.task)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Addtask(%v) is invalid task.Cannot be added", tt.task)
			}

		})
	}

}
