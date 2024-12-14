package task

import (
	"net/http"
	"strings"
	"to-do-list/utils"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandlerImpl) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.BindJSON(&request); err != nil {
		utils.SendError(c, http.StatusBadRequest, "description is required")
		return
	}

	task, err := h.UseCase.CreateTask(request.Description)
	if err != nil && strings.Contains(err.Error(), "description") {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		utils.InternalError(c)
		return
	}

	utils.SendSuccess(c, http.StatusCreated, task)
}
