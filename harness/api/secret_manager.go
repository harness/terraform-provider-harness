package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/api/unpublished"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
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

func (c *SecretClient) ListSecretManagers() ([]*unpublished.SecretManager, error) {
	req, err := c.APIClient.NewAuthorizedGetRequest("/gateway/api/secrets/list-configs")
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add(helpers.QueryParameters.AccountId.String(), c.APIClient.AccountId)
	req.URL.RawQuery = query.Encode()

	resp, err := c.APIClient.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responsePackage := &unpublished.Package{}
	err = json.NewDecoder(resp.Body).Decode(responsePackage)

	if err != nil {
		return nil, err
	}

	secretManagers := []*unpublished.SecretManager{}
	err = json.Unmarshal(*responsePackage.Resource, &secretManagers)
	if err != nil {
		return nil, err
	}

	return secretManagers, nil
}

// Currently there is no way to find the Id of the default secret manager
// directly through the API. Instead, this method creates a temporary secret
// without specifying which secret manager to use. Once it's created we can
// then read back the id of the secret manager that is automatically assigned.
func (c *SecretClient) GetDefaultSecretManagerId() (string, error) {
	input := &graphql.CreateSecretInput{
		EncryptedText: &graphql.EncryptedTextInput{},
	}
	input.EncryptedText.Name = "__temp__" + utils.RandStringBytes(6)
	input.EncryptedText.Value = "test"

	secret, err := c.CreateEncryptedText(input)
	if err != nil {
		return "", err
	}

	if secret == nil {
		return "", errors.New("could not create secret")
	}

	defer func() {
		err := c.DeleteSecret(secret.Id, graphql.SecretTypes.EncryptedText)
		if err != nil {
			panic(err)
		}
	}()

	return secret.SecretManagerId, nil
}
