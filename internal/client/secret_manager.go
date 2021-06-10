package client

import (
	"fmt"
)

const (
	standardSecretsManagerFields = `
		id
		name
	`
)

// CRUD
func (ac *SecretClient) GetSecretManagerById(id string) (*SecretManager, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($secretManagerId: String!) {
			secretManager(secretManagerId: $secretManagerId) {
				%s
			}
		}`, standardSecretsManagerFields),
		Variables: map[string]interface{}{
			"secretManagerId": id,
		},
	}

	res := &struct {
		SecretManager SecretManager
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, res)

	if err != nil {
		return nil, err
	}

	return &res.SecretManager, nil
}

func (ac *SecretClient) GetSecretManagerByName(name string) (*SecretManager, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			secretManagerByName(name: $name) {
				%s
			}
		}`, standardSecretsManagerFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		SecretManagerByName SecretManager
	}{}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return nil, err
	}

	return &res.SecretManagerByName, nil
}
