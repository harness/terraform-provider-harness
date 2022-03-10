package cloudprovider

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCloudProviderSpot() *schema.Resource {

	providerSchema := map[string]*schema.Schema{
		"account_id": {
			Description: "The Spot account ID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"token_secret_name": {
			Description: "The name of the Harness secret containing the spot account token",
			Type:        schema.TypeString,
			Required:    true,
		},
	}

	// usage_scope is not supported because the scope will always be inherited from `token_secret_name`
	commonSchema := commonCloudProviderSchema()
	delete(commonSchema, "usage_scope")

	helpers.MergeSchemas(commonSchema, providerSchema)

	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating a Spot cloud provider."),
		CreateContext: resourceCloudProviderSpotCreateOrUpdate,
		ReadContext:   resourceCloudProviderSpotRead,
		UpdateContext: resourceCloudProviderSpotCreateOrUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudProviderSpotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	cp := &cac.SpotInstCloudProvider{}
	if err := c.ConfigAsCodeClient.GetCloudProviderById(d.Id(), cp); err != nil {
		return diag.FromErr(err)
	} else if cp.IsEmpty() {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readCloudProviderSpot(c, d, cp)
}

func readCloudProviderSpot(c *cd.ApiClient, d *schema.ResourceData, cp *cac.SpotInstCloudProvider) diag.Diagnostics {
	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("account_id", cp.AccountId)

	if cp.Token != nil {
		d.Set("token_secret_name", cp.Token.Name)
	}

	return nil
}

func resourceCloudProviderSpotCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cd.ApiClient)

	var input *cac.SpotInstCloudProvider
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.SpotInstCloudProvider).(*cac.SpotInstCloudProvider)
	} else {
		input = &cac.SpotInstCloudProvider{}
		if err = c.ConfigAsCodeClient.GetCloudProviderById(d.Id(), input); err != nil {
			return diag.FromErr(err)
		} else if input.IsEmpty() {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	input.Name = d.Get("name").(string)
	input.AccountId = d.Get("account_id").(string)

	if token := d.Get("token_secret_name").(string); token != "" {
		input.Token = &cac.SecretRef{
			Name: token,
		}
	}

	cp, err := c.ConfigAsCodeClient.UpsertSpotInstCloudProvider(input)

	if err != nil {
		return diag.FromErr(err)
	}

	return readCloudProviderSpot(c, d, cp)
}
