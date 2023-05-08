package apiKey

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceApiKey() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating ApiKeys.",

		ReadContext:   resourceApiKeyRead,
		CreateContext: resourceApiKeyCreateOrUpdate,
		UpdateContext: resourceApiKeyCreateOrUpdate,
		DeleteContext: resourceApiKeyDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the ApiKey",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags for the API Key",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"apikey_type": {
				Description:  "Type of the API Key",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"USER", "SERVICE_ACCOUNT"}, false),
			},
			"parent_id": {
				Description: "Parent Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
			},
			"default_time_to_expire_token": {
				Description: "Expiry time of the apiKey",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"account_id": {
				Description: "Account Identifier for the Entity",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}

	return resource
}

func resourceApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	type_ := d.Get("apikey_type").(string)
	parentId := d.Get("parent_id").(string)

	resp, httpResp, err := c.ApiKeyApi.GetAggregatedApiKey(ctx, c.AccountId, type_, parentId, id, &nextgen.ApiKeyApiGetAggregatedApiKeyOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readApiKey(d, resp.Data.ApiKey)

	return nil
}

func resourceApiKeyCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoApiKey
	var httpResp *http.Response

	id := d.Id()
	apiKey := buildApiKey(d)

	if id == "" {
		resp, httpResp, err = c.ApiKeyApi.CreateApiKey(ctx, c.AccountId, &nextgen.ApiKeyApiCreateApiKeyOpts{Body: optional.NewInterface(apiKey)})
	} else {
		resp, httpResp, err = c.ApiKeyApi.UpdateApiKey(ctx, c.AccountId, &nextgen.ApiKeyApiUpdateApiKeyOpts{Body: optional.NewInterface(apiKey)})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readApiKey(d, resp.Data)

	return nil
}

func resourceApiKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	type_ := d.Get("apikey_type").(string)
	parentId := d.Get("parent_id").(string)
	_, httpResp, err := c.ApiKeyApi.DeleteApiKey(ctx, c.AccountId, type_, parentId, d.Id(), &nextgen.ApiKeyApiDeleteApiKeyOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildApiKey(d *schema.ResourceData) *nextgen.ApiKey {
	apiKey := &nextgen.ApiKey{}

	if attr, ok := d.GetOk("identifier"); ok {
		apiKey.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		apiKey.Name = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		apiKey.Description = attr.(string)
	}

	if attr, ok := d.GetOk("tags"); ok {
		apiKey.Tags = attr.(map[string]string)
	}

	if attr, ok := d.GetOk("apikey_type"); ok {
		apiKey.ApiKeyType = attr.(string)
	}

	if attr, ok := d.GetOk("parent_id"); ok {
		apiKey.ParentIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("default_time_to_expire_token"); ok {
		apiKey.DefaultTimeToExpireToken = int64(attr.(int))
	}

	if attr, ok := d.GetOk("account_id"); ok {
		apiKey.AccountIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("org_id"); ok {
		apiKey.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		apiKey.ProjectIdentifier = attr.(string)
	}
	return apiKey
}

func readApiKey(d *schema.ResourceData, apiKey *nextgen.ApiKey) {
	d.SetId(apiKey.Identifier)
	d.Set("name", apiKey.Name)
	d.Set("description", apiKey.Description)
	d.Set("tags", apiKey.Tags)
	d.Set("apikey_type", apiKey.ApiKeyType)
	d.Set("parent_id", apiKey.ParentIdentifier)
	d.Set("default_time_to_expire_token", apiKey.DefaultTimeToExpireToken)
	d.Set("account_id", apiKey.AccountIdentifier)
	d.Set("project_id", apiKey.ProjectIdentifier)
	d.Set("org_id", apiKey.OrgIdentifier)
}
