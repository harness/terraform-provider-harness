package client

import "fmt"

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

// CRUD
func (ac *ApplicationClient) GetApplicationById(id string) (*Application, error) {
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
		Application Application
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.Application, nil
}

func (ac *ApplicationClient) GetApplicationByName(name string) (*Application, error) {
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
		ApplicationByName Application
	}{}
	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.ApplicationByName, nil
}

func (ac *ApplicationClient) CreateApplication(input *Application) (*Application, error) {

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
		CreateApplication CreateApplicationPayload
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
			"app": &DeleteApplicationInput{
				ApplicationId: id,
			},
		},
	}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &struct{}{})

	return err
}

// Get the client for interacting with Harness Applications
func (c *ApiClient) Applications() *ApplicationClient {
	return &ApplicationClient{
		APIClient: c,
	}
}

func (ac *ApplicationClient) UpdateApplication(input *UpdateApplicationInput) (*Application, error) {

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
		UpdateApplication UpdateApplicationPayload
	}{}

	err := ac.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return res.UpdateApplication.Application, nil
}
