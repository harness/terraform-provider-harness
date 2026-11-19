package agent

import (
	"context"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
)

// regenerateAgentCredentials regenerates the agent's private key (the agent token) via
// POST /gitops/api/v1/agents/{identifier}/credentials and returns the new key. A response that
// carries no key is an error: the server has already rotated, so reporting success would leave
// the caller with a key that no longer works.
func regenerateAgentCredentials(ctx context.Context, session *internal.Session, identifier, orgID, projectID string) (string, error) {
	c, ctx := session.GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	opts := &nextgen.AgentsApiAgentServiceForServerRegenerateCredentialsWithScopeOpts{}
	if orgID != "" {
		opts.OrgIdentifier = optional.NewString(orgID)
	}
	if projectID != "" {
		opts.ProjectIdentifier = optional.NewString(projectID)
	}

	agent, httpResp, err := c.AgentApi.AgentServiceForServerRegenerateCredentialsWithScope(ctx, identifier, c.AccountId, opts)
	if err != nil {
		status, detail := "", err.Error()
		if httpResp != nil {
			status = httpResp.Status
		}
		var swaggerErr nextgen.GenericSwaggerError
		if errors.As(err, &swaggerErr) && len(swaggerErr.Body()) > 0 {
			detail = string(swaggerErr.Body())
		}
		return "", fmt.Errorf("regenerating agent credentials for %q failed with status %s: %s", identifier, status, detail)
	}
	if agent.Credentials == nil || agent.Credentials.PrivateKey == "" {
		return "", fmt.Errorf("regenerating agent credentials for %q returned no private key; the previous key may already be invalid, so check the agent in Harness and rotate again", identifier)
	}
	return agent.Credentials.PrivateKey, nil
}
