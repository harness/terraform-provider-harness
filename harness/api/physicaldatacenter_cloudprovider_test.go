package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetPhysicalDatacenterCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetPhysicalDataCEnterCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreatePhysicalDataCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetPhysicalDatacenterCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetPhysicalDatacenterCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createPhysicalDataCenterCloudProvider(name string) (*graphql.PhysicalDataCenterCloudProvider, error) {
	c := getClient()

	input := &graphql.PhysicalDataCenterCloudProvider{}
	input.Name = name

	cp, err := c.CloudProviders().CreatePhysicalDataCenterCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
