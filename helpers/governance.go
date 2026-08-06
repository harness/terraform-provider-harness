package helpers

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// HandleEmptyCreateUpdateResponse returns a diagnostic when a create/update API call
// succeeds but returns no entity data, typically due to OPA policy denial.
func HandleEmptyCreateUpdateResponse(governance *nextgen.GovernanceMetadata, resourceKind string) diag.Diagnostics {
	if governance != nil && governance.Deny {
		return diag.Errorf("%s", formatGovernanceDenial(governance, resourceKind))
	}
	return diag.Errorf("%s operation blocked: API returned empty response. Check governance policies in the Harness UI.", resourceKind)
}

func formatGovernanceDenial(governance *nextgen.GovernanceMetadata, resourceKind string) string {
	msg := fmt.Sprintf("%s operation denied by OPA policy evaluation", resourceKind)

	if governance.Id != "" {
		msg += fmt.Sprintf(". Evaluation ID: %s — view details under Settings > Governance > Policy Evaluations in the Harness UI", governance.Id)
	}

	return msg
}
