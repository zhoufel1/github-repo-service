package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/zhoufel1/github-projects-service/internal/models"
	"net/http"
)

func requestRepos(username string) ([]models.Repository, error) {
	requestURL := fmt.Sprintf("http://api.github.com/users/%s/repos", username)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var repos []models.Repository
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}
	return repos, nil
}
