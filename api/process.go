package api

import "fmt"

func processReleases(releases []Release) error {
	for i := 0; i < len(releases); i++ {
		fmt.Printf("%s\n", releases[i].Name)
	}
	return nil
}
