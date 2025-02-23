//go:build release

package main

import (
	"github.com/joho/godotenv"
	"github.com/tnqbao/gau-to-do-list/routes"
	"log"
)

func main() {
	err := godotenv.Load("/gau_to_do_List/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := routes.SetupRouter()

	router.Run(":8088")
}
