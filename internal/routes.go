package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()
	router.GET("/repos", getRepos)
	router.POST("/repos", updateRepos)
}

func getRepos(c *gin.Context) {
	store, err := NewStore()
	if err != nil {
		log.Fatal(err)
		return
	}
	repos := store.Retrieve()
	_, err = json.Marshal(repos)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func updateRepos(c *gin.Context) {
	return
}
