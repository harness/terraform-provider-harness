package cd

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/cd/unpublished"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
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
	err := ac.ApiClient.ExecuteGraphQLQuery(query, res)

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

	err := ac.ApiClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return nil, err
	}

	return &res.SecretManagerByName, nil
}

// WARNING: This method requires the use of a bearer token which isn't supported in most scenarios.
func (c *SecretClient) ListSecretManagers() ([]*unpublished.SecretManager, error) {
	req, err := c.ApiClient.NewAuthorizedGetRequest("/secrets/list-configs")
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)
	req.URL.RawQuery = query.Encode()

	resp, err := c.ApiClient.Configuration.HTTPClient.Do(req)
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

var defaultSecretManagerId string
var defaultSecretManagerLookupError error
var configSecretManagerId sync.Once

// Currently there is no way to find the Id of the default secret manager
// directly through the API. Instead, this method creates a temporary secret
// without specifying which secret manager to use. Once it's created we can
// then read back the id of the secret manager that is automatically assigned.
func (c *SecretClient) GetDefaultSecretManagerId() (string, error) {
	configSecretManagerId.Do(func() {
		var secret *graphql.EncryptedText

		input := &graphql.CreateSecretInput{
			EncryptedText: &graphql.EncryptedTextInput{},
		}
		input.EncryptedText.Name = "__temp__" + utils.RandStringBytes(6)
		input.EncryptedText.Value = "test"

		secret, defaultSecretManagerLookupError = c.CreateEncryptedText(input)
		if defaultSecretManagerLookupError != nil {
			return
		}

		if secret == nil {
			defaultSecretManagerLookupError = errors.New("could not create secret")
			return
		}

		defaultSecretManagerId = secret.SecretManagerId

		_ = c.DeleteSecret(secret.Id, graphql.SecretTypes.EncryptedText)
	})

	return defaultSecretManagerId, defaultSecretManagerLookupError
}
