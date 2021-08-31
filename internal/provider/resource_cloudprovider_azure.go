package provider

import (
	"context"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudProviderAzure() *schema.Resource {

	providerSchema := map[string]*schema.Schema{
		"environment_type": {
			Description:  fmt.Sprintf("The type of environment. Valid options are %s", cac.AzureEnvironmentTypesSlice),
			Type:         schema.TypeString,
			Optional:     true,
			Default:      cac.AzureEnvironmentTypes.AzureGlobal.String(),
			ValidateFunc: validation.StringInSlice(cac.AzureEnvironmentTypesSlice, false),
		},
		"client_id": {
			Description: "The client id for the Azure application",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tenant_id": {
			Description: "The tenant id for the Azure application",
			Type:        schema.TypeString,
			Required:    true,
		},
		"key": {
			Description: "The Name of the Harness secret containing the key for the Azure application",
			Type:        schema.TypeString,
			Required:    true,
		},
	}

	helpers.MergeSchemas(commonCloudProviderSchema(), providerSchema)

	return &schema.Resource{
		Description:   "Resource for creating an Azure cloud provider",
		CreateContext: resourceCloudProviderAzureCreate,
		ReadContext:   resourceCloudProviderAzureRead,
		UpdateContext: resourceCloudProviderAzureUpdate,
		DeleteContext: resourceCloudProviderAzureDelete,

		Schema: providerSchema,
	}
}

func resourceCloudProviderAzureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	name := d.Get("name").(string)

	cp := &cac.AzureCloudProvider{}
	err := c.ConfigAsCode().GetCloudProviderByName(name, cp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("environment_type", cp.AzureEnvironmentType)
	d.Set("client_id", cp.ClientId)
	d.Set("tenant_id", cp.TenantId)
	d.Set("key", cp.Key.Name)

	scope, err := flattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}

func resourceCloudProviderAzureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	input := cac.NewEntity(cac.ObjectTypes.AzureCloudProvider).(*cac.AzureCloudProvider)
	input.Name = d.Get("name").(string)
	input.AzureEnvironmentType = cac.AzureEnvironmentType(d.Get("environment_type").(string))
	input.ClientId = d.Get("client_id").(string)
	input.TenantId = d.Get("tenant_id").(string)
	input.Key = &cac.SecretRef{
		Name: d.Get("key").(string),
	}

	restrictions, err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	input.UsageRestrictions = restrictions

	cp, err := c.ConfigAsCode().UpsertAzureCloudProvider(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cp.Id)

	return nil
}

func resourceCloudProviderAzureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := cac.NewEntity(cac.ObjectTypes.AzureCloudProvider).(*cac.AzureCloudProvider)
	cp.Name = d.Get("name").(string)
	cp.AzureEnvironmentType = cac.AzureEnvironmentType(d.Get("environment_type").(string))
	cp.ClientId = d.Get("client_id").(string)
	cp.TenantId = d.Get("tenant_id").(string)
	cp.Key = &cac.SecretRef{
		Name: d.Get("key").(string),
	}

	usageRestrictions, err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	cp.UsageRestrictions = usageRestrictions

	_, err = c.ConfigAsCode().UpsertAzureCloudProvider(cp)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudProviderAzureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	err := c.CloudProviders().DeleteCloudProvider(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
