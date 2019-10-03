package main

import (
	"fmt"

	"github.com/malvahq/ci-docker-hugo/api"
)

func main() {
	err := api.UpdateReleaseInfo()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

}
