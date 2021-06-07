package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEncryptedSecretByName(t *testing.T) {
	client := getClient()
	expectedName := "somesecret"
	s, err := client.Secrets().GetEncryptedSecretByName("somesecret")
	require.NoError(t, err)
	require.Equal(t, s.Name, expectedName)
}
