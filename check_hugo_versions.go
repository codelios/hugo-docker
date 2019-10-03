package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// HugoURL represents an URL to get hugo releases
	HugoURL = "https://api.github.com/repos/gohugoio/hugo/releases"
)

func checkHugoVersions() error {
	resp, err := http.Get(HugoURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func main() {
	err := checkHugoVersions()
	if err != nil {
		return
	}
}
