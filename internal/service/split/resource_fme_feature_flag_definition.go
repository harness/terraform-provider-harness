package split

import (
	"context"
	"encoding/json"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMEFeatureFlagDefinition manages a Split feature flag definition in one environment.
func ResourceFMEFeatureFlagDefinition() *schema.Resource {
	return &schema.Resource{
		Description: "Create, update, and remove a Harness FME (Split) feature flag definition in an environment. `definition` is JSON matching Split's definition payload (see Split API). Import id format: `org_id/project_id/environment_id/flag_name`.",

		CreateContext: resourceFMEFeatureFlagDefinitionCreate,
		ReadContext:   resourceFMEFeatureFlagDefinitionRead,
		UpdateContext: resourceFMEFeatureFlagDefinitionUpdate,
		DeleteContext: resourceFMEFeatureFlagDefinitionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMEFeatureFlagDefinitionImport,
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
			"flag_name": {
				Description: "Feature flag (split) name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"definition": {
				Description: "JSON object for the split definition (treatments, defaultTreatment, defaultRule, trafficAllocation, rules, etc.).",
				Type:        schema.TypeString,
				Required:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if k != "definition" {
						return false
					}
					return splitDefinitionJSONSemanticallyEqual(old, new)
				},
			},
			"definition_id": {
				Description: "Split definition id from the API when available.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func fmeFeatureFlagDefinitionID(orgID, projectID, envID, flagName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, envID, flagName)
}

func resourceFMEFeatureFlagDefinitionImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	orgID, projectID, envID, flagName, err := ParseImportID4(d.Id())
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
	if err := d.Set("flag_name", flagName); err != nil {
		return nil, err
	}
	d.SetId(fmeFeatureFlagDefinitionID(orgID, projectID, envID, flagName))
	return []*schema.ResourceData{d}, nil
}

func splitDefinitionRequestFromString(s string) (splitsdk.SplitDefinitionRequest, error) {
	var req splitsdk.SplitDefinitionRequest
	if err := json.Unmarshal([]byte(s), &req); err != nil {
		return req, err
	}
	return req, nil
}

func resourceFMEFeatureFlagDefinitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	req, err := splitDefinitionRequestFromString(d.Get("definition").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	flagName := d.Get("flag_name").(string)
	envID := d.Get("environment_id").(string)
	if _, err := client.Splits.CreateDefinition(wsID, flagName, envID, req); err != nil {
		return diag.FromErr(err)
	}
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	d.SetId(fmeFeatureFlagDefinitionID(orgID, projectID, envID, flagName))
	return resourceFMEFeatureFlagDefinitionRead(ctx, d, meta)
}

func resourceFMEFeatureFlagDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	flagName := d.Get("flag_name").(string)
	envID := d.Get("environment_id").(string)
	def, err := client.Splits.GetDefinition(wsID, flagName, envID)
	if err != nil {
		return diag.FromErr(err)
	}
	if def == nil {
		d.SetId("")
		return nil
	}
	defStr, err := splitDefinitionRequestJSONForState(def)
	if err != nil {
		return diag.FromErr(err)
	}
	prior, _ := d.Get("definition").(string)
	merged, err := splitDefinitionMergePresentationFromPrior(defStr, prior)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("definition", merged); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("definition_id", def.ID); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMEFeatureFlagDefinitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	req, err := splitDefinitionRequestFromString(d.Get("definition").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	flagName := d.Get("flag_name").(string)
	envID := d.Get("environment_id").(string)
	if _, err := client.Splits.UpdateDefinitionFull(wsID, flagName, envID, req); err != nil {
		return diag.FromErr(err)
	}
	return resourceFMEFeatureFlagDefinitionRead(ctx, d, meta)
}

func resourceFMEFeatureFlagDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	flagName := d.Get("flag_name").(string)
	envID := d.Get("environment_id").(string)
	if err := client.Splits.RemoveDefinition(wsID, flagName, envID); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
