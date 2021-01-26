package routes

import (
	"log"

	"github-repo-service/internal/fetch"
	"github-repo-service/internal/models"
	"github-repo-service/internal/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/repos", getRepos)
	router.POST("/repos", updateRepos)
	return router
}

func getRepos(c *gin.Context) {
	store, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
		return
	}
	var repos []models.Repository
	if !store.IsInitiallyFetched() {
		username := os.Getenv("GITHUB_USERNAME")
		repos, err := fetch.RequestRepos(username)
		if err != nil {
			log.Fatal(err)
			return
		}
		store.Create(repos)
		store.SetFetched()
		c.JSON(http.StatusOK, repos)
	} else {
		repos = store.Retrieve()
		c.JSON(http.StatusOK, repos)
	}
}

func updateRepos(c *gin.Context) {
	return
}
