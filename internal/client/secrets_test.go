package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

func createEncryptedTextSecret(name string, value string) (*EncryptedText, error) {
	client := getClient()

	input := &CreateSecretInput{
		SecretType: SecretTypes.EncryptedText,
		EncryptedText: &EncryptedTextInput{
			Name:            name,
			SecretManagerId: os.Getenv(envvar.HarnessAccountId),
			Value:           value,
		},
	}

	return client.Secrets().CreateEncryptedText(input)
}

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

func deleteEncryptedTextSecret(secret *EncryptedText) error {
	client := getClient()

	deleteInput := &DeleteSecretInput{
		SecretId:   secret.Id,
		SecretType: secret.SecretType,
	}
	return client.Secrets().DeleteSecret(deleteInput)
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
	deleteEncryptedTextSecret(testSecret)
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
	deleteEncryptedTextSecret(testSecret)
}

func TestUpdateSecret(t *testing.T) {
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
	input := &UpdateSecretInput{
		SecretId: expectedSecret.Id,
		EncryptedText: &UpdateEncryptedText{
			Name: updatedName,
		},
	}
	updatedSecret, err := client.Secrets().UpdateEncryptedText(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, updatedSecret)
	require.Equal(t, updatedName, updatedSecret.Name)

	// Cleanup
	deleteEncryptedTextSecret(updatedSecret)
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
	expectedId := "2WnPVgLGSZW6KbApZuxeaw"
	s, err := client.Secrets().GetEncryptedFileById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
}

func TestGetSSHCredentialById_SSHAuth(t *testing.T) {
	client := getClient()
	expectedId := "ujr-0fn-S1-QGtDCezcgxw"
	s, err := client.Secrets().GetSSHCredentialById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
	require.Equal(t, SSHAuthenticationTypes.SSHAuthentication, s.AuthenticationType)
	require.Nil(t, s.KerberosAuthentication)
}

func TestGetSSHCredentialById_KerberosAuth(t *testing.T) {
	client := getClient()
	expectedId := "sxE5iOK-Sciw0Uszboy_YQ"
	s, err := client.Secrets().GetSSHCredentialById(expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedId, s.Id)
	require.Equal(t, SSHAuthenticationTypes.KerberosAuthentication, s.AuthenticationType)
	require.Nil(t, s.SSHAuthentication)
}

func TestGetSSHCredentialByName_SSHAuth(t *testing.T) {
	client := getClient()
	expectedName := "somesshkey_sshauth"
	s, err := client.Secrets().GetSSHCredentialByName(expectedName)
	require.NoError(t, err)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, SSHAuthenticationTypes.SSHAuthentication, s.AuthenticationType)
	require.Nil(t, s.KerberosAuthentication)
}

func TestGetSSHCredentialByName_KerberosAuth(t *testing.T) {
	client := getClient()
	expectedName := "somesshkey_kerberosauth"
	s, err := client.Secrets().GetSSHCredentialByName(expectedName)
	require.NoError(t, err)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, SSHAuthenticationTypes.KerberosAuthentication, s.AuthenticationType)
	require.Nil(t, s.SSHAuthentication)
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

func TestCreateEncryptedText(t *testing.T) {
	client := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))

	input := &CreateSecretInput{
		EncryptedText: &EncryptedTextInput{
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
	require.Equal(t, SecretTypes.EncryptedText, s.SecretType)
	require.Equal(t, client.AccountId, s.SecretManagerId)
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)
}

func TestDeleteSecret_EncryptedText(t *testing.T) {

	// Setup
	client := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))

	// Create a secret first
	input := &CreateSecretInput{
		SecretType: SecretTypes.EncryptedText,
		EncryptedText: &EncryptedTextInput{
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
	deleteInput := &DeleteSecretInput{
		SecretId:   s.Id,
		SecretType: SecretTypes.EncryptedText,
	}

	err = client.Secrets().DeleteSecret(deleteInput)

	// Verify secret no longer exists
	require.NoError(t, err)

	s, err = client.Secrets().GetEncryptedTextById(s.Id)
	require.Error(t, err)
	require.Nil(t, s)

}
