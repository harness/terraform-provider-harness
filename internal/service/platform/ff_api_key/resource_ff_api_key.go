package ff_api_key

import (
	"context"
	"github.com/harness/terraform-provider-harness/internal/service/platform/feature_flag"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFFApiKey() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating an environment SDK key for Feature Flags.",

		ReadContext:   resourceFFApiKeyRead,
		DeleteContext: resourceFFApiKeyDelete,
		CreateContext: resourceFFApiKeyCreate,
		Importer:      nil,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the SDK API Key",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the SDK API Key",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Description of the SDK API Key",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"api_key": {
				Description: "The value of the SDK API Key",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"expired_at": {
				Description: "Expiration datetime of the SDK API Key",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"env_id": {
				Description: "Environment Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"org_id": {
				Description: "Organization Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Project Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"type": {
				Description: "Type of SDK. Valid values are `Server` or `Client`.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}

	return resource
}

type ApiKeyQueryParameters struct {
	Identifier     string
	ProjectId      string
	EnvironmentId  string
	OrganizationId string
}

type ApiKeyOpts struct {
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type_       string `json:"type"`
	ExpiredAt   int    `json:"expiredAt,omitempty"`
}

func resourceFFApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	qp := buildFFApiKeyQueryParameters(d)

	resp, httpResp, err := c.APIKeysApi.GetAPIKey(ctx, id, qp.ProjectId, qp.EnvironmentId, c.AccountId, qp.OrganizationId)

	if err != nil {
		return feature_flag.HandleCFApiError(err, d, httpResp)
	}

	readFFApiKey(d, &resp, qp)

	return nil
}

func resourceFFApiKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	qp := buildFFApiKeyQueryParameters(d)
	opts := buildFFApiKeyOpts(d)

	var err error
	var resp nextgen.CfApiKey
	var httpResp *http.Response

	resp, httpResp, err = c.APIKeysApi.AddAPIKey(ctx, c.AccountId, qp.OrganizationId, qp.EnvironmentId, qp.ProjectId, opts)

	if err != nil {
		// handle conflict
		if httpResp != nil && httpResp.StatusCode == 409 {
			return diag.Errorf("An api key with identifier [%s] orgIdentifier [%s] project [%s]  and environment [%s] already exists", d.Get("identifier").(string), qp.OrganizationId, qp.ProjectId, qp.EnvironmentId)
		}
		return feature_flag.HandleCFApiError(err, d, httpResp)
	}

	readFFApiKey(d, &resp, qp)

	return nil
}

func resourceFFApiKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}
	qp := buildFFApiKeyQueryParameters(d)

	httpResp, err := c.APIKeysApi.DeleteAPIKey(ctx, d.Id(), qp.ProjectId, qp.EnvironmentId, c.AccountId, qp.OrganizationId)
	if err != nil {
		return feature_flag.HandleCFApiError(err, d, httpResp)
	}

	return nil
}

func readFFApiKey(d *schema.ResourceData, apiKey *nextgen.CfApiKey, qp *ApiKeyQueryParameters) {
	d.SetId(apiKey.Identifier)
	d.Set("identifier", apiKey.Identifier)
	d.Set("name", apiKey.Name)
	if d.IsNewResource() {
		d.Set("api_key", apiKey.ApiKey)
	}
	d.Set("type", apiKey.Type_)

	d.Set("project_id", qp.ProjectId)
	d.Set("org_id", qp.OrganizationId)
	d.Set("env_id", qp.EnvironmentId)

}

func buildFFApiKeyQueryParameters(d *schema.ResourceData) *ApiKeyQueryParameters {
	return &ApiKeyQueryParameters{
		Identifier:     d.Get("identifier").(string),
		ProjectId:      d.Get("project_id").(string),
		EnvironmentId:  d.Get("env_id").(string),
		OrganizationId: d.Get("org_id").(string),
	}
}

func buildFFApiKeyOpts(d *schema.ResourceData) *nextgen.APIKeysApiAddAPIKeyOpts {
	opts := &ApiKeyOpts{
		Identifier:  d.Get("identifier").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type_:       d.Get("type").(string),
		ExpiredAt:   d.Get("expired_at").(int),
	}

	return &nextgen.APIKeysApiAddAPIKeyOpts{
		Body: optional.NewInterface(opts),
	}

}
