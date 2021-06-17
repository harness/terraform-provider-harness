package api

import (
	"fmt"

	"github.com/micahlmartin/terraform-provider-harness/harness/api/graphql"
)

func (c *SecretClient) CreateSSHCredential(sshInput *graphql.SSHCredential) (*graphql.SSHCredential, error) {

	// Set defaults
	input := &graphql.CreateSecretInput{
		SecretType:    graphql.SecretTypes.SSHCredential,
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
			Secret graphql.SSHCredential
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateSecret.Secret, nil
}

func (c *SecretClient) UpdateSSHCredential(id string, cred *graphql.SSHCredential) (*graphql.SSHCredential, error) {

	input := &graphql.UpdateSecretInput{
		SecretId:      id,
		SecretType:    graphql.SecretTypes.SSHCredential,
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
			Secret graphql.SSHCredential
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateSecret.Secret, nil
}

func (c *SecretClient) GetSSHCredentialById(id string) (*graphql.SSHCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				%s
			}
		}`, graphql.SecretTypes.SSHCredential, sshCredentialFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret graphql.SSHCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	if err := res.Secret.SetSSHAuthenticationType(); err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) GetSSHCredentialByName(name string) (*graphql.SSHCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on SSHCredential {
					%s
				}
			}
		}`, graphql.SecretTypes.SSHCredential, sshCredentialFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName graphql.SSHCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	if err := res.SecretByName.SetSSHAuthenticationType(); err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

// Determines which SSH authentication type is used

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
