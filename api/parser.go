package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// HugoURL represents an URL to get hugo releases
	HugoURL = "https://api.github.com/repos/gohugoio/hugo/releases"
)

// Release abstracts a small set of fields from the actual API
type Release struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

func ParseAPIResponse(bytes []byte) ([]Release, error) {
	var releases []Release
	json.Unmarshal(bytes, &releases)
	return releases, nil
}

func getAPIResponse() ([]byte, error) {
	resp, err := http.Get(HugoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

// GetReleases retrieves the release of hugo library
func GetReleases() ([]Release, error) {
	response, err := getAPIResponse()
	if err != nil {
		return nil, err
	}
	return ParseAPIResponse(response)
}
