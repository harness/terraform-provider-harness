package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetAzureCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetAzureCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetAzureCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetAzureCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateAzureCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	secret, err := c.Secrets().GetEncryptedTextByName(expectedName)
	require.NoError(t, err)
	c.Secrets().DeleteSecret(secret.Id, secret.SecretType)
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

	updatedCP, err := c.CloudProviders().UpdateAzureCloudProvider(cp.Id, input)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedName, updatedCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createAzureCloudProvider(name string) (*graphql.AzureCloudProvider, error) {

	c := getClient()
	expectedName := name

	secret, err := createEncryptedTextSecret(expectedName, TestEnvVars.AzureClientSecret.Get())
	if err != nil {
		return nil, err
	}

	input := &graphql.AzureCloudProvider{}
	input.Name = expectedName
	input.ClientId = TestEnvVars.AzureClientId.Get()
	input.KeySecretId = secret.Id
	input.TenantId = TestEnvVars.AzureTenantId.Get()

	return c.CloudProviders().CreateAzureCloudProvider(input)
}
