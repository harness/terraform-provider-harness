package chaos

import "github.com/harness/harness-go-sdk/harness/chaos/graphql/model"

// WithChaosHubDescription sets the description for the Chaos Hub
func WithChaosHubDescription(description string) func(*model.ChaosHubRequest) *model.ChaosHubRequest {
	return func(req *model.ChaosHubRequest) *model.ChaosHubRequest {
		req.Description = &description
		return req
	}
}

// WithChaosHubTags sets the tags for the Chaos Hub
func WithChaosHubTags(tags []string) func(*model.ChaosHubRequest) *model.ChaosHubRequest {
	return func(req *model.ChaosHubRequest) *model.ChaosHubRequest {
		req.Tags = tags
		return req
	}
}

// WithChaosHubRepoName sets the repository name for the Chaos Hub
func WithChaosHubRepoName(repoName string) func(*model.ChaosHubRequest) *model.ChaosHubRequest {
	return func(req *model.ChaosHubRequest) *model.ChaosHubRequest {
		repoNameCopy := repoName // Create a copy to ensure we're not taking the address of a loop variable
		req.RepoName = &repoNameCopy
		return req
	}
}

// WithChaosHubConnectorScope sets the connector scope for the Chaos Hub
func WithChaosHubConnectorScope(scope model.ConnectorScope) func(*model.ChaosHubRequest) *model.ChaosHubRequest {
	return func(req *model.ChaosHubRequest) *model.ChaosHubRequest {
		req.ConnectorScope = scope
		return req
	}
}
