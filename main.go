package main

import (
	// "fmt"
	"log"
	// "os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	// username, present := os.LookupEnv("GITHUB_USERNAME")
	// if !present {
	// 	log.Fatal("GITHUB_USERNAME not set ")
	// }
	// repos, err := fetch.RequestRepos(username)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// storeHandler, err := store.NewRepoStore()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// results := storeHandler.Retrieve()
	// for _, repo := range results {
	// 	fmt.Println(repo.FullName)
	// }
}
