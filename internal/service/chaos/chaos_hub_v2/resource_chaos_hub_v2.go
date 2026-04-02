package chaos_hub_v2

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosHubV2() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Chaos Hub V2.",

		CreateContext: resourceChaosHubV2Create,
		ReadContext:   resourceChaosHubV2Read,
		UpdateContext: resourceChaosHubV2Update,
		DeleteContext: resourceChaosHubV2Delete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceChaosHubV2Import,
		},
		Schema: resourceChaosHubV2Schema(),
	}
}

// resourceChaosHubV2Create creates a new chaos hub
func resourceChaosHubV2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Build the create request
	req := buildCreateChaosHubV2Request(d)

	// Extract identifiers
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	log.Printf("[DEBUG] Creating chaos hub with identity: %s, name: %s", req.Identity, req.Name)

	// Make the API call
	resp, httpResp, err := c.DefaultApi.CreateChaosHub(
		ctx,
		req,
		c.AccountId,
		orgID,
		projectID,
	)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	// Set the ID to the hub identity
	d.SetId(resp.Identity)

	log.Printf("[DEBUG] Created chaos hub with identity: %s, hub_id: %s", resp.Identity, resp.HubId)

	return resourceChaosHubV2Read(ctx, d, meta)
}

// resourceChaosHubV2Read reads the chaos hub details
func resourceChaosHubV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	hubIdentity := d.Id()
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	log.Printf("[DEBUG] Reading chaos hub with identity: %s", hubIdentity)

	// Get the hub using the REST API
	hub, httpResp, err := c.DefaultApi.GetChaosHub(
		ctx,
		c.AccountId,
		orgID,
		projectID,
		hubIdentity,
	)
	if err != nil {
		// Use graceful destroy handling for hub read errors
		// This handles 404 and certain 500 errors (resource not found/inconsistent state)
		return helpers.HandleChaosReadApiErrorWithGracefulDestroy(err, d, httpResp, []string{
			"hub not found",
			"no matching hub",
		})
	}

	// Set the resource data
	if err := setChaosHubV2Data(d, &hub, c.AccountId, orgID, projectID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set chaos hub data: %v", err))
	}

	return nil
}

// resourceChaosHubV2Update updates an existing chaos hub
func resourceChaosHubV2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	hubIdentity := d.Id()
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	// Build the update request
	req := buildUpdateChaosHubV2Request(d)

	log.Printf("[DEBUG] Updating chaos hub with identity: %s", hubIdentity)

	// Update the hub using the REST API
	hub, httpResp, err := c.DefaultApi.UpdateChaosHub(
		ctx,
		req,
		c.AccountId,
		orgID,
		projectID,
		hubIdentity,
	)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Updated chaos hub with identity: %s", hub.Identity)

	return resourceChaosHubV2Read(ctx, d, meta)
}

// resourceChaosHubV2Delete deletes a chaos hub
func resourceChaosHubV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	hubIdentity := d.Id()
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	log.Printf("[DEBUG] Deleting chaos hub with identity: %s", hubIdentity)

	_, httpResp, err := c.DefaultApi.DeleteHub(
		ctx,
		c.AccountId,
		hubIdentity,
		&chaos.DefaultApiDeleteHubOpts{
			OrganizationIdentifier: optional.NewString(orgID),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)
	if err != nil {
		// Handle graceful errors during delete (API constraints)
		// Only handle "at least one hub required" - template errors should fail properly
		diags := helpers.HandleChaosReadApiErrorWithGracefulDestroy(err, d, httpResp, []string{
			"at least one hub is required",
			"at least one hub required",
		})
		// If the helper cleared the state (SetId("")), we're done
		if d.Id() == "" {
			log.Printf("[DEBUG] Hub delete handled gracefully (API constraint): %s", hubIdentity)
			return diags
		}
		// Otherwise, it's a real error
		return diags
	}

	log.Printf("[DEBUG] Deleted chaos hub with identity: %s", hubIdentity)

	// Clear the ID from state
	d.SetId("")
	return nil
}

