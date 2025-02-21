//go:build !dev

package main

import (
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

	router.Run(":8088")
}
