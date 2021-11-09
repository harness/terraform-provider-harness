package cd

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetPhysicalDatacenterCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetPhysicalDataCEnterCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreatePhysicalDataCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetPhysicalDatacenterCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetPhysicalDatacenterCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestUpdatePhysicalDataCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedname := fmt.Sprintf("%s_updated", expectedName)
	cp, err := createPhysicalDataCenterCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)

	updateInput := &graphql.UpdatePhysicalDataCenterCloudProviderInput{
		Name: updatedname,
	}

	updatedCP, err := c.CloudProviderClient.UpdatePhysicalDataCenterCloudProvider(cp.Id, updateInput)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedname, updatedCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createPhysicalDataCenterCloudProvider(name string) (*graphql.PhysicalDataCenterCloudProvider, error) {
	c := getClient()

	input := &graphql.PhysicalDataCenterCloudProvider{}
	input.Name = name

	cp, err := c.CloudProviderClient.CreatePhysicalDataCenterCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
