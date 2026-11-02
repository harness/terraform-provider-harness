package agent

import (
	"context"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
)

// regenerateAgentCredentials regenerates the agent's private key (the agent token) via
// POST /gitops/api/v1/agents/{identifier}/credentials and returns the new key.
func regenerateAgentCredentials(ctx context.Context, session *internal.Session, identifier, orgID, projectID string) (string, error) {
	c, ctx := session.GetPlatformClientWithContext(ctx)

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
	if agent.Credentials == nil {
		return "", nil
	}
	return agent.Credentials.PrivateKey, nil
}
