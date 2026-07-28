package cluster_orchestrator

// This file is intentionally in package cluster_orchestrator (white-box), not
// cluster_orchestrator_test like clusterorch_config_test.go: it calls the
// unexported readClusterOrchConfig directly. It cannot be merged into
// clusterorch_config_test.go because that file imports internal/acctest, which
// imports internal/provider, which imports this package - moving it here would
// create an import cycle. Deterministic; no live API / TF_ACC required.

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestReadClusterOrchConfig_NilConfig guards against a panic when the Cluster
// Orchestrator Details API returns a null/empty "config" (e.g. partial create,
// disabled orchestrator, or an API gap). Config is a *ClusterOrchConfig in the
// SDK model, so it can be nil; readClusterOrchConfig and its helpers must not
// dereference it unconditionally.
func TestReadClusterOrchConfig_NilConfig(t *testing.T) {
	tests := []struct {
		name string
		orch *nextgen.ClusterOrchestrator
	}{
		{
			name: "config_null",
			orch: &nextgen.ClusterOrchestrator{
				ID:       "orch-nil-config",
				Disabled: true,
				Config:   nil,
			},
		},
		{
			name: "config_present_minimal",
			orch: &nextgen.ClusterOrchestrator{
				ID:       "orch-minimal-config",
				Disabled: false,
				Config:   &nextgen.ClusterOrchConfig{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, ResourceClusterOrchestratorConfig().Schema, map[string]interface{}{})

			defer func() {
				if rec := recover(); rec != nil {
					t.Fatalf("readClusterOrchConfig panicked (bug confirmed): %v", rec)
				}
			}()

			readClusterOrchConfig(d, tt.orch)

			if d.Id() != tt.orch.ID {
				t.Errorf("expected id %q, got %q", tt.orch.ID, d.Id())
			}
		})
	}
}
