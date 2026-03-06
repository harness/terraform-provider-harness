package default_images

import (
	"context"
	"fmt"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceDefaultImages() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving Harness default execution images for CI, IACM, or IDP.",
		ReadContext: dataSourceDefaultImagesRead,
		Schema: map[string]*schema.Schema{
			"kind": {
				Description:  "The service kind. Supported values: `ci`, `iacm`, `idp`.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"ci", "iacm", "idp"}, false),
			},
			"infra_type": {
				Description: "The infrastructure type passed to the execution config API (e.g. `K8`, `VM`). Defaults to `K8`.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "K8",
			},
			"type": {
				Description:  "The configuration type to retrieve. Use `default` for Harness default images or `customer` for customer-configured overrides. Defaults to `default`.",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "default",
				ValidateFunc: validation.StringInSlice([]string{"default", "customer"}, false),
			},
			"images": {
				Description: "Map of image field names to image tag values.",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceDefaultImagesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	kind := d.Get("kind").(string)
	infraType := d.Get("infra_type").(string)
	cfgType := d.Get("type").(string)

	s := meta.(*internal.Session)
	var data map[string]string

	switch kind {
	case "ci":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		if cfgType == "customer" {
			resp, err := c.CiExecutionConfigApi.GetCustomerConfig(authCtx, infraType, false)
			if err != nil {
				return diag.FromErr(err)
			}
			data = map[string]string(resp.Data)
		} else {
			resp, err := c.CiExecutionConfigApi.GetDefaultConfig(authCtx, infraType)
			if err != nil {
				return diag.FromErr(err)
			}
			data = map[string]string(resp.Data)
		}
	case "iacm":
		c, authCtx := s.GetPlatformClientWithContext(ctx)
		if cfgType == "customer" {
			resp, err := c.IacmExecutionConfigApi.GetCustomerConfig(authCtx, infraType, false)
			if err != nil {
				return diag.FromErr(err)
			}
			data = map[string]string(resp.Data)
		} else {
			resp, err := c.IacmExecutionConfigApi.GetDefaultConfig(authCtx, infraType)
			if err != nil {
				return diag.FromErr(err)
			}
			data = map[string]string(resp.Data)
		}
	case "idp":
		c, authCtx := s.GetIDPClientWithContext(ctx)
		if cfgType == "customer" {
			resp, err := c.ExecutionConfigApi.GetCustomerConfig(authCtx, infraType, false)
			if err != nil {
				return diag.FromErr(err)
			}
			data = map[string]string(resp.Data)
		} else {
			resp, err := c.ExecutionConfigApi.GetDefaultConfig(authCtx, infraType)
			if err != nil {
				return diag.FromErr(err)
			}
			data = map[string]string(resp.Data)
		}
	default:
		return diag.FromErr(fmt.Errorf("unsupported kind %q", kind))
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", kind, infraType, cfgType))

	images := make(map[string]interface{}, len(data))
	for k, v := range data {
		images[k] = v
	}
	if err := d.Set("images", images); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
