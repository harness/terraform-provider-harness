package cd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
)

// Helper type for accessing all delegate related crud methods
type DelegateClient struct {
	ApiClient *ApiClient
}

func (c *DelegateClient) WaitForDelegate(ctx context.Context, delegateName string, timeout time.Duration) (*graphql.Delegate, error) {

	pollInterval := time.Second * 10
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		delegate, err := c.GetDelegateByName(delegateName)
		if err != nil || delegate != nil {
			return delegate, err
		}

		select {
		case <-ctx.Done():
			c.ApiClient.Log.Errorf("Timedout waitingWaiting for delegate %s", delegateName)
			return nil, ctx.Err()
		case <-time.After(pollInterval):
			c.ApiClient.Log.Infof("Waiting for delegate %s", delegateName)
		}
	}
}

func (c *DelegateClient) DeleteDelegate(delegateId string) error {

	query := &GraphQLQuery{
		Query: `mutation($input: DeleteDelegateInput!) {
			deleteDelegate(input: $input) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"input": struct {
				AccountId  string `json:"accountId"`
				DelegateId string `json:"delegateId"`
			}{
				AccountId:  c.ApiClient.Configuration.AccountId,
				DelegateId: delegateId,
			},
		},
	}

	res := &struct {
		deleteDelegate struct {
			clientMutationId string
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

func (c *DelegateClient) UpdateDelegateApprovalStatus(input *graphql.DelegateApprovalRejectInput) (*graphql.Delegate, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: DelegateApproveRejectInput!) {
			delegateApproveReject(input: $input) {
				delegate {
					%[1]s
				}
			}
		}`, standardDelegateFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		DelegateApproveReject struct {
			Delegate graphql.Delegate
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.DelegateApproveReject.Delegate, nil
}

func (c *DelegateClient) GetDelegatesByName(name string) ([]*graphql.Delegate, error) {
	delegateList, _, err := c.ListDelegatesWithFilters(1, 0, name, "", "")

	if err != nil {
		return nil, err
	}

	if len(delegateList) == 0 {
		return nil, nil
	}

	return delegateList, nil
}

func (c *DelegateClient) GetDelegateById(id string) (*graphql.Delegate, error) {

	hasmore := true
	offset := 0
	limit := 100

	for hasmore {

		delegateList, pagination, err := c.ListDelegatesWithFilters(limit, offset, "", "", "")

		if err != nil {
			return nil, err
		}

		if len(delegateList) == 0 {
			return nil, nil
		}

		for _, delegate := range delegateList {
			if delegate.UUID == id {
				return delegate, nil
			}
		}

		hasmore = pagination.HasMore
		offset += limit
	}

	return nil, nil
}

func (c *DelegateClient) GetDelegateByHostName(hostName string) (*graphql.Delegate, error) {

	hasmore := true
	offset := 0
	limit := 100

	for hasmore {

		delegateList, pagination, err := c.ListDelegatesWithFilters(limit, offset, "", "", "")

		if err != nil {
			return nil, err
		}

		if len(delegateList) == 0 {
			return nil, nil
		}

		for _, delegate := range delegateList {
			if delegate.HostName == hostName {
				return delegate, nil
			}
		}

		hasmore = pagination.HasMore
		offset += limit
	}

	return nil, nil
}

func (c *DelegateClient) GetDelegateByName(name string) (*graphql.Delegate, error) {
	delegateList, _, err := c.ListDelegatesWithFilters(1, 0, name, "", "")

	if err != nil {
		return nil, err
	}

	if len(delegateList) == 0 {
		return nil, nil
	}

	return delegateList[0], nil
}

func (c *DelegateClient) ListDelegatesWithFilters(limit int, offset int, name string, status graphql.DelegateStatus, delegateType graphql.DelegateType) ([]*graphql.Delegate, *graphql.PageInfo, error) {

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
	return c.ListDelegatesWithFilters(limit, offset, "", "", "")
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
