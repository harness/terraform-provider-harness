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
		CreateContext: resourcePipelineCreateOrUpdate,
		Importer:      helpers.ProjectResourceImporter,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account identifier of the GitOps project.",
				Type:        schema.TypeString,
				Required:    true,
			},
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
			"webhook_identifier": {
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
			"webhook_name": {
				Description: "Contains parameters for importing a pipeline",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)
	resource.Schema["tags"].Description = resource.Schema["tags"].Description + " These should match the tag value passed in the YAML; if this parameter is null or not passed, the tags specified in YAML should also be null."
	return resource
}

func resourcePipelineCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	if attr, ok := d.GetOk("webhook_identifier"); ok {
		webhook_identifier = attr.(string)
	}
	if attr, ok := d.GetOk("webhook_name"); ok {
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
			HarnessAccount: optional.NewString(accountIdentifier),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

	} else if len(orgIdentifier) > 0 {
		_, httpResp, err := c.OrgGitxWebhooksApiService.CreateOrgGitxWebhook(ctx, orgIdentifier, &nextgen.OrgGitxWebhooksApiCreateOrgGitxWebhookOpts{
			HarnessAccount: optional.NewString(accountIdentifier),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	} else {
		_, httpResp, err := c.GitXWebhooksApiService.CreateGitxWebhook(ctx, &nextgen.GitXWebhooksApiCreateGitxWebhookOpts{
			HarnessAccount: optional.NewString(accountIdentifier),
			Body:           optional.NewInterface(payload),
		})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	// if resp == nil {
	// 	d.SetId("")
	// 	d.MarkNewResource()
	// 	return nil
	// }
	// setProjectDetails(d, &resp)

	return nil
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Implement resource read logic if needed
	return nil
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Implement resource update logic if needed
	return nil
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Implement resource delete logic if needed
	return nil
}
