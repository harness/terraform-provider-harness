package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
	"github.com/micahlmartin/terraform-provider-harness/internal/common"
)

type GraphQLResponse struct {
	Data             GraphQLResponseData      `json:"data"`
	Metadata         interface{}              `json:"metadata"`
	Resource         string                   `json:"resource"`
	ResponseMessages []GraphQLResponseMessage `json:"responseMessages"`
	Errors           []GraphQLError           `json:"errors"`
}

type GraphQLError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int      `json:"line"`
		Column int      `json:"column"`
		Path   []string `json:"path"`
	} `json:"column"`
}

type GraphQLResponseMessage struct {
	Code         string   `json:"code"`
	Level        string   `json:"level"`
	Message      string   `json:"message"`
	Exception    string   `json:"exception"`
	FailureTypes []string `json:"failureTypes"`
}

type GraphQLResponseData struct {
	Application       *Application              `json:"application"`
	ApplicationByName *Application              `json:"applicationByName"`
	CreateApplication *CreateApplicationPayload `json:"createApplication"`
	DeleteApplication *DeleteApplicationPayload `json:"deleteApplication"`
}

type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func (m *GraphQLResponseMessage) ToError() error {
	return fmt.Errorf("%s %s: %s", m.Level, m.Code, m.Message)
}

func (m *GraphQLError) ToError() error {
	return errors.New(m.Message)
}

// Creates a new client for interacting with the GraphQL API
func (client *ApiClient) GraphQL() *graphql.Client {
	url := fmt.Sprintf("%s%s?accountId=%s", client.Endpoint, common.DEFAULT_GRAPHQL_API_URL, client.AccountId)
	return graphql.NewClient(url)
}

// Create new request for making GraphQL Api calls
func (client *ApiClient) NewGraphQLRequest(query *GraphQLQuery) (*http.Request, error) {
	var requestBody bytes.Buffer

	// JSON encode our body payload
	if err := json.NewEncoder(&requestBody).Encode(query); err != nil {
		return nil, err
	}

	fmt.Println(requestBody.String())

	req, err := http.NewRequest(http.MethodPost, getGraphQLUrl(), &requestBody)

	if err != nil {
		return nil, err
	}

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(common.ACCOUNT_ID_QUERY_PARAM, client.AccountId)
	req.URL.RawQuery = q.Encode()

	// Configure additional headers
	req.Header.Set(common.HTTP_HEADER_X_API_KEY, client.APIKey)
	req.Header.Set(common.HTTP_HEADER_CONTENT_TYPE, common.HTTP_HEADER_APPLICATION_JSON)
	req.Header.Set(common.HTTP_HEADER_ACCEPT, common.HTTP_HEADER_APPLICATION_JSON)

	fmt.Println(req.URL)
	return req, nil
}

// Executes a GraphQL query
func (client *ApiClient) ExecuteGraphQLQuery(query *GraphQLQuery) (*GraphQLResponse, error) {
	req, err := client.NewGraphQLRequest(query)

	if err != nil {
		return nil, err
	}

	res, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}

	log.Println(buf.String())

	var responseObj GraphQLResponse

	// Unmarshal into our response object
	if err := json.NewDecoder(&buf).Decode(&responseObj); err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	// Check if there are any errors
	if responseObj.ResponseMessages != nil {
		return nil, responseObj.ResponseMessages[0].ToError()
	}

	if responseObj.Errors != nil && len(responseObj.Errors) > 0 {
		return nil, responseObj.Errors[0].ToError()
	}

	return &responseObj, nil
}

// Returns fully qualified path to the GraphQL Api
func getGraphQLUrl() string {
	return fmt.Sprintf("%s%s", common.DEFAULT_API_URL, common.DEFAULT_GRAPHQL_API_URL)
}
