package api

import (
	"fmt"

	"github.com/micahlmartin/terraform-provider-harness/harness/api/graphql"
)

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
