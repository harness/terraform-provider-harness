package client

import (
	"errors"
	"fmt"
)

const encryptedSecretFields = `
id
name
secretManagerId
secretType
usageScope {
	appEnvScopes {
		application {
			filterType
			appId
		}
		environment {
			filterType
			envId
		}
	}
}
inheritScopesFromSM
scopedToAccount
`

const sshCredentialFields = `
... on SSHCredential {
	sshAuthentication: authenticationType {
		... on SSHAuthentication  {
			port
			userName
		}
	}
	kerberosAuthentication: authenticationType {
		... on KerberosAuthentication {
			port
			principal
			realm
		}
	}
	id
	name
	secretType
	usageScope {
			appEnvScopes {
				application {
					filterType
					appId
				}
				environment {
					filterType
					envId
				}
			}
		}
}
`

const winRMCredentialFields = `
	id
	authenticationScheme
	domain
	id
	name
	port
	secretType
	skipCertCheck
	usageScope {
		appEnvScopes {
			application {
				filterType
				appId
			}
			environment {
				filterType
				envId
			}
		}
	}
	useSSL
	userName
`

func (c *SecretClient) GetEncryptedTextByName(name string) (*EncryptedText, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on EncryptedText {
					%s
				}
			}
		}`, SecretTypes.EncryptedText, encryptedSecretFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName EncryptedText
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

func (c *SecretClient) GetEncryptedFileByName(name string) (*EncryptedFile, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on EncryptedFile {
					%s
				}
			}
		}`, SecretTypes.EncryptedFile, encryptedSecretFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName EncryptedFile
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

func (c *SecretClient) GetEncryptedTextById(id string) (*EncryptedText, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				... on EncryptedText {
					%s
				}
			}
		}`, SecretTypes.EncryptedText, encryptedSecretFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret EncryptedText
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) GetEncryptedFileById(id string) (*EncryptedFile, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				... on EncryptedFile {
					%s
				}
			}
		}`, SecretTypes.EncryptedFile, encryptedSecretFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret EncryptedFile
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) GetSSHCredentialById(id string) (*SSHCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				... on SSHCredential {
					%s
				}
			}
		}`, SecretTypes.SSHCredential, sshCredentialFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret SSHCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	if err := res.Secret.setSSHAuthenticationType(); err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) GetSSHCredentialByName(name string) (*SSHCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on SSHCredential {
					%s
				}
			}
		}`, SecretTypes.SSHCredential, sshCredentialFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName SSHCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	if err := res.SecretByName.setSSHAuthenticationType(); err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

func (c *SecretClient) GetWinRMCredentialByName(name string) (*WinRMCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on WinRMCredential {
					%s
				}
			}
		}`, SecretTypes.WinRMCredential, winRMCredentialFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName WinRMCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

func (c *SecretClient) GetWinRMCredentialById(id string) (*WinRMCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				... on WinRMCredential {
					%s
				}
			}
		}`, SecretTypes.WinRMCredential, winRMCredentialFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret WinRMCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) CreateEncryptedText(input *CreateSecretInput) (*EncryptedText, error) {

	// Set defaults
	input.SecretType = SecretTypes.EncryptedText
	if input.EncryptedText.SecretManagerId == "" {
		input.EncryptedText.SecretManagerId = c.APIClient.AccountId
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($secret: CreateSecretInput!) {
			createSecret(input: $secret) {
				secret {
				... on EncryptedText {
						%s
					}
				}
			}
		}`, encryptedSecretFields),
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		CreateSecret struct {
			Secret EncryptedText
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateSecret.Secret, nil
}

func (c *SecretClient) UpdateEncryptedText(input *UpdateSecretInput) (*EncryptedText, error) {

	input.SecretType = SecretTypes.EncryptedText

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($secret: UpdateSecretInput!) {
			updateSecret(input: $secret) {
				secret {
				... on EncryptedText {
						%s
					}
				}
			}
		}`, encryptedSecretFields),
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		UpdateSecret struct {
			Secret EncryptedText
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateSecret.Secret, nil
}

func (c *SecretClient) DeleteSecret(input *DeleteSecretInput) error {

	query := &GraphQLQuery{
		Query: `mutation($secret: DeleteSecretInput!) {
			deleteSecret(input: $secret) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		DeleteSecret struct {
			ClientMutationId string
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

func CreateEncryptedFile() {}

func (auth *SSHAuthentication) IsValid() bool {
	return auth.Username != ""
}

func (auth *KerberosAuthentication) IsValid() bool {
	return auth.Principal != "" && auth.Realm != ""
}

// Determines which SSH authentication type is used
func (c *SSHCredential) setSSHAuthenticationType() error {

	if c.SSHAuthentication != nil && c.SSHAuthentication.IsValid() {
		c.AuthenticationType = SSHAuthenticationTypes.SSHAuthentication
		c.KerberosAuthentication = nil
	} else if c.KerberosAuthentication != nil && c.KerberosAuthentication.IsValid() {
		c.AuthenticationType = SSHAuthenticationTypes.KerberosAuthentication
		c.SSHAuthentication = nil
	} else {
		return errors.New("invalid SSH Authentication type")
	}

	return nil
}

// // Sets up defaults for Encrypted Text creation
// func (i *EncryptedTextInput) setDefaults() {

// 	// Create a default scope for Prod and Non-prod
// 	if i.UsageScope == nil {

// 		scopes := make([]*AppEnvScope, 2)

// 		scopes[0] = &AppEnvScope{
// 			Application: &AppScopeFilter{
// 				FilterType: ApplicationFilterTypes.All,
// 			},
// 			Environment: &EnvScopeFilter{
// 				FilterType: EnvironmentFilterTypes.NonProduction,
// 			},
// 		}

// 		scopes[1] = &AppEnvScope{
// 			Application: &AppScopeFilter{
// 				FilterType: ApplicationFilterTypes.All,
// 			},
// 			Environment: &EnvScopeFilter{
// 				FilterType: EnvironmentFilterTypes.Production,
// 			},
// 		}

// 		i.UsageScope = &UsageScope{
// 			AppEnvScopes: scopes,
// 		}
// 	}
// }
