package main

import (
	"fmt"

	"github.com/malvahq/ci-docker-hugo/api"
)

func main() {
	releases, err := api.GetAllReleases()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Retrieved %d releases\n", len(releases))
	for i := 0; i < len(releases); i++ {
		fmt.Printf("%s\n", releases[i].Name)
	}

}
