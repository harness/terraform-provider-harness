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
