package client

import "fmt"

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
