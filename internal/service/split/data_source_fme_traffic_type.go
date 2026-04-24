package split

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceFMETrafficType looks up a Split traffic type by name for a Harness org and project.
func DataSourceFMETrafficType() *schema.Resource {
	return &schema.Resource{
		Description: "Look up a Harness FME (Split) traffic type by name. The workspace is resolved from `org_id` and `project_id` via Workspaces.ResolveWorkspaceID.",

		ReadContext: dataSourceFMETrafficTypeRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Harness organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Harness project identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Traffic type name in Split (e.g. `user`).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"traffic_type_id": {
				Description: "The Split traffic type ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"display_attribute_id": {
				Description: "Display attribute ID when set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMETrafficTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags != nil && diags.HasError() {
		return diags
	}

	tt, err := TrafficTypeByOrganizationProjectAndName(ctx, sessionFromMeta(meta), client, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("traffic_type_id", tt.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_attribute_id", tt.DisplayAttributeID); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(tt.ID)
	return nil
}
