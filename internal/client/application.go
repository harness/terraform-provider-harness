package client

type ApplicationClient struct {
	APIClient *ApiClient
}

func (ac *ApplicationClient) GetApplicationById(id string) (*Application, error) {
	query := &GraphQLQuery{
		Query: `query($applicationId: String!) {
			application(applicationId: $applicationId) {
				id
      	name
				createdBy {
					id
					name
					email
				}
      	tags {
        	name
        	value
      	}	
			}
		}`,
		Variables: map[string]interface{}{
			"applicationId": id,
		},
	}

	res, err := ac.APIClient.ExecuteGraphQLQuery(query)

	if err != nil {
		return nil, err
	}

	return res.Data.Application, nil
}

func (ac *ApplicationClient) GetApplicationByName(name string) (*Application, error) {
	query := &GraphQLQuery{
		Query: `query($name: String!) {
			applicationByName(name: $name) {
				id
      	name
				createdBy {
					id
					name
					email
				}
      	tags {
        	name
        	value
      	}	
			}
		}`,
		Variables: map[string]interface{}{
			"name": name,
		},
	}

	res, err := ac.APIClient.ExecuteGraphQLQuery(query)

	if err != nil {
		return nil, err
	}

	return res.Data.ApplicationByName, nil
}

func (ac *ApplicationClient) CreateApplication(input *CreateApplicationInput) (*Application, error) {

	query := &GraphQLQuery{
		Query: `mutation createapp($app: CreateApplicationInput!) {
			createApplication(input: $app) {
				clientMutationId
				application {
					name
					id
				}
			}
		}`,
		Variables: map[string]interface{}{
			"app": &input,
		},
	}

	res, err := ac.APIClient.ExecuteGraphQLQuery(query)

	if err != nil {
		return nil, err
	}

	return res.Data.CreateApplication.Application, nil
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

	_, err := ac.APIClient.ExecuteGraphQLQuery(query)

	return err
}

// Get the client for interacting with Harness Applications
func (c *ApiClient) Applications() *ApplicationClient {
	return &ApplicationClient{
		APIClient: c,
	}
}
