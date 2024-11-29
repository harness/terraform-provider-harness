package agent

import (
	"context"
	"fmt"

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
		Description: "Datasource for fetching a Harness GitOps Agent.",

		ReadContext: dataSourceGitopsAgentRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent.",
				Type:        schema.TypeString,
				Computed:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release.",
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
							Description: "The kubernetes namespace where the agent should be installed.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"high_availability": {
							Description: "Indicates if the agent is deployed in HA mode.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"is_namespaced": {
							Description: "Indicates if the agent is namespaced.",
							Type:        schema.TypeBool,
							Optional:    true,
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
			"with_credentials": {
				Description: "Specify whether to retrieve the gitops agent's token. (The field agent_token will be populated only if the agent has never connected to Harness before). For retrieval of this information, the user associated to the token being used must have Gitops Agent Edit permissions",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_authenticated": {
				Description: "This computed field specifies if the referenced agent ever successfully connected and was authenticated to harness. Note that this is different from whether the agent is currently connected. <b>Set with_credentials to true to allow computing of this field.</b> For retrieval of this information, the user associated to the token being used must have Gitops Agent Edit permissions",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"prefixed_identifier": {
				Description: "Prefixed identifier of the GitOps agent. Agent identifier prefixed with scope of the agent",
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
		WithCredentials:   optional.NewBool(d.Get("with_credentials").(bool)),
	})

	if err != nil && (httpResp == nil || httpResp.StatusCode != 404) {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil || httpResp.StatusCode == 404 {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	readAgent(d, &resp)
	err = readAgentDataSourceOnlyFields(d, &resp)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readAgentDataSourceOnlyFields(d *schema.ResourceData, agent *nextgen.V1Agent) error {
	err := d.Set("with_credentials", d.Get("with_credentials"))
	if err != nil {
		return fmt.Errorf("error setting with_credentials field while reading response: %w", err)
	}
	if d.Get("with_credentials").(bool) {
		err = d.Set("is_authenticated", agent.Credentials != nil && agent.Credentials.PrivateKey == "")
		if err != nil {
			return fmt.Errorf("error setting is_authenticated field while reading response: %w", err)
		}
	}
	return nil
}
