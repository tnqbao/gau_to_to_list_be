package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/api/task"
	"github.com/tnqbao/gau-to-do-list/middlewares"
	"github.com/tnqbao/gau-to-do-list/models"
	"sync"
)

var (
	listTask = &[]models.Task{}
	mu       sync.Mutex
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(func(c *gin.Context) {
		c.Set("listTask", listTask)
		c.Set("listTaskMutex", &mu)
		c.Next()
	})

	endpoint := r.Group("/")
	{
		endpoint.GET("/tasks/:id", task.GetTaskByID)
		endpoint.GET("/tasks", task.GetAllTask)

		endpoint.POST("/tasks", task.CreateTask)

		endpoint.PUT("/tasks/:id", task.UpdateTaskById)

		endpoint.DELETE("/tasks/:id", task.DeleteTaskById)
	}
	return r
}
