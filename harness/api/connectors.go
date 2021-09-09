package api

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

type ConnectorClient struct {
	APIClient *Client
}

// Get the client for interacting with Harness Applications
func (c *Client) Connectors() *ConnectorClient {
	return &ConnectorClient{
		APIClient: c,
	}
}

func (c *ConnectorClient) GetGitConnectorById(id string) (*graphql.GitConnector, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($connectorId: String!) {
			connector(connectorId: $connectorId) {
				... on GitConnector {
					%s
				}
			}
		}`, gitConnectorFields),
		Variables: map[string]interface{}{
			"connectorId": id,
		},
	}

	res := struct {
		Connector graphql.GitConnector
	}{}
	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		if strings.Contains(err.Error(), "Connector does not exist") {
			return nil, nil
		}
		return nil, err
	}

	return &res.Connector, nil
}

func (c *ConnectorClient) CreateGitConnector(connector *graphql.GitConnectorInput) (*graphql.GitConnector, error) {

	// Set defaults
	input := &graphql.CreateConnectorInput{
		ConnectorType: graphql.ConnectorTypes.Git,
		GitConnector:  connector,
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($connector: CreateConnectorInput!) {
			createConnector(input: $connector) {
				connector {
				... on GitConnector {
						%s
					}
				}
			}
		}`, gitConnectorFields),
		Variables: map[string]interface{}{
			"connector": &input,
		},
	}

	res := &struct {
		CreateConnector struct {
			Connector graphql.GitConnector
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateConnector.Connector, nil
}

func (c *ConnectorClient) UpdateGitConnector(id string, connector *graphql.GitConnectorInput) (*graphql.GitConnector, error) {

	// Set defaults
	input := &graphql.UpdateConnectorInput{ConnectorId: id}
	input.ConnectorType = graphql.ConnectorTypes.Git
	input.GitConnector = connector

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($connector: UpdateConnectorInput!) {
			updateConnector(input: $connector) {
				connector {
				... on GitConnector {
						%s
					}
				}
			}
		}`, gitConnectorFields),
		Variables: map[string]interface{}{
			"connector": &input,
		},
	}

	res := &struct {
		UpdateConnector struct {
			Connector graphql.GitConnector
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateConnector.Connector, nil
}

func (c *ConnectorClient) GetGitConnectorByName(name string) (*graphql.GitConnector, error) {
	limit := 10
	offset := 0
	hasMore := true

	for hasMore {
		connectors, _, err := c.ListGitConnectors(limit, offset)
		if err != nil {
			return nil, err
		}

		for _, connector := range connectors {
			if connector.Name == name {
				return connector, nil
			}
		}

		hasMore = len(connectors) == limit
		offset += limit
	}

	return nil, nil
}

func (c *ConnectorClient) ListGitConnectors(limit int, offset int) ([]*graphql.GitConnector, *graphql.PageInfo, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($limit: Int!, $offset: Int) {
			connectors(limit: $limit, offset: $offset) {
				nodes {
					... on GitConnector {
						%[1]s
					}
				}
				%[2]s
			}
		}`, gitConnectorFields, paginationFields),
		Variables: map[string]interface{}{
			"limit":  limit,
			"offset": offset,
		},
	}

	res := struct {
		Connectors struct {
			Nodes    []*graphql.GitConnector
			PageInfo *graphql.PageInfo
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.Connectors.Nodes, res.Connectors.PageInfo, nil
}

func (c *ConnectorClient) DeleteConnector(id string) error {

	query := &GraphQLQuery{
		Query: `mutation($connector: DeleteConnectorInput!) {
			deleteConnector(input: $connector) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"connector": struct {
				ConnectorId string `json:"connectorId"`
			}{
				ConnectorId: id,
			},
		},
	}

	res := &struct {
		DeleteConnector struct {
			ClientMutationId string
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

var gitConnectorFields = `
	URL
	branch
	createdAt
	createdBy {
		id 
		name
		email
		id
	}
	customCommitDetails{
		authorName
		authorEmailId
		commitMessage
	}
	delegateSelectors
	description
	generateWebhookUrl
	id
	name
	passwordSecretId
	sshSettingId
	urlType
	usageScope {
		appEnvScopes {
			application {
				appId
				filterType
			}
			environment {
				envId
				filterType
			}
		}
	}
	userName
	webhookUrl
`
