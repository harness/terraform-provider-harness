package cd

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetAzureCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetAzureCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetAzureCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetAzureCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateAzureCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	secret, err := c.SecretClient.GetEncryptedTextByName(expectedName)
	require.NoError(t, err)
	err = c.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}

func TestUpdateAzureCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", expectedName)

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)

	input := &graphql.UpdateAzureCloudProviderInput{
		Name: updatedName,
	}

	updatedCP, err := c.CloudProviderClient.UpdateAzureCloudProvider(cp.Id, input)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedName, updatedCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createAzureCloudProvider(name string) (*graphql.AzureCloudProvider, error) {

	c := getClient()
	expectedName := name

	secret, err := createEncryptedTextSecret(expectedName, helpers.TestEnvVars.AzureClientSecret.Get())
	if err != nil {
		return nil, err
	}

	input := &graphql.AzureCloudProvider{}
	input.Name = expectedName
	input.ClientId = helpers.TestEnvVars.AzureClientId.Get()
	input.KeySecretId = secret.Id
	input.TenantId = helpers.TestEnvVars.AzureTenantId.Get()

	return c.CloudProviderClient.CreateAzureCloudProvider(input)
}
