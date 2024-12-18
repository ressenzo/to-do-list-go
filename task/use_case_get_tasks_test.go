package task

import "testing"

func TestGetTasks(t *testing.T) {
	t.Run("get all", func(t *testing.T) {

		tasks, err := useCase.GetTasks()

		if err != nil {
			t.Error("should not get error")
		}

		if len(tasks) == 0 {
			t.Error("should get tasks")
		}
	})
}
