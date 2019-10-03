package main_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAPIResponse(t *testing.T) {
	_, err := ioutil.ReadFile("api_hugo_output.txt")
	assert.Nil(t, err)
}
