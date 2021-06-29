package api

import (
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

const (
	secretFileId = "2WnPVgLGSZW6KbApZuxeaw"
)

func getExampleUsageScopes() *graphql.UsageScope {
	var scopes []*graphql.AppEnvScope

	scope := &graphql.AppEnvScope{
		Application: &graphql.AppScopeFilter{
			FilterType: graphql.ApplicationFilterTypes.All,
		},
		Environment: &graphql.EnvScopeFilter{
			FilterType: graphql.EnvironmentFilterTypes.NonProduction,
		},
	}
	scopes = append(scopes, scope)

	return &graphql.UsageScope{
		AppEnvScopes: scopes,
	}
}

func deleteSecret(id string, secretType graphql.SecretType) error {
	client := getClient()

	return client.Secrets().DeleteSecret(id, secretType)
}
