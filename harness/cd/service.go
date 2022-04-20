package cd

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

type ServiceClient struct {
	ApiClient *ApiClient
}

func (ac *ServiceClient) ListServices(limit int, offset int, filters []*graphql.ServiceFilter) ([]*graphql.Service, *graphql.PageInfo, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($filters: [ServiceFilter]) {
			services(limit: %[3]d, offset: %[4]d. filters: $filters) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, standardServiceFields, paginationFields, limit, offset),
		Variables: map[string]interface{}{
			"filters": filters,
		},
	}

	res := struct {
		Services struct {
			Nodes    []*graphql.Service
			PageInfo *graphql.PageInfo
		}
	}{}

	err := ac.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.Services.Nodes, res.Services.PageInfo, nil
}

func (ac *ServiceClient) ListServicesByApplicationId(appId string, limit int, offset int) ([]*graphql.Service, *graphql.PageInfo, error) {
	filter := []*graphql.ServiceFilter{
		{
			Application: &graphql.IdFilter{
				Operator: graphql.IdOperatorTypes.Equals,
				Values:   []string{appId},
			},
		},
	}
	return ac.ListServices(limit, offset, filter)
}

func (ac *ServiceClient) GetServiceById(id string) (*graphql.Service, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($id: String!) {
			service(serviceId: $id) {
				%s
			}
		}`, standardServiceFields),
		Variables: map[string]interface{}{
			"id": id,
		},
	}

	res := struct {
		Service graphql.Service
	}{}
	err := ac.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil
	}

	return &res.Service, nil
}

func (ac *ServiceClient) GetServiceByName(appId string, name string) (*graphql.Service, error) {
	limit := 25
	offset := 0
	hasMore := true

	for hasMore {
		services, pagination, err := ac.ListServicesByApplicationId(appId, limit, offset)
		if err != nil {
			return nil, err
		}

		for _, service := range services {
			if strings.EqualFold(service.Name, name) {
				return service, nil
			}
		}

		hasMore = pagination.HasMore
		offset += limit
	}

	return nil, nil
}

const (
	standardServiceFields = `
	name
	artifactType
	createdAt
	createdBy {
		email
		externalUserId
		id
		name
	}
	deploymentType
	description
	id
	name
	tags {
		name
		value
	}
`
)
