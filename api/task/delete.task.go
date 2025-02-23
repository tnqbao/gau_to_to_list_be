package task

import (
	"github.com/tnqbao/gau-to-do-list/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTaskById(c *gin.Context) {
	taskList, valid := utils.HandleGetListTaskFromContext(c)
	if !valid {
		return
	}

	id := c.Param("id")
	idParseToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": "invalid task id"})
		return
	}

	for i, t := range *taskList {
		if t.ID == idParseToInt {
			*taskList = append((*taskList)[:i], (*taskList)[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "task deleted!"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": "404", "error": "task not found"})
}
