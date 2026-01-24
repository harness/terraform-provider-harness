package experiment

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosExperiment() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for creating chaos experiments from experiment templates.",

		CreateContext: resourceExperimentCreate,
		ReadContext:   resourceExperimentRead,
		UpdateContext: resourceExperimentUpdate,
		DeleteContext: resourceExperimentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceExperimentImport,
		},

		Schema: resourceChaosExperimentSchema(),
	}
}

// resourceExperimentCreate creates a new experiment from a template
func resourceExperimentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Extract required fields
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	templateIdentity := d.Get("template_identity").(string)
	hubIdentity := d.Get("hub_identity").(string)
	name := d.Get("name").(string)
	infraRef := d.Get("infra_ref").(string)

	log.Printf("[DEBUG] Creating experiment from template: %s in hub: %s", templateIdentity, hubIdentity)

	// Build request
	req := chaos.TypesCreateExperimentFromTemplateRequest{
		AccountIdentifier:      c.AccountId,
		OrganizationIdentifier: orgID,
		ProjectIdentifier:      projectID,
		Name:                   name,
		InfraRef:               infraRef,
	}

	// Add optional fields
	if v, ok := d.GetOk("identity"); ok {
		req.Identity = v.(string)
		log.Printf("[DEBUG] Using provided identity: %s", req.Identity)
	} else {
		// Generate identity from name
		// API requires: lowercase letters, numbers, and dashes only
		// Convert to lowercase and replace invalid characters with dashes
		identity := strings.ToLower(name)
		// Replace any character that's not lowercase letter, number, or dash with dash
		reg := regexp.MustCompile(`[^a-z0-9-]+`)
		identity = reg.ReplaceAllString(identity, "-")
		// Remove leading/trailing dashes
		identity = strings.Trim(identity, "-")
		req.Identity = identity
		log.Printf("[DEBUG] Generated identity from name: %s -> %s", name, req.Identity)
	}

	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		req.Tags = expandTags(v.(*schema.Set).List())
	}

	// Add import_type (CRITICAL)
	importType := d.Get("import_type").(string)
	var importTypePtr *chaos.MongodbImportType
	if importType == "LOCAL" {
		local := chaos.LOCAL_MongodbImportType
		importTypePtr = &local
		log.Printf("[DEBUG] Using LOCAL import type (full copy)")
	} else {
		ref := chaos.REFERENCE_MongodbImportType
		importTypePtr = &ref
		log.Printf("[DEBUG] Using REFERENCE import type (template reference)")
	}
	req.ImportType = importTypePtr

	// Extract hub scope (where the template lives) - separate from experiment scope
	hubOrgID := d.Get("hub_org_id").(string)
	hubProjectID := d.Get("hub_project_id").(string)

	// Build optional parameters
	// CRITICAL: Query params define HUB scope (where template lives), NOT experiment scope
	// Hub identity should be passed as-is, without any prefix
	opts := &chaos.ExperimenttemplateApiCreateExperimentFromTemplateOpts{
		HubIdentity:            optional.NewString(hubIdentity),  // No prefix!
		OrganizationIdentifier: optional.NewString(hubOrgID),     // Hub's org (may be empty for account-level)
		ProjectIdentifier:      optional.NewString(hubProjectID), // Hub's project (may be empty for org/account-level)
	}

	log.Printf("[DEBUG] Hub location - org: %q, project: %q, identity: %q", hubOrgID, hubProjectID, hubIdentity)
	log.Printf("[DEBUG] Experiment location - org: %q, project: %q", orgID, projectID)

	if v, ok := d.GetOk("revision"); ok {
		opts.Revision = optional.NewString(v.(string))
		log.Printf("[DEBUG] Using template revision: %s", v.(string))
	}

	// Call CreateExperimentFromTemplate API
	log.Printf("[DEBUG] Calling CreateExperimentFromTemplate API")
	resp, httpResp, err := c.ExperimenttemplateApi.CreateExperimentFromTemplate(
		ctx,
		req,
		c.AccountId,
		templateIdentity,
		opts,
	)

	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	// Extract experiment ID from response (CRITICAL: Use Id, not Identity)
	var experimentID string
	if resp.Data != nil && resp.Data.Id != "" {
		experimentID = resp.Data.Id
		log.Printf("[DEBUG] Experiment created with ID: %s", experimentID)
	} else {
		return diag.Errorf("failed to get experiment ID from response")
	}

	// Store identity separately (different from ID)
	if resp.Data != nil && resp.Data.Identity != "" {
		d.Set("identity", resp.Data.Identity)
		log.Printf("[DEBUG] Experiment identity: %s", resp.Data.Identity)
	}

	// Set ID: org_id/project_id/experiment_id (CRITICAL: Use experiment_id, not identity)
	d.SetId(fmt.Sprintf("%s/%s/%s", orgID, projectID, experimentID))
	log.Printf("[DEBUG] Set resource ID: %s", d.Id())

	// Store manifest if LOCAL import (only available in create response)
	if importType == "LOCAL" && resp.Data != nil && resp.Data.Manifest != "" {
		d.Set("manifest", resp.Data.Manifest)
		log.Printf("[DEBUG] Stored manifest for LOCAL import, length: %d", len(resp.Data.Manifest))
	}

	// Read back to populate all computed fields
	return resourceExperimentRead(ctx, d, meta)
}

