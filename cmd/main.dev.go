//go:build dev

package main

import (
	"github.com/joho/godotenv"
	"github.com/tnqbao/gau-to-do-list/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := routes.SetupRouter()

	router.Run(":8088")
}
