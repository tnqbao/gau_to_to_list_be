package task

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/models"
	"github.com/tnqbao/gau-to-do-list/utils"
	"net/http"
	"strconv"
)

func UpdateTaskById(c *gin.Context) {
	taskList, valid := utils.HandleGetListTaskFromContext(c)
	if !valid {
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": "invalid task id"})
		return
	}

	for i, t := range *taskList {
		if t.ID == id {
			if (*taskList)[i].Completed == task.Completed {
				c.JSON(http.StatusOK, gin.H{"status": 200, "message": "nothing changed", "task": (*taskList)[i]})
				return
			}
			(*taskList)[i].Completed = task.Completed
			c.JSON(http.StatusOK, gin.H{"status": 200, "message": "task updated!", "task": (*taskList)[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": "task not found"})
}
