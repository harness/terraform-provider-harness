package cd

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

type TagClient struct {
	ApiClient *ApiClient
}

func (c *TagClient) AttachTag(input *graphql.AttachTagInput) (*graphql.AttachTagPayload, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: AttachTagInput!) {
			attachTag(input: $input) {
				clientMutationId
				tagLink {
					%[1]s
				}
			}
		}`, tagLinkFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		AttachTag graphql.AttachTagPayload
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.AttachTag, nil
}

func (c *TagClient) DetachTag(input *graphql.DetachTagInput) error {

	query := &GraphQLQuery{
		Query: `mutation($input: DetachTagInput!) {
			detachTag(input: $input) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		DetachTag struct {
			ClientMutationId string
		}
	}{}

	return c.ApiClient.ExecuteGraphQLQuery(query, &res)
}

const tagLinkFields = `
appId
entityId
entityType
name
value
`
