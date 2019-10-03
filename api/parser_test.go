package api_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/malvahq/ci-docker-hugo/api"
	"github.com/stretchr/testify/assert"
)

func TestParseAPIResponse(t *testing.T) {
	body, err := ioutil.ReadFile("api_hugo_output.txt")
	assert.Nil(t, err)
	releases, err := api.ParseAPIResponse(body)
	assert.Nil(t, err)
	fmt.Printf("Length: %d\n", len(releases))
	assert.NotNil(t, releases)
	for i := 0; i < len(releases); i++ {
		fmt.Printf("%s\n", releases[i].Name)
	}
}
