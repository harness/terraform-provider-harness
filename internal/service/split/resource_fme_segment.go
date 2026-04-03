package split

import (
	"context"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMESegment manages a classic Split segment at workspace scope.
func ResourceFMESegment() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) segment. Import id format: `org_id/project_id/segment_name`.",

		CreateContext: resourceFMESegmentCreate,
		ReadContext:   resourceFMESegmentRead,
		DeleteContext: resourceFMESegmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMESegmentImport,
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
				Description: "Segment name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Segment description. Set when the segment is created; changing this value forces replacement (destroy and recreate), not an in-place Split API update.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceFMESegmentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, segName, err := ParseImportID3(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("name", segName); err != nil {
		return nil, err
	}
	d.SetId(segName)

	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := ResolveWorkspaceIDFromOrgProject(ctx, meta, client, orgID, projectID)
	if err != nil {
		return nil, err
	}
	seg, err := client.Segments.Get(wsID, segName)
	if err != nil {
		return nil, err
	}
	if seg == nil {
		return nil, fmt.Errorf("cannot import segment %q: Split API returned no segment", segName)
	}
	ttID := ""
	if seg.TrafficType != nil {
		ttID = seg.TrafficType.ID
	}
	if ttID == "" {
		return nil, fmt.Errorf("cannot import segment %q: Split API returned no traffic type id (required attribute traffic_type_id)", segName)
	}
	if err := d.Set("traffic_type_id", ttID); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceFMESegmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	req := splitsdk.SegmentRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
	_, err := client.Segments.Create(wsID, ttID, req)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(d.Get("name").(string))
	return resourceFMESegmentRead(ctx, d, meta)
}

func resourceFMESegmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	seg, err := client.Segments.Get(wsID, d.Id())
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
	switch {
	case seg.TrafficType != nil && seg.TrafficType.ID != "":
		if err := d.Set("traffic_type_id", seg.TrafficType.ID); err != nil {
			return diag.FromErr(err)
		}
	default:
		// Split list/get may omit trafficType; keep configured state.
		if err := d.Set("traffic_type_id", d.Get("traffic_type_id")); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceFMESegmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	if err := client.Segments.Delete(wsID, d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
