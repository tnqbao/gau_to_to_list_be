package task

import (
	"github.com/tnqbao/gau_to_do_list_be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_to_do_list_be/models"
)

func CreateTask(c *gin.Context) {

	taskList, valid := utils.HandleGetListTaskFromContext(c)
	if !valid {
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": err.Error()})
		return
	}

	if len(*taskList) > 0 {
		task.ID = (*taskList)[len(*taskList)-1].ID + 1
	} else {
		task.ID = 1
	}

	*taskList = append(*taskList, task)

	c.JSON(http.StatusCreated, gin.H{"status": "201", "message": "task created!", "taskID": task.ID})
}
