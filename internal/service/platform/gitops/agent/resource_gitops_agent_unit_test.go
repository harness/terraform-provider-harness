package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgentTokenSensitive(t *testing.T) {
	assert.True(t, ResourceGitopsAgent().Schema["agent_token"].Sensitive)
	assert.True(t, DataSourceGitopsAgent().Schema["agent_token"].Sensitive)
}
