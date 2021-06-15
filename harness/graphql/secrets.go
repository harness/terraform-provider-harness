package graphql

import (
	"errors"
	"fmt"
)

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

// Get the client for interacting with Harness Applications
func (c *ApiClient) Secrets() *SecretClient {
	return &SecretClient{
		APIClient: c,
	}
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

func (c *SecretClient) DeleteSecret(secretId string, secretType string) error {

	query := &GraphQLQuery{
		Query: `mutation($secret: DeleteSecretInput!) {
			deleteSecret(input: $secret) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"secret": &DeleteSecretInput{
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

func (c *SecretClient) DeleteSecretObj(secret *Secret) error {
	if secret == nil {
		return errors.New("could not delete secret. object is nil")
	}

	return c.DeleteSecret(secret.Id, secret.SecretType)
}
