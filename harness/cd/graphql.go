package cd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/hashicorp/go-retryablehttp"
)

const DefaultGraphQLApiUrl = "/api/graphql"

func (m *GraphQLResponseMessage) ToError() error {
	return fmt.Errorf("%s %s: %s", m.Level, m.Code, m.Message)
}

func (m *GraphQLError) ToError() error {
	return errors.New(m.Message)
}

// Create new request for making GraphQL Api calls
func (client *ApiClient) NewGraphQLRequest(query *GraphQLQuery) (*retryablehttp.Request, error) {
	var requestBody bytes.Buffer

	// JSON encode our body payload
	if err := json.NewEncoder(&requestBody).Encode(query); err != nil {
		return nil, err
	}

	req, err := client.NewAuthorizedPostRequest(DefaultGraphQLApiUrl, &requestBody)

	if err != nil {
		return nil, err
	}

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(helpers.QueryParameters.AccountId.String(), client.Configuration.AccountId)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Executes a GraphQL query
func (c *ApiClient) ExecuteGraphQLQuery(query *GraphQLQuery, responseObj interface{}) error {
	req, err := c.NewGraphQLRequest(query)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			c.Log.Info(r)
		}
	}()

	gqlResponse := &GraphQLStandardResponse{}
	err = c.getJson(req, gqlResponse)

	if err != nil {
		return err
	}

	// Check if there are any errors
	if gqlResponse.ResponseMessages != nil {
		return gqlResponse.ResponseMessages[0].ToError()
	}

	if gqlResponse.Errors != nil && len(gqlResponse.Errors) > 0 {
		return gqlResponse.Errors[0].ToError()
	}

	// Unmarshal into designated response object
	err = json.Unmarshal(*gqlResponse.Data, responseObj)
	if err != nil {
		return err
	}

	return nil
}

const paginationFields = `
pageInfo {
	limit
	offset
	total
	hasMore
}
`
