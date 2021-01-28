package main

import (
	"github-repo-service/internal/routes"
	"github.com/joho/godotenv"
	"log"
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
