package task

import (
	"github.com/gin-gonic/gin"
)

// Routers for task endpoint
func Route(r *gin.RouterGroup) {
	task := r.Group("/task")
	{
		task.GET("/", Ping)
	}
}
