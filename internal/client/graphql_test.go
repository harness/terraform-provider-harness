package client

import (
	"fmt"
	"testing"

	"github.com/micahlmartin/terraform-provider-harness/internal/httphelpers"
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
	require.Equal(t, client.APIKey, req.Header.Get(httphelpers.HeaderApiKey))
	require.Equal(t, httphelpers.HeaderApplicationJson, req.Header.Get(httphelpers.HeaderContentType))
	require.Equal(t, httphelpers.HeaderApplicationJson, req.Header.Get(httphelpers.HeaderAccept))
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

	// Execute query
	res, err := client.ExecuteGraphQLQuery(query)

	// Validate
	require.NoError(t, err)
	require.Len(t, res.ResponseMessages, 0)
}

func TestUnauthorizedGraphQLQuery(t *testing.T) {
	// Setup
	client := getUnauthorizedClient()
	query := &GraphQLQuery{
		Query: `query {}`,
	}

	// Execute query
	res, err := client.ExecuteGraphQLQuery(query)

	// Validate
	require.Error(t, err)
	require.Nil(t, res)
	require.EqualError(t, err, "ERROR INVALID_TOKEN: Token is not valid.")
}
