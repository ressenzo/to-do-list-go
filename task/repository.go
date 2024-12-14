package task

import (
	"to-do-list/schemas"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTasks() ([]*schemas.Task, error)
	CreateTask(schemas.Task) (*schemas.Task, error)
}

type TaskRepositoryImpl struct {
	connection string
}

func NewTaskRepository(connection string) TaskRepository {
	return &TaskRepositoryImpl{connection: connection}
}

func (r *TaskRepositoryImpl) GetTasks() ([]*schemas.Task, error) {
	tasks := make([]*schemas.Task, 2)
	tasks[0] = &schemas.Task{Id: uuid.NewString()[:8], Description: "Task 1"}
	tasks[1] = &schemas.Task{Id: uuid.NewString()[:8], Description: "Task 2"}

	return tasks, nil
}

func (r *TaskRepositoryImpl) CreateTask(task schemas.Task) (*schemas.Task, error) {
	return &task, nil
}
