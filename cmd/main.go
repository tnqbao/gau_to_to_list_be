//go:build !dev

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tnqbao/gau-to-do-list/models"
	"github.com/tnqbao/gau-to-do-list/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/gau_blog/.env.blog")
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
