package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CONTENT_TYPE_KEY string = "Content-Type"
const CONTENT_TYPE_VALUE string = "application/json"

func InternalError(c *gin.Context) {
	c.Header(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong",
		"code":    http.StatusInternalServerError,
	})
}

func SendError(c *gin.Context, code int, msg string) {
	c.Header(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
	c.JSON(code, gin.H{
		"message": msg,
		"code":    code,
	})
}

func SendSuccess(c *gin.Context, code int, data interface{}) {
	c.Header(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
	c.JSON(code, gin.H{
		"content": data,
	})
}
