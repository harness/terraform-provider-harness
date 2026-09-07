package helpers

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
)

func TestHandleEmptyCreateUpdateResponse(t *testing.T) {
	t.Run("denied by OPA policy with evaluation id", func(t *testing.T) {
		diags := HandleEmptyCreateUpdateResponse(&nextgen.GovernanceMetadata{
			Deny:    true,
			Id:      "68029671",
			Message: "Harness Secret Manager is not allowed.",
		}, "secret")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}

		got := diags[0].Summary
		want := "secret operation denied by OPA policy evaluation. Evaluation ID: 68029671 — view details under Settings > Governance > Policy Evaluations in the Harness UI"
		if got != want {
			t.Fatalf("unexpected summary:\ngot:  %q\nwant: %q", got, want)
		}
	})

	t.Run("denied by OPA policy without evaluation id", func(t *testing.T) {
		diags := HandleEmptyCreateUpdateResponse(&nextgen.GovernanceMetadata{
			Deny: true,
		}, "connector")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}

		got := diags[0].Summary
		want := "connector operation denied by OPA policy evaluation"
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

func TestCheckGovernanceMetadata(t *testing.T) {
	t.Run("nil metadata returns no error", func(t *testing.T) {
		diags := CheckGovernanceMetadata(nil, "user group")
		if diags != nil {
			t.Fatalf("expected nil diagnostics, got: %v", diags)
		}
	})

	t.Run("deny=false returns no error", func(t *testing.T) {
		diags := CheckGovernanceMetadata(&nextgen.GovernanceMetadata{
			Deny:   false,
			Status: "pass",
		}, "user group")
		if diags != nil {
			t.Fatalf("expected nil diagnostics, got: %v", diags)
		}
	})

	t.Run("deny=true with deny messages returns formatted error", func(t *testing.T) {
		diags := CheckGovernanceMetadata(&nextgen.GovernanceMetadata{
			Deny: true,
			Details: []nextgen.PolicySetMetadata{
				{
					PolicyMetadata: []nextgen.PolicyMetadata{
						{
							PolicyName:   "P_UG_Combined",
							Status:       "error",
							DenyMessages: []string{"identifier must match '<name>_team'"},
						},
					},
				},
			},
		}, "user group")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}
		got := diags[0].Summary
		if got != "OPA policy evaluation failed for user group: [P_UG_Combined] identifier must match '<name>_team'" {
			t.Fatalf("unexpected summary: %q", got)
		}
	})

	t.Run("deny=true with no messages returns fallback error", func(t *testing.T) {
		diags := CheckGovernanceMetadata(&nextgen.GovernanceMetadata{
			Deny: true,
		}, "user group")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}
		got := diags[0].Summary
		if got != "OPA policy evaluation failed for user group: Policy evaluation denied the request" {
			t.Fatalf("unexpected summary: %q", got)
		}
	})

	t.Run("deny=true with multiple policies concatenates messages", func(t *testing.T) {
		diags := CheckGovernanceMetadata(&nextgen.GovernanceMetadata{
			Deny: true,
			Details: []nextgen.PolicySetMetadata{
				{
					PolicyMetadata: []nextgen.PolicyMetadata{
						{
							PolicyName:   "P_UG_Naming",
							Status:       "error",
							DenyMessages: []string{"identifier must match '<name>_team'"},
						},
						{
							PolicyName:   "P_UG_Min",
							Status:       "warning",
							DenyMessages: []string{"cannot have zero members"},
						},
					},
				},
			},
		}, "user group")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}
		got := diags[0].Summary
		if got != "OPA policy evaluation failed for user group: [P_UG_Naming] identifier must match '<name>_team'; [P_UG_Min] cannot have zero members" {
			t.Fatalf("unexpected summary: %q", got)
		}
	})

	t.Run("policies with passing status are not included in messages", func(t *testing.T) {
		diags := CheckGovernanceMetadata(&nextgen.GovernanceMetadata{
			Deny: true,
			Details: []nextgen.PolicySetMetadata{
				{
					PolicyMetadata: []nextgen.PolicyMetadata{
						{
							PolicyName:   "P_UG_Pass",
							Status:       "pass",
							DenyMessages: []string{"should not appear"},
						},
						{
							PolicyName:   "P_UG_Fail",
							Status:       "error",
							DenyMessages: []string{"identifier must match '<name>_team'"},
						},
					},
				},
			},
		}, "user group")
		if !diags.HasError() {
			t.Fatal("expected error diagnostic")
		}
		got := diags[0].Summary
		if got != "OPA policy evaluation failed for user group: [P_UG_Fail] identifier must match '<name>_team'" {
			t.Fatalf("unexpected summary: %q", got)
		}
	})
}
