package app_project

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

func DatasourceGitopsAppProjectMapping() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing the Harness GitOps Application Project Mappings.",

		ReadContext: datasourceGitopsAppProjectMappingRead,
		Importer:    helpers.GitopsAppProjectMappingImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent's Application Project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization identifier of the GitOps agent's Application Project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps agent's Application Project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"agent_id": {
				Description: "Agent identifier for which the ArgoCD and Harness project mapping is to be created.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps Application Project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"argo_project_name": {
				Description: "ArgoCD Project name which is to be mapped to the Harness project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
	return resource
}

func datasourceGitopsAppProjectMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	argo_proj_name := d.Get("argo_project_name").(string)
	if identifier == argo_proj_name {
		identifier = ""
	}
	resp, httpResp, err := c.ProjectMappingsApi.AppProjectMappingServiceGetAppProjectMappingV2(ctx, agentIdentifier, identifier, &nextgen.ProjectMappingsApiAppProjectMappingServiceGetAppProjectMappingV2Opts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		ArgoProjectName:   optional.NewString(argo_proj_name),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	readAppProjectMapping(d, &resp)
	return nil
}
