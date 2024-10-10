package provider

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProvider() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Provider.",

		ReadContext:   resourceProviderRead,
		UpdateContext: resourceProviderCreateOrUpdate,
		DeleteContext: resourceProviderDelete,
		CreateContext: resourceProviderCreateOrUpdate,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "The identifier of the provider entity.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"name": {
				Description: "The name of the provider entity.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the provider entity.",
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Computed:    false,
			},
			"type": {
				Description: "The type of the provider entity.",
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"last_modified_at": {
				Description: "The last modified time of the provider entity.",
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"spec": {
				Description: "Contains parameters related to the provider entity.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "The type of the provider entity.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"domain": {
							Description: "Host domain of the provider.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    false,
							Required:    false,
						},
						"secret_manager_ref": {
							Description: "Secret Manager Ref to store the access/refresh tokens",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    false,
							Required:    false,
						},
						"delegate_selectors": {
							Description: "Delegate selectors to fetch the access token",
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    false,
							Required:    false,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"client_id": {
							Description: "Client Id of the OAuth app to connect",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    false,
							Required:    false,
						},
						"client_secret_ref": {
							Description: "Client Secret Ref of the OAuth app to connect",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    false,
							Required:    false,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	resp, httpResp, err := c.ProviderApi.GetProvider(ctx, id, c.AccountId)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readProvider(d, resp.Data)

	return nil
}

func resourceProviderCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var httpResp *http.Response

	id := d.Id()

	if id == "" {
		var resp nextgen.ProviderCreateResponse
		provider := createProvider(d)
		providerParams := providerCreateParam(provider)
		resp, httpResp, err = c.ProviderApi.CreateProvider(ctx, c.AccountId, &providerParams)

		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

		readCreateProvider(d, resp.Data)
	} else {
		var resp nextgen.ProviderUpdateResponse
		provider := updateProvider(d)
		providerParams := providerUpdateParam(provider)
		resp, httpResp, err = c.ProviderApi.UpdateProvider(ctx, id, c.AccountId, &providerParams)

		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}

		readUpdateProvider(d, resp.Data)
	}
	return nil
}

func resourceProviderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var resp nextgen.ProviderDeleteResponse
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.ProviderApi.DeleteProvider(ctx, d.Id(), c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readDeleteProvider(d, resp.Data)
	return nil
}

func createProvider(d *schema.ResourceData) *nextgen.ProviderCreateRequest {
	specData := d.Get("spec").([]interface{})
	var spec json.RawMessage

	if len(specData) > 0 {
		specMap := specData[0].(map[string]interface{})

		delegateInterfaces, ok := specMap["delegate_selectors"].([]interface{})
		if !ok {
			log.Fatalf("delegate_selectors is not a []interface{}")
		}

		var delegateSelectors []string
		for _, v := range delegateInterfaces {
			str, ok := v.(string)
			if !ok {
				log.Fatalf("One of the delegate_selectors is not a string")
			}
			delegateSelectors = append(delegateSelectors, str)
		}

		if specMap["type"] == "BITBUCKET_SERVER" {
			bitbucketSpec := &nextgen.BitbucketServerSpec{
				Domain:            specMap["domain"].(string),
				DelegateSelectors: delegateSelectors,
				SecretManagerRef:  specMap["secret_manager_ref"].(string),
				ClientId:          specMap["client_id"].(string),
				ClientSecretRef:   specMap["client_secret_ref"].(string),
				Type_:             nextgen.ProviderTypes.BitbucketServer,
			}
			spec, _ = json.Marshal(bitbucketSpec)
		}

	}

	return &nextgen.ProviderCreateRequest{
		Identifier:  d.Get("identifier").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Spec:        spec,
	}
}

func updateProvider(d *schema.ResourceData) *nextgen.ProviderUpdateRequest {
	specData := d.Get("spec").([]interface{})
	var spec json.RawMessage

	if len(specData) > 0 {
		specMap := specData[0].(map[string]interface{})
		marshalledSpec, err := json.Marshal(specMap)
		if err != nil {
			log.Printf("Error marshaling spec: %v", err)
			return nil
		}

		spec = marshalledSpec
	}

	return &nextgen.ProviderUpdateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Spec:        spec,
	}
}

func readCreateProvider(d *schema.ResourceData, provider *nextgen.ProviderCreateApiResponse) {
	d.SetId(provider.Identifier)
	d.Set("identifier", provider.Identifier)
}

func readUpdateProvider(d *schema.ResourceData, provider *nextgen.ProviderUpdateApiResponse) {
	d.SetId(provider.Identifier)
	d.Set("identifier", provider.Identifier)
}

func readDeleteProvider(d *schema.ResourceData, provider *nextgen.ProviderDeleteApiResponse) {
	d.SetId(provider.Identifier)
	d.Set("identifier", provider.Identifier)
}

func providerCreateParam(provider *nextgen.ProviderCreateRequest) nextgen.ProviderApiCreateProviderOpts {
	return nextgen.ProviderApiCreateProviderOpts{
		Body: optional.NewInterface(provider),
	}
}

func providerUpdateParam(provider *nextgen.ProviderUpdateRequest) nextgen.ProviderApiUpdateProviderOpts {
	return nextgen.ProviderApiUpdateProviderOpts{
		Body: optional.NewInterface(provider),
	}
}

func readProvider(d *schema.ResourceData, so *nextgen.Provider) {
	d.SetId(so.Identifier)
	d.Set("identifier", so.Identifier)
	d.Set("name", so.Name)
	d.Set("description", so.Description)
	d.Set("type", so.Type_)
	d.Set("last_modified_at", so.LastModifiedAt)
	var specData map[string]interface{}
	err := json.Unmarshal(so.Spec, &specData)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}
	d.Set("spec", []interface{}{
		map[string]interface{}{
			"type":               so.Type_,
			"domain":             specData["domain"],
			"secret_manager_ref": specData["secretManagerRef"],
			"client_id":          specData["clientId"],
			"client_secret_ref":  specData["clientSecretRef"],
			"delegate_selectors": specData["delegateSelectors"],
		},
	})
}
