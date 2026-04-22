package split

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMERuleBasedSegmentEnvironmentAssociation enables a rule-based segment in an environment and manages its definition JSON.
func ResourceFMERuleBasedSegmentEnvironmentAssociation() *schema.Resource {
	return &schema.Resource{
		Description: "Enable a workspace rule-based segment in an FME environment and manage its `RuleBasedSegmentDefinition` JSON. The segment must exist (`harness_fme_rule_based_segment`). Import id format: `org_id/project_id/environment_id/segment_name`.",

		CreateContext: resourceFMERuleBasedSegmentEnvironmentAssociationCreate,
		ReadContext:   resourceFMERuleBasedSegmentEnvironmentAssociationRead,
		UpdateContext: resourceFMERuleBasedSegmentEnvironmentAssociationUpdate,
		DeleteContext: resourceFMERuleBasedSegmentEnvironmentAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMERuleBasedSegmentEnvironmentAssociationImport,
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
				Description: "Rule-based segment name (must match an existing workspace segment).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"definition_json": {
				Description: "JSON body for `RuleBasedSegmentDefinition` in this environment. After `EnableInEnvironment`, this is applied via update. Semantically equivalent JSON (e.g. different key order from `jsonencode` vs API refresh) is not treated as a diff.",
				Type:        schema.TypeString,
				Required:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == new {
						return true
					}
					return ruleBasedSegmentDefinitionJSONSemanticallyEqual(old, new)
				},
			},
		},
	}
}

// rbsEnvAssocDefinitionJSONForState builds definition_json for Terraform state. Split's list-in-environment
// payload often omits title/comment even after UpdateDefinition; merge from prior state/config and workspace Get.
func rbsEnvAssocDefinitionJSONForState(ctx context.Context, client *splitsdk.APIClient, wsID, segName string, entry *splitsdk.RuleBasedSegmentEnvironmentEntry, priorJSON string) (string, error) {
	_ = ctx
	body := ruleBasedSegmentDefinitionFromEnvEntry(entry)
	var prior splitsdk.RuleBasedSegmentDefinition
	if priorJSON != "" {
		_ = json.Unmarshal([]byte(priorJSON), &prior)
		normalizeRuleBasedSegmentDefinition(&prior)
	}
	if body.Title == "" && prior.Title != "" {
		body.Title = prior.Title
	}
	if body.Comment == "" && prior.Comment != "" {
		body.Comment = prior.Comment
	}
	if wsDef, err := client.RuleBasedSegments.Get(wsID, segName); err == nil && wsDef != nil {
		if body.Title == "" {
			body.Title = wsDef.Title
		}
		if body.Comment == "" {
			body.Comment = wsDef.Comment
		}
	}
	normalizeRuleBasedSegmentDefinition(&body)
	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func resourceFMERuleBasedSegmentEnvironmentAssociationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := ResolveWorkspaceIDFromOrgProject(ctx, meta, client, orgID, projectID)
	if err != nil {
		return nil, err
	}
	entries, err := client.RuleBasedSegments.ListInEnvironment(wsID, envID)
	if err != nil {
		return nil, err
	}
	entry := findRuleBasedSegmentEnvEntry(entries, segName)
	if entry == nil {
		return nil, fmt.Errorf("rule-based segment %q not found in environment %q", segName, envID)
	}
	jsonStr, err := rbsEnvAssocDefinitionJSONForState(ctx, client, wsID, segName, entry, "")
	if err != nil {
		return nil, err
	}
	if err := d.Set("definition_json", jsonStr); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func ruleBasedSegmentPresentInEnvironment(ctx context.Context, client *splitsdk.APIClient, wsID, envID, segName string) (*splitsdk.RuleBasedSegmentEnvironmentEntry, error) {
	const maxAttempts = 12
	const poll = 350 * time.Millisecond
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			t := time.NewTimer(poll)
			select {
			case <-ctx.Done():
				t.Stop()
				return nil, ctx.Err()
			case <-t.C:
			}
		}
		entries, err := client.RuleBasedSegments.ListInEnvironment(wsID, envID)
		if err != nil {
			return nil, err
		}
		if entry := findRuleBasedSegmentEnvEntry(entries, segName); entry != nil {
			return entry, nil
		}
	}
	return nil, nil
}

func resourceFMERuleBasedSegmentEnvironmentAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	var body splitsdk.RuleBasedSegmentDefinition
	if err := json.Unmarshal([]byte(d.Get("definition_json").(string)), &body); err != nil {
		return diag.FromErr(err)
	}
	if _, err := client.RuleBasedSegments.EnableInEnvironment(envID, segName); err != nil {
		return diag.FromErr(err)
	}
	if _, err := client.RuleBasedSegments.UpdateDefinition(wsID, envID, segName, body); err != nil {
		return diag.FromErr(err)
	}
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	d.SetId(segmentEnvAssociationID(orgID, projectID, envID, segName))
	return resourceFMERuleBasedSegmentEnvironmentAssociationRead(ctx, d, meta)
}

func resourceFMERuleBasedSegmentEnvironmentAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	entry, err := ruleBasedSegmentPresentInEnvironment(ctx, client, wsID, envID, segName)
	if err != nil {
		return diag.FromErr(err)
	}
	if entry == nil {
		d.SetId("")
		return nil
	}
	priorJSON := ""
	if v, ok := d.GetOk("definition_json"); ok {
		priorJSON = v.(string)
	}
	jsonStr, err := rbsEnvAssocDefinitionJSONForState(ctx, client, wsID, segName, entry, priorJSON)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("definition_json", jsonStr); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceFMERuleBasedSegmentEnvironmentAssociationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	if !d.HasChange("definition_json") {
		return resourceFMERuleBasedSegmentEnvironmentAssociationRead(ctx, d, meta)
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	var body splitsdk.RuleBasedSegmentDefinition
	if err := json.Unmarshal([]byte(d.Get("definition_json").(string)), &body); err != nil {
		return diag.FromErr(err)
	}
	if _, err := client.RuleBasedSegments.UpdateDefinition(wsID, envID, segName, body); err != nil {
		return diag.FromErr(err)
	}
	return resourceFMERuleBasedSegmentEnvironmentAssociationRead(ctx, d, meta)
}

func resourceFMERuleBasedSegmentEnvironmentAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	envID := d.Get("environment_id").(string)
	segName := d.Get("segment_name").(string)
	if err := client.RuleBasedSegments.DeleteInEnvironment(envID, segName); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
