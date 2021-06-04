package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiVersionRequest(t *testing.T) {

	client := getClient()

	resp, err := client.GetAPIVersion()

	if err != nil {
		t.Error(err)
	}

	assert.True(t, resp.Resource.RuntimeInfo.Primary)
}
