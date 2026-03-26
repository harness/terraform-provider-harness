package split

import (
	"context"
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFMERuleBasedSegment manages a Split rule-based segment at workspace scope.
func ResourceFMERuleBasedSegment() *schema.Resource {
	return &schema.Resource{
		Description: "Create and delete a Harness FME (Split) rule-based segment at workspace scope. To enable the segment in an environment and manage its definition JSON, use `harness_fme_rule_based_segment_environment_association`. Import id format: `org_id/project_id/segment_name`.",

		CreateContext: resourceFMERuleBasedSegmentCreate,
		ReadContext:   resourceFMERuleBasedSegmentRead,
		DeleteContext: resourceFMERuleBasedSegmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceFMERuleBasedSegmentImport,
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
				Description: "Rule-based segment name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceFMERuleBasedSegmentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
	ttID, err := ruleBasedSegmentResolveTrafficTypeID(client, wsID, name)
	if err != nil {
		return nil, err
	}
	if err := d.Set("traffic_type_id", ttID); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceFMERuleBasedSegmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	name := d.Get("name").(string)
	ttID := d.Get("traffic_type_id").(string)
	if _, err := client.RuleBasedSegments.Create(wsID, ttID, splitsdk.RuleBasedSegmentCreateRequest{Name: name}); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(name)
	return resourceFMERuleBasedSegmentRead(ctx, d, meta)
}

func ruleBasedSegmentDefinitionFromEnvEntry(e *splitsdk.RuleBasedSegmentEnvironmentEntry) splitsdk.RuleBasedSegmentDefinition {
	if e == nil {
		return splitsdk.RuleBasedSegmentDefinition{}
	}
	return splitsdk.RuleBasedSegmentDefinition{
		Title:            e.Title,
		Comment:          e.Comment,
		Rules:            e.Rules,
		ExcludedKeys:     e.ExcludedKeys,
		ExcludedSegments: e.ExcludedSegments,
	}
}

func findRuleBasedSegmentEnvEntry(entries []splitsdk.RuleBasedSegmentEnvironmentEntry, segmentName string) *splitsdk.RuleBasedSegmentEnvironmentEntry {
	for i := range entries {
		e := &entries[i]
		if e.Name == segmentName || e.ID == segmentName {
			return e
		}
	}
	return nil
}

func resourceFMERuleBasedSegmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	segName := d.Id()
	def, err := client.RuleBasedSegments.Get(wsID, segName)
	if err != nil {
		return diag.FromErr(err)
	}
	if def == nil {
		d.SetId("")
		return nil
	}
	return nil
}

func resourceFMERuleBasedSegmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return diags
	}
	wsID, diags := ResolveWorkspaceIDFromSchema(ctx, d, meta)
	if diags.HasError() {
		return diags
	}
	name := d.Id()
	if err := client.RuleBasedSegments.Delete(wsID, name); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
