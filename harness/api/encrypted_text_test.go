package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/envvar"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateEncryptedText(t *testing.T) {
	client := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))

	input := &graphql.CreateSecretInput{
		EncryptedText: &graphql.EncryptedTextInput{
			Name:            expectedName,
			SecretManagerId: os.Getenv(envvar.HarnessAccountId),
			Value:           "someval",
			UsageScope:      getExampleUsageScopes(),
		},
	}

	s, err := client.Secrets().CreateEncryptedText(input)
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.SecretTypes.EncryptedText, s.SecretType)
	require.Equal(t, client.AccountId, s.SecretManagerId)
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, graphql.ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, graphql.EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)
}

func TestDeleteSecret_EncryptedText(t *testing.T) {

	// Setup
	client := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))

	// Create a secret first
	input := &graphql.CreateSecretInput{
		SecretType: graphql.SecretTypes.EncryptedText,
		EncryptedText: &graphql.EncryptedTextInput{
			Name:            expectedName,
			SecretManagerId: os.Getenv(envvar.HarnessAccountId),
			Value:           "someval",
		},
	}

	// Verify secret created successfully
	s, err := client.Secrets().CreateEncryptedText(input)
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)

	// Delete secret
	err = client.Secrets().DeleteSecret(s.Id, s.SecretType)

	// Verify secret no longer exists
	require.NoError(t, err)

	s, err = client.Secrets().GetEncryptedTextById(s.Id)
	require.Error(t, err)
	require.Nil(t, s)

}

func TestGetEncryptedTextByName(t *testing.T) {

	// Setup
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	expectedValue := t.Name()

	// Create a secret for us to fetch by name
	expectedSecret, err := createEncryptedTextSecret(expectedName, expectedValue)
	require.NoError(t, err)
	require.NotNil(t, expectedSecret)

	// Get secret
	client := getClient()
	testSecret, err := client.Secrets().GetEncryptedTextByName(expectedName)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, testSecret)
	require.Equal(t, expectedName, testSecret.Name)
	require.Equal(t, expectedSecret.Id, testSecret.Id)
	require.Nil(t, testSecret.UsageScope)

	// Cleanup
	err = deleteSecret(testSecret.Id, testSecret.SecretType)
	require.NoError(t, err)
}

func TestGetEncryptedTextById(t *testing.T) {
	// Setup
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	expectedValue := t.Name()

	// Create a secret for us to fetch by name
	expectedSecret, err := createEncryptedTextSecret(expectedName, expectedValue)
	require.NoError(t, err)
	require.NotNil(t, expectedSecret)

	// Get secret
	client := getClient()
	testSecret, err := client.Secrets().GetEncryptedTextById(expectedSecret.Id)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, testSecret)
	require.Equal(t, expectedName, testSecret.Name)
	require.Equal(t, expectedSecret.Id, testSecret.Id)

	// Cleanup
	err = deleteSecret(testSecret.Id, testSecret.SecretType)
	require.NoError(t, err)
}

func TestUpdateEncryptedTextSecret(t *testing.T) {
	// Setup
	initialName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", initialName)
	expectedValue := t.Name()

	// Create a secret for us to fetch by name
	expectedSecret, err := createEncryptedTextSecret(initialName, expectedValue)
	require.NoError(t, err)
	require.NotNil(t, expectedSecret)
	require.Equal(t, initialName, expectedSecret.Name)

	// Update secret
	client := getClient()
	input := &graphql.UpdateSecretInput{
		SecretId: expectedSecret.Id,
		EncryptedText: &graphql.UpdateEncryptedText{
			Name: updatedName,
		},
	}
	updatedSecret, err := client.Secrets().UpdateEncryptedText(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, updatedSecret)
	require.Equal(t, updatedName, updatedSecret.Name)

	// Cleanup
	err = deleteSecret(updatedSecret.Id, updatedSecret.SecretType)
	require.NoError(t, err)
}

func createEncryptedTextSecret(name string, value string) (*graphql.EncryptedText, error) {
	client := getClient()

	input := &graphql.CreateSecretInput{
		SecretType: graphql.SecretTypes.EncryptedText,
		EncryptedText: &graphql.EncryptedTextInput{
			Name:            name,
			SecretManagerId: os.Getenv(envvar.HarnessAccountId),
			Value:           value,
		},
	}

	return client.Secrets().CreateEncryptedText(input)
}
