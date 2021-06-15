package graphql

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
	"github.com/micahlmartin/terraform-provider-harness/harness/httphelpers"
)

func (m *GraphQLResponseMessage) ToError() error {
	return fmt.Errorf("%s %s: %s", m.Level, m.Code, m.Message)
}

func (m *GraphQLError) ToError() error {
	return errors.New(m.Message)
}

// Creates a new client for interacting with the GraphQL API
func (client *ApiClient) GraphQL() *graphql.Client {
	url := fmt.Sprintf("%s%s?%s=%s", client.Endpoint, DefaultGraphQLApiUrl, QueryParamAccountId, client.AccountId)
	return graphql.NewClient(url)
}

// Create new request for making GraphQL Api calls
func (client *ApiClient) NewGraphQLRequest(query *GraphQLQuery) (*http.Request, error) {
	var requestBody bytes.Buffer

	// JSON encode our body payload
	if err := json.NewEncoder(&requestBody).Encode(query); err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] GraphQL Query: %s", requestBody.String())

	req, err := http.NewRequest(http.MethodPost, getGraphQLUrl(), &requestBody)

	if err != nil {
		return nil, err
	}

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(QueryParamAccountId, client.AccountId)
	req.URL.RawQuery = q.Encode()

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderApiKey, client.APIKey)
	req.Header.Set(httphelpers.HeaderContentType, httphelpers.HeaderApplicationJson)
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)

	return req, nil
}

// Executes a GraphQL query
func (client *ApiClient) ExecuteGraphQLQuery(query *GraphQLQuery, responseObj interface{}) error {
	req, err := client.NewGraphQLRequest(query)

	if err != nil {
		return err
	}

	res, err := client.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}

	log.Printf("[DEBUG] GraphQL response: %s", buf.String())

	gqlResponse := &GraphQLStandardResponse{}

	// Unmarshal into our response object
	if err := json.NewDecoder(&buf).Decode(&gqlResponse); err != nil {
		return fmt.Errorf("error decoding response: %s", err)
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

// Returns fully qualified path to the GraphQL Api
func getGraphQLUrl() string {
	return fmt.Sprintf("%s%s", DefaultApiUrl, DefaultGraphQLApiUrl)
}
