package cluster_orchestrator

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func capturePanic(fn func()) (panicked bool, value interface{}) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			value = r
		}
	}()
	fn()
	return panicked, value
}

func TestCCM32488Class_ReadClusterOrchConfigNilNested(t *testing.T) {
	schemaMap := ResourceClusterOrchestratorConfig().Schema

	cases := []struct {
		name string
		orch *nextgen.ClusterOrchestrator
	}{
		{
			name: "config null",
			orch: &nextgen.ClusterOrchestrator{
				ID:     "orch-1",
				Config: nil,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, schemaMap, map[string]interface{}{
				"orchestrator_id": "orch-1",
			})
			panicked, val := capturePanic(func() {
				readClusterOrchConfig(d, tc.orch)
			})
			if panicked {
				t.Errorf("readClusterOrchConfig panicked (bug confirmed): %v", val)
			} else {
				t.Log("RESULT: PASS (no panic)")
			}
		})
	}
}
