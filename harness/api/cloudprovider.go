package api

import (
	"encoding/json"
	"fmt"
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
