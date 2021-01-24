package main

import (
	"fmt"
	"github.com/zhoufel1/github-projects-service/internal/fetch"
	"log"
)

func main() {
	repos, err := fetch.RequestRepos("zhoufel1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repos[1].FullName)
}