// resourceChaosHubV2Import handles the import of a chaos hub resource
func resourceChaosHubV2Import(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	// Parse the import ID which can be in one of these formats:
	// 1. Org level: "org-id/project-id/hub-identity"
	// 2. Account level: "hub-identity"
	importID := d.Id()
	parts := strings.Split(importID, "/")

	var hubIdentity, orgID, projectID string

	switch len(parts) {
	case 1:
		// Account level: "hub-identity"
		hubIdentity = parts[0]
	case 2:
		// Org level: "org-id/hub-identity"
		orgID = parts[0]
		hubIdentity = parts[1]
	case 3:
		// Project level: "org-id/project-id/hub-identity"
		orgID = parts[0]
		projectID = parts[1]
		hubIdentity = parts[2]
	default:
		return nil, fmt.Errorf("invalid import ID format. Expected \"org-id/project-id/hub-identity\" or \"org-id/hub-identity\" or \"hub-identity\", got: %s", importID)
	}

	if hubIdentity == "" {
		return nil, fmt.Errorf("hub identity cannot be empty")
	}

	log.Printf("[DEBUG] Importing chaos hub with identity: %s, org: %s, project: %s", hubIdentity, orgID, projectID)

	// Set the ID in the format that our Read function expects
	d.SetId(hubIdentity)

	// Set the individual ID fields
	if orgID != "" {
		if err := d.Set("org_id", orgID); err != nil {
			return nil, fmt.Errorf("failed to set org_id: %v", err)
		}
	}
	if projectID != "" {
		if err := d.Set("project_id", projectID); err != nil {
			return nil, fmt.Errorf("failed to set project_id: %v", err)
		}
	}
	if err := d.Set("identity", hubIdentity); err != nil {
		return nil, fmt.Errorf("failed to set identity: %v", err)
	}

	// Get the hub details using the REST API
	hub, httpResp, err := c.DefaultApi.GetChaosHub(
		ctx,
		c.AccountId,
		orgID,
		projectID,
		hubIdentity,
	)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			return nil, fmt.Errorf("chaos hub not found with identity: %s (account_id: %s, org_id: %s, project_id: %s)",
				hubIdentity, c.AccountId, orgID, projectID)
		}
		return nil, fmt.Errorf("failed to get chaos hub details for identity: %s: %v", hubIdentity, err)
	}

	log.Printf("[DEBUG] Found chaos hub: %+v", hub)

	// Call the read function to populate the rest of the state
	log.Printf("[DEBUG] Calling read function to populate state...")
	diags := resourceChaosHubV2Read(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading chaos hub: %v", diags)
	}

	log.Printf("[DEBUG] Successfully imported chaos hub: %s", d.Id())
	return []*schema.ResourceData{d}, nil
}

// buildCreateChaosHubV2Request builds the request object for create operations
func buildCreateChaosHubV2Request(d *schema.ResourceData) chaos.Chaoshubv2CreateHubRequest {
	req := chaos.Chaoshubv2CreateHubRequest{
		Identity:     d.Get("identity").(string),
		Name:         d.Get("name").(string),
		ConnectorRef: d.Get("connector_ref").(string),
	}

	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	if v, ok := d.GetOk("repo_branch"); ok {
		req.RepoBranch = v.(string)
	}

	if v, ok := d.GetOk("repo_name"); ok {
		req.RepoName = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		tagsList := v.([]interface{})
		tags := make([]string, len(tagsList))
		for i, tag := range tagsList {
			tags[i] = tag.(string)
		}
		req.Tags = tags
	}

	return req
}

// buildUpdateChaosHubV2Request builds the request object for update operations
func buildUpdateChaosHubV2Request(d *schema.ResourceData) chaos.Chaoshubv2UpdateHubRequest {
	req := chaos.Chaoshubv2UpdateHubRequest{
		Name: d.Get("name").(string),
	}

	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		tagsList := v.([]interface{})
		tags := make([]string, len(tagsList))
		for i, tag := range tagsList {
			tags[i] = tag.(string)
		}
		req.Tags = tags
	}

	return req
}

// setChaosHubV2Data sets the chaos hub data in the resource
func setChaosHubV2Data(d *schema.ResourceData, hub *chaos.Chaoshubv2GetHubResponse, accountID, orgID, projectID string) error {
	d.Set("identity", hub.Identity)
	d.Set("name", hub.Name)
	d.Set("hub_id", hub.HubId)
	d.Set("account_id", hub.AccountID)
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)
	d.Set("is_default", hub.IsDefault)
	d.Set("is_removed", hub.IsRemoved)

	if hub.Description != "" {
		d.Set("description", hub.Description)
	}

	if hub.RepoBranch != "" {
		d.Set("repo_branch", hub.RepoBranch)
	}

	if hub.RepoName != "" {
		d.Set("repo_name", hub.RepoName)
	}

	if hub.RepoUrl != "" {
		d.Set("repo_url", hub.RepoUrl)
	}

	if hub.ConnectorId != "" {
		d.Set("connector_id", hub.ConnectorId)
	}

	if len(hub.Tags) > 0 {
		d.Set("tags", hub.Tags)
	}

	if hub.CreatedAt > 0 {
		d.Set("created_at", hub.CreatedAt)
	}

	if hub.UpdatedAt > 0 {
		d.Set("updated_at", hub.UpdatedAt)
	}

	if hub.LastSyncedAt > 0 {
		d.Set("last_synced_at", hub.LastSyncedAt)
	}

	if hub.CreatedBy != "" {
		d.Set("created_by", hub.CreatedBy)
	}

	if hub.UpdatedBy != "" {
		d.Set("updated_by", hub.UpdatedBy)
	}

	if hub.ActionTemplateCount > 0 {
		d.Set("action_template_count", hub.ActionTemplateCount)
	}

	if hub.ExperimentTemplateCount > 0 {
		d.Set("experiment_template_count", hub.ExperimentTemplateCount)
	}

	if hub.FaultTemplateCount > 0 {
		d.Set("fault_template_count", hub.FaultTemplateCount)
	}

	if hub.ProbeTemplateCount > 0 {
		d.Set("probe_template_count", hub.ProbeTemplateCount)
	}

	return nil
}
