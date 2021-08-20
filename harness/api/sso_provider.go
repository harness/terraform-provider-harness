package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

type SSOClient struct {
	APIClient *Client
}

// Get the client for interacting with Harness Applications
func (c *Client) SSO() *SSOClient {
	return &SSOClient{
		APIClient: c,
	}
}

func (c *SSOClient) GetSSOProviderById(id string) (*graphql.SSOProvider, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			ssoProvider(ssoProviderId: "%[1]s") {
				%[2]s
			}
		}`, id, ssoProviderFields),
	}

	res := struct {
		SSOProvider *graphql.SSOProvider
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return res.SSOProvider, nil
}

func (c *SSOClient) GetSSOProviderByName(name string) (*graphql.SSOProvider, error) {

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		providers, pagination, err := c.ListSSOProviders(limit, offset)
		if err != nil {
			return nil, err
		}

		for _, provider := range providers {
			if provider.Name == name {
				return provider, nil
			}
		}

		hasMore = pagination.HasMore
		offset += 1
	}

	return nil, fmt.Errorf("could not find SSO Provider with name: %s", name)
}

func (c *SSOClient) ListSSOProviders(limit int, offset int) ([]*graphql.SSOProvider, *graphql.PageInfo, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($limit: Int!, $offset: Int) {
			ssoProviders(limit: $limit, offset: $offset) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, ssoProviderFields, paginationFields),
		Variables: map[string]interface{}{
			"limit":  limit,
			"offset": offset,
		},
	}

	res := struct {
		SSOProviders struct {
			Nodes    []*graphql.SSOProvider
			PageInfo *graphql.PageInfo
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.SSOProviders.Nodes, res.SSOProviders.PageInfo, nil
}

const (
	ssoProviderFields = `
	id
	name
	ssoType
`
)
