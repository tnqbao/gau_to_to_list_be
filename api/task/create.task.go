package task

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/models"
)

func CreateTask(c *gin.Context) {
	listTask := c.MustGet("listTask").([]models.Task)
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	listTask = append(listTask, task)
	c.JSON(200, gin.H{"message": "Task created!", "task": task})

}
