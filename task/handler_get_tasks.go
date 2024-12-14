package task

import (
	"net/http"
	"to-do-list/utils"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandlerImpl) GetTasks(c *gin.Context) {
	tasks, err := h.UseCase.GetTasks()

	if err != nil {
		utils.InternalError(c)
		return
	}

	utils.SendSuccess(c, http.StatusOK, tasks)
}
