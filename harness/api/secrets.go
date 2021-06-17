package api

import (
	"errors"
	"fmt"

	"github.com/micahlmartin/terraform-provider-harness/harness/api/graphql"
)

// Get the client for interacting with Harness Applications
func (c *Client) Secrets() *SecretClient {
	return &SecretClient{
		APIClient: c,
	}
}

func (c *SecretClient) GetEncryptedFileByName(name string) (*graphql.EncryptedFile, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on EncryptedFile {
					%s
				}
			}
		}`, graphql.SecretTypes.EncryptedFile, encryptedSecretFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName graphql.EncryptedFile
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

func (c *SecretClient) GetEncryptedFileById(id string) (*graphql.EncryptedFile, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				... on EncryptedFile {
					%s
				}
			}
		}`, graphql.SecretTypes.EncryptedFile, encryptedSecretFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret graphql.EncryptedFile
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) GetWinRMCredentialByName(name string) (*graphql.WinRMCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretByName(name: $name, secretType: %s) {
				... on WinRMCredential {
					%s
				}
			}
		}`, graphql.SecretTypes.WinRMCredential, winRMCredentialFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretByName graphql.WinRMCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return &res.SecretByName, nil
}

func (c *SecretClient) GetWinRMCredentialById(id string) (*graphql.WinRMCredential, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			secret(secretId: $id, secretType: %s) {
				... on WinRMCredential {
					%s
				}
			}
		}`, graphql.SecretTypes.WinRMCredential, winRMCredentialFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := &struct {
		Secret graphql.WinRMCredential
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return &res.Secret, nil
}

func (c *SecretClient) DeleteSecret(secretId string, secretType string) error {

	query := &GraphQLQuery{
		Query: `mutation($secret: DeleteSecretInput!) {
			deleteSecret(input: $secret) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"secret": &graphql.DeleteSecretInput{
				SecretId:   secretId,
				SecretType: secretType,
			},
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

func (c *SecretClient) DeleteSecretObj(secret *graphql.Secret) error {
	if secret == nil {
		return errors.New("could not delete secret. object is nil")
	}

	return c.DeleteSecret(secret.Id, secret.SecretType)
}

const (
	winRMCredentialFields = `
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
)
