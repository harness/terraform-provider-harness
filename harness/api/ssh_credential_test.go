package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateSSHCredential_SSHAuthentication_inlinesshkey(t *testing.T) {

	var (
		client               = getClient()
		expectedName         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		passphraseSecretName = fmt.Sprintf("inlinesshkey_%s", utils.RandStringBytes(12))
	)

	// Create secret for ssh password
	passphraseSecret, err := createEncryptedTextSecret(passphraseSecretName, "foo")
	require.NoError(t, err)

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.SSH,
		Name:                 expectedName,
		SSHAuthentication: &graphql.SSHAuthentication{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &graphql.SSHAuthenticationMethod{
				SSHCredentialType: graphql.SSHCredentialTypes.SSHKey,
				InlineSSHKey: &graphql.InlineSSHKey{
					PassphraseSecretId: passphraseSecret.Id,
					SSHKeySecretFileId: secretFileId,
				},
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	s, err := client.Secrets().CreateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.SSHAuthentication.Port, 22)
	require.Equal(t, s.SSHAuthentication.Username, "testuser")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, graphql.ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, graphql.EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	err = client.Secrets().DeleteSecret(s.Id, s.SecretType)
	require.NoError(t, err)
}

func TestCreateSSHCredential_SSHAuthentication_serverpassword(t *testing.T) {

	var (
		client               = getClient()
		expectedName         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		passphraseSecretName = fmt.Sprintf("inlinesshkey_%s", utils.RandStringBytes(12))
	)

	// Create secret for ssh password
	passphraseSecret, err := createEncryptedTextSecret(passphraseSecretName, "foo")
	require.NoError(t, err)

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.SSH,
		Name:                 expectedName,
		SSHAuthentication: &graphql.SSHAuthentication{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &graphql.SSHAuthenticationMethod{
				SSHCredentialType: graphql.SSHCredentialTypes.Password,
				ServerPassword: &graphql.SSHPassword{
					PasswordSecretId: passphraseSecret.Id,
				},
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	s, err := client.Secrets().CreateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.SSHAuthentication.Port, 22)
	require.Equal(t, s.SSHAuthentication.Username, "testuser")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, graphql.ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, graphql.EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	err = client.Secrets().DeleteSecret(s.Id, s.SecretType)
	require.NoError(t, err)
}

func TestCreateSSHCredential_SSHAuthentication_keyfile(t *testing.T) {

	var (
		client               = getClient()
		expectedName         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		passphraseSecretName = fmt.Sprintf("inlinesshkey_%s", utils.RandStringBytes(12))
	)

	// Create secret for ssh password
	passphraseSecret, err := createEncryptedTextSecret(passphraseSecretName, "foo")
	require.NoError(t, err)

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.SSH,
		Name:                 expectedName,
		SSHAuthentication: &graphql.SSHAuthentication{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &graphql.SSHAuthenticationMethod{
				SSHCredentialType: graphql.SSHCredentialTypes.SSHKeyFilePath,
				SSHKeyFile: &graphql.SSHKeyFile{
					PassphraseSecretId: passphraseSecret.Id,
					Path:               "/some/path",
				},
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	s, err := client.Secrets().CreateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.SSHAuthentication.Port, 22)
	require.Equal(t, s.SSHAuthentication.Username, "testuser")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, graphql.ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, graphql.EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	err = client.Secrets().DeleteSecret(s.Id, s.SecretType)
	require.NoError(t, err)
}

func TestCreateSSHCredential_KerberosAuth_password(t *testing.T) {

	var (
		client               = getClient()
		expectedName         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		passphraseSecretName = fmt.Sprintf("inlinesshkey_%s", utils.RandStringBytes(12))
	)

	// Create secret for ssh password
	passphraseSecret, err := createEncryptedTextSecret(passphraseSecretName, "foo")
	require.NoError(t, err)

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.Kerberos,
		Name:                 expectedName,
		KerberosAuthentication: &graphql.KerberosAuthentication{
			Port:      9292,
			Principal: "someuser",
			Realm:     "somerealm",
			TGTGenerationMethod: &graphql.TGTGenerationMethod{
				KerberosPassword: &graphql.KerberosPassword{
					PasswordSecretId: passphraseSecret.Id,
				},
				TGTGenerationUsing: graphql.TGTGenerationUsingOptions.Password,
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	s, err := client.Secrets().CreateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.KerberosAuthentication.Port, 9292)
	require.Equal(t, s.KerberosAuthentication.Principal, "someuser")
	require.Equal(t, s.KerberosAuthentication.Realm, "somerealm")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, graphql.ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, graphql.EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	err = client.Secrets().DeleteSecret(s.Id, s.SecretType)
	require.NoError(t, err)
}

func TestCreateSSHCredential_KerberosAuth_keytabfile(t *testing.T) {

	var (
		client       = getClient()
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	)

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.Kerberos,
		Name:                 expectedName,
		KerberosAuthentication: &graphql.KerberosAuthentication{
			Port:      9292,
			Principal: "someuser",
			Realm:     "somerealm",
			TGTGenerationMethod: &graphql.TGTGenerationMethod{
				TGTGenerationUsing: graphql.TGTGenerationUsingOptions.KeyTabFile,
				KeyTabFile: &graphql.KeyTabFile{
					FilePath: "/some/path",
				},
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	s, err := client.Secrets().CreateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, graphql.SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.KerberosAuthentication.Port, 9292)
	require.Equal(t, s.KerberosAuthentication.Principal, "someuser")
	require.Equal(t, s.KerberosAuthentication.Realm, "somerealm")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, graphql.ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, graphql.EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	err = client.Secrets().DeleteSecret(s.Id, s.SecretType)
	require.NoError(t, err)
}

func TestUpdateSSHCredential(t *testing.T) {
	// Setup
	initialName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", initialName)

	// Create a secret for us to fetch by name
	expectedSecret, err := createSSHCredential_sshAuth(initialName)
	require.NoError(t, err)
	require.NotNil(t, expectedSecret)
	require.Equal(t, initialName, expectedSecret.Name)

	// Update secret
	client := getClient()
	input := &graphql.SSHCredential{
		Name: updatedName,
	}

	updatedSecret, err := client.Secrets().UpdateSSHCredential(expectedSecret.Id, input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, updatedSecret)
	require.Equal(t, updatedName, updatedSecret.Name)

	// Cleanup
	err = deleteSecret(updatedSecret.Id, updatedSecret.SecretType)
	require.NoError(t, err)
}

func TestGetSSHCredentialById_SSHAuth(t *testing.T) {

	// Setup
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	cred, err := createSSHCredential_sshAuth(name)
	require.NoError(t, err)
	require.NotNil(t, cred)

	client := getClient()
	s, err := client.Secrets().GetSSHCredentialById(cred.Id)

	// Verify
	require.NoError(t, err)
	require.Equal(t, cred.Id, s.Id)
	require.Equal(t, graphql.SSHAuthenticationTypes.SSHAuthentication, s.AuthenticationType)
}

func TestGetSSHCredentialById_kerberosAuth(t *testing.T) {

	// Setup
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	cred, err := createSSHCredential_kerberosAuth(name)
	require.NoError(t, err)
	require.NotNil(t, cred)

	client := getClient()
	s, err := client.Secrets().GetSSHCredentialById(cred.Id)

	// Verify
	require.NoError(t, err)
	require.Equal(t, cred.Id, s.Id)
	require.Equal(t, graphql.SSHAuthenticationTypes.KerberosAuthentication, s.AuthenticationType)
}

func TestGetSSHCredentialByName_SSHAuth(t *testing.T) {
	// Setup
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	cred, err := createSSHCredential_sshAuth(name)
	require.NoError(t, err)
	require.NotNil(t, cred)

	client := getClient()
	s, err := client.Secrets().GetSSHCredentialByName(cred.Name)

	// Verify
	require.NoError(t, err)
	require.Equal(t, cred.Id, s.Id)
	require.Equal(t, graphql.SSHAuthenticationTypes.SSHAuthentication, s.AuthenticationType)
}

func TestGetSSHCredentialByName_KerberosAuth(t *testing.T) {
	// Setup
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	cred, err := createSSHCredential_kerberosAuth(name)
	require.NoError(t, err)
	require.NotNil(t, cred)

	client := getClient()
	s, err := client.Secrets().GetSSHCredentialByName(cred.Name)

	// Verify
	require.NoError(t, err)
	require.Equal(t, cred.Id, s.Id)
	require.Equal(t, graphql.SSHAuthenticationTypes.KerberosAuthentication, s.AuthenticationType)
}

func createSSHCredential_sshAuth(name string) (*graphql.SSHCredential, error) {

	passphraseSecret, err := createEncryptedTextSecret(name, "foo")
	if err != nil {
		return nil, err
	}

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.SSH,
		Name:                 name,
		SSHAuthentication: &graphql.SSHAuthentication{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &graphql.SSHAuthenticationMethod{
				SSHCredentialType: graphql.SSHCredentialTypes.Password,
				ServerPassword: &graphql.SSHPassword{
					PasswordSecretId: passphraseSecret.Id,
				},
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	client := getClient()
	s, err := client.Secrets().CreateSSHCredential(input)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func createSSHCredential_kerberosAuth(name string) (*graphql.SSHCredential, error) {

	passphraseSecret, err := createEncryptedTextSecret(name, "foo")
	if err != nil {
		return nil, err
	}

	input := &graphql.SSHCredential{
		AuthenticationScheme: graphql.SSHAuthenticationSchemes.Kerberos,
		Name:                 name,
		KerberosAuthentication: &graphql.KerberosAuthentication{
			Port:      9292,
			Principal: "someuser",
			Realm:     "somerealm",
			TGTGenerationMethod: &graphql.TGTGenerationMethod{
				TGTGenerationUsing: graphql.TGTGenerationUsingOptions.Password,
				KerberosPassword: &graphql.KerberosPassword{
					PasswordSecretId: passphraseSecret.Id,
				},
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	client := getClient()
	s, err := client.Secrets().CreateSSHCredential(input)
	if err != nil {
		return nil, err
	}

	return s, nil
}
