package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSecretManagerById(t *testing.T) {
	client := getClient()
	sm, err := client.Secrets().GetSecretManagerById(client.AccountId)
	require.NoError(t, err)
	require.Equal(t, sm.Id, client.AccountId)
}

func TestGetSecretManagerByName(t *testing.T) {
	client := getClient()
	sm, err := client.Secrets().GetSecretManagerByName("Harness Secrets Manager")
	require.NoError(t, err)
	require.Equal(t, sm.Id, client.AccountId)
}

func TestGetSecretManagerByName_NoManagerFound(t *testing.T) {
	client := getClient()
	sm, err := client.Secrets().GetSecretManagerByName("Bad Name")
	require.Error(t, err)
	require.Nil(t, sm)
}

func TestListSecretManagers(t *testing.T) {
	t.Skip("This endpoint requires the use of a bearer token.")
	client := getClient()
	managers, err := client.Secrets().ListSecretManagers()
	require.NoError(t, err)
	require.NotEmpty(t, managers)
}

func TestGetDefaultSecretManager(t *testing.T) {
	client := getClient()
	smId, err := client.Secrets().GetDefaultSecretManagerId()
	require.NoError(t, err)
	require.NotEmpty(t, smId)
}