// resourceExperimentRead reads the experiment details
func resourceExperimentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID: org_id/project_id/experiment_id (CRITICAL: This is experiment_id, not identity)
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 {
		return diag.Errorf("invalid ID format, expected: org_id/project_id/experiment_id, got: %s", d.Id())
	}

	orgID := parts[0]
	projectID := parts[1]
	experimentID := parts[2] // This is the actual experiment ID from the API

	log.Printf("[DEBUG] Reading experiment: %s (org: %s, project: %s)", experimentID, orgID, projectID)

	// Call GetChaosV2Experiment API (CRITICAL: Use experiment_id, not identity)
	experiment, httpResp, err := c.DefaultApi.GetChaosV2Experiment(
		ctx,
		c.AccountId,
		orgID,
		projectID,
		experimentID, // Use experiment_id from resource ID
	)

	if err != nil {
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}

	// Check if experiment is soft-deleted
	if experiment.IsRemoved {
		log.Printf("[DEBUG] Experiment %s is soft-deleted, removing from state", experimentID)
		d.SetId("")
		return nil
	}

	// ========== SET ALL FIELDS ==========

	// Scope
	d.Set("org_id", experiment.OrgID)
	d.Set("project_id", experiment.ProjectID)

	// Basic info
	d.Set("name", experiment.Name)
	d.Set("identity", experiment.Identity)
	d.Set("description", experiment.Description)
	if len(experiment.Tags) > 0 {
		d.Set("tags", flattenTags(experiment.Tags))
	}

	// Infrastructure
	d.Set("infra_id", experiment.InfraID)
	if experiment.InfraType != nil {
		d.Set("infra_type", string(*experiment.InfraType))
	}

	// Computed fields
	d.Set("experiment_id", experiment.ExperimentID)
	d.Set("is_custom_experiment", experiment.IsCustomExperiment)

	// Experiment type
	if experiment.ExperimentType != nil {
		d.Set("experiment_type", string(*experiment.ExperimentType))
	}

	// Fault IDs
	if len(experiment.FaultIDs) > 0 {
		d.Set("fault_ids", experiment.FaultIDs)
	}

	// Scheduling
	d.Set("cron_syntax", experiment.CronSyntax)
	d.Set("is_cron_enabled", experiment.IsCronEnabled)
	d.Set("is_single_run_cron_enabled", experiment.IsSingleRunCronEnabled)

	// Execution history
	d.Set("last_executed_at", experiment.LastExecutedAt)
	d.Set("total_experiment_runs", experiment.TotalExperimentRuns)

	// Network
	d.Set("target_network_map_id", experiment.TargetNetworkMapID)

	// Metadata
	d.Set("created_at", experiment.CreatedAt)
	d.Set("created_by", experiment.CreatedBy)
	d.Set("updated_at", experiment.UpdatedAt)
	d.Set("updated_by", experiment.UpdatedBy)

	// Import type and manifest (infer from response)
	if experiment.TemplateDetails != nil {
		// REFERENCE import - has template details
		d.Set("import_type", "REFERENCE")
		log.Printf("[DEBUG] Experiment uses REFERENCE import (template reference)")

		templateDetails := []map[string]interface{}{
			{
				"identity":      experiment.TemplateDetails.Identity,
				"hub_reference": experiment.TemplateDetails.HubReference,
				"reference":     experiment.TemplateDetails.Reference,
				"revision":      experiment.TemplateDetails.Revision,
			},
		}
		d.Set("template_details", templateDetails)

		// Also set top-level hub_identity and template_identity for convenience
		d.Set("hub_identity", experiment.TemplateDetails.HubReference)
		d.Set("template_identity", experiment.TemplateDetails.Identity)
		if experiment.TemplateDetails.Revision != "" {
			d.Set("revision", experiment.TemplateDetails.Revision)
		}
	} else {
		// LOCAL import - manifest stored during create (not available in read API)
		// Keep existing manifest value from state
		if manifestVal, ok := d.GetOk("manifest"); ok && manifestVal.(string) != "" {
			d.Set("import_type", "LOCAL")
			log.Printf("[DEBUG] Experiment uses LOCAL import (manifest preserved from state)")
		}
	}

	log.Printf("[DEBUG] Successfully read experiment: %s", experimentID)
	return nil
}

