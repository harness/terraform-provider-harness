package cd

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
)

func (c *ApplicationClient) UpdateGitSyncConfig(config *graphql.UpdateApplicationGitSyncConfigInput) (*graphql.GitSyncConfig, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: UpdateApplicationGitSyncConfigInput!) {
			updateApplicationGitSyncConfig(input: $input) {
				gitSyncConfig {
					%s
				}
			}
		}`, getGitSyncConfigFields()),
		Variables: map[string]interface{}{
			"input": &config,
		},
	}

	res := &struct {
		UpdateApplicationGitSyncConfig struct {
			GitSyncConfig graphql.GitSyncConfig
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateApplicationGitSyncConfig.GitSyncConfig, nil
}

func (c *ApplicationClient) RemoveGitSyncConfig(applicationId string) error {
	query := &GraphQLQuery{
		Query: `mutation($input: RemoveApplicationGitSyncConfigInput!) {
			removeApplicationGitSyncConfig(input: $input) {
				application {
					id
				}
			}
		}`,
		Variables: map[string]interface{}{
			"input": map[string]interface{}{
				"applicationId": applicationId,
			},
		},
	}

	res := &struct {
		RemoveApplicationGitSyncConfig struct {
			ClientMutationId string
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

func getGitSyncConfigFields() string {
	return fmt.Sprintf(`
	branch
	repositoryName
	syncEnabled
	gitConnector {
		%s
	}
`, gitConnectorFields)
}
