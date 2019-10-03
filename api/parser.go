package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	// HugoURL represents an URL to get hugo releases
	HugoURL = "https://api.github.com/repos/gohugoio/hugo/releases"

	// PerPage represents the number of results returned per page
	PerPage = 50

	// NiceInterval represents the time to sleep before making a sucessive request
	NiceInterval = 10 * time.Second

	// ReleasesJSON represents the JSON file that contains the release info
	ReleasesJSON = "hugo_releases.json"
)

var (
	// ErrNoReleasesRetrieved represents an error when the API returns 0 releases
	ErrNoReleasesRetrieved = errors.New("No releases were retrieved from remote api. Probably rate limited  ? ")
)

// Release abstracts a small set of fields from the actual API
type Release struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

// ParseAPIResponse parses API response to retrieve release information
func ParseAPIResponse(bytes []byte) ([]Release, error) {
	var releases []Release
	json.Unmarshal(bytes, &releases)
	return releases, nil
}

func populateQueryParameters(q url.Values, page int) error {
	// https://developer.github.com/v3/#pagination
	q.Add("page", strconv.FormatInt(int64(page), 10))
	q.Add("per_page", strconv.FormatInt(PerPage, 10))
	return nil
}

func getAPIResponse(page int) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", HugoURL, nil)
	if err != nil {
		return nil, nil, err
	}
	q := req.URL.Query()
	populateQueryParameters(q, page)
	req.URL.RawQuery = q.Encode()
	finalURL := req.URL.String()
	fmt.Printf("Querying %s\n", finalURL)
	req.Header.Set("User-Agent", "Golang_GitHub_ReleaseChecker/1.0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		fmt.Printf("About to sleep for %s\n", NiceInterval)
		time.Sleep(NiceInterval)
	}()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return resp, body, nil
}

// GetReleases retrieves the release of hugo library
func getReleases(page int) ([]Release, error) {
	response, content, err := getAPIResponse(page)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", response.Header)
	return ParseAPIResponse(content)
}

// GetAllReleases retrieves all the release of hugo library
func GetAllReleases() ([]Release, error) {
	allReleases := make([]Release, 0)
	page := 0
	for true {
		releases, err := getReleases(page)
		if err != nil {
			return nil, err
		}
		if len(releases) == 0 {
			break
		}
		allReleases = append(allReleases, releases...)
		page++
	}
	return allReleases, nil
}

func checkReleasesJSON() bool {
	if _, err := os.Stat(ReleasesJSON); err == nil {
		// TODO: Check if timestamp is out of date as well
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func updateLatestReleaseInfo() error {
	releases, err := GetAllReleases()
	if err != nil {
		return err
	}
	if len(releases) == 0 {
		return ErrNoReleasesRetrieved
	}
	jsonString, err := json.Marshal(releases)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ReleasesJSON, jsonString, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func getLocalReleases() ([]Release, error) {
	content, err := ioutil.ReadFile(ReleasesJSON)
	if err != nil {
		return nil, err
	}
	return ParseAPIResponse(content)
}


// UpdateReleaseInfo updates the release information once a day from remote
func UpdateReleaseInfo() error {
	exists := checkReleasesJSON()
	if !exists {
		err := updateLatestReleaseInfo()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Refreshing local data")
	}
	releases, err := getLocalReleases()
	if err != nil {
		return err
	}
	fmt.Printf("Retrieved %d releases\n", len(releases))
	err = processReleases(releases)
	if err != nil {
		return err
	}
	return nil
}
