package cd

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
)

// Helper type for accessing all delegate related crud methods
type DelegateClient struct {
	ApiClient *ApiClient
}

func (c *DelegateClient) GetDelegateByName(name string) (*graphql.Delegate, error) {
	delegateList, _, err := c.GetDelegateWithFilters(1, 0, name, "", "")

	if err != nil {
		return nil, err
	}

	if len(delegateList) == 0 {
		return nil, nil
	}

	return delegateList[0], nil
}

func (c *DelegateClient) GetDelegateWithFilters(limit int, offset int, name string, status graphql.DelegateStatus, delegateType graphql.DelegateType) ([]*graphql.Delegate, *graphql.PageInfo, error) {

	filters := []string{}

	if name != "" {
		filters = append(filters, fmt.Sprintf(`delegateName:"%s"`, name))
	}

	if status != "" {
		filters = append(filters, fmt.Sprintf(`delegateStatus: %s`, status))
	}

	if delegateType != "" {
		filters = append(filters, fmt.Sprintf(`delegateType: %s`, delegateType))
	}

	filterStr := strings.Join(filters, ", ")

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query {
			delegateList(limit: %[3]d, offset: %[4]d, filters: [{
				accountId: "%[5]s",
				%[6]s
			}]) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, standardDelegateFields, paginationFields, limit, offset, c.ApiClient.Configuration.AccountId, filterStr),
	}

	res := struct {
		DelegateList struct {
			Nodes    []*graphql.Delegate `json:"nodes"`
			PageInfo *graphql.PageInfo
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.DelegateList.Nodes, res.DelegateList.PageInfo, nil
}

func (c *DelegateClient) ListDelegates(limit int, offset int) ([]*graphql.Delegate, *graphql.PageInfo, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query {
			delegateList(limit: %[3]d, offset: %[4]d) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, standardDelegateFields, paginationFields, limit, offset),
	}

	res := struct {
		DelegateList struct {
			Nodes    []*graphql.Delegate `json:"nodes"`
			PageInfo *graphql.PageInfo
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.DelegateList.Nodes, res.DelegateList.PageInfo, nil
}

const standardDelegateFields = `
ip
hostName
delegateName
uuid
accountId
delegateProfileId
delegateType
description
hostName
lastHeartBeat
polllingModeEnabled
status
version
`
