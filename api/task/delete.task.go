package task

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_to_do_list_be/utils"
	"net/http"
	"strconv"
)

func DeleteTaskById(c *gin.Context) {
	taskList, valid := utils.HandleGetListTaskFromContext(c)
	if !valid {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "error": "failed to get task list"})
		return
	}

	mutex, valid := utils.HandleGetMutexFromContext(c)
	id := c.Param("id")
	idParseToInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": "invalid task id"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, t := range *taskList {
		if t.ID == idParseToInt {
			*taskList = append((*taskList)[:i], (*taskList)[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "task deleted!"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"status": "404", "error": "task not found"})
}
