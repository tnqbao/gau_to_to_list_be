package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau_to_do_list_be/models"
	"net/http"
	"sync"
)

func HandleGetListTaskFromContext(c *gin.Context) (*[]models.Task, bool) {
	listTaskPtr, exists := c.Get("listTask")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "error": "Internal Server Error"})
		fmt.Println("listTask not found in context")
		return nil, false
	}

	taskList, ok := listTaskPtr.(*[]models.Task)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "error": "listTask type assertion failed"})
		fmt.Println("listTask type assertion failed")
		return nil, false
	}
	return taskList, true
}

func HandleGetMutexFromContext(c *gin.Context) (*sync.Mutex, bool) {
	mutexInterface, exists := c.Get("listTaskMutex")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "error": "Internal Server Error"})
		fmt.Println("mutex not found in context")
		return nil, false
	}

	mutex, ok := mutexInterface.(*sync.Mutex)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "500", "error": "mutex type assertion failed"})
		fmt.Println("mutex type assertion failed")
		return nil, false
	}
	return mutex, true
}
