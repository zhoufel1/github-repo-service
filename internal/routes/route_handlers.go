package routes

import (
	"github-repo-service/internal/fetch"
	"github-repo-service/internal/models"
	"github-repo-service/internal/store"
	"log"

	"github.com/gin-gonic/gin"

	"net/http"
	"os"
)

func getReposHandler(c *gin.Context) {
	db, err := store.NewReposDB()
	if err != nil {
		log.Fatal(err)
	}
	var repos []models.Repository
	if !db.IsInitialized {
		repos, err := fetchRepos()
		if err != nil {
			log.Fatal(err)
		}
		db.Initialize(repos)
	}
	repos = db.RetrieveRepos()
	c.JSON(http.StatusOK, repos)
}

func updateReposHandler(c *gin.Context) {
	db, err := store.NewReposDB()
	if err != nil {
		log.Fatal(err)
	}
	repos, err := fetchRepos()
	if err != nil {
		log.Fatal(err)
	}
	db.UpdateRepos(repos)

}

func fetchRepos() ([]models.Repository, error) {
	username := os.Getenv("GITHUB_USERNAME")
	repos, err := fetch.RequestRepos(username)
	if err != nil {
		return nil, err
	}
	return repos, nil
}
