package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/stretchr/testify/require"
)

func TestNewGraphQLRequest(t *testing.T) {

	// Setup
	client := getUnauthorizedClient()

	query := &GraphQLQuery{
		Query: `{}`,
	}

	// Execute
	req, err := client.NewGraphQLRequest(query)

	// Validate
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("accountId=%s", client.AccountId), req.URL.RawQuery)
	require.Equal(t, client.Endpoint, fmt.Sprintf("%s://%s", req.URL.Scheme, req.Host))
	require.Equal(t, client.APIKey, req.Header.Get(helpers.HTTPHeaders.ApiKey.String()))
	require.Equal(t, helpers.HTTPHeaders.ApplicationJson.String(), req.Header.Get(helpers.HTTPHeaders.ContentType.String()))
	require.Equal(t, helpers.HTTPHeaders.ApplicationJson.String(), req.Header.Get(helpers.HTTPHeaders.Accept.String()))
}

func TestExecuteGraphQLQuery(t *testing.T) {

	// Setup
	client := getClient()
	query := &GraphQLQuery{
		Query: `{
			applications(limit: 1) {
				nodes {
					id
					name
				}
			}
		}`,
	}

	res := &struct {
		Applications graphql.Applications
	}{}
	// Execute query
	err := client.ExecuteGraphQLQuery(query, &res)

	// Validate
	require.NoError(t, err)
	require.Len(t, res.Applications.Nodes, 1)
}

func TestUnauthorizedGraphQLQuery(t *testing.T) {
	t.Skip("Need to fix the response code https://harness.atlassian.net/browse/SWAT-5062")
	// Setup
	client := getUnauthorizedClient()
	_, _, err := client.Applications().ListApplications(1, 0)

	// Validate
	require.Error(t, err)
	require.EqualError(t, err, "ERROR INVALID_TOKEN: Token is not valid.")
}
