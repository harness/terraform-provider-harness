package helpers

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

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
