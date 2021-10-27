package cd

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/terraform-provider-harness/helpers"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCloudProviderTanzu() *schema.Resource {

	providerSchema := map[string]*schema.Schema{
		"endpoint": {
			Description: "The url of the Tanzu platform.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_validation": {
			Description: "Skip validation of Tanzu configuration.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"username": {
			Description:   "The username to use to authenticate to Tanzu.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"username_secret_name"},
		},
		"username_secret_name": {
			Description:   "The name of the Harness secret containing the username to authenticate to Tanzu with.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"username"},
		},
		"password_secret_name": {
			Description: "The name of the Harness secret containing the password to use to authenticate to Tanzu.",
			Type:        schema.TypeString,
			Required:    true,
		},
	}

	// usage_scope is not supported because the scope will always be inherited from `token_secret_name`
	commonSchema := commonCloudProviderSchema()
	delete(commonSchema, "usage_scope")

	helpers.MergeSchemas(commonSchema, providerSchema)

	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating a Tanzu cloud provider."),
		CreateContext: resourceCloudProviderTanzuCreateOrUpdate,
		ReadContext:   resourceCloudProviderTanzuRead,
		UpdateContext: resourceCloudProviderTanzuCreateOrUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudProviderTanzuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := &cac.PcfCloudProvider{}
	if err := c.ConfigAsCode().GetCloudProviderById(d.Id(), cp); err != nil {
		return diag.FromErr(err)
	} else if cp.IsEmpty() {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readCloudProviderTanzu(c, d, cp)

}

func readCloudProviderTanzu(c *api.Client, d *schema.ResourceData, cp *cac.PcfCloudProvider) diag.Diagnostics {
	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("endpoint", cp.EndpointUrl)
	d.Set("skip_validation", cp.SkipValidation)
	d.Set("username", cp.Username)

	if cp.UsernameSecretId != nil {
		d.Set("username_secret_name", cp.UsernameSecretId.Name)
	}

	d.Set("password_secret_name", cp.Password.Name)

	return nil
}

func resourceCloudProviderTanzuCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.PcfCloudProvider
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.PcfCloudProvider).(*cac.PcfCloudProvider)
	} else {
		input = &cac.PcfCloudProvider{}
		if err = c.ConfigAsCode().GetCloudProviderById(d.Id(), input); err != nil {
			return diag.FromErr(err)
		} else if input.IsEmpty() {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	input.Name = d.Get("name").(string)
	input.EndpointUrl = d.Get("endpoint").(string)
	input.SkipValidation = d.Get("skip_validation").(bool)
	input.Username = d.Get("username").(string)

	if attr := d.Get("username_secret_name").(string); attr != "" {
		input.UsernameSecretId = &cac.SecretRef{
			Name: attr,
		}
	}

	input.Password = &cac.SecretRef{
		Name: d.Get("password_secret_name").(string),
	}

	cp, err := c.ConfigAsCode().UpsertPcfCloudProvider(input)

	if err != nil {
		return diag.FromErr(err)
	}

	return readCloudProviderTanzu(c, d, cp)
}
