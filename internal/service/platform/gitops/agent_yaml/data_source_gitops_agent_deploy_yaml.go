package agent_yaml

import (
	"context"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGitopsAgentDeployYaml() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for fetching a Harness Gitops Agents.",

		ReadContext: dataSourceGitopsAgentDeployYamlRead,

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
			"namespace": {
				Description: "The k8s namespace that the GitOps agent resides in.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "Deployment YAML of the GitOps agent.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
	return resource
}

func dataSourceGitopsAgentDeployYamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	agentIdentifier := d.Get("identifier").(string)

	resp, httpResp, err := c.AgentApi.AgentServiceForServerGetDeployYaml(ctx, agentIdentifier, c.AccountId, &nextgen.AgentsApiAgentServiceForServerGetDeployYamlOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		Namespace:         optional.NewString(d.Get("namespace").(string)),
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
	readAgentYaml(agentIdentifier, d, resp)
	return nil
}

func readAgentYaml(agentIdentifier string, d *schema.ResourceData, yaml string) {
	d.SetId(agentIdentifier)
	d.Set("yaml", yaml)
}
