package main

import (
	"github-repo-service/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	router := routes.SetupRouter()
	router.Run(":8080")
}
