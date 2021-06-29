package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetPcfCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPcfCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetPcfCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetPcfCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPcfCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetPcfCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreatesPcfCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createPcfCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createPcfCloudProvider(name string) (*graphql.PcfCloudProvider, error) {
	c := getClient()

	input := &graphql.PcfCloudProvider{}
	input.Name = name
	input.EndpointUrl = "https://example.com"
	input.PasswordSecretId = "abc123"
	input.SkipValidation = true
	input.UserName = "foo123"

	cp, err := c.CloudProviders().CreatePcfCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
