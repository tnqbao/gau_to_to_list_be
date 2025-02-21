package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/api/task"
	"github.com/tnqbao/gau-to-do-list/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	endpoint := r.Group("/")
	{
		endpoint.POST("/task", task.CreateTask)
	}
	return r
}
