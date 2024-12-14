package task

import (
	"errors"
	"to-do-list/schemas"

	"github.com/google/uuid"
)

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
