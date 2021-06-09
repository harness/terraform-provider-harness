package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	secretFileId = "2WnPVgLGSZW6KbApZuxeaw"
)

func getExampleUsageScopes() *UsageScope {
	var scopes []*AppEnvScope

	scope := &AppEnvScope{
		Application: &AppScopeFilter{
			FilterType: ApplicationFilterTypes.All,
		},
		Environment: &EnvScopeFilter{
			FilterType: EnvironmentFilterTypes.NonProduction,
		},
	}
	scopes = append(scopes, scope)

	return &UsageScope{
		AppEnvScopes: scopes,
	}
}

func TestGetEncryptedFileName(t *testing.T) {
	client := getClient()
	expectedName := "secretfile"
	s, err := client.Secrets().GetEncryptedFileByName(expectedName)
	require.NoError(t, err)
	require.Equal(t, expectedName, s.Name)
}

func TestGetEncryptedFileById(t *testing.T) {
	client := getClient()
	expectedId := secretFileId
	s, err := client.Secrets().GetEncryptedFileById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
}

func TestGetWinRMCredentialById(t *testing.T) {
	client := getClient()
	expectedId := "CpiGVJOHSEmzSc39KFVBJg"
	s, err := client.Secrets().GetWinRMCredentialById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
	require.Equal(t, WinRMAuthenticationTypes.NTLM, s.AuthenticationScheme)
}

func TestGetWinRMCredentialByName(t *testing.T) {
	client := getClient()
	expectedName := "winrm_ntlm"
	s, err := client.Secrets().GetWinRMCredentialByName(expectedName)
	require.NoError(t, err)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, WinRMAuthenticationTypes.NTLM, s.AuthenticationScheme)
}

func deleteSecret(id string, secretType string) error {
	client := getClient()

	return client.Secrets().DeleteSecret(id, secretType)
}
