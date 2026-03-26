package split

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceFMELargeSegment looks up a Split large segment by name for a Harness org and project.
func DataSourceFMELargeSegment() *schema.Resource {
	return &schema.Resource{
		Description: "Look up a Harness FME (Split) large segment by name. The workspace is resolved from `org_id` and `project_id` via Workspaces.ResolveWorkspaceID. " +
			"After create, Get can return 404 briefly; this data source retries for a short period.",

		ReadContext: dataSourceFMELargeSegmentRead,

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
				Description: "Large segment name in Split.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"large_segment_id": {
				Description: "The Split large segment ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Large segment description.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"traffic_type_id": {
				Description: "Split traffic type ID for this large segment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceFMELargeSegmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags != nil && diags.HasError() {
		return diags
	}

	seg, err := LargeSegmentByOrganizationProjectAndName(ctx, sessionFromMeta(meta), client,
		d.Get("org_id").(string), d.Get("project_id").(string), d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if seg == nil {
		return diag.Errorf("large segment %q not found", d.Get("name").(string))
	}

	if err := d.Set("large_segment_id", seg.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", seg.Description); err != nil {
		return diag.FromErr(err)
	}
	ttID := ""
	if seg.TrafficType != nil {
		ttID = seg.TrafficType.ID
	}
	if err := d.Set("traffic_type_id", ttID); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(seg.ID)
	return nil
}
