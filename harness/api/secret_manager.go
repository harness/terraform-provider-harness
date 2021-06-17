package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

const (
	standardSecretsManagerFields = `
		id
		name
	`
)

// CRUD
func (ac *SecretClient) GetSecretManagerById(id string) (*graphql.SecretManager, error) {
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
		SecretManager graphql.SecretManager
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, res)

	if err != nil {
		return nil, err
	}

	return &res.SecretManager, nil
}

func (ac *SecretClient) GetSecretManagerByName(name string) (*graphql.SecretManager, error) {
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
		SecretManagerByName graphql.SecretManager
	}{}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return nil, err
	}

	return &res.SecretManagerByName, nil
}
