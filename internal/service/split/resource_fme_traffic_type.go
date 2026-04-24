package split

import (
	"context"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMETrafficType creates and deletes a Split traffic type (no update API).
func ResourceFMETrafficType() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) traffic type. Import id format: `org_id/project_id/traffic_type_id`.",

		CreateContext: resourceFMETrafficTypeCreate,
		ReadContext:   resourceFMETrafficTypeRead,
		DeleteContext: resourceFMETrafficTypeDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMETrafficTypeImport,
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
			"name": {
				Description: "Traffic type name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"display_attribute_id": {
				Description: "Optional display attribute ID.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if k != "display_attribute_id" {
						return false
					}
					// Split often defaults omitted display attribute to "name".
					if old == "" && new == "name" {
						return true
					}
					if old == "name" && new == "" {
						return true
					}
					return old == new
				},
			},
			"traffic_type_id": {
				Description: "The Split traffic type ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMETrafficTypeImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, ttID, err := ParseImportID3(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	d.SetId(ttID)

	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := ResolveWorkspaceIDFromOrgProject(ctx, meta, client, orgID, projectID)
	if err != nil {
		return nil, err
	}
	tt, err := client.TrafficTypes.FindByID(wsID, ttID)
	if err != nil {
		return nil, err
	}
	if tt == nil {
		return nil, fmt.Errorf("traffic type %q not found in workspace for org_id %q project_id %q", ttID, orgID, projectID)
	}
	if err := d.Set("name", tt.Name); err != nil {
		return nil, err
	}
	if err := d.Set("display_attribute_id", tt.DisplayAttributeID); err != nil {
		return nil, err
	}
	if err := d.Set("traffic_type_id", tt.ID); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceFMETrafficTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	req := splitsdk.CreateRequest{
		Name:               d.Get("name").(string),
		DisplayAttributeID: d.Get("display_attribute_id").(string),
	}
	tt, err := client.TrafficTypes.Create(wsID, req)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(tt.ID)
	return resourceFMETrafficTypeRead(ctx, d, meta)
}

func resourceFMETrafficTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	tt, err := client.TrafficTypes.FindByID(wsID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if tt == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("name", tt.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_attribute_id", tt.DisplayAttributeID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("traffic_type_id", tt.ID); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMETrafficTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	if err := client.TrafficTypes.Delete(wsID, d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
