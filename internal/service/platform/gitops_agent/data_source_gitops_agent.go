package gitops_agent

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
		Description: "Datasource for creating a Harness Gitops Agents.",

		ReadContext: dataSourceGitopsAgentRead,

		Schema: map[string]*schema.Schema{
			"account_identifier": {
				Description: "account identifier of the agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_identifier": {
				Description: "project identifier of the agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_identifier": {
				Description: "organization identifier of the agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "identifier of the agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "name of the agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "description of the agent.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "type of the agent.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tags": {
				Description: "tags for the agent.",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"metadata": {
				Description: "tags for the agent.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Description: "namespace of the agent.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"high_availability": {
							Description: "If the agent should be high availability.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					}},
			},
		},
	}
	return resource
}

func dataSourceGitopsAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("identifier").(string)

	resp, httpResp, err := c.AgentServiceApi.AgentServiceGet(ctx, agentIdentifier, &nextgen.AgentServiceApiAgentServiceGetOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(d.Get("org_identifier").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_identifier").(string)),
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
