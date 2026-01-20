package ip_allowlist

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

func ResourceIPAllowlist() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing IP allowlist configs.",
		ReadContext:   resourceIPAllowlistRead,
		CreateContext: resourceIPAllowlistCreateOrUpdate,
		UpdateContext: resourceIPAllowlistCreateOrUpdate,
		DeleteContext: resourceIPAllowlistDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Description: "CIDR range or IP address to allow.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"allowed_source_type": {
				Description: "List of sources to allow. Valid values are `UI` and `API`.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validation.StringInSlice(nextgen.AllowedSourceTypeValues, false)},
			},
			"enabled": {
				Description: "Whether the allowlist config is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"created": {
				Description: "Creation timestamp for the config.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"updated": {
				Description: "Last update timestamp for the config.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}

	helpers.SetCommonResourceSchema(resource.Schema)

	return resource
}

func resourceIPAllowlistRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	resp, httpResp, err := c.IPAllowlistApiService.GetIpAllowlistConfig(ctx, id, &nextgen.IPAllowlistApiGetIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.IpAllowlistConfig == nil {
		d.SetId("")
		return nil
	}

	readIPAllowlistConfig(d, &resp)

	return nil
}

func resourceIPAllowlistCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	req := nextgen.IpAllowlistConfigRequest{
		IpAllowlistConfig: buildIPAllowlistConfig(d),
	}

	var err error
	var resp nextgen.IpAllowlistConfigResponse
	var httpResp *http.Response

	if id == "" {
		resp, httpResp, err = c.IPAllowlistApiService.CreateIpAllowlistConfig(ctx, req, &nextgen.IPAllowlistApiCreateIpAllowlistConfigOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	} else {
		resp, httpResp, err = c.IPAllowlistApiService.UpdateIpAllowlistConfig(ctx, id, req, &nextgen.IPAllowlistApiUpdateIpAllowlistConfigOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.IpAllowlistConfig == nil {
		return nil
	}

	readIPAllowlistConfig(d, &resp)

	return nil
}

func resourceIPAllowlistDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		return nil
	}

	httpResp, err := c.IPAllowlistApiService.DeleteIpAllowlistConfig(ctx, id, &nextgen.IPAllowlistApiDeleteIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildIPAllowlistConfig(d *schema.ResourceData) *nextgen.IpAllowlistConfig {
	config := &nextgen.IpAllowlistConfig{
		Identifier: d.Get("identifier").(string),
		Name:       d.Get("name").(string),
		IpAddress:  d.Get("ip_address").(string),
		Enabled:    d.Get("enabled").(bool),
	}

	if attr, ok := d.GetOk("description"); ok {
		config.Description = attr.(string)
	}

	if attr, ok := d.GetOk("tags"); ok {
		config.Tags = helpers.ExpandTags(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("allowed_source_type"); ok {
		config.AllowedSourceType = expandAllowedSourceTypes(attr.(*schema.Set).List())
	}

	return config
}

func readIPAllowlistConfig(d *schema.ResourceData, resp *nextgen.IpAllowlistConfigResponse) {
	config := resp.IpAllowlistConfig

	d.SetId(config.Identifier)
	d.Set("identifier", config.Identifier)
	d.Set("name", config.Name)
	d.Set("description", config.Description)
	d.Set("ip_address", config.IpAddress)
	d.Set("enabled", config.Enabled)
	d.Set("allowed_source_type", flattenAllowedSourceTypes(config.AllowedSourceType))
	d.Set("tags", helpers.FlattenTags(config.Tags))
	d.Set("created", resp.Created)
	d.Set("updated", resp.Updated)
}

func expandAllowedSourceTypes(values []interface{}) []nextgen.AllowedSourceType {
	if len(values) == 0 {
		return nil
	}

	result := make([]nextgen.AllowedSourceType, 0, len(values))
	for _, value := range values {
		result = append(result, nextgen.AllowedSourceType(value.(string)))
	}

	return result
}

func flattenAllowedSourceTypes(values []nextgen.AllowedSourceType) []string {
	if len(values) == 0 {
		return nil
	}

	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.String())
	}

	return result
}
