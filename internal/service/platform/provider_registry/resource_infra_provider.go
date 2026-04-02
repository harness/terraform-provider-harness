package provider_registry

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInfraProvider() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing Terraform/OpenTofu Providers in the IaCM Provider Registry.",
		ReadContext:   resourceInfraProviderRead,
		CreateContext: resourceInfraProviderCreate,
		UpdateContext: resourceInfraProviderUpdate,
		DeleteContext: resourceInfraProviderDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"type": {
				Description: "Provider type (e.g., aws, azurerm, google).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Description of the provider.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": {
				Description: "Unique identifier of the provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account": {
				Description: "Account that owns the provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created": {
				Description: "Timestamp when the provider was created.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"updated": {
				Description: "Timestamp when the provider was last updated.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"versions": {
				Description: "List of provider versions.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Description: "Version number.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"synced": {
							Description: "Whether the version is synced.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"files": {
							Description: "List of uploaded files for this version.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceInfraProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProvider(
		ctx,
		id,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readProvider(d, &resp)
	return nil
}

func resourceInfraProviderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		d.MarkNewResource()
	}
	createProvider := buildCreateProviderRequestBody(d)
	providerType := d.Get("type").(string)
	log.Printf("[DEBUG] Creating provider with type %s and body %v", providerType, createProvider)
	provRes, httpRes, err := c.ProviderRegistryApi.ProviderRegistryCreateProvider(ctx, createProvider, c.AccountId, providerType)

	if err != nil {
		return parseError(err, httpRes)
	}
	setProviderId(d, &provRes)
	resourceInfraProviderRead(ctx, d, m)
	return nil
}

func resourceInfraProviderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		return nil
	}
	httpRes, err := c.ProviderRegistryApi.ProviderRegistryDeleteProvider(ctx, id, c.AccountId)
	if err != nil {
		return parseError(err, httpRes)
	}
	return nil
}

func resourceInfraProviderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Update is not supported for provider registry resources. Provider type cannot be changed and description updates are not supported by the API.")
}

func setProviderId(d *schema.ResourceData, provider *nextgen.AnsibleDataInfo) {
	d.SetId(provider.Id)
}

func readProvider(d *schema.ResourceData, provider *nextgen.GetProviderResponse) {
	d.SetId(provider.Id)
	d.Set("id", provider.Id)
	d.Set("account", provider.Account)
	d.Set("created", provider.Created)
	d.Set("description", provider.Description)
	d.Set("type", provider.Type_)
	d.Set("updated", provider.Updated)

	if len(provider.Versions) > 0 {
		versions := make([]interface{}, len(provider.Versions))
		for i, v := range provider.Versions {
			versionMap := map[string]interface{}{
				"version": v.Version,
				"synced":  v.Synced,
				"files":   v.Files,
			}
			versions[i] = versionMap
		}
		d.Set("versions", versions)
	}
}

func buildCreateProviderRequestBody(d *schema.ResourceData) nextgen.CreateProviderRequestBody {
	provider := nextgen.CreateProviderRequestBody{}

	if desc, ok := d.GetOk("description"); ok {
		provider.Description = desc.(string)
	}

	return provider
}

func parseError(err error, httpResp *http.Response) diag.Diagnostics {
	if httpResp != nil && httpResp.StatusCode == 401 {
		return diag.Errorf("%s\nHint:\n1) Please check if token has expired or is wrong.\n2) Harness Provider is misconfigured. For firstgen resources please give the correct api_key and for nextgen resources please give the correct platform_api_key.", httpResp.Status)
	}
	if httpResp != nil && httpResp.StatusCode == 403 {
		return diag.Errorf("%s\nHint:\n1) Please check if the token has required permission for this operation.\n2) Please check if the token has expired or is wrong.", httpResp.Status)
	}

	se, ok := err.(nextgen.GenericSwaggerError)
	if !ok {
		diag.FromErr(err)
	}

	iacmErrBody := se.Body()
	iacmErr := nextgen.IacmError{}
	jsonErr := json.Unmarshal(iacmErrBody, &iacmErr)
	if jsonErr != nil {
		return diag.Errorf("%s", err.Error())
	}

	return diag.Errorf("%s\nHint:\n1) %s", httpResp.Status, iacmErr.Message)
}
