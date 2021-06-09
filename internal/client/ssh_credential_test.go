package client

import (
	"fmt"
	"testing"

	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
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

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.SSH,
		Name:                 expectedName,
		SSHAuthentication: &SSHAuthenticationInput{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &SSHAuthenticationMethod{
				SSHCredentialType: SSHCredentialTypes.SSHKey,
				InlineSSHKey: &InlineSSHKey{
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
	require.Equal(t, SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.SSHAuthentication.Port, 22)
	require.Equal(t, s.SSHAuthentication.Username, "testuser")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	client.Secrets().DeleteSecret(s.Id, s.SecretType)
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

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.SSH,
		Name:                 expectedName,
		SSHAuthentication: &SSHAuthenticationInput{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &SSHAuthenticationMethod{
				SSHCredentialType: SSHCredentialTypes.Password,
				ServerPassword: &SSHPassword{
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
	require.Equal(t, SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.SSHAuthentication.Port, 22)
	require.Equal(t, s.SSHAuthentication.Username, "testuser")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	client.Secrets().DeleteSecret(s.Id, s.SecretType)
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

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.SSH,
		Name:                 expectedName,
		SSHAuthentication: &SSHAuthenticationInput{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &SSHAuthenticationMethod{
				SSHCredentialType: SSHCredentialTypes.SSHKeyFilePath,
				SSHKeyFile: &SSHKeyFile{
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
	require.Equal(t, SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.SSHAuthentication.Port, 22)
	require.Equal(t, s.SSHAuthentication.Username, "testuser")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	client.Secrets().DeleteSecret(s.Id, s.SecretType)
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

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.Kerberos,
		Name:                 expectedName,
		KerberosAuthentication: &KerberosAuthentication{
			Port:      9292,
			Principal: "someuser",
			Realm:     "somerealm",
			TGTGenerationMethod: &TGTGenerationMethod{
				KerberosPassword: &KerberosPassword{
					PasswordSecretId: passphraseSecret.Id,
				},
				TGTGenerationUsing: TGTGenerationUsingOptions.Password,
			},
		},
		UsageScope: getExampleUsageScopes(),
	}

	s, err := client.Secrets().CreateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, expectedName, s.Name)
	require.Equal(t, SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.KerberosAuthentication.Port, 9292)
	require.Equal(t, s.KerberosAuthentication.Principal, "someuser")
	require.Equal(t, s.KerberosAuthentication.Realm, "somerealm")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	client.Secrets().DeleteSecret(s.Id, s.SecretType)
}

func TestCreateSSHCredential_KerberosAuth_keytabfile(t *testing.T) {

	var (
		client       = getClient()
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	)

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.Kerberos,
		Name:                 expectedName,
		KerberosAuthentication: &KerberosAuthentication{
			Port:      9292,
			Principal: "someuser",
			Realm:     "somerealm",
			TGTGenerationMethod: &TGTGenerationMethod{
				TGTGenerationUsing: TGTGenerationUsingOptions.KeyTabFile,
				KeyTabFile: &KeyTabFile{
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
	require.Equal(t, SecretTypes.SSHCredential, s.SecretType)
	require.Equal(t, s.KerberosAuthentication.Port, 9292)
	require.Equal(t, s.KerberosAuthentication.Principal, "someuser")
	require.Equal(t, s.KerberosAuthentication.Realm, "somerealm")
	require.NotNil(t, s.UsageScope)
	require.Len(t, s.UsageScope.AppEnvScopes, 1)
	require.Equal(t, ApplicationFilterTypes.All, s.UsageScope.AppEnvScopes[0].Application.FilterType)
	require.Equal(t, EnvironmentFilterTypes.NonProduction, s.UsageScope.AppEnvScopes[0].Environment.FilterType)

	// Cleanup
	client.Secrets().DeleteSecret(s.Id, s.SecretType)
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
	input := &UpdateSSHCredential{
		Name: updatedName,
	}

	updatedSecret, err := client.Secrets().UpdateSSHCredential(input)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, updatedSecret)
	require.Equal(t, updatedName, updatedSecret.Name)

	// Cleanup
	deleteSecret(updatedSecret.Id, updatedSecret.SecretType)
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
	require.Equal(t, SSHAuthenticationTypes.SSHAuthentication, s.AuthenticationType)
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
	require.Equal(t, SSHAuthenticationTypes.KerberosAuthentication, s.AuthenticationType)
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
	require.Equal(t, SSHAuthenticationTypes.SSHAuthentication, s.AuthenticationType)
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
	require.Equal(t, SSHAuthenticationTypes.KerberosAuthentication, s.AuthenticationType)
}

func createSSHCredential_sshAuth(name string) (*SSHCredential, error) {

	passphraseSecret, err := createEncryptedTextSecret(name, "foo")
	if err != nil {
		return nil, err
	}

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.SSH,
		Name:                 name,
		SSHAuthentication: &SSHAuthenticationInput{
			Port:     22,
			Username: "testuser",
			SSHAuthenticationMethod: &SSHAuthenticationMethod{
				SSHCredentialType: SSHCredentialTypes.Password,
				ServerPassword: &SSHPassword{
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

func createSSHCredential_kerberosAuth(name string) (*SSHCredential, error) {

	passphraseSecret, err := createEncryptedTextSecret(name, "foo")
	if err != nil {
		return nil, err
	}

	input := &SSHCredentialInput{
		AuthenticationScheme: SSHAuthenticationSchemes.Kerberos,
		Name:                 name,
		KerberosAuthentication: &KerberosAuthentication{
			Port:      9292,
			Principal: "someuser",
			Realm:     "somerealm",
			TGTGenerationMethod: &TGTGenerationMethod{
				TGTGenerationUsing: TGTGenerationUsingOptions.Password,
				KerberosPassword: &KerberosPassword{
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
