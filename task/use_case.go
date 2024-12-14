package task

import (
	"errors"
	"to-do-list/schemas"

	"github.com/google/uuid"
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

func (u *TaskUseCaseImpl) GetTasks() ([]*schemas.Task, error) {
	return u.repo.GetTasks()
}

func (u *TaskUseCaseImpl) CreateTask(description string) (*schemas.Task, error) {
	if description == "" {
		return nil, errors.New("description can not be empty")
	}

	if len(description) > 50 {
		return nil, errors.New("description can not be greater than 50")
	}

	task := schemas.Task{Id: uuid.NewString()[:8], Description: description}
	return u.repo.CreateTask(task)
}
