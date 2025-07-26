package chaos_hub

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChaosHubImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	c := meta.(*internal.Session).ChaosClient

	// Parse the import ID which can be in one of these formats:
	// 1. Account level: "hub-id"
	// 2. Org level: "org-id/hub-id"
	// 3. Project level: "org-id/project-id/hub-id"
	importID := d.Id()
	parts := strings.Split(importID, "/")

	var hubID, orgID, projectID string

	switch len(parts) {
	case 1:
		// Account level: "hub-id"
		hubID = parts[0]
	case 2:
		// Org level: "org-id/hub-id"
		orgID = parts[0]
		hubID = parts[1]
	case 3:
		// Project level: "org-id/project-id/hub-id"
		orgID = parts[0]
		projectID = parts[1]
		hubID = parts[2]
	default:
		return nil, fmt.Errorf("invalid import ID format. Expected \"<hub-id>\", \"<org-id>/<hub-id>\", or \"<org-id>/<project-id>/<hub-id>\"")
	}

	if hubID == "" {
		return nil, fmt.Errorf("hub ID cannot be empty")
	}

	// Create a client for the Chaos Hub API
	client := chaos.NewChaosHubClient(c)

	// Get the account ID from the provider config
	accountID := c.AccountId
	if accountID == "" {
		return nil, fmt.Errorf("account ID must be configured in the provider")
	}

	// Create identifiers for the request
	identifiers := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if orgID != "" {
		identifiers.OrgIdentifier = orgID
	}
	if projectID != "" {
		identifiers.ProjectIdentifier = projectID
	}

	log.Printf("[DEBUG] Importing hub with ID: %s, org: %s, project: %s", hubID, orgID, projectID)

	// Get the full hub details using the ID
	hub, err := client.Get(ctx, identifiers, hubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chaos hub details for ID: %s: %v", hubID, err)
	}

	// Create a ScopedIdentifiersRequest for the hub
	scopedIdentifiers := ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set org and project identifiers if they exist
	if orgID != "" {
		scopedIdentifiers.OrgIdentifier = &orgID
	}
	if projectID != "" {
		scopedIdentifiers.ProjectIdentifier = &projectID
	}

	// Set the resource ID using the hub's ID
	d.SetId(hub.ID)

	// Set the resource attributes
	d.Set("name", hub.Name)
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)
	d.Set("description", hub.Description)
	d.Set("connector_id", hub.ConnectorID)
	d.Set("connector_scope", hub.ConnectorScope)
	d.Set("repo_branch", hub.RepoBranch)
	d.Set("repo_name", hub.RepoName)
	d.Set("is_default", hub.IsDefault)
	d.Set("created_at", hub.CreatedAt)
	d.Set("updated_at", hub.UpdatedAt)
	d.Set("last_synced_at", hub.LastSyncedAt)
	d.Set("is_available", hub.IsAvailable)
	d.Set("total_experiments", hub.TotalExperiments)
	d.Set("total_faults", hub.TotalFaults)

	// Set tags if they exist
	if len(hub.Tags) > 0 {
		d.Set("tags", hub.Tags)
	}

	return []*schema.ResourceData{d}, nil
}

// ScopedIdentifiersRequest represents the identifiers for a scoped resource
type ScopedIdentifiersRequest struct {
	AccountIdentifier string  `json:"accountIdentifier"`
	OrgIdentifier     *string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier *string `json:"projectIdentifier,omitempty"`
}

// ChaosHubRequestWrapper wraps the model.ChaosHubRequest to include additional fields
type ChaosHubRequestWrapper struct {
	model.ChaosHubRequest
	IsDefault bool `json:"isDefault,omitempty"`
}

func resourceChaosHubCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient

	connectorScope := model.ConnectorScopeProject
	if v, ok := d.GetOk("connector_scope"); ok {
		connectorScope = model.ConnectorScope(v.(string))
	}
	req := &ChaosHubRequestWrapper{
		ChaosHubRequest: model.ChaosHubRequest{
			HubName:        d.Get("name").(string),
			ConnectorID:    d.Get("connector_id").(string),
			RepoBranch:     d.Get("repo_branch").(string),
			ConnectorScope: connectorScope,
		},
		IsDefault: d.Get("is_default").(bool),
	}

	if v, ok := d.GetOk("repo_name"); ok {
		repoName := v.(string)
		req.RepoName = &repoName
	}

	if v, ok := d.GetOk("description"); ok {
		desc := v.(string)
		req.Description = &desc
	}

	if v, ok := d.GetOk("tags"); ok {
		tags := make([]string, len(v.([]interface{})))
		for i, tag := range v.([]interface{}) {
			tags[i] = tag.(string)
		}
		req.Tags = tags
	}

	identifiers := getIdentifiers(d, c.AccountId)
	hubClient := chaos.NewChaosHubClient(c)

	// Convert ScopedIdentifiersRequest to model.IdentifiersRequest
	modelIdentifiers := model.IdentifiersRequest{
		AccountIdentifier: identifiers.AccountIdentifier,
	}
	if identifiers.OrgIdentifier != nil {
		orgID := *identifiers.OrgIdentifier
		modelIdentifiers.OrgIdentifier = orgID
	}
	if identifiers.ProjectIdentifier != nil {
		projectID := *identifiers.ProjectIdentifier
		modelIdentifiers.ProjectIdentifier = projectID
	}

	// Call the Create method with the required parameters
	hub, err := hubClient.Create(
		ctx,
		req.HubName,
		req.RepoBranch,
		req.ConnectorID,
		modelIdentifiers,
		// Add optional parameters using functional options
		func(r *model.ChaosHubRequest) *model.ChaosHubRequest {
			if req.RepoName != nil {
				r.RepoName = req.RepoName
			}
			if req.ConnectorScope != "" {
				r.ConnectorScope = req.ConnectorScope
			}
			if req.Description != nil {
				r.Description = req.Description
			}
			if req.Tags != nil {
				r.Tags = req.Tags
			}
			return r
		},
	)
	if err != nil {
		return diag.Errorf("failed to create chaos hub: %v", err)
	}

	d.SetId(hub.ID)
	log.Printf("[DEBUG] Created Chaos Hub with ID: %s", d.Id())

	return resourceChaosHubRead(ctx, d, meta)
}

func resourceChaosHubRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Reading Chaos Hub with ID: %s", d.Id())

	c := meta.(*internal.Session).ChaosClient

	hubID := d.Id()
	identifiers := getIdentifiers(d, c.AccountId)
	hubClient := chaos.NewChaosHubClient(c)

	// Convert ScopedIdentifiersRequest to model.IdentifiersRequest
	modelIdentifiers := model.IdentifiersRequest{
		AccountIdentifier: identifiers.AccountIdentifier,
	}
	if identifiers.OrgIdentifier != nil {
		orgID := *identifiers.OrgIdentifier
		modelIdentifiers.OrgIdentifier = orgID
	}
	if identifiers.ProjectIdentifier != nil {
		projectID := *identifiers.ProjectIdentifier
		modelIdentifiers.ProjectIdentifier = projectID
	}

	hub, err := hubClient.Get(ctx, modelIdentifiers, hubID)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			log.Printf("[WARN] Chaos Hub %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to read chaos hub: %v", err)
	}

	d.Set("name", hub.Name)
	d.Set("org_id", identifiers.OrgIdentifier)
	d.Set("project_id", identifiers.ProjectIdentifier)
	d.Set("id", hub.ID)
	d.Set("connector_id", hub.ConnectorID)
	d.Set("connector_scope", hub.ConnectorScope.String())
	d.Set("repo_branch", hub.RepoBranch)
	d.Set("is_default", hub.IsDefault)
	d.Set("created_at", hub.CreatedAt)
	d.Set("updated_at", hub.UpdatedAt)
	d.Set("last_synced_at", hub.LastSyncedAt)
	d.Set("is_available", hub.IsAvailable)
	d.Set("total_experiments", hub.TotalExperiments)
	d.Set("total_faults", hub.TotalFaults)

	if hub.RepoName != nil {
		d.Set("repo_name", *hub.RepoName)
	}
	if hub.Description != nil {
		d.Set("description", *hub.Description)
	}
	if len(hub.Tags) > 0 {
		d.Set("tags", hub.Tags)
	}

	return nil
}

func resourceChaosHubUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Updating Chaos Hub with ID: %s", d.Id())

	c := meta.(*internal.Session).ChaosClient
	identifiers := getIdentifiers(d, c.AccountId)
	hubID := d.Id()

	hubClient := chaos.NewChaosHubClient(c)

	// Convert ScopedIdentifiersRequest to model.IdentifiersRequest
	modelIdentifiers := model.IdentifiersRequest{
		AccountIdentifier: identifiers.AccountIdentifier,
	}
	if identifiers.OrgIdentifier != nil {
		orgID := *identifiers.OrgIdentifier
		modelIdentifiers.OrgIdentifier = orgID
	}
	if identifiers.ProjectIdentifier != nil {
		projectID := *identifiers.ProjectIdentifier
		modelIdentifiers.ProjectIdentifier = projectID
	}

	// Get the current state
	currentHub, err := hubClient.Get(ctx, modelIdentifiers, hubID)
	if err != nil {
		return diag.Errorf("failed to get current chaos hub: %v", err)
	}

	// Initialize request with required fields
	req := &model.ChaosHubRequest{
		HubName:     d.Get("name").(string),
		RepoBranch:  d.Get("repo_branch").(string),
		ConnectorID: d.Get("connector_id").(string),
	}

	// Handle optional fields
	if v, ok := d.GetOk("description"); ok {
		desc := v.(string)
		req.Description = &desc
	}

	if v, ok := d.GetOk("repo_name"); ok {
		repoName := v.(string)
		req.RepoName = &repoName
	}

	if v, ok := d.GetOk("connector_scope"); ok {
		req.ConnectorScope = model.ConnectorScope(v.(string))
	} else {
		req.ConnectorScope = model.ConnectorScopeProject
	}

	// Handle tags
	if v, ok := d.GetOk("tags"); ok {
		tagsRaw := v.([]interface{})
		tags := make([]string, len(tagsRaw))
		for i, v := range tagsRaw {
			tags[i] = v.(string)
		}
		req.Tags = tags
	}

	// Perform the update
	updatedHub, err := hubClient.Update(
		ctx,
		hubID,
		req.HubName,
		req.RepoBranch,
		req.ConnectorID,
		modelIdentifiers,
		// Use functional options for additional fields
		func(r *model.ChaosHubRequest) *model.ChaosHubRequest {
			if req.Description != nil {
				r.Description = req.Description
			}
			if req.RepoName != nil {
				r.RepoName = req.RepoName
			}
			if req.ConnectorScope != "" {
				r.ConnectorScope = req.ConnectorScope
			}
			if req.Tags != nil {
				r.Tags = req.Tags
			}
			return r
		},
	)

	if err != nil {
		return diag.Errorf("failed to update chaos hub: %v", err)
	}

	// Update the resource ID if the name has changed
	if updatedHub.Name != currentHub.Name {
		log.Printf("[DEBUG] Updating resource ID from %s to %s due to name change", d.Id(), updatedHub.ID)
		d.SetId(updatedHub.ID)
	}

	return resourceChaosHubRead(ctx, d, meta)
}

func resourceChaosHubDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Deleting Chaos Hub with ID: %s", d.Id())

	c := meta.(*internal.Session).ChaosClient

	identifiers := getIdentifiers(d, c.AccountId)
	hubID := d.Id()

	log.Printf("[DEBUG] Parsed ID - Account: %s, Org: %v, Project: %v, HubID: %s",
		identifiers.AccountIdentifier,
		identifiers.OrgIdentifier,
		identifiers.ProjectIdentifier,
		hubID)

	hubClient := chaos.NewChaosHubClient(c)

	// Convert ScopedIdentifiersRequest to model.IdentifiersRequest
	modelIdentifiers := model.IdentifiersRequest{
		AccountIdentifier: identifiers.AccountIdentifier,
	}
	if identifiers.OrgIdentifier != nil {
		orgID := *identifiers.OrgIdentifier
		modelIdentifiers.OrgIdentifier = orgID
	}
	if identifiers.ProjectIdentifier != nil {
		projectID := *identifiers.ProjectIdentifier
		modelIdentifiers.ProjectIdentifier = projectID
	}

	// Try to get the hub first to check if it exists
	_, err := hubClient.Get(ctx, modelIdentifiers, hubID)
	if err != nil {
		// If we get a "not found" error, the resource is already deleted
		if strings.Contains(strings.ToLower(err.Error()), "not found") ||
			strings.Contains(strings.ToLower(err.Error()), "no matching documents") {
			log.Printf("[WARN] Chaos Hub %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to get chaos hub for deletion: %v", err)
	}

	// If we get here, the hub exists, so proceed with deletion
	deleted, err := hubClient.Delete(
		ctx,
		hubID,
		modelIdentifiers,
	)

	if err != nil {
		// If we get a "not found" error, the resource might have been deleted out of band
		if strings.Contains(strings.ToLower(err.Error()), "not found") ||
			strings.Contains(strings.ToLower(err.Error()), "no matching documents") {
			log.Printf("[WARN] Chaos Hub %s not found during deletion, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to delete chaos hub: %v", err)
	}

	if !deleted {
		return diag.Errorf("failed to delete chaos hub: unknown error")
	}

	// Clear the ID from state
	d.SetId("")
	return nil
}

// Helper functions
func getIdentifiers(d *schema.ResourceData, accountID string) ScopedIdentifiersRequest {
	identifiers := ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}

	if v, ok := d.GetOk("org_id"); ok {
		orgID := v.(string)
		identifiers.OrgIdentifier = &orgID
	}

	if v, ok := d.GetOk("project_id"); ok {
		projectID := v.(string)
		identifiers.ProjectIdentifier = &projectID
	}

	return identifiers
}

func generateID(identifiers ScopedIdentifiersRequest, hubID string) string {
	parts := []string{identifiers.AccountIdentifier}

	// Always include org and project, even if empty
	orgID := ""
	if identifiers.OrgIdentifier != nil {
		orgID = *identifiers.OrgIdentifier
	}
	parts = append(parts, orgID)

	projectID := ""
	if identifiers.ProjectIdentifier != nil {
		projectID = *identifiers.ProjectIdentifier
	}
	parts = append(parts, projectID)

	// Ensure hubID is not empty
	if hubID == "" {
		log.Printf("[WARN] Generating ID with empty hubID")
	}

	parts = append(parts, hubID)
	return strings.Join(parts, "/")
}

func parseID(id string, accountID string) (ScopedIdentifiersRequest, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 {
		return ScopedIdentifiersRequest{}, "", fmt.Errorf("invalid ID format: expected org_id/project_id/hub_id, got: %s", id)
	}

	identifiers := ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}
	if parts[0] != "" {
		identifiers.OrgIdentifier = &parts[0]
	}
	if parts[1] != "" {
		identifiers.ProjectIdentifier = &parts[1]
	}

	return identifiers, parts[3], nil
}
