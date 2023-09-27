package agent

import (
	"context"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitopsAgent() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for fetching a Harness Gitops Agents.",

		ReadContext: dataSourceGitopsAgentRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Organization identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the GitOps agent.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the GitOps agent.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "Default: \"AGENT_TYPE_UNSET\"\nEnum: \"AGENT_TYPE_UNSET\" \"CONNECTED_ARGO_PROVIDER\" \"MANAGED_ARGO_PROVIDER\"",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags for the GitOps agents. These can be used to search or filter the GitOps agents.",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"metadata": {
				Description: "Metadata of the agent.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Description: "The k8s namespace that this agent resides in.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"high_availability": {
							Description: "Indicates if the deployment should be deployed using the deploy-ha.yaml",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
			"agent_token": {
				Description: "Agent token to be used for authentication of the agent with Harness.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"operator": {
				Description: "The Operator to use for the Harness GitOps agent. Enum: \"ARGO\" \"FLAMINGO\"",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
	return resource
}

func dataSourceGitopsAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("identifier").(string)

	resp, httpResp, err := c.AgentApi.AgentServiceForServerGet(ctx, agentIdentifier, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	readAgent(d, &resp)
	return nil
}
