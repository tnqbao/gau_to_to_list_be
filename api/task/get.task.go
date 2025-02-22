package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/models"
)

func GetTaskByID(c *gin.Context) {
	listTask, ok := c.MustGet("listTask").([]models.Task)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": "Internal Server Error"})
		return
	}

	id := c.Param("id")
	idParseToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": "Invalid task ID"})
		return
	}

	for _, t := range listTask {
		if t.ID == idParseToInt {
			c.JSON(http.StatusOK, gin.H{"status": 200, "task": t})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": "Task not found"})
}
