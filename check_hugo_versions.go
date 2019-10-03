package main

import (
	"github.com/malvahq/ci-docker-hugo/api"
)

func main() {
	_, err := api.GetReleases()
	if err != nil {
		return
	}

}
