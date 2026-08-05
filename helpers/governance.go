package helpers

import (
	"fmt"
	"strings"

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

func CheckGovernanceMetadata(gm *nextgen.UserGroupGovernanceMetadata, entityType string) diag.Diagnostics {
	if gm == nil || !gm.Deny {
		return nil
	}

	messages := extractDenyMessages(gm)
	if len(messages) == 0 {
		messages = append(messages, "Policy evaluation denied the request")
	}

	return diag.Errorf("OPA policy evaluation failed for %s: %s", entityType, strings.Join(messages, "; "))
}

func extractDenyMessages(gm *nextgen.UserGroupGovernanceMetadata) []string {
	var messages []string
	for _, detail := range gm.Details {
		for _, policy := range detail.PolicyMetadata {
			if policy.Status == "error" || policy.Status == "warning" {
				for _, msg := range policy.DenyMessages {
					messages = append(messages, fmt.Sprintf("[%s] %s", policy.PolicyName, msg))
				}
			}
		}
	}
	return messages
}
