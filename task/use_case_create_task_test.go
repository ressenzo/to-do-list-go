package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	t.Run("description is empty", func(t *testing.T) {
		description := ""
		_, err := useCase.CreateTask(description)
		expected := "description can not be empty"

		if err == nil {
			t.Errorf("expected error '%v', but didn't get", expected)
		}
	})

	t.Run("description length is greater than 50", func(t *testing.T) {
		description51Length := "111111111111111111111111111111111111111111111111111"
		_, err := useCase.CreateTask(description51Length)
		expected := "description can not be greater than 50"

		if err == nil {
			t.Errorf("expected error '%v', but didn't get", expected)
		}
	})

	t.Run("valid values", func(t *testing.T) {
		description := "Task"
		task, err := useCase.CreateTask(description)

		if err != nil {
			t.Error("should not get error")
		}

		if task.Description == "" {
			t.Error("task description should not be empty")
		}

		if task.Id == "" {
			t.Error("task id should not be empty")
		}

		if len(task.Id) != 8 {
			t.Error("task id length should be 8")
		}
	})
}
