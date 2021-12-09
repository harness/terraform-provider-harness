package cd

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/stretchr/testify/require"
)

func TestNewGraphQLRequest(t *testing.T) {

	// Setup
	client := GetUnauthorizedClient()

	query := &GraphQLQuery{
		Query: `{}`,
	}

	// Execute
	req, err := client.NewGraphQLRequest(query)

	// Validate
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("accountId=%s", client.Configuration.AccountId), req.URL.RawQuery)
	require.Equal(t, client.Configuration.Endpoint, fmt.Sprintf("%s://%s/gateway/api", req.URL.Scheme, req.Host))
	require.Equal(t, client.Configuration.APIKey, req.Header.Get(helpers.HTTPHeaders.ApiKey.String()))
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
	client := GetUnauthorizedClient()
	_, _, err := client.ApplicationClient.ListApplications(1, 0)

	// Validate
	require.Error(t, err)
	require.EqualError(t, err, "ERROR INVALID_TOKEN: Token is not valid.")
}
