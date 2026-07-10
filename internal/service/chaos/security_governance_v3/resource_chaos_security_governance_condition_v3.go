package security_governance_v3

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosSecurityGovernanceConditionV3() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing a Harness Chaos Security Governance Condition (V3 / REST API). " +
			"A condition defines which faults and infrastructure a governance rule applies to.\n\n" +
			"## Usage notes\n\n" +
			"- `infra_type` (required, immutable) accepts `Kubernetes`, `KubernetesV2`, `Linux`, `Windows`, `CloudFoundry`, `Container`.\n" +
			"- `fault_spec` is always required.\n" +
			"- Use `k8s_spec` when `infra_type` is `Kubernetes` or `KubernetesV2`; use `machine_spec` when `infra_type` is `Linux` or `Windows`. " +
			"For `CloudFoundry`/`Container`, only `fault_spec` applies.\n" +
			"- All `operator` fields accept `EQUAL_TO` or `NOT_EQUAL_TO`.\n\n" +
			"## Behavior notes\n\n" +
			"- `fault_spec.fault_type`: the legacy value `FAULT_NAME` is normalized to `FAULT` to stay " +
			"consistent with the Harness UI, GraphQL, and existing conditions. If you set `FAULT_NAME`, " +
			"it is stored and read back as `FAULT`. Use `FAULT` or `FAULT_GROUP`.\n\n" +
			"## Import\n\n" +
			"Import uses the 3-part ID `org_id/project_id/condition_id`.\n",
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
	orgID, projectID, conditionID, err := parseScopedImportIDV3(d.Id(), "condition-id")
	if err != nil {
		return nil, err
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