// resourceExperimentUpdate updates the experiment
func resourceExperimentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Experiments created from templates have limited update support
	// Most fields are ForceNew, so changes will trigger recreation

	log.Printf("[DEBUG] Update called for experiment: %s", d.Id())

	// Check if only updatable fields changed
	if d.HasChanges("name", "description", "tags") {
		// For now, return error - updates not supported
		// Future: Implement via SaveChaosV2Experiment if needed
		return diag.Errorf("experiments created from templates cannot be updated directly. Changes to template_identity, hub_identity, infra_ref, or revision will trigger recreation.")
	}

	// If we reach here, no changes or only computed fields changed
	return resourceExperimentRead(ctx, d, meta)
}

// resourceExperimentDelete deletes the experiment
func resourceExperimentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 {
		return diag.Errorf("invalid ID format")
	}

	orgID := parts[0]
	projectID := parts[1]
	experimentIdentity := parts[2]

	log.Printf("[DEBUG] Deleting experiment: %s (org: %s, project: %s)", experimentIdentity, orgID, projectID)

	// Call DeleteChaosV2Experiment API
	_, httpResp, err := c.DefaultApi.DeleteChaosV2Experiment(
		ctx,
		c.AccountId,
		orgID,
		projectID,
		experimentIdentity,
	)

	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	d.SetId("")
	log.Printf("[DEBUG] Successfully deleted experiment: %s", experimentIdentity)
	return nil
}

// resourceExperimentImport imports an existing experiment
func resourceExperimentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Import format: org_id/project_id/experiment_identity
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid import ID format, expected: org_id/project_id/experiment_identity, got: %s", d.Id())
	}

	log.Printf("[DEBUG] Importing experiment with ID: %s", d.Id())

	// Set scope fields
	d.Set("org_id", parts[0])
	d.Set("project_id", parts[1])

	// Read to populate all fields
	diags := resourceExperimentRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read experiment during import: %v", diags)
	}

	log.Printf("[DEBUG] Successfully imported experiment: %s", d.Id())
	return []*schema.ResourceData{d}, nil
}

// Helper functions

// expandTags converts Terraform tags to SDK format
func expandTags(tags []interface{}) []string {
	result := make([]string, 0, len(tags))
	for _, tag := range tags {
		if tag != nil {
			result = append(result, tag.(string))
		}
	}
	return result
}

// flattenTags converts SDK tags to Terraform format
func flattenTags(tags []string) *schema.Set {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		result = append(result, tag)
	}
	return schema.NewSet(schema.HashString, result)
}
