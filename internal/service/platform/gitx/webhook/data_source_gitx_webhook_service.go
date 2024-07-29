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

func DataSourceWebhook() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness pipeline.",
		ReadContext: dataSourceWebhookRead,
		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Identifier field of the organization the resource is scoped to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Identifier field of the project the resource is scoped to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identifier": {
				Description: "GitX webhook identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "GitX webhook name",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)
	resource.Schema["tags"].Description = resource.Schema["tags"].Description + " These should match the tag value passed in the YAML; if this parameter is null or not passed, the tags specified in YAML should also be null."
	return resource
}

func dataSourceWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		if len(resp.WebhookIdentifier) <= 0 {
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
		if len(resp.WebhookIdentifier) <= 0 {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		setWebhookUpdateDetails(d, c.AccountId, orgIdentifier, projectIdentifier, &resp)
	}

	return nil
}
