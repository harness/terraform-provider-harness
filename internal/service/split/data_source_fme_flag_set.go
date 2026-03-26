package split

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceFMEFlagSet looks up a Split flag set by name for a Harness org and project.
func DataSourceFMEFlagSet() *schema.Resource {
	return &schema.Resource{
		Description: "Look up a Harness FME (Split) flag set by name. The workspace is resolved from `org_id` and `project_id` via Workspaces.ResolveWorkspaceID. " +
			"After create, the list API can lag briefly; this data source retries for a short period.",

		ReadContext: dataSourceFMEFlagSetRead,

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
				Description: "Flag set name in Split.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"flag_set_id": {
				Description: "The Split flag set ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Flag set description.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMEFlagSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags != nil && diags.HasError() {
		return diags
	}

	fs, err := FlagSetByOrganizationProjectAndName(ctx, sessionFromMeta(meta), client, d.Get("org_id").(string), d.Get("project_id").(string), d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("flag_set_id", fs.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", fs.Description); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fs.ID)
	return nil
}
