package task

import (
	"encoding/json"
	"github.com/tnqbao/gau-to-do-list/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTaskByID(c *gin.Context) {
	taskList, valid := utils.HandleGetListTaskFromContext(c)
	if !valid {
		return
	}

	id := c.Param("id")
	idParseToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": "invalid task id"})
		return
	}

	for _, t := range *taskList {
		if t.ID == idParseToInt {
			c.JSON(http.StatusOK, gin.H{"status": 200, "task": t})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": "task not found"})
}

func GetAllTask(c *gin.Context) {
	taskList, valid := utils.HandleGetListTaskFromContext(c)
	if !valid {
		return
	}

	jsonData, err := json.Marshal(taskList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": "Failed to marshal task list"})
		return
	}
	c.JSON(http.StatusOK, json.RawMessage(jsonData))
}
