package client

import (
	"fmt"
	"testing"

	"github.com/micahlmartin/terraform-provider-harness/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestNewGraphQLRequest(t *testing.T) {
	client := getUnauthorizedClient()

	query := &GraphQLQuery{
		Query: `{}`,
	}
	req, err := client.NewGraphQLRequest(query)

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("accountId=%s", client.AccountId), req.URL.RawQuery)
	assert.Equal(t, client.Endpoint, fmt.Sprintf("%s://%s", req.URL.Scheme, req.Host))
	assert.Equal(t, client.APIKey, req.Header.Get(common.HTTP_HEADER_X_API_KEY))
	assert.Equal(t, common.HTTP_HEADER_APPLICATION_JSON, req.Header.Get(common.HTTP_HEADER_CONTENT_TYPE))
	assert.Equal(t, common.HTTP_HEADER_APPLICATION_JSON, req.Header.Get(common.HTTP_HEADER_ACCEPT))
}

func TestExecuteGraphQLQuery(t *testing.T) {
	client := getClient()
	query := &GraphQLQuery{
		Query: `{}`,
	}

	res, err := client.ExecuteGraphQLQuery(query)

	assert.Nil(t, err)
	assert.Len(t, res.ResponseMessages, 0)
}

func TestUnauthorizedGraphQLQuery(t *testing.T) {
	client := getUnauthorizedClient()
	query := &GraphQLQuery{
		Query: `{}`,
	}
	res, err := client.ExecuteGraphQLQuery(query)

	assert.Nil(t, res)
	assert.EqualError(t, err, "ERROR INVALID_TOKEN: Token is not valid.")
}
