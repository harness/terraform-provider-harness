package api

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

// Helper type for accessing all application related crud methods
type ApplicationClient struct {
	APIClient *Client
}

// Get the client for interacting with Harness Applications
func (c *Client) Applications() *ApplicationClient {
	return &ApplicationClient{
		APIClient: c,
	}
}

// CRUD
func (ac *ApplicationClient) GetApplicationById(id string) (*graphql.Application, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($applicationId: String!) {
			application(applicationId: $applicationId) {
				%s
			}
		}`, standardApplicationFields),
		Variables: map[string]interface{}{
			"applicationId": id,
		},
	}

	res := struct {
		Application graphql.Application
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	// Need to fix https://harness.atlassian.net/browse/PL-19934
	if err != nil {
		if strings.Contains(err.Error(), "User not authorized") {
			if cacApp, err := ac.APIClient.ConfigAsCode().GetApplicationById(id); err != nil {
				return nil, err
			} else if cacApp.IsEmpty() {
				return nil, nil
			}
		}
		return nil, err
	}

	return &res.Application, nil
}

func (ac *ApplicationClient) GetApplicationByName(name string) (*graphql.Application, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($name: String!) {
			applicationByName(name: $name) {
				%s	
			}
		}`, standardApplicationFields),
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res := &struct {
		ApplicationByName graphql.Application
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	// Need to fix https://harness.atlassian.net/browse/PL-19934
	if err != nil {
		// if strings.Contains(err.Error(), "User not authorized") {
		// 	if cacApp, err := ac.APIClient.ConfigAsCode().GetApplicationByName(name); err != nil {
		// 		return nil, err
		// 	} else if cacApp == nil {
		// 		return nil, nil
		// 	}
		// }
		return nil, err
	}

	return &res.ApplicationByName, nil
}

func (ac *ApplicationClient) CreateApplication(input *graphql.Application) (*graphql.Application, error) {

	query := &GraphQLQuery{
		Query: `mutation createapp($app: CreateApplicationInput!) {
			createApplication(input: $app) {
				clientMutationId
				application {
					id
					name
					description
				}
			}
		}`,
		Variables: map[string]interface{}{
			"app": &input,
		},
	}

	res := &struct {
		CreateApplication graphql.CreateApplicationPayload
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return res.CreateApplication.Application, nil
}

func (ac *ApplicationClient) DeleteApplication(id string) error {

	query := &GraphQLQuery{
		Query: `mutation deleteApp($app: DeleteApplicationInput!) {
			deleteApplication(input: $app) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"app": &graphql.DeleteApplicationInput{
				ApplicationId: id,
			},
		},
	}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &struct{}{})

	return err
}

func (ac *ApplicationClient) UpdateApplication(input *graphql.UpdateApplicationInput) (*graphql.Application, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation updateapp($app: UpdateApplicationInput!) {
			updateApplication(input: $app) {
				clientMutationId
				application {
					%s
				}
			}
		}`, standardApplicationFields),
		Variables: map[string]interface{}{
			"app": &input,
		},
	}

	res := struct {
		UpdateApplication graphql.UpdateApplicationPayload
	}{}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return res.UpdateApplication.Application, nil
}

func (ac *ApplicationClient) ListApplications(limit int, offset int) ([]*graphql.Application, *graphql.PageInfo, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query {
			applications(limit: %[3]d, offset: %[4]d) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, standardApplicationFields, paginationFields, limit, offset),
	}

	res := struct {
		Applications struct {
			Nodes    []*graphql.Application
			PageInfo *graphql.PageInfo
		}
	}{}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.Applications.Nodes, res.Applications.PageInfo, nil
}

const (
	standardApplicationFields = `
	id
	name
	description
	createdBy {
		id
		name
		email
	}
	gitSyncConfig {
		branch
		gitConnector {
			id
			name
			branch
		}
		repositoryName
		syncEnabled
	}
	tags {
		name
		value
	}	
	`
)
