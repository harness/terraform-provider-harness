package cd

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
)

type SSOClient struct {
	ApiClient *ApiClient
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

	if err := c.ApiClient.ExecuteGraphQLQuery(query, &res); err != nil {
		if strings.Contains(err.Error(), helpers.SSO_PROVIDER_NOT_FOUND) {
			return nil, nil
		}
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

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

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
