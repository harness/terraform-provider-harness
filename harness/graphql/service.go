package graphql

// Get the client for interacting with Harness Applications
// func (c *ApiClient) Services() *ServiceClient {
// 	return &ServiceClient{
// 		APIClient: c,
// 	}
// }

// func (c *ServiceClient) GetServiceById(id string) (*Service, error) {

// 	query := &GraphQLQuery{
// 		Query: fmt.Sprintf(`query($id: String!) {
// 			service(secretId: $id) {
// 				%s
// 			}
// 		}`, serviceFields),
// 		Variables: map[string]interface{}{
// 			"id": id,
// 		},
// 	}

// 	res := &struct {
// 		Service EncryptedFile
// 	}{}

// 	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res.Secret, nil
// }

// const (
// 	serviceFields = `
// 	id
// 	name
// 	artifactType
// 	deploymentType
// 	createdAt
// 	createdBy {
// 		id
// 		name
// 		email
// 	}
// 	artifactSources{
// 		id
// 		name
// 		createdAt
// 	}
// `
// )
