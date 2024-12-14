package task

import "github.com/gin-gonic/gin"

type TaskHandler interface {
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
}

type TaskHandlerImpl struct {
	UseCase TaskUseCase
}

func NewTaskHandler(useCase TaskUseCase) TaskHandler {
	return &TaskHandlerImpl{UseCase: useCase}
}
