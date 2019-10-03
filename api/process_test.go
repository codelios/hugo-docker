package api_test

import (
	"testing"

	"github.com/malvahq/ci-docker-hugo/api"
	"github.com/stretchr/testify/assert"
)

func TestGetDockerBranchName(t *testing.T) {
	assert.Equal(t, "0.32.1", api.GetDockerBranchName("v0.32.1"))
}
