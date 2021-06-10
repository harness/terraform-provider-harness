package client

import (
	"errors"
	"fmt"
)

func (c *SecretClient) CreateSSHCredential(sshInput *SSHCredential) (*SSHCredential, error) {

	// Set defaults
	input := &CreateSecretInput{
		SecretType:    SecretTypes.SSHCredential,
		SSHCredential: sshInput,
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($secret: CreateSecretInput!) {
			createSecret(input: $secret) {
				secret {
				... on SSHCredential {
						%s
					}
				}
			}
		}`, sshCredentialFields),
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		CreateSecret struct {
			Secret SSHCredential
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateSecret.Secret, nil
}

func (c *SecretClient) UpdateSSHCredential(id string, cred *SSHCredential) (*SSHCredential, error) {

	input := &UpdateSecretInput{
		SecretId:      id,
		SecretType:    SecretTypes.SSHCredential,
		SSHCredential: cred,
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($secret: UpdateSecretInput!) {
			updateSecret(input: $secret) {
				secret {
					%s
				}
			}
		}`, sshCredentialFields),
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		UpdateSecret struct {
			Secret SSHCredential
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateSecret.Secret, nil
}

func (c *SecretClient) GetSSHCredentialById(id string) (*SSHCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				%s
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

func (auth *SSHAuthentication) isValid() bool {
	return auth.Username != ""
}

func (auth *KerberosAuthentication) isValid() bool {
	return auth.Principal != "" && auth.Realm != ""
}

// Determines which SSH authentication type is used
func (c *SSHCredential) setSSHAuthenticationType() error {

	if c.SSHAuthentication != nil && c.SSHAuthentication.isValid() {
		c.AuthenticationType = SSHAuthenticationTypes.SSHAuthentication
		c.KerberosAuthentication = nil
	} else if c.KerberosAuthentication != nil && c.KerberosAuthentication.isValid() {
		c.AuthenticationType = SSHAuthenticationTypes.KerberosAuthentication
		c.SSHAuthentication = nil
	} else {
		return errors.New("invalid SSH Authentication type")
	}

	return nil
}

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
