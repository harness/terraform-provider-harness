package security_governance_v3

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosSecurityGovernanceConditionV3() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing a Harness Chaos Security Governance Condition (V3 / REST API).",
		CreateContext: resourceChaosSecurityGovernanceConditionV3Create,
		ReadContext:   resourceChaosSecurityGovernanceConditionV3Read,
		UpdateContext: resourceChaosSecurityGovernanceConditionV3Update,
		DeleteContext: resourceChaosSecurityGovernanceConditionV3Delete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceChaosSecurityGovernanceConditionV3Import,
		},
		Schema: resourceChaosSecurityGovernanceConditionV3Schema(),
	}
}

// resourceChaosSecurityGovernanceConditionV3Import imports an existing condition.
// Expected format: <org_id>/<project_id>/<condition_id>
func resourceChaosSecurityGovernanceConditionV3Import(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	importID := d.Id()
	parts := strings.Split(importID, "/")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid import ID format. Expected \"<org-id>/<project-id>/<condition-id>\", got: %s", importID)
	}

	orgID := parts[0]
	projectID := parts[1]
	conditionID := parts[2]

	if orgID == "" || projectID == "" || conditionID == "" {
		return nil, fmt.Errorf("org_id, project_id, and condition_id cannot be empty")
	}

	log.Printf("[DEBUG] Importing chaos security governance condition: %s (org: %s, project: %s)", conditionID, orgID, projectID)

	d.SetId(conditionID)
	if err := d.Set("org_id", orgID); err != nil {
		return nil, fmt.Errorf("failed to set org_id: %w", err)
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, fmt.Errorf("failed to set project_id: %w", err)
	}

	diags := resourceChaosSecurityGovernanceConditionV3Read(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read condition during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}
