package agent

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAgentTokenSensitive(t *testing.T) {
	assert.True(t, ResourceGitopsAgent().Schema["agent_token"].Sensitive)
	assert.True(t, DataSourceGitopsAgent().Schema["agent_token"].Sensitive)
}

func TestStoreAgentTokenSchema(t *testing.T) {
	s, ok := ResourceGitopsAgent().Schema["store_agent_token"]
	require.True(t, ok)
	assert.Equal(t, schema.TypeBool, s.Type)
	assert.True(t, s.Optional)
	assert.Equal(t, true, s.Default)

	_, onDataSource := DataSourceGitopsAgent().Schema["store_agent_token"]
	assert.False(t, onDataSource)
}

func TestStoreAgentTokenTruePersistsToken(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, map[string]interface{}{
		"identifier":        "agent",
		"name":              "agent",
		"type":              "MANAGED_ARGO_PROVIDER",
		"store_agent_token": true,
	})

	readGitopsAgentResource(d, testAgentWithPrivateKey("secret-token"))
	assert.Equal(t, "secret-token", d.Get("agent_token"))
}

func TestStoreAgentTokenOmittedPersistsToken(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, map[string]interface{}{
		"identifier": "agent",
		"name":       "agent",
		"type":       "MANAGED_ARGO_PROVIDER",
	})

	readGitopsAgentResource(d, testAgentWithPrivateKey("secret-token"))
	assert.Equal(t, "secret-token", d.Get("agent_token"))
}

func TestStoreAgentTokenFalseOmitsToken(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, map[string]interface{}{
		"identifier":        "agent",
		"name":              "agent",
		"type":              "MANAGED_ARGO_PROVIDER",
		"store_agent_token": false,
	})

	readGitopsAgentResource(d, testAgentWithPrivateKey("secret-token"))
	assert.Empty(t, d.Get("agent_token"))
}

func TestStoreAgentTokenFalseScrubsExistingToken(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, map[string]interface{}{
		"identifier":        "agent",
		"name":              "agent",
		"type":              "MANAGED_ARGO_PROVIDER",
		"store_agent_token": false,
		"agent_token":       "already-in-state",
	})

	readGitopsAgentResource(d, testAgentWithPrivateKey(""))
	assert.Empty(t, d.Get("agent_token"))
}

func TestStoreAgentTokenTruePreservesExistingToken(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, map[string]interface{}{
		"identifier":        "agent",
		"name":              "agent",
		"type":              "MANAGED_ARGO_PROVIDER",
		"store_agent_token": true,
		"agent_token":       "already-in-state",
	})

	readGitopsAgentResource(d, testAgentWithPrivateKey(""))
	assert.Equal(t, "already-in-state", d.Get("agent_token"))
}

func TestReadAgentDataSourceStillStoresToken(t *testing.T) {
	d := schema.TestResourceDataRaw(t, DataSourceGitopsAgent().Schema, map[string]interface{}{
		"identifier": "agent",
	})

	readAgent(d, testAgentWithPrivateKey("secret-token"))
	assert.Equal(t, "secret-token", d.Get("agent_token"))
}

func testAgentWithPrivateKey(token string) *nextgen.V1Agent {
	agent := &nextgen.V1Agent{
		Identifier:        "agent",
		Name:              "agent",
		AccountIdentifier: "acc",
		Metadata:          &nextgen.V1AgentMetadata{Namespace: "ns"},
	}
	if token != "" {
		agent.Credentials = &nextgen.V1AgentCredentials{PrivateKey: token}
	}
	return agent
}
