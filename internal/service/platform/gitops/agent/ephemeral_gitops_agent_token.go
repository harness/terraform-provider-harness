package agent

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ ephemeral.EphemeralResource              = &ephemeralGitopsAgentToken{}
	_ ephemeral.EphemeralResourceWithConfigure = &ephemeralGitopsAgentToken{}
)

// NewEphemeralGitopsAgentToken returns the ephemeral resource that reads a GitOps agent token
// without persisting it to Terraform state or plan files.
func NewEphemeralGitopsAgentToken() ephemeral.EphemeralResource {
	return &ephemeralGitopsAgentToken{}
}

type ephemeralGitopsAgentToken struct {
	session *internal.Session
}

type ephemeralGitopsAgentTokenModel struct {
	Identifier types.String `tfsdk:"identifier"`
	OrgId      types.String `tfsdk:"org_id"`
	ProjectId  types.String `tfsdk:"project_id"`
	AgentToken types.String `tfsdk:"agent_token"`
}

func (e *ephemeralGitopsAgentToken) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_platform_gitops_agent_token"
}

func (e *ephemeralGitopsAgentToken) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Ephemeral resource for reading the token of a Harness GitOps Agent. The token is available to the configuration during the run but is never written to state or plan files. The token is only returned while the agent has never connected to Harness, and the user associated to the API key must have GitOps Agent Edit permissions.",
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Description: "Identifier of the GitOps agent.",
				Required:    true,
			},
			"org_id": schema.StringAttribute{
				Description: "Organization identifier of the GitOps agent.",
				Optional:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "Project identifier of the GitOps agent.",
				Optional:    true,
			},
			"agent_token": schema.StringAttribute{
				Description: "Agent token to be used for authentication of the agent with Harness.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (e *ephemeralGitopsAgentToken) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	session, ok := req.ProviderData.(*internal.Session)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("Expected *internal.Session, got %T.", req.ProviderData))
		return
	}
	e.session = session
}

func (e *ephemeralGitopsAgentToken) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data ephemeralGitopsAgentTokenModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if e.session == nil {
		resp.Diagnostics.AddError("Provider not configured", "The Harness provider session is not available.")
		return
	}

	c, ctx := e.session.GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())

	agent, httpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, data.Identifier.ValueString(), c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(data.OrgId.ValueString()),
		ProjectIdentifier: optional.NewString(data.ProjectId.ValueString()),
		WithCredentials:   optional.NewBool(true),
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			resp.Diagnostics.AddError("GitOps agent not found", fmt.Sprintf("No GitOps agent with identifier %q was found.", data.Identifier.ValueString()))
			return
		}
		resp.Diagnostics.AddError("Error reading GitOps agent", err.Error())
		return
	}

	if agent.Credentials == nil || agent.Credentials.PrivateKey == "" {
		resp.Diagnostics.AddError(
			"Agent token not available",
			"Harness returns the agent token only while the agent has never connected. The agent may already be connected, or the API key may lack GitOps Agent Edit permissions.",
		)
		return
	}

	data.AgentToken = types.StringValue(agent.Credentials.PrivateKey)
	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
