package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudProviderSpot() *schema.Resource {

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

	helpers.MergeSchemas(commonCloudProviderSchema(), providerSchema)

	return &schema.Resource{
		Description:   "Resource for creating a GCP cloud provider",
		CreateContext: resourceCloudProviderSpotCreate,
		ReadContext:   resourceCloudProviderSpotRead,
		UpdateContext: resourceCloudProviderSpotUpdate,
		DeleteContext: resourceCloudProviderSpotDelete,

		Schema: providerSchema,
	}
}

func resourceCloudProviderSpotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	name := d.Get("name").(string)

	cp := &cac.SpotInstCloudProvider{}
	err := c.ConfigAsCode().GetCloudProviderByName(name, cp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("account_id", cp.AccountId)

	if cp.Token != nil {
		d.Set("token_secret_name", cp.Token.Name)
	}

	scope, err := flattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}

func resourceCloudProviderSpotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	input := cac.NewEntity(cac.ObjectTypes.SpotInstCloudProvider).(*cac.SpotInstCloudProvider)
	input.Name = d.Get("name").(string)
	input.AccountId = d.Get("account_id").(string)

	if token := d.Get("token_secret_name").(string); token != "" {
		input.Token = &cac.SecretRef{
			Name: token,
		}
	}

	restrictions, err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.UsageRestrictions = restrictions

	cp, err := c.ConfigAsCode().UpsertSpotInstCloudProvider(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cp.Id)

	return nil
}

func resourceCloudProviderSpotUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := cac.NewEntity(cac.ObjectTypes.SpotInstCloudProvider).(*cac.SpotInstCloudProvider)
	cp.Name = d.Get("name").(string)
	cp.AccountId = d.Get("account_id").(string)

	if token := d.Get("token_secret_name").(string); token != "" {
		cp.Token = &cac.SecretRef{
			Name: token,
		}
	}

	usageRestrictions, err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	cp.UsageRestrictions = usageRestrictions

	_, err = c.ConfigAsCode().UpsertSpotInstCloudProvider(cp)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudProviderSpotDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	err := c.CloudProviders().DeleteCloudProvider(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
