package cloudprovider

import (
	"context"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/harness-io/terraform-provider-harness/helpers"
	"github.com/harness-io/terraform-provider-harness/internal/service/cd/usagescope"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCloudProviderGcp() *schema.Resource {

	providerSchema := map[string]*schema.Schema{
		"skip_validation": {
			Description:   "Skip validation of GCP configuration.",
			Type:          schema.TypeBool,
			Optional:      true,
			Default:       false,
			ConflictsWith: []string{"secret_file_id"},
		},
		"delegate_selectors": {
			Description: "Delegate selectors to use for this provider.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			ConflictsWith: []string{"secret_file_id"},
		},
		"secret_file_id": {
			Description:   "The id of the secret containing the GCP credentials",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"delegate_selectors", "usage_scope"},
		},
	}

	helpers.MergeSchemas(commonCloudProviderSchema(), providerSchema)

	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating a GCP cloud provider."),
		CreateContext: resourceCloudProviderGcpCreateOrUpdate,
		ReadContext:   resourceCloudProviderGcpRead,
		UpdateContext: resourceCloudProviderGcpCreateOrUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudProviderGcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	cp := &cac.GcpCloudProvider{}
	if err := c.CDClient.ConfigAsCodeClient.GetCloudProviderById(d.Id(), cp); err != nil {
		return diag.FromErr(err)
	} else if cp.IsEmpty() {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readCloudProviderGcp(c, d, cp)
}

func readCloudProviderGcp(c *sdk.Session, d *schema.ResourceData, cp *cac.GcpCloudProvider) diag.Diagnostics {
	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("skip_validation", cp.SkipValidation)
	d.Set("delegate_selectors", cp.DelegateSelectors)

	if cp.ServiceAccountKeyFileContent != nil {
		d.Set("secret_file_id", cp.ServiceAccountKeyFileContent.Name)
	}

	scope, err := usagescope.FlattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}

func resourceCloudProviderGcpCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	var input *cac.GcpCloudProvider
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.GcpCloudProvider).(*cac.GcpCloudProvider)
	} else {
		input = &cac.GcpCloudProvider{}
		if err = c.CDClient.ConfigAsCodeClient.GetCloudProviderById(d.Id(), input); err != nil {
			return diag.FromErr(err)
		} else if input.IsEmpty() {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	input.Name = d.Get("name").(string)
	input.SkipValidation = d.Get("skip_validation").(bool)

	if selectors := d.Get("delegate_selectors").([]interface{}); len(selectors) > 0 {
		input.UseDelegateSelectors = true
		input.DelegateSelectors = utils.ExpandDelegateSelectors(selectors)
	}

	if secretId := d.Get("secret_file_id").(string); secretId != "" {
		input.ServiceAccountKeyFileContent = &cac.SecretRef{
			Name: secretId,
		}
	}

	if input.UsageRestrictions == nil {
		input.UsageRestrictions = &cac.UsageRestrictions{}
	}

	if err := usagescope.ExpandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List(), input.UsageRestrictions); err != nil {
		return diag.FromErr(err)
	}

	cp, err := c.CDClient.ConfigAsCodeClient.UpsertGcpCloudProvider(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readCloudProviderGcp(c, d, cp)
}
