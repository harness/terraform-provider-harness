// graphql.go
package chaos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

const (
	// DefaultGraphQLEndpoint is the default GraphQL endpoint
	DefaultGraphQLEndpoint = "/query"
)

// GraphQLQuery represents a GraphQL query
type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLError represents a GraphQL error
type GraphQLError struct {
	Message    string                 `json:"message"`
	Path       []string               `json:"path,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// GraphQLResponse represents a GraphQL response
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors,omitempty"`
}

// ToError converts a GraphQLError to a Go error
func (e *GraphQLError) ToError() error {
	if e == nil {
		return nil
	}
	if e.Message == "" {
		return fmt.Errorf("graphql error: %v", e.Path)
	}
	return fmt.Errorf("graphql error: %s (path: %v)", e.Message, e.Path)
}

// ExecuteGraphQL executes a GraphQL query and unmarshals the response
func (c *APIClient) ExecuteGraphQL(ctx context.Context, query string, variables map[string]interface{}, response interface{}) error {
	// Recover from panics
	defer func() {
		if r := recover(); r != nil {
			c.Logger.WithFields(logrus.Fields{
				"stack": string(debug.Stack()),
			}).Errorf("recovered from panic in ExecuteGraphQL: %v", r)
		}
	}()

	// Create the request body
	reqBody, err := json.Marshal(GraphQLQuery{
		Query:     query,
		Variables: variables,
	})
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	// Create the request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.cfg.BasePath+DefaultGraphQLEndpoint,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Harness-Account", c.AccountId)

	// Execute the request
	resp, err := c.cfg.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the GraphQL response
	var graphQLResp GraphQLResponse
	if err := json.Unmarshal(body, &graphQLResp); err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Check for GraphQL errors
	if len(graphQLResp.Errors) > 0 {
		// Return the first error for simplicity
		return graphQLResp.Errors[0].ToError()
	}

	// Unmarshal the data into the response object if provided
	if response != nil {
		if err := json.Unmarshal(graphQLResp.Data, response); err != nil {
			return fmt.Errorf("error unmarshaling response data: %w", err)
		}
	}

	return nil
}
