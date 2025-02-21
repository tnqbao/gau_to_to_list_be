//go:build dev

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tnqbao/gau-to-do-list/models"
	"github.com/tnqbao/gau-to-do-list/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := routes.SetupRouter()

	listTask := []models.Task{}

	router.Use(func(c *gin.Context) {
		c.Set("listTask", listTask)
		c.Next()
	})

	router.Run(":8088")
}
