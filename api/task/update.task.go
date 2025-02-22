package task

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/models"
	"net/http"
	"strconv"
)

func UpdateTaskById(c *gin.Context) {
	listTask, ok := c.MustGet("listTask").([]models.Task)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": "Internal Server Error"})
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": err.Error()})
		return
	}

	id := c.Param("id")
	idParseToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": "Invalid task ID"})
		return
	}

	for i, t := range listTask {
		if t.ID == idParseToInt {
			listTask[i].Completed = task.Completed
			c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Task updated!", "task": listTask[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": "Task not found"})
}
