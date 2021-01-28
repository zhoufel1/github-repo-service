package routes

import (
	"github-repo-service/internal/fetch"
	"github-repo-service/internal/models"
	"github-repo-service/internal/store"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func getReposHandler(c *gin.Context) {
	db, err := store.NewRepoDB()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.ResponseJSON{
				Code:    500,
				Message: "Failed to connect to database.",
			})
		log.Fatal(err)
	}
	var repos []models.Repository
	if !db.IsInitialized {
		repos, err := fetchRepos()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.ResponseJSON{
					Code:    500,
					Message: "Failed to fetch repos.",
				})
			log.Fatal(err)
		}
		db.Initialize(repos)
	}
	repos = db.RetrieveRepos()
	c.JSON(http.StatusOK, repos)
}

func updateReposHandler(c *gin.Context) {
	db, err := store.NewRepoDB()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.ResponseJSON{
				Code:    500,
				Message: "Failed to connect to database.",
			})
		log.Fatal(err)
	}
	repos, err := fetchRepos()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.ResponseJSON{
				Code:    500,
				Message: "Failed to fetch repos.",
			})
		log.Fatal(err)
	}
	db.UpdateRepos(repos)
	c.JSON(
		http.StatusOK,
		models.ResponseJSON{
			Code:    200,
			Message: "Repositories successfully updated.",
		})

}

func fetchRepos() ([]models.Repository, error) {
	username := os.Getenv("GITHUB_USERNAME")
	repos, err := fetch.RequestRepos(username)
	if err != nil {
		return nil, err
	}
	return repos, nil
}
