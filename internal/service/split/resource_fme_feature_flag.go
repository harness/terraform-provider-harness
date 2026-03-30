package split

import (
	"context"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMEFeatureFlag manages a Split feature flag (split) and optional tag associations.
func ResourceFMEFeatureFlag() *schema.Resource {
	return &schema.Resource{
		Description: "Create, update, and delete a Harness FME (Split) feature flag. Import id format: `org_id/project_id/flag_name`.",

		CreateContext: resourceFMEFeatureFlagCreate,
		ReadContext:   resourceFMEFeatureFlagRead,
		UpdateContext: resourceFMEFeatureFlagUpdate,
		DeleteContext: resourceFMEFeatureFlagDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMEFeatureFlagImport,
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
				Description: "Feature flag name (Split split name).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tag names to associate with the flag (applied after create/update).",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"flag_id": {
				Description: "Split internal flag ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMEFeatureFlagImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
	sp, err := client.Splits.Get(wsID, name)
	if err != nil {
		return nil, err
	}
	if sp == nil {
		return nil, fmt.Errorf("cannot import feature flag %q: Split API returned no split", name)
	}
	ttID := ""
	if sp.TrafficType != nil {
		ttID = sp.TrafficType.ID
	}
	if ttID == "" {
		return nil, fmt.Errorf("cannot import feature flag %q: Split API returned no traffic type id (required attribute traffic_type_id)", name)
	}
	if err := d.Set("traffic_type_id", ttID); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func tagListFromSet(d *schema.ResourceData) []string {
	v := d.Get("tags").(*schema.Set)
	out := make([]string, 0, v.Len())
	for _, x := range v.List() {
		out = append(out, x.(string))
	}
	return out
}

func resourceFMEFeatureFlagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	ttID := d.Get("traffic_type_id").(string)
	req := splitsdk.SplitCreateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
	sp, err := client.Splits.Create(wsID, ttID, req)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(sp.Name)
	if tags := tagListFromSet(d); len(tags) > 0 {
		if err := client.Tags.AssociateTags(wsID, sp.Name, splitsdk.ObjectTypeSplit, tags); err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceFMEFeatureFlagRead(ctx, d, meta)
}

func resourceFMEFeatureFlagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	sp, err := client.Splits.Get(wsID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if sp == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("name", sp.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", sp.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("flag_id", sp.ID); err != nil {
		return diag.FromErr(err)
	}
	switch {
	case sp.TrafficType != nil && sp.TrafficType.ID != "":
		if err := d.Set("traffic_type_id", sp.TrafficType.ID); err != nil {
			return diag.FromErr(err)
		}
	default:
		if err := d.Set("traffic_type_id", d.Get("traffic_type_id")); err != nil {
			return diag.FromErr(err)
		}
	}
	tagNames := make([]interface{}, 0, len(sp.Tags))
	for _, t := range sp.Tags {
		tagNames = append(tagNames, t.Name)
	}
	if err := d.Set("tags", schema.NewSet(schema.HashString, tagNames)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMEFeatureFlagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	if d.HasChange("description") {
		if _, err := client.Splits.UpdateDescription(wsID, d.Id(), d.Get("description").(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("tags") {
		tags := tagListFromSet(d)
		if err := client.Tags.AssociateTags(wsID, d.Id(), splitsdk.ObjectTypeSplit, tags); err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceFMEFeatureFlagRead(ctx, d, meta)
}

func resourceFMEFeatureFlagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	if err := client.Splits.Delete(wsID, d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
