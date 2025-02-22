package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/models"
)

func CreateTask(c *gin.Context) {
	listTask, ok := c.MustGet("listTask").([]models.Task)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "error": "Internal Server Error"})
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": err.Error()})
		return
	}

	if len(listTask) > 0 {
		task.ID = listTask[len(listTask)-1].ID + 1
	} else {
		task.ID = 1
	}

	listTask = append(listTask, task)
	c.JSON(http.StatusCreated, gin.H{"status": "201", "message": "Task created!", "task": task})
}
