package repo_webhook

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepoWebhook() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Harness Repo Webhook.",
		ReadContext:   resourceRepoWebhookRead,
		CreateContext: resourceRepoWebhookCreateOrUpdate,
		UpdateContext: resourceRepoWebhookCreateOrUpdate,
		DeleteContext: resourceRepoWebhookDelete,
		Importer:      helpers.RepoRuleResourceImporter,

		Schema: createSchema(),
	}

	helpers.SetMultiLevelDatasourceSchemaWithoutCommonFields(resource.Schema)

	return resource
}

func resourceRepoWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := d.Get("repo_identifier").(string)

	webhook, resp, err := c.WebhookApi.GetWebhook(
		ctx,
		c.AccountId,
		repoIdentifier,
		d.Id(),
		&code.WebhookApiGetWebhookOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if err != nil {
		return helpers.HandleReadApiError(err, d, resp)
	}

	readRepoWebhook(d, &webhook, orgID.Value(), projectID.Value(), repoIdentifier)
	return nil
}

func resourceRepoWebhookCreateOrUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)
	var err error
	var webhook code.OpenapiWebhookType
	var resp *http.Response

	id := d.Id()
	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := d.Get("repo_identifier").(string)

	if id != "" {
		// update webhook
		body := buildRepoWebhookBodyForUpdate(d)
		webhook, resp, err = c.WebhookApi.UpdateWebhook(
			ctx,
			c.AccountId,
			repoIdentifier,
			id,
			&code.WebhookApiUpdateWebhookOpts{
				Body:              optional.NewInterface(body),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	} else {
		// create webhook
		body := buildRepoWebhookBodyForCreate(d)
		webhook, resp, err = c.WebhookApi.CreateWebhook(
			ctx,
			c.AccountId,
			repoIdentifier,
			&code.WebhookApiCreateWebhookOpts{
				Body:              optional.NewInterface(body),
				OrgIdentifier:     orgID,
				ProjectIdentifier: projectID,
			},
		)
	}
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	readRepoWebhook(d, &webhook, orgID.Value(), projectID.Value(), repoIdentifier)
	return nil
}

func resourceRepoWebhookDelete(
	ctx context.Context,
	d *schema.ResourceData,
	meta interface{},
) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetCodeClientWithContext(ctx)

	orgID := helpers.BuildField(d, "org_id")
	projectID := helpers.BuildField(d, "project_id")
	repoIdentifier := d.Get("repo_identifier").(string)
	resp, err := c.WebhookApi.DeleteWebhook(
		ctx,
		c.AccountId,
		repoIdentifier,
		d.Id(),
		&code.WebhookApiDeleteWebhookOpts{
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		},
	)
	if err != nil {
		return helpers.HandleApiError(err, d, resp)
	}

	return nil
}

func buildRepoWebhookBodyForCreate(d *schema.ResourceData) *code.OpenapiCreateWebhookRequest {
	return &code.OpenapiCreateWebhookRequest{
		Description: d.Get("description").(string),
		Enabled:     d.Get("enabled").(bool),
		Secret:      d.Get("secret").(string),
		Identifier:  d.Get("identifier").(string),
		Insecure:    d.Get("insecure").(bool),
		Triggers:    convertToTrigger(d.Get("triggers").([]interface{})),
		Url:         d.Get("url").(string),
	}
}

func buildRepoWebhookBodyForUpdate(d *schema.ResourceData) *code.OpenapiUpdateWebhookRequest {
	return &code.OpenapiUpdateWebhookRequest{
		Description: d.Get("description").(string),
		Enabled:     d.Get("enabled").(bool),
		Secret:      d.Get("secret").(string),
		Identifier:  d.Get("identifier").(string),
		Insecure:    d.Get("insecure").(bool),
		Triggers:    convertToTrigger(d.Get("triggers").([]interface{})),
		Url:         d.Get("url").(string),
	}
}

func readRepoWebhook(d *schema.ResourceData, webhook *code.OpenapiWebhookType, orgId string, projectId string, repoIdentifier string) {
	d.SetId(webhook.Identifier)
	d.Set("org_id", orgId)
	d.Set("project_id", projectId)
	d.Set("repo_identifier", repoIdentifier)
	d.Set("identifier", webhook.Identifier)
	d.Set("created", webhook.Created)
	d.Set("created_by", webhook.CreatedBy)
	d.Set("description", webhook.Description)
	d.Set("url", webhook.Url)
	d.Set("has_secret", webhook.HasSecret)
	d.Set("triggers", webhook.Triggers)
	d.Set("enabled", webhook.Enabled)
	d.Set("insecure", webhook.Insecure)
}

func convertToTrigger(in []interface{}) []code.EnumWebhookTrigger {
	list := make([]code.EnumWebhookTrigger, len(in))

	for i, v := range in {
		list[i] = code.EnumWebhookTrigger(v.(string))
	}

	return list
}

func createSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"repo_identifier": {
			Description: "Identifier of the repository.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"identifier": {
			Description: "Identifier of the webhook.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description of the webhook.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"url": {
			Description: "URL that's called by the webhook.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"secret": {
			Description: "Webhook secret which will be used to sign the webhook payload.",
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},
		"triggers": {
			Description: "List of triggers of the webhook (keep empty for all triggers).",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"enabled": {
			Description: "Webhook enabled.",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"insecure": {
			Description: "Allow insecure connections for provided webhook URL.",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"created_by": {
			Description: "ID of the user who created the webhook.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created": {
			Description: "Timestamp when the webhook was created.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"has_secret": {
			Description: "Created webhook has secret encoding.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
	}
}
