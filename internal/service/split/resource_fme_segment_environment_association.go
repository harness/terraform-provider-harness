package split

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMESegmentEnvironmentAssociation activates or deactivates a classic segment in a Split environment.
func ResourceFMESegmentEnvironmentAssociation() *schema.Resource {
	return &schema.Resource{
		Description: "Activate a Split segment in an environment (create) or remove activation (destroy). Import id format: `org_id/project_id/environment_id/segment_name`.",

		CreateContext: resourceFMESegmentEnvironmentAssociationCreate,
		ReadContext:   resourceFMESegmentEnvironmentAssociationRead,
		DeleteContext: resourceFMESegmentEnvironmentAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMESegmentEnvironmentAssociationImport,
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
		},
	}
}

func segmentEnvAssociationID(orgID, projectID, envID, segName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, envID, segName)
}

func resourceFMESegmentEnvironmentAssociationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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

func resourceFMESegmentEnvironmentAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	if _, err := client.Segments.Activate(envID, segName); err != nil {
		return diag.FromErr(err)
	}
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	d.SetId(segmentEnvAssociationID(orgID, projectID, envID, segName))
	return resourceFMESegmentEnvironmentAssociationRead(ctx, d, meta)
}

// splitSegmentEnvKeysErrLooksNotFound is true when GetSegmentKeys failed because the segment is not active
// in the environment (or not yet visible). We retry only on these errors after Activate.
func splitSegmentEnvKeysErrLooksNotFound(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "404") || strings.Contains(s, "not found") || strings.Contains(s, "Not Found")
}

// segmentActivatedInEnvironment reports whether the classic segment is active in the environment.
// ListSegmentsAll returns API segment ids when present, which do not match Terraform's segment_name;
// GetSegmentKeys uses env id + segment name in the URL, which matches Activate/Deactivate and import.
func segmentActivatedInEnvironment(ctx context.Context, client *split.APIClient, wsID, envID, segName string) (bool, error) {
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
		_, _, err := client.Environments.GetSegmentKeys(wsID, envID, segName, 0, 1)
		if err == nil {
			return true, nil
		}
		if splitSegmentEnvKeysErrLooksNotFound(err) {
			continue
		}
		return false, err
	}
	return false, nil
}

func resourceFMESegmentEnvironmentAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	active, err := segmentActivatedInEnvironment(ctx, client, wsID, envID, segName)
	if err != nil {
		return diag.FromErr(err)
	}
	if !active {
		d.SetId("")
		return nil
	}
	return nil
}

func resourceFMESegmentEnvironmentAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	if err := client.Segments.Deactivate(envID, segName); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
