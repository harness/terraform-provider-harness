package helpers

import (
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func TestHandleEmptyCreateUpdateResponse(t *testing.T) {
	t.Run("denied by OPA policy with evaluation id and policy name in details", func(t *testing.T) {
		diags := HandleEmptyCreateUpdateResponse(&nextgen.GovernanceMetadata{
			Deny:    true,
			Id:      "68029671",
			Message: "Harness Secret Manager is not allowed.",
			DetailsList: []nextgen.PolicySetMetadata{
				{
					PolicySetName: "Secret - Restrict Harness Built-in Services",
					PolicyMetadataList: []nextgen.PolicyMetadata{
						{
							PolicyName: "Secret_Block_Harness_Secret_Manager",
							Status:     "FAILED",
						},
					},
				},
			},
		}, "secret")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}

		got := diags[0].Summary
		if !strings.Contains(got, "Secret_Block_Harness_Secret_Manager") {
			t.Fatalf("expected policy name in message, got: %q", got)
		}
		if !strings.Contains(got, "68029671") {
			t.Fatalf("expected evaluation ID in message, got: %q", got)
		}
		if strings.Contains(got, "OPA policy \"68029671\"") {
			t.Fatalf("evaluation ID must not be labeled as policy name, got: %q", got)
		}
		if !strings.Contains(got, "Harness Secret Manager is not allowed.") {
			t.Fatalf("expected denial message, got: %q", got)
		}
	})

	t.Run("denied without details list", func(t *testing.T) {
		diags := HandleEmptyCreateUpdateResponse(&nextgen.GovernanceMetadata{
			Deny:    true,
			Id:      "68029671",
			Message: "Harness Secret Manager is not allowed.",
		}, "secret")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}

		got := diags[0].Summary
		want := "secret operation denied by OPA policy evaluation: Harness Secret Manager is not allowed. Evaluation ID: 68029671 — view details under Settings > Governance > Policy Evaluations in the Harness UI"
		if got != want {
			t.Fatalf("unexpected summary:\ngot:  %q\nwant: %q", got, want)
		}
	})

	t.Run("empty response without governance metadata", func(t *testing.T) {
		diags := HandleEmptyCreateUpdateResponse(nil, "connector")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}
		if got := diags[0].Summary; got != "connector operation blocked: API returned empty response. Check governance policies in the Harness UI." {
			t.Fatalf("unexpected summary: %q", got)
		}
	})
}
