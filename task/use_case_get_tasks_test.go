package task

import "testing"

func Test_GetTasks_GetAll(t *testing.T) {
	tasks, err := useCase.GetTasks()

	if err != nil {
		t.Error("should not get error")
	}

	if len(tasks) == 0 {
		t.Error("should get tasks")
	}
}
