package split

import (
	"context"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMELargeSegment manages a Split large segment.
func ResourceFMELargeSegment() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) large segment at workspace scope. Use `harness_fme_large_segment_environment_association` for each environment where the segment should have a definition. Import id format: `org_id/project_id/segment_name`.",

		CreateContext: resourceFMELargeSegmentCreate,
		ReadContext:   resourceFMELargeSegmentRead,
		DeleteContext: resourceFMELargeSegmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMELargeSegmentImport,
		},

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "Harness organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"project_id": {
				Description: "Harness project identifier.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"traffic_type_id": {
				Description: "Split traffic type ID.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Large segment name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Large segment description. Set when the segment is created; changing this value forces replacement (destroy and recreate), not an in-place Split API update.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"large_segment_id": {
				Description: "Split large segment id when returned by the API.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMELargeSegmentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, name, err := ParseImportID3(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("name", name); err != nil {
		return nil, err
	}
	d.SetId(name)

	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := ResolveWorkspaceIDFromOrgProject(ctx, meta, client, orgID, projectID)
	if err != nil {
		return nil, err
	}
	seg, err := client.LargeSegments.Get(wsID, name)
	if err != nil {
		return nil, err
	}
	if seg == nil {
		return nil, fmt.Errorf("cannot import large segment %q: Split API returned no segment", name)
	}
	ttID := ""
	if seg.TrafficType != nil {
		ttID = seg.TrafficType.ID
	}
	if ttID == "" {
		return nil, fmt.Errorf("cannot import large segment %q: Split API returned no traffic type id (required attribute traffic_type_id)", name)
	}
	if err := d.Set("traffic_type_id", ttID); err != nil {
		return nil, err
	}
	if err := d.Set("description", seg.Description); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceFMELargeSegmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	req := splitsdk.LargeSegmentCreateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
	seg, err := client.LargeSegments.Create(wsID, ttID, req)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(seg.Name)
	if err := d.Set("large_segment_id", seg.ID); err != nil {
		return diag.FromErr(err)
	}
	return resourceFMELargeSegmentRead(ctx, d, meta)
}

func resourceFMELargeSegmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	seg, err := client.LargeSegments.Get(wsID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if seg == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("name", seg.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", seg.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("large_segment_id", seg.ID); err != nil {
		return diag.FromErr(err)
	}
	switch {
	case seg.TrafficType != nil && seg.TrafficType.ID != "":
		if err := d.Set("traffic_type_id", seg.TrafficType.ID); err != nil {
			return diag.FromErr(err)
		}
	default:
		if err := d.Set("traffic_type_id", d.Get("traffic_type_id")); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceFMELargeSegmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	name := d.Id()
	if err := client.LargeSegments.Delete(wsID, name); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
