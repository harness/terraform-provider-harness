package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSecretManagerById(t *testing.T) {
	client := getClient()
	sm, err := client.SecretClient.GetSecretManagerById(client.Configuration.AccountId)
	require.NoError(t, err)
	require.Equal(t, sm.Id, client.Configuration.AccountId)
}

func TestGetSecretManagerByName(t *testing.T) {
	client := getClient()
	sm, err := client.SecretClient.GetSecretManagerByName("Harness Secrets Manager")
	require.NoError(t, err)
	require.Equal(t, sm.Id, client.Configuration.AccountId)
}

func TestGetSecretManagerByName_NoManagerFound(t *testing.T) {
	client := getClient()
	sm, err := client.SecretClient.GetSecretManagerByName("Bad Name")
	require.Error(t, err)
	require.Nil(t, sm)
}

func TestListSecretManagers(t *testing.T) {
	t.Skip("This endpoint requires the use of a bearer token.")
	client := getClient()
	managers, err := client.SecretClient.ListSecretManagers()
	require.NoError(t, err)
	require.NotEmpty(t, managers)
}

func TestGetDefaultSecretManager(t *testing.T) {
	client := getClient()
	smId, err := client.SecretClient.GetDefaultSecretManagerId()
	require.NoError(t, err)
	require.NotEmpty(t, smId)
}
