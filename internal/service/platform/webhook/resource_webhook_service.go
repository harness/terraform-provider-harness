package webhook

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceWebhook() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness pipeline.",

		ReadContext:   resourceWebhookRead,
		UpdateContext: resourceWebhookUpdate,
		DeleteContext: resourceWebhookDelete,
		CreateContext: resourceWebhookCreate,
		Importer:      helpers.GitWebhookResourceImporter,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Org identifier of the GitOps project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the GitOps repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "If true, returns Pipeline YAML with Templates applied on it.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"repo_name": {
				Description: "If true, returns Pipeline YAML with Templates applied on it.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connector_ref": {
				Description: "Pipeline YAML after resolving Templates (returned as a String).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"folder_paths": {
				Description: "Flag to set if importing from Git",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Description: "Contains parameters for importing a pipeline",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_enabled": {
				Description: "Contains parameters for importing a pipeline",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)
	resource.Schema["tags"].Description = resource.Schema["tags"].Description + " These should match the tag value passed in the YAML; if this parameter is null or not passed, the tags specified in YAML should also be null."
	return resource
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var repo_name, connector_ref, webhook_identifier, webhook_name, orgIdentifier, projectIdentifier string

	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("repo_name"); ok {
		repo_name = attr.(string)
	}
	if attr, ok := d.GetOk("connector_ref"); ok {
		connector_ref = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		webhook_identifier = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		webhook_name = attr.(string)
	}

	var folder_paths []string
	if sr, ok := d.GetOk("folder_paths"); ok {

		if path, ok := sr.([]interface{}); ok {
			for _, repo := range path {
				folder_paths = append(folder_paths, repo.(string))
			}
		}
	}

	// Prepare JSON payload
	payload := map[string]interface{}{
		"repo_name":          repo_name,
		"connector_ref":      connector_ref,
		"webhook_identifier": webhook_identifier,
		"webhook_name":       webhook_name,
		"folder_paths":       folder_paths,
	}

	if len(orgIdentifier) > 0 && len(projectIdentifier) > 0 {
		_, httpResp, err := c.ProjectGitxWebhooksApiService.CreateProjectGitxWebhook(ctx, orgIdentifier, projectIdentifier, &nextgen.ProjectGitxWebhooksApiCreateProjectGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

	} else if len(orgIdentifier) > 0 {
		_, httpResp, err := c.OrgGitxWebhooksApiService.CreateOrgGitxWebhook(ctx, orgIdentifier, &nextgen.OrgGitxWebhooksApiCreateOrgGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	} else {
		_, httpResp, err := c.GitXWebhooksApiService.CreateGitxWebhook(ctx, &nextgen.GitXWebhooksApiCreateGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	}

	setWebhookDetails(d, c.AccountId, orgIdentifier, projectIdentifier, webhook_identifier, webhook_name, repo_name, connector_ref)

	return nil
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var webhook_identifier, orgIdentifier, projectIdentifier string

	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		webhook_identifier = attr.(string)
	}

	if len(orgIdentifier) > 0 && len(projectIdentifier) > 0 {
		resp, httpResp, err := c.ProjectGitxWebhooksApiService.GetProjectGitxWebhook(ctx, orgIdentifier, projectIdentifier, webhook_identifier, &nextgen.ProjectGitxWebhooksApiGetProjectGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		if len(resp.WebhookIdentifier) < 0 {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		setWebhookUpdateDetails(d, c.AccountId, orgIdentifier, projectIdentifier, &resp)

	} else if len(orgIdentifier) > 0 {
		resp, httpResp, err := c.OrgGitxWebhooksApiService.GetOrgGitxWebhook(ctx, orgIdentifier, webhook_identifier, &nextgen.OrgGitxWebhooksApiGetOrgGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		if len(resp.WebhookIdentifier) > 0 {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		setWebhookUpdateDetails(d, c.AccountId, orgIdentifier, projectIdentifier, &resp)
	} else {
		resp, httpResp, err := c.GitXWebhooksApiService.GetGitxWebhook(ctx, webhook_identifier, &nextgen.GitXWebhooksApiGetGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		if len(resp.WebhookIdentifier) < 0 {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		setWebhookUpdateDetails(d, c.AccountId, orgIdentifier, projectIdentifier, &resp)
	}

	return nil
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var repo_name, connector_ref, webhook_identifier, webhook_name, orgIdentifier, projectIdentifier, accountIdentifier string

	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("account_id"); ok {
		accountIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("repo_name"); ok {
		repo_name = attr.(string)
	}
	if attr, ok := d.GetOk("connector_ref"); ok {
		connector_ref = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		webhook_identifier = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		webhook_name = attr.(string)
	}

	var folder_paths []string
	if sr, ok := d.GetOk("folder_paths"); ok {

		if path, ok := sr.([]interface{}); ok {
			for _, repo := range path {
				folder_paths = append(folder_paths, repo.(string))
			}
		}
	}

	// Prepare JSON payload
	payload := map[string]interface{}{
		"repo_name":          repo_name,
		"connector_ref":      connector_ref,
		"webhook_identifier": webhook_identifier,
		"webhook_name":       webhook_name,
		"folder_paths":       folder_paths,
	}

	if len(orgIdentifier) > 0 && len(projectIdentifier) > 0 {
		_, httpResp, err := c.ProjectGitxWebhooksApiService.UpdateProjectGitxWebhook(ctx, orgIdentifier, projectIdentifier, webhook_identifier, &nextgen.ProjectGitxWebhooksApiUpdateProjectGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

	} else if len(orgIdentifier) > 0 {
		_, httpResp, err := c.OrgGitxWebhooksApiService.UpdateOrgGitxWebhook(ctx, orgIdentifier, webhook_identifier, &nextgen.OrgGitxWebhooksApiUpdateOrgGitxWebhookOpts{
			HarnessAccount: optional.NewString(accountIdentifier),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	} else {
		_, httpResp, err := c.GitXWebhooksApiService.UpdateGitxWebhook(ctx, webhook_identifier, &nextgen.GitXWebhooksApiUpdateGitxWebhookOpts{
			HarnessAccount: optional.NewString(accountIdentifier),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	}

	setWebhookDetails(d, accountIdentifier, orgIdentifier, projectIdentifier, webhook_identifier, webhook_name, repo_name, connector_ref)

	return nil
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	var webhook_identifier, orgIdentifier, projectIdentifier string

	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("identifier"); ok {
		webhook_identifier = attr.(string)
	}

	if len(orgIdentifier) > 0 && len(projectIdentifier) > 0 {
		httpResp, err := c.ProjectGitxWebhooksApiService.DeleteProjectGitxWebhook(ctx, orgIdentifier, projectIdentifier, webhook_identifier, &nextgen.ProjectGitxWebhooksApiDeleteProjectGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

	} else if len(orgIdentifier) > 0 {
		httpResp, err := c.OrgGitxWebhooksApiService.DeleteOrgGitxWebhook(ctx, orgIdentifier, webhook_identifier, &nextgen.OrgGitxWebhooksApiDeleteOrgGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	} else {
		httpResp, err := c.GitXWebhooksApiService.DeleteGitxWebhook(ctx, webhook_identifier, &nextgen.GitXWebhooksApiDeleteGitxWebhookOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	}

	return nil
}

func setWebhookDetails(d *schema.ResourceData, account_id string, orgIdentifier string, projectIdentifier string, webhook_identifier string, webhook_name string, repo_name string, connector_ref string) {
	d.SetId(webhook_identifier)
	d.Set("account_id", account_id)
	d.Set("org_id", orgIdentifier)
	d.Set("project_id", projectIdentifier)
	d.Set("identifier", webhook_identifier)
	d.Set("name", webhook_name)
	d.Set("repo_name", repo_name)
	d.Set("connector_ref", connector_ref)
}

func setWebhookUpdateDetails(d *schema.ResourceData, account_id string, orgIdentifier string, projectIdentifier string, resp *nextgen.GitXWebhookResponse) {
	d.SetId(resp.WebhookIdentifier)
	d.Set("account_id", account_id)
	d.Set("identifier", resp.WebhookIdentifier)
	d.Set("name", resp.WebhookName)
	d.Set("repo_name", resp.RepoName)
	d.Set("connector_ref", resp.ConnectorRef)
	d.Set("folder_paths", resp.FolderPaths)
	d.Set("org_id", orgIdentifier)
	d.Set("project_id", projectIdentifier)
}
