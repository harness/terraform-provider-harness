package cd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/hashicorp/go-retryablehttp"
)

const DefaultGraphQLApiUrl = "/graphql"

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

	log.Printf("[DEBUG] GraphQL: Query %s", requestBody.String())

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
func (client *ApiClient) ExecuteGraphQLQuery(query *GraphQLQuery, responseObj interface{}) error {
	req, err := client.NewGraphQLRequest(query)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	res, err := client.Configuration.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return errors.New("unauthorized")
	}

	defer res.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return fmt.Errorf("error reading body: %s", err)
	}

	log.Printf("[TRACE] GraphQL response: %s", buf.String())

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

const paginationFields = `
pageInfo {
	limit
	offset
	total
	hasMore
}
`
