package cd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/cd/unpublished"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
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

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

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

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

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

	if resp.IsEmpty() {
		return nil, nil
	}

	return resp, nil
}

func (c *SecretClient) GetEncryptedTextById(id string) (*graphql.EncryptedText, error) {
	resp := &graphql.EncryptedText{}
	err := c.getSecretById(id, graphql.SecretTypes.EncryptedText, getEncryptedTextFields(), resp)
	if err != nil {
		return nil, err
	}

	if resp.IsEmpty() {
		return nil, nil
	}

	return resp, nil
}

// WARNING: This method requires the use of a bearer token which isn't supported in most scenarios.
func (c *SecretClient) ListEncryptedTextSecrets(limit int, offset int) ([]*unpublished.EncryptedText, *graphql.PageInfo, error) {
	req, err := c.ApiClient.NewAuthorizedGetRequest("/secrets/list-secrets-page")
	if err != nil {
		return nil, nil, err
	}

	query := req.URL.Query()
	query.Add(helpers.QueryParameters.AccountId.String(), c.ApiClient.Configuration.AccountId)
	query.Add(helpers.QueryParameters.Type.String(), "SECRET_TEXT")
	query.Add(helpers.QueryParameters.Limit.String(), strconv.Itoa(limit))
	query.Add(helpers.QueryParameters.Offset.String(), strconv.Itoa(offset))
	req.URL.RawQuery = query.Encode()

	resp, err := c.ApiClient.Configuration.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	responsePackage := &unpublished.Package{}
	err = json.NewDecoder(resp.Body).Decode(responsePackage)

	if err != nil {
		return nil, nil, err
	}

	resource := &unpublished.EncryptedTextResource{}
	err = json.Unmarshal(*responsePackage.Resource, resource)
	if err != nil {
		return nil, nil, err
	}

	pageInfo := &graphql.PageInfo{}

	offset, _ = strconv.Atoi(resource.Offset)
	pageInfo.Offset = offset

	limit, _ = strconv.Atoi(resource.Limit)
	pageInfo.Limit = limit

	return resource.Secrets, pageInfo, nil
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
