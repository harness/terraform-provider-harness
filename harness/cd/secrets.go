package cd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
)

type SecretClient struct {
	ApiClient *ApiClient
}

func (c *SecretClient) getSecretByName(name string, secretType graphql.SecretType, fields string, respObj interface{}) error {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			secretByName(name: "%[1]s", secretType: %[2]s) {
				%[3]s
			}
		}`, name, secretType, fields),
	}

	res := &struct {
		SecretByName interface{}
	}{
		SecretByName: respObj,
	}

	if err := c.ApiClient.ExecuteGraphQLQuery(query, &res); err != nil {
		if strings.Contains(err.Error(), "No secret exists") {
			return nil
		}
	}

	return nil
}

func (c *SecretClient) getSecretById(id string, secretType graphql.SecretType, fields string, respObj interface{}) error {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			secret(secretId: "%[1]s", secretType: %[2]s) {
				%[3]s
			}
		}`, id, secretType, fields),
	}

	res := &struct {
		Secret *json.RawMessage
	}{}

	if err := c.ApiClient.ExecuteGraphQLQuery(query, &res); err != nil {
		if strings.Contains(err.Error(), "No secret exists") {
			return nil
		}
	}

	err := json.Unmarshal(*res.Secret, respObj)
	if err != nil {
		return err
	}

	return nil
}

func (c *SecretClient) DeleteSecret(secretId string, secretType graphql.SecretType) error {

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

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

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
	commonSecretFields = `
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
	`
)
