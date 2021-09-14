package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudProviderDataCenter() *schema.Resource {

	providerSchema := commonCloudProviderSchema()

	return &schema.Resource{
		Description:   "Resource for creating a physical data center cloud provider",
		CreateContext: resourceCloudProviderDataCenterCreateOrUpdate,
		ReadContext:   resourceCloudProviderDataCenterRead,
		UpdateContext: resourceCloudProviderDataCenterCreateOrUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudProviderDataCenterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := &cac.PhysicalDatacenterCloudProvider{}
	if err := c.ConfigAsCode().GetCloudProviderById(d.Id(), cp); err != nil {
		return diag.FromErr(err)
	} else if cp.IsEmpty() {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readCloudProviderDataCenter(c, d, cp)
}

func readCloudProviderDataCenter(c *api.Client, d *schema.ResourceData, cp *cac.PhysicalDatacenterCloudProvider) diag.Diagnostics {

	d.SetId(cp.Id)
	d.Set("name", cp.Name)

	scope, err := flattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}

func resourceCloudProviderDataCenterCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.PhysicalDatacenterCloudProvider
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.PhysicalDataCenterCloudProvider).(*cac.PhysicalDatacenterCloudProvider)
	} else {
		input = &cac.PhysicalDatacenterCloudProvider{}
		if err = c.ConfigAsCode().GetCloudProviderById(d.Id(), input); err != nil {
			return diag.FromErr(err)
		} else if input.IsEmpty() {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	input.Name = d.Get("name").(string)
	if err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List(), input.UsageRestrictions); err != nil {
		return diag.FromErr(err)
	}

	cp, err := c.ConfigAsCode().UpsertPhysicalDataCenterCloudProvider(input)

	if err != nil {
		return diag.FromErr(err)
	}

	return readCloudProviderDataCenter(c, d, cp)
}
