package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphqlListCloudProviders(t *testing.T) {
	c := getClient()

	cpList, info, err := c.CloudProviderClient.ListCloudProviders(5, 0)

	require.NoError(t, err)
	require.NotNil(t, cpList)
	require.NotNil(t, info)
}
