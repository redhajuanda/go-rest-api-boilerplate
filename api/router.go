package api

import (
	"github.com/gin-gonic/gin"
	"github.com/redhajuanda/rest-api-boilerplate/api/task"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		task.Route(api)
	}
}
