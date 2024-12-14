package task

import "to-do-list/schemas"

func (u *TaskUseCaseImpl) GetTasks() ([]*schemas.Task, error) {
	return u.repo.GetTasks()
}
