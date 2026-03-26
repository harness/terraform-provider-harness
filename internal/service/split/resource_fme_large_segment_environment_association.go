package split

import (
	"context"
	"time"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMELargeSegmentEnvironmentAssociation ensures a large segment has a definition in a Split environment.
func ResourceFMELargeSegmentEnvironmentAssociation() *schema.Resource {
	return &schema.Resource{
		Description: "Create or remove a large segment definition in an FME environment (workspace-level segment must already exist). Use one resource per environment. Import id format: `org_id/project_id/environment_id/segment_name`.",

		CreateContext: resourceFMELargeSegmentEnvironmentAssociationCreate,
		ReadContext:   resourceFMELargeSegmentEnvironmentAssociationRead,
		DeleteContext: resourceFMELargeSegmentEnvironmentAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMELargeSegmentEnvironmentAssociationImport,
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
				Description: "Large segment name (workspace-level segment must exist).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceFMELargeSegmentEnvironmentAssociationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
	d.SetId(segmentEnvAssociationID(orgID, projectID, envID, segName))
	return []*schema.ResourceData{d}, nil
}

func resourceFMELargeSegmentEnvironmentAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	if err := client.LargeSegments.CreateDefinitionForEnvironment(envID, segName); err != nil {
		return diag.FromErr(err)
	}
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	d.SetId(segmentEnvAssociationID(orgID, projectID, envID, segName))
	return resourceFMELargeSegmentEnvironmentAssociationRead(ctx, d, meta)
}

// largeSegmentDefinedInEnvironment is true when the large segment appears in ListInEnvironment for the workspace and env.
func largeSegmentDefinedInEnvironment(ctx context.Context, client *split.APIClient, wsID, envID, segName string) (bool, error) {
	const maxAttempts = 12
	const poll = 350 * time.Millisecond
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			t := time.NewTimer(poll)
			select {
			case <-ctx.Done():
				t.Stop()
				return false, ctx.Err()
			case <-t.C:
			}
		}
		list, err := client.LargeSegments.ListInEnvironment(wsID, envID)
		if err != nil {
			return false, err
		}
		for _, e := range list {
			if e.Name == segName {
				return true, nil
			}
		}
	}
	return false, nil
}

func resourceFMELargeSegmentEnvironmentAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	present, err := largeSegmentDefinedInEnvironment(ctx, client, wsID, envID, segName)
	if err != nil {
		return diag.FromErr(err)
	}
	if !present {
		d.SetId("")
		return nil
	}
	return nil
}

func resourceFMELargeSegmentEnvironmentAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	if err := client.LargeSegments.DeleteInEnvironment(envID, segName); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
