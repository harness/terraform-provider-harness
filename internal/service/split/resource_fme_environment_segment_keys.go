package split

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMEEnvironmentSegmentKeys manages keys for a classic segment in a Split environment.
func ResourceFMEEnvironmentSegmentKeys() *schema.Resource {
	return &schema.Resource{
		Description: "Replace keys for a classic segment in an environment (`replace=true` semantics). Import id format: `org_id/project_id/environment_id/segment_name`.",

		CreateContext: resourceFMEEnvironmentSegmentKeysCreate,
		ReadContext:   resourceFMEEnvironmentSegmentKeysRead,
		UpdateContext: resourceFMEEnvironmentSegmentKeysUpdate,
		DeleteContext: resourceFMEEnvironmentSegmentKeysDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMEEnvironmentSegmentKeysImport,
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
			"environment_id": {
				Description: "Split environment ID.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"segment_name": {
				Description: "Classic segment name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"keys": {
				Description: "Segment keys (full replace on each apply).",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func fmeSegmentKeysID(orgID, projectID, envID, segName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, envID, segName)
}

func resourceFMEEnvironmentSegmentKeysImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, envID, segName, err := ParseImportID4(d.Id())
	if err != nil {
		return nil, err
	}
	if err := d.Set("org_id", orgID); err != nil {
		return nil, err
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, err
	}
	if err := d.Set("environment_id", envID); err != nil {
		return nil, err
	}
	if err := d.Set("segment_name", segName); err != nil {
		return nil, err
	}
	d.SetId(fmeSegmentKeysID(orgID, projectID, envID, segName))
	return []*schema.ResourceData{d}, nil
}

func keysFromSet(d *schema.ResourceData) []string {
	v := d.Get("keys").(*schema.Set)
	out := make([]string, 0, v.Len())
	for _, x := range v.List() {
		out = append(out, x.(string))
	}
	sort.Strings(out)
	return out
}

func resourceFMEEnvironmentSegmentKeysCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	keys := keysFromSet(d)
	if err := client.Environments.AddSegmentKeysWithReplace(wsID, envID, segName, true, keys); err != nil {
		return diag.FromErr(err)
	}
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	d.SetId(fmeSegmentKeysID(orgID, projectID, envID, segName))
	return resourceFMEEnvironmentSegmentKeysRead(ctx, d, meta)
}

func resourceFMEEnvironmentSegmentKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	keys, err := client.Environments.GetSegmentKeysAll(wsID, envID, segName)
	if err != nil {
		return diag.FromErr(err)
	}
	sort.Strings(keys)
	interf := make([]interface{}, len(keys))
	for i, k := range keys {
		interf[i] = k
	}
	if err := d.Set("keys", schema.NewSet(schema.HashString, interf)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMEEnvironmentSegmentKeysUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	keys := keysFromSet(d)
	if err := client.Environments.AddSegmentKeysWithReplace(wsID, envID, segName, true, keys); err != nil {
		return diag.FromErr(err)
	}
	return resourceFMEEnvironmentSegmentKeysRead(ctx, d, meta)
}

func resourceFMEEnvironmentSegmentKeysDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	keys, err := client.Environments.GetSegmentKeysAll(wsID, envID, segName)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(keys) > 0 {
		if err := client.Environments.RemoveSegmentKeys(wsID, envID, segName, keys); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}
