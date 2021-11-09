package cd

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetSpotInstCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, secret, err := createSpotInstCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetSpotInstCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}

func TestGetSpotInstCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, secret, err := createSpotInstCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetSpotInstCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}

func TestCreateSpotInstCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, secret, err := createSpotInstCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}

func TestUpdateSpotInstCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", expectedName)

	cp, secret, err := createSpotInstCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	input := &graphql.UpdateSpotInstCloudProviderInst{
		AccountId:     helpers.TestEnvVars.SpotInstAccountId.Get(),
		TokenSecretId: secret.Id,
		Name:          updatedName,
	}

	updatedCP, err := c.CloudProviderClient.UpdateSpotInstCloudProvider(cp.Id, input)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedName, updatedCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.SecretClient.DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}

func createSpotInstCloudProvider(name string) (*graphql.SpotInstCloudProvider, *graphql.EncryptedText, error) {
	c := getClient()

	secret, err := createEncryptedTextSecret(name, helpers.TestEnvVars.SpotInstToken.Get())
	if err != nil {
		return nil, nil, err
	}

	input := &graphql.SpotInstCloudProvider{}
	input.Name = name
	input.AccountId = helpers.TestEnvVars.SpotInstAccountId.Get()
	input.TokenSecretId = secret.Id

	cp, err := c.CloudProviderClient.CreateSpotInstCloudProvider(input)
	if err != nil {
		return nil, nil, err
	}

	return cp, secret, nil
}
