package connector

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCCM32488Class_KubernetesCloudCostReadNoPanic(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceConnectorKubernetesCloudCost().Schema, map[string]interface{}{})
	panicked, val := capturePanic(func() {
		_ = readConnectorKubernetesCloudCost(d, &nextgen.ConnectorInfo{
			K8sClusterCloudCost: &nextgen.CeKubernetesClusterConfigDto{
				ConnectorRef:    "ref",
				FeaturesEnabled: []string{"VISIBILITY"},
			},
		})
	})
	if panicked {
		t.Errorf("unexpected panic: %v", val)
	}
}
