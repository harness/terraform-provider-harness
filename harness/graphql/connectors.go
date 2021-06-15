package graphql

import "fmt"

// Get the client for interacting with Harness Applications
func (c *ApiClient) Connectors() *ConnectorClient {
	return &ConnectorClient{
		APIClient: c,
	}
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

func (c *ConnectorClient) GetGitConnectorById(id string) (*GitConnector, error) {

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
		Connector GitConnector
	}{}
	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.Connector, nil
}

func (c *ConnectorClient) CreateGitConnector(connector *GitConnectorInput) (*GitConnector, error) {

	// Set defaults
	input := &CreateConnectorInput{
		ConnectorType: ConnectorTypes.Git,
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
			Connector GitConnector
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateConnector.Connector, nil
}

func (c *ConnectorClient) UpdateGitConnector(id string, connector *GitConnectorInput) (*GitConnector, error) {

	// Set defaults
	input := &UpdateConnectorInput{ConnectorId: id}
	input.ConnectorType = ConnectorTypes.Git
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
			Connector GitConnector
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

func (d *CustomCommitDetails) IsEmpty() bool {
	if d == nil {
		return true
	}

	return (d.AuthorEmailId + d.AuthorName + d.CommitMessage) == ""
}
