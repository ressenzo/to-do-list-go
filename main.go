package main

import (
	"to-do-list/router"
	"to-do-list/task"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()

	taskRepo := task.NewTaskRepository("")
	taskUseCase := task.NewTaskUseCase(taskRepo)
	taskHandler := task.NewTaskHandler(taskUseCase)

	router.InitializeRoutes(ginRouter, taskHandler)
}
