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

func ResourceGitopsAppProjectMapping() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing Harness GitOps Application Project Mappings.",

		CreateContext: resourceGitopsAppProjectMappingCreate,
		ReadContext:   resourceGitopsAppProjectMappingRead,
		UpdateContext: resourceGitopsAppProjectMappingUpdate,
		DeleteContext: resourceGitopsAppProjectMappingDelete,
		Importer:      helpers.GitopsAppProjectMappingImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps agent's Application Project.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
				ForceNew:    true,
			},
			"identifier": {
				Description: "Identifier of the GitOps Application Project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"argo_project_name": {
				Description: "ArgoCD Project name which is to be mapped to the Harness project.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
	return resource
}

func resourceGitopsAppProjectMappingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	createAppProjectMappingRequest := buildCreateAppProjectMappingRequest(d)
	agentIdentifier := d.Get("agent_id").(string)
	resp, httpResp, err := c.ProjectMappingsApi.AppProjectMappingServiceCreateV2(ctx, *createAppProjectMappingRequest, agentIdentifier)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	readAppProjectMapping(d, &resp)
	return nil
}

func resourceGitopsAppProjectMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	argo_proj_name := d.Get("argo_project_name").(string)
	// During import we are using argo_project_name as identifier not the actual identifier which is mongo id
	// So we are not fetching mapping by mongo id but by argo_project_name, agent_id, account_id, org_id and project_id.
	// argo_project_name, agent_id, account_id, org_id and project_id uniquely identify mapping.
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

func resourceGitopsAppProjectMappingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	updateAppProjectMappingRequest := buildUpdateAppProjectMappingRequest(d)
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	resp, httpResp, err := c.ProjectMappingsApi.AppProjectMappingServiceUpdateV2(ctx, *updateAppProjectMappingRequest, agentIdentifier, identifier)

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

func resourceGitopsAppProjectMappingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := d.Get("agent_id").(string)
	identifier := d.Get("identifier").(string)
	_, httpResp, err := c.ProjectMappingsApi.AppProjectMappingServiceDeleteV2(ctx, agentIdentifier, identifier, &nextgen.ProjectMappingsApiAppProjectMappingServiceDeleteV2Opts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildCreateAppProjectMappingRequest(d *schema.ResourceData) *nextgen.V1AppProjectMappingCreateRequestV2 {
	var appProjectMappingRequest nextgen.V1AppProjectMappingCreateRequestV2

	if attr, ok := d.GetOk("account_id"); ok {
		appProjectMappingRequest.AccountIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		appProjectMappingRequest.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		appProjectMappingRequest.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("agent_id"); ok {
		appProjectMappingRequest.AgentIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("argo_project_name"); ok {
		appProjectMappingRequest.ArgoProjectName = attr.(string)
	}

	return &appProjectMappingRequest
}

func buildUpdateAppProjectMappingRequest(d *schema.ResourceData) *nextgen.V1AppProjectMappingQueryV2 {
	var appProjectMappingRequest nextgen.V1AppProjectMappingQueryV2

	if attr, ok := d.GetOk("account_id"); ok {
		appProjectMappingRequest.AccountIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		appProjectMappingRequest.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		appProjectMappingRequest.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("agent_id"); ok {
		appProjectMappingRequest.AgentIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("argo_project_name"); ok {
		appProjectMappingRequest.ArgoProjectName = attr.(string)
	}

	return &appProjectMappingRequest
}

func readAppProjectMapping(d *schema.ResourceData, mapping *nextgen.V1AppProjectMappingV2) {
	d.SetId(mapping.Identifier)
	d.Set("identifier", mapping.Identifier)
	d.Set("account_id", mapping.AccountIdentifier)
	d.Set("org_id", mapping.OrgIdentifier)
	d.Set("project_id", mapping.ProjectIdentifier)
	d.Set("agent_id", mapping.AgentIdentifier)
	d.Set("argo_project_name", mapping.ArgoProjectName)
}
