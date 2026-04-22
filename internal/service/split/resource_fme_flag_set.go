package split

import (
	"context"
	"fmt"
	"strings"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMEFlagSet creates and deletes a Split flag set (no update API).
func ResourceFMEFlagSet() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) flag set. Import id format: `org_id/project_id/flag_set_id`.",

		CreateContext: resourceFMEFlagSetCreate,
		ReadContext:   resourceFMEFlagSetRead,
		DeleteContext: resourceFMEFlagSetDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMEFlagSetImport,
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
				Description: "Flag set name. Split requires `^[a-z0-9][_a-z0-9]*`, max 50 characters (lowercase letters, digits, underscores; no hyphens).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Flag set description.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"flag_set_id": {
				Description: "The Split flag set ID (same as `id`).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceFMEFlagSetImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, fsID, err := ParseImportID3(d.Id())
	if err != nil {
		return nil, err
	}
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("split client: %v", diags)
	}
	fs, err := client.FlagSets.FindByID(fsID)
	if err != nil {
		return nil, err
	}
	if fs == nil {
		return nil, fmt.Errorf("cannot import flag set: Split API returned no flag set for id %q", fsID)
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("name", fs.Name); err != nil {
		return nil, err
	}
	if err := d.Set("description", fs.Description); err != nil {
		return nil, err
	}
	if err := d.Set("flag_set_id", fs.ID); err != nil {
		return nil, err
	}
	d.SetId(fsID)
	return []*schema.ResourceData{d}, nil
}

// splitFlagSetErrLooksNotFound is true when FindByID/Delete failed because the flag set no longer exists.
func splitFlagSetErrLooksNotFound(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "404") || strings.Contains(s, "not found") || strings.Contains(s, "Not Found")
}

func resourceFMEFlagSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	req := splitsdk.FlagSetRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Workspace:   &splitsdk.WorkspaceIDRef{Type: "workspace", ID: wsID},
	}
	fs, err := client.FlagSets.Create(req)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fs.ID)
	return resourceFMEFlagSetRead(ctx, d, meta)
}

func resourceFMEFlagSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	fs, err := client.FlagSets.FindByID(d.Id())
	if err != nil {
		if splitFlagSetErrLooksNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	if fs == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("name", fs.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", fs.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("flag_set_id", fs.ID); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMEFlagSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	if err := client.FlagSets.Delete(d.Id()); err != nil {
		if splitFlagSetErrLooksNotFound(err) {
			return nil
		}
		return diag.FromErr(err)
	}
	return nil
}
