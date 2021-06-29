package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *SecretClient) CreateEncryptedText(input *graphql.CreateSecretInput) (*graphql.EncryptedText, error) {

	// Set defaults
	input.SecretType = graphql.SecretTypes.EncryptedText

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($secret: CreateSecretInput!) {
			createSecret(input: $secret) {
				secret {
					%[1]s
				}
			}
		}`, getEncryptedTextFields()),
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		CreateSecret struct {
			Secret graphql.EncryptedText
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateSecret.Secret, nil
}

func (c *SecretClient) UpdateEncryptedText(input *graphql.UpdateSecretInput) (*graphql.EncryptedText, error) {

	input.SecretType = graphql.SecretTypes.EncryptedText

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($secret: UpdateSecretInput!) {
			updateSecret(input: $secret) {
				secret {
					%[1]s
				}
			}
		}`, getEncryptedTextFields()),
		Variables: map[string]interface{}{
			"secret": &input,
		},
	}

	res := &struct {
		UpdateSecret struct {
			Secret graphql.EncryptedText
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateSecret.Secret, nil
}

func (c *SecretClient) GetEncryptedTextByName(name string) (*graphql.EncryptedText, error) {
	resp := &graphql.EncryptedText{}
	err := c.getSecretByName(name, graphql.SecretTypes.EncryptedText, getEncryptedTextFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SecretClient) GetEncryptedTextById(id string) (*graphql.EncryptedText, error) {
	resp := &graphql.EncryptedText{}
	err := c.getSecretById(id, graphql.SecretTypes.EncryptedText, getEncryptedTextFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getEncryptedTextFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on EncryptedText {
			inheritScopesFromSM
			scopedToAccount
			secretManagerId
	}`, commonSecretFields)
}
