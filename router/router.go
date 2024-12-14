package router

import (
	"to-do-list/task"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, t task.TaskHandler) {
	basePath := "/api/v1"

	routes := r.Group(basePath)
	{
		routes.GET("/tasks", t.GetTasks)
		routes.POST("/tasks", t.CreateTask)
	}

	r.Run(":8080")
}
