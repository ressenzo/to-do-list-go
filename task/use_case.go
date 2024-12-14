package task

import (
	"to-do-list/schemas"
)

type TaskUseCase interface {
	GetTasks() ([]*schemas.Task, error)
	CreateTask(description string) (*schemas.Task, error)
}

type TaskUseCaseImpl struct {
	repo TaskRepository
}

func NewTaskUseCase(repo TaskRepository) *TaskUseCaseImpl {
	return &TaskUseCaseImpl{repo: repo}
}
