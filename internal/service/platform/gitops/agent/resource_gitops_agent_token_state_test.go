package agent

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func newAgentData(t *testing.T, skip bool, existingToken string) *schema.ResourceData {
	t.Helper()
	raw := map[string]interface{}{
		"identifier":                "a1",
		"name":                      "a1",
		"type":                      "MANAGED_ARGO_PROVIDER",
		"skip_agent_token_in_state": skip,
	}
	d := schema.TestResourceDataRaw(t, ResourceGitopsAgent().Schema, raw)
	if existingToken != "" {
		require.NoError(t, d.Set("agent_token", existingToken))
	}
	return d
}

func agentWithKey(key string) *nextgen.V1Agent {
	return &nextgen.V1Agent{
		Identifier:  "a1",
		Name:        "a1",
		Metadata:    &nextgen.V1AgentMetadata{},
		Credentials: &nextgen.V1AgentCredentials{PrivateKey: key},
	}
}

func TestReadAgent_DefaultStoresToken(t *testing.T) {
	d := newAgentData(t, false, "")
	readAgent(d, agentWithKey("KEY"))
	require.Equal(t, "KEY", d.Get("agent_token"))
}

func TestReadAgent_DefaultKeepsExistingTokenWhenResponseHasNone(t *testing.T) {
	d := newAgentData(t, false, "OLD")
	readAgent(d, agentWithKey(""))
	require.Equal(t, "OLD", d.Get("agent_token"))
}

func TestReadAgent_SkipDoesNotStoreToken(t *testing.T) {
	d := newAgentData(t, true, "")
	readAgent(d, agentWithKey("KEY"))
	require.Equal(t, "", d.Get("agent_token"))
}

func TestReadAgent_SkipScrubsTokenAlreadyInState(t *testing.T) {
	d := newAgentData(t, true, "OLD")
	readAgent(d, agentWithKey(""))
	require.Equal(t, "", d.Get("agent_token"))
}

func TestAgentTokenIsSensitive(t *testing.T) {
	require.True(t, ResourceGitopsAgent().Schema["agent_token"].Sensitive)
	require.True(t, DataSourceGitopsAgent().Schema["agent_token"].Sensitive)
}
