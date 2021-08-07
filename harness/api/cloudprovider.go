package api

import (
	"encoding/json"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

type CloudProviderClient struct {
	APIClient *Client
}

func (c *Client) CloudProviders() *CloudProviderClient {
	return &CloudProviderClient{
		APIClient: c,
	}
}

func (c *CloudProviderClient) DeleteCloudProvider(id string) error {

	query := &GraphQLQuery{
		Query: `mutation($input: DeleteCloudProviderInput!) {
			deleteCloudProvider(input: $input) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"input": struct {
				CloudProviderId string `json:"cloudProviderId"`
			}{
				CloudProviderId: id,
			},
		},
	}

	res := &struct {
		DeleteCloudProvider struct {
			ClientMutationId string
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

const (
	commonCloudProviderFields = `
		id
		name
		description
		createdAt
		createdBy {
			id
			name
		}
		type
		isContinuousEfficiencyEnabled
	`

	ceHealthStatusFields = `
		ceHealthStatus {
			clusterHealthStatusList {
				clusterId
				clusterName
				errors
				lastEventTimestamp
				messages
			}
			isCEConnector
			isHealthy
			messages
		}
	`
)

func (c *CloudProviderClient) getCloudProviderById(id string, fields string, respObj interface{}) error {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			cloudProvider(cloudProviderId: "%[1]s") {
				%[2]s
			}
		}`, id, fields),
	}

	res := &struct {
		CloudProvider interface{}
	}{
		CloudProvider: respObj,
	}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return err
	}

	return nil
}

func (c *CloudProviderClient) updateCloudProvider(input interface{}, fields string, respObj interface{}) error {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: UpdateCloudProviderInput!) {
			updateCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, fields),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		UpdateCloudProvider struct {
			CloudProvider *json.RawMessage
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*res.UpdateCloudProvider.CloudProvider, respObj)
	if err != nil {
		return err
	}

	return nil
}

func (ac *CloudProviderClient) ListCloudProviders(limit int, offset int) ([]*graphql.CloudProvider, *graphql.PageInfo, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query {
			cloudProviders(limit: %[3]d, offset: %[4]d) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, commonCloudProviderFields, paginationFields, limit, offset),
	}

	res := struct {
		CloudProviders struct {
			Nodes    []*graphql.CloudProvider
			PageInfo *graphql.PageInfo
		}
	}{}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.CloudProviders.Nodes, res.CloudProviders.PageInfo, nil
}

func (c *CloudProviderClient) getCloudProviderByName(name string, fields string, respObj interface{}) error {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			cloudProviderByName(name: "%[1]s") {
				%[2]s
			}
		}`, name, fields),
	}

	res := &struct {
		CloudProviderByName interface{}
	}{
		CloudProviderByName: respObj,
	}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return err
	}

	return nil
}

func (c *CloudProviderClient) createCloudProvider(input interface{}, fields string, respObj interface{}) error {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, fields),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider *json.RawMessage
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*res.CreateCloudProvider.CloudProvider, respObj)
	if err != nil {
		return err
	}

	return nil
}
