package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiVersionRequest(t *testing.T) {

	// Setup
	client := getClient()

	// Execute
	resp, err := client.GetAPIVersion()

	// Verify
	require.NoError(t, err)
	require.True(t, resp.Resource.RuntimeInfo.Primary)
}
